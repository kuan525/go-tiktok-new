package db

import (
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/errno"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreateAt    time.Time      `json:"create_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}

func AddNewComment(comment *Comment) error {
	if ok, _ := CheckUserExistById(comment.UserId); !ok {
		return errno.UserIsNotExistErr
	}
	if ok, _ := CheckVideoExistById(comment.VideoId); !ok {
		return errno.VideoIsNotExistErr
	}
	if err := DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentById(commentId int64) error {
	if ok, _ := CheckCommentExist(commentId); !ok {
		return errno.CommentIsNotExistErr
	}
	comment := Comment{}
	if err := DB.Where("id = ?", commentId).Delete(comment).Error; err != nil {
		return err
	}
	return nil
}

func CheckCommentExist(commentId int64) (bool, error) {
	comment := &Comment{}
	if err := DB.Where("id = ?", commentId).Find(comment).Error; err != nil {
		return false, err
	}

	if comment.ID == 0 {
		return false, nil
	}
	return true, nil
}

func GetCommentListByVideoId(videoId int64) ([]*Comment, error) {
	var commentList []*Comment
	if ok, _ := CheckVideoExistById(videoId); !ok {
		return commentList, errno.VideoIsNotExistErr
	}
	err := DB.Model(&Comment{}).Where("video_id = ?", videoId).Find(commentList).Error
	if err != nil {
		return commentList, err
	}
	return commentList, err
}

func GetCommentCountByVideoId(videoId int64) (int64, error) {
	var sum int64
	err := DB.Model(&Comment{}).Where("video_id = ?", videoId).Count(&sum).Error
	if err != nil {
		return sum, err
	}
	return sum, nil
}
