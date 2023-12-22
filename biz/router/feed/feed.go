package feed

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/feed"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feedMw(), feed.Feed)...)
		}
	}
}
