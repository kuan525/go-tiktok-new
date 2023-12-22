package favorite

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

func _publishMw() []app.HandlerFunc {
	return nil
}

func _favoriteActionMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _favoriteListMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _favoriteMw() []app.HandlerFunc {
	return nil
}

func _listMw() []app.HandlerFunc {
	return nil
}

func _actionMw() []app.HandlerFunc {
	return nil
}
