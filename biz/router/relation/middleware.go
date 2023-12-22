package relation

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

func _relationMw() []app.HandlerFunc {
	return nil
}

func _actionMw() []app.HandlerFunc {
	return nil
}

func _relationActionMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _followMw() []app.HandlerFunc {
	return nil
}

func _listMw() []app.HandlerFunc {
	return nil
}

func _relationFollowListMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _followerMw() []app.HandlerFunc {
	return nil
}

func _relationFollowerListMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _friendMw() []app.HandlerFunc {
	return nil
}

func _relationFriendListMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}
