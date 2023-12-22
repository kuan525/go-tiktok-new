package feed

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-tiktok-new/biz/mw/jwt"
)

func rootMw() []app.HandlerFunc {
	return nil
}

func _douyinMw() []app.HandlerFunc {
	return nil
}

func _feedMw() []app.HandlerFunc {
	return nil
}

func feedMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if len(token) == 0 {
			return
		} else {
			jwt.JwtMiddleware.MiddlewareFunc()
			return
		}
	}
}
