package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go-tiktok-new/biz/dal"
	"go-tiktok-new/biz/mw/jwt"
	"go-tiktok-new/biz/mw/minio"
	"go-tiktok-new/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/reverseproxy"
)

func Init() {
	dal.Init()
	jwt.Init()
	minio.Init()
}

// minioReverseProxy 资源反向代理
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(constants.MinioReverseProxyHost)
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func main() {
	Init()

	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts(constants.HertzMainHost),
	)
	h.GET("src/*name", minioReverseProxy)

	register(h)

	h.Spin()
}
