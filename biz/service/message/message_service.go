package message

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/social/message"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

type MessageService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewMessageService(ctx context.Context, c *app.RequestContext) *MessageService {
	return &MessageService{
		ctx: ctx,
		c:   c,
	}
}

func (m *MessageService) GetMessageChat(req *message.DouyinMessageChatRequest) ([]*message.Message, error) {
	messages := make([]*message.Message, 0)
	fromUserId, _ := m.c.Get("current_user_id")
	toUserId := req.ToUserId
	preMsgTime := req.PreMsgTime
	dbMseeages, err := db.GetMessageByIdPair(fromUserId.(int64), toUserId, utils.MillTimeStampToTime(preMsgTime))
	if err != nil {
		return messages, err
	}

	for _, dbMessage := range dbMseeages {
		messages = append(messages, &message.Message{
			Id:         dbMessage.ID,
			ToUserId:   dbMessage.ToUserId,
			FromUserId: dbMessage.FromUserId,
			Content:    dbMessage.Content,
			CreateTime: dbMessage.CreateAt.UnixNano() / 1000000,
		})
	}
	return messages, nil
}

func (m *MessageService) MessageAction(req *message.DouyinMessageActionRequest) error {
	fromUserId, _ := m.c.Get("current_user_id")
	toUserId := req.ToUserId
	content := req.Content

	ok, err := db.AddNewMessage(&db.Messages{
		FromUserId: fromUserId.(int64),
		ToUserId:   toUserId,
		Content:    content,
	})

	if err != nil {
		return err
	}
	if !ok {
		return errno.MessageAddFailedErr
	}

	return nil
}
