package feed

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/basic/feed"
	feed_service "go-tiktok-new/biz/service/feed"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// Feed
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, feed.DouyinFeedResponse{
			StatusMsg:  resp.StatusMsg,
			StatusCode: resp.StatusCode,
		})
		return
	}

	resp, err := feed_service.NewFeedService(ctx, c).Feed(&req)
	if err != nil {
		bresp := utils.BuildBaseResp(err)
		resp.StatusMsg = bresp.StatusMsg
		resp.StatusCode = bresp.StatusCode
		c.JSON(consts.StatusOK, resp)
	}
	resp.StatusMsg = errno.SuccessMsg
	resp.StatusCode = errno.SuccessCode
	c.JSON(consts.StatusOK, resp)
}
