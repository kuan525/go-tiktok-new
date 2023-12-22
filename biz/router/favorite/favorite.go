package favorite

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/favorite"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_favorite := _douyin.Group("favorite", _favoriteMw()...)
			{
				_action := _favorite.Group("/action", _actionMw()...)
				_action.POST("/", append(_favoriteActionMw(), favorite.FavoriteAction)...)
			}
			{
				_list := _favorite.Group("/list", _listMw()...)
				_list.GET("/", append(_favoriteActionMw(), favorite.FavoriteList)...)
			}
		}
	}
}
