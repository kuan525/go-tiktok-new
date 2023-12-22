package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"go-tiktok-new/biz/mw/jwt"
)

func rootMw() []app.HandlerFunc {
	return nil
}

func _douyinMw() []app.HandlerFunc {
	return nil
}

func _userMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	return nil
}

func _userLoginMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.LoginHandler,
	}
}

func _registerMw() []app.HandlerFunc {
	return nil
}

func _userRegisterMw() []app.HandlerFunc {
	return nil
}
