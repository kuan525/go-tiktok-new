package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/router/comment"
	"go-tiktok-new/biz/router/favorite"
	"go-tiktok-new/biz/router/feed"
	"go-tiktok-new/biz/router/message"
	"go-tiktok-new/biz/router/publish"
	"go-tiktok-new/biz/router/relation"
	"go-tiktok-new/biz/router/user"
)

func GenerateRegister(r *server.Hertz) {
	message.Register(r)
	relation.Register(r)
	comment.Register(r)
	favorite.Register(r)
	feed.Register(r)
	publish.Register(r)
	user.Register(r)
}
