package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/basic/user"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identity      = "userId"
)

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte("tiktok secret key"),
		// 中间件会在查询字符串（query parameters）和表单中（form body）查找名为 "token" 的字段
		TokenLookup: "query:token,from:token",
		Timeout:     24 * time.Hour,
		// 存储用户唯一标识
		IdentityKey: identity,

		// 进行身份验证
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.DouyinUserLoginRequest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			user, err := db.QueryUser(loginRequest.Username)
			if ok := utils.VerifyPassword(loginRequest.Password, user.Password); !ok {
				err = errno.PasswordIsNotVerified
				return nil, err
			}
			if err != nil {
				return nil, err
			}
			c.Set("userId", user.ID)
			return user.ID, nil
		},

		// 生成负载Payload部分，这个data是Authenticator的返回值
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identity: v,
				}
			}
			return jwt.MapClaims{}
		},
		//登陆成功后，将令牌放在上下文中，这个token是在PayloadFunc运行玩之后内部生成的
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success, token is issued clientIP: "+c.ClientIP())
			c.Set("token", token)
		},
		//验证用户的权限，接受一个jwt令牌的负载，验证其中的data是否有效
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//处理json数据时，会将int等数据弄成float64类型，string则为string
			if v, ok := data.(float64); ok {
				currentUserId := int64(v)
				c.Set("current_user_id", currentUserId)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		// 用户未授权时，返回一个错误信息
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  message,
			})
		},
		// 根据发生的错误情况，返回一个相对应的消息状态，这个消息被用来通知客户端发生了什么错误，帮助排查问题
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			resp := utils.BuildBaseResp(e)
			return resp.StatusMsg
		},
	})
}
