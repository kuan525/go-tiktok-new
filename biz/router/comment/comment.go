package comment

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/comment"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_commentActionMw(), comment.CommentAction)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_commentListMw(), comment.CommentList)...)
			}
		}
	}
}
