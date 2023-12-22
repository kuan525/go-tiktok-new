package message

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler/message"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_message := _douyin.Group("/message", _messageMw()...)
			{
				_action := _message.Group("/action", _actionMw()...)
				_action.POST("/", append(_messageActionMw(), message.MessageAction)...)
			}
			{
				_chat := _message.Group("/chat", _chatMw()...)
				_chat.GET("/", append(_messageChatMw(), message.MessageChat)...)
			}
		}
	}
}
