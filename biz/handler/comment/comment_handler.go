package comment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/interact/comment"
	comment_service "go-tiktok-new/biz/service/comment"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// CommentAction
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, comment.DouyinCommentActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_comment, err := comment_service.NewCommentService(ctx, c).AddNewComment(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, comment.DouyinCommentActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, comment.DouyinCommentActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Comment:    _comment,
	})
}

// CommentList
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := comment_service.NewCommentService(ctx, c).CommentList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, comment.DouyinCommentListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}
