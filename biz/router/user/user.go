package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/user"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_userMw(), user.User)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userLoginMw(), user.UserLogin)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userRegisterMw(), user.UserRegister)...)
			}
		}
	}
}
