package publish

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/basic/publish"
	service "go-tiktok-new/biz/service/publish"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// PublishAction
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(errno.ParamErr)
		c.JSON(consts.StatusBadRequest, publish.DouyinPublishActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		resp := utils.BuildBaseResp(errno.ParamErr)
		c.JSON(consts.StatusOK, publish.DouyinPublishListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	}

	err = service.NewPublishService(ctx, c).PublishAction(&req, file.Filename)
	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, publish.DouyinPublishActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}

// PublishList
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewPublishService(ctx, c).PublishList(&req)
	if err != nil {
		bresp := utils.BuildBaseResp(err)
		resp.StatusCode = bresp.StatusCode
		resp.StatusMsg = bresp.StatusMsg
		c.JSON(consts.StatusOK, resp)
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.SuccessMsg
	c.JSON(consts.StatusOK, resp)
}
