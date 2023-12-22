package comment

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

func _commentMw() []app.HandlerFunc {
	return nil
}

func _actionMw() []app.HandlerFunc {
	return nil
}

func _commentActionMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _listMw() []app.HandlerFunc {
	return nil
}

func _commentListMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}
