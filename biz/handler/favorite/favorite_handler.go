package favorite

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/interact/favorite"
	favorite_service "go-tiktok-new/biz/service/favorite"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// FavoriteAction
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	ok, err := favorite_service.NewFavoriteService(ctx, c).FavoriteAction(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	if !ok {
		resp := utils.BuildBaseResp(errno.FavoriteActionErr)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// FavoriteList
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	favoriteList, err := favorite_service.NewFavoriteService(ctx, c).GetGavoriteList(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := favorite.DouyinFavoriteListResponse{
		VideoList:  favoriteList,
		StatusMsg:  errno.SuccessMsg,
		StatusCode: errno.SuccessCode,
	}
	c.JSON(consts.StatusOK, resp)
}
