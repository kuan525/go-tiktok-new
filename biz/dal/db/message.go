package db

import (
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/errno"
	"time"
)

type Messages struct {
	ID         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Content    string    `json:"content"`
	CreateAt   time.Time `json:"create_at"`
}

func (Messages) TableName() string {
	return constants.MessageTableName
}

func AddNewMessage(message *Messages) (bool, error) {
	exist, err := QueryUserById(message.FromUserId)
	if exist == nil || err != nil {
		return false, errno.UserIsNotExistErr
	}
	exist, err = QueryUserById(message.ToUserId)
	if exist == nil || err != nil {
		return false, errno.UserIsNotExistErr
	}
	if err = DB.Create(message).Error; err != nil {
		return false, err
	}
	return true, err
}

func GetMessageByIdPair(userId1, userId2 int64, preMsgTime time.Time) ([]Messages, error) {
	exist, err := QueryUserById(userId1)
	if exist == nil || err != nil {
		return nil, errno.UserIsNotExistErr
	}
	exist, err = QueryUserById(userId2)
	if exist == nil || err != nil {
		return nil, errno.UserIsNotExistErr
	}

	var messages []Messages
	err = DB.Where("to_user_id = ? AND from_user_id = ? AND created_at > ?", userId1, userId2, preMsgTime).
		Or("to_user_id = ? AND from_user_id = ? AND created_at > ?", userId2, userId1, preMsgTime).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetLatestMessageByIdPair(userId1, userId2 int64) (*Messages, error) {
	exist, err := QueryUserById(userId1)
	if exist == nil || err != nil {
		return nil, errno.UserIsNotExistErr
	}
	exist, err = QueryUserById(userId2)
	if exist == nil || err != nil {
		return nil, errno.UserIsNotExistErr
	}
	var message Messages
	err = DB.Where("to_user_id = ? AND from_user_id = ?", userId1, userId2).
		Or("to_user_id = ? AND from_user_id = ?", userId2, userId1).
		Last(&message).Error
	if err == nil {
		return &message, nil
	} else {
		if err.Error() == "record not found" {
			return nil, nil
		} else {
			return nil, err
		}
	}
}
