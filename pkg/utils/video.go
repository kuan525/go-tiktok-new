package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go-tiktok-new/biz/mw/minio"
	"strings"
)

func NewFileName(userId, time int64) string {
	return fmt.Sprintf("%d.%d", userId, time)
}

// URLConvert 将数据库中的路径转换为前端可访问的完整url
func URLConvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = string(c.URI().Scheme())
	u.Host = string(c.URI().Host())
	u.Path = "/src" + u.Path
	return u.String()
}
