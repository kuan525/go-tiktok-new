package message

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/social/message"
	message_service "go-tiktok-new/biz/service/message"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// MessageChat
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
			StatusCode:  resp.StatusCode,
			StatusMsg:   resp.StatusMsg,
			MessageList: []*message.Message{},
		})
		return
	}

	messages, err := message_service.NewMessageService(ctx, c).GetMessageChat(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
			StatusCode:  resp.StatusCode,
			StatusMsg:   resp.StatusMsg,
			MessageList: messages,
		})
		return
	}

	c.JSON(consts.StatusOK, message.DouyinMessageChatResponse{
		StatusCode:  errno.SuccessCode,
		StatusMsg:   errno.SuccessMsg,
		MessageList: messages,
	})
}

// MessageAction
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	err = message_service.NewMessageService(ctx, c).MessageAction(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, message.DouyinMessageActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}
