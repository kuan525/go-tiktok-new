package publish

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/publish"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_publishActionMw(), publish.PublishAction)...)
			}
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_publishListMw(), publish.PublishList)...)
			}
		}
	}
}
