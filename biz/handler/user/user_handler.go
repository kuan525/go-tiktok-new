package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/basic/user"
	"go-tiktok-new/biz/mw/jwt"
	service "go-tiktok-new/biz/service/user"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// UserRegister user registration api
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_, err = service.NewUserService(ctx, c).UserRegister(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	token := c.GetString("token")
	v, _ := c.Get("user_id")
	userId := v.(int64)
	c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     userId,
	})
}

// UserLogin user login api
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get("user_id")
	userId := v.(int64)
	token := c.GetString("token")
	c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     userId,
	})
}

// User get user info
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	u, err := service.NewUserService(ctx, c).UserInfo(&req)

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, user.DouyinUserResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		User:       u,
	})
}
