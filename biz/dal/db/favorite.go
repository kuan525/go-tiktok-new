package db

import (
	"go-tiktok-new/biz/mw/redis"
	"go-tiktok-new/pkg/constants"
	"gorm.io/gorm"
	"time"
)

var rdFavorite redis.Favorite

type Favorites struct {
	ID       int64          `json:"id"`
	UserId   int64          `json:"user_id"`
	VideoId  int64          `json:"video_id"`
	CreateAt time.Time      `json:"create_at"`
	DeleteAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Favorites) TableName() string {
	return constants.FavoritesTableName
}

func AddNewFavorite(favorite *Favorites) (bool, error) {
	err := DB.Create(favorite).Error
	if err != nil {
		return false, err
	}

	if rdFavorite.CheckLiked(favorite.VideoId) {
		rdFavorite.AddLiked(favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckLike(favorite.UserId) {
		rdFavorite.AddLiked(favorite.UserId, favorite.VideoId)
	}

	return true, nil
}

func DeleteFavorite(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).
		Delete(favorite).Error
	if err != nil {
		return false, err
	}

	if rdFavorite.CheckLiked(favorite.VideoId) {
		rdFavorite.DelLiked(favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckLiked(favorite.UserId) {
		rdFavorite.DelLiked(favorite.UserId, favorite.VideoId)
	}
	return true, nil
}

func QueryFavoriteExist(userId, videoId int64) (bool, error) {
	if rdFavorite.CheckLiked(videoId) {
		return rdFavorite.ExistLiked(userId, videoId), nil
	}
	if rdFavorite.CheckLike(userId) {
		return rdFavorite.ExistLiked(userId, videoId), nil
	}

	var sum int64
	err := DB.Model(&Favorites{}).Where("video_id = ? AND user_id = ?", videoId, userId).
		Count(&sum).Error
	if err != nil {
		return false, err
	}
	if sum == 0 {
		return false, nil
	}
	return true, nil
}

func QueryTotalFavoritedByAuthorId(authorId int64) (int64, error) {
	var sum int64
	err := DB.Model(&Favorites{}).Joins("JOIN videos ON likes.video_id = videos.id").
		Where("videos.author_id = ?", authorId).Count(&sum).Error
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func getFavoriteIdList(userId int64) ([]int64, error) {
	var favoriteAction []Favorites
	err := DB.Where("user_id = ?", userId).Find(&favoriteAction).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favoriteAction {
		result = append(result, v.VideoId)
	}
	return result, nil
}

func GetFavoriteIdList(userId int64) ([]int64, error) {
	if rdFavorite.CheckLike(userId) {
		return rdFavorite.GetLiked(userId), nil
	}
	return getFavoriteIdList(userId)
}

func GetFavoriteCountByUserId(userId int64) (int64, error) {
	if rdFavorite.CheckLiked(userId) {
		return rdFavorite.CountLiked(userId)
	}

	videos, err := getFavoriteIdList(userId)
	if err != nil {
		return 0, err
	}

	go func(user int64, videos []int64) {
		for _, videos := range videos {
			rdFavorite.AddLiked(user, videos)
		}
	}(userId, videos)

	return int64(len(videos)), nil
}

func getFavoriterIdList(videoId int64) ([]int64, error) {
	var favoriteActions []Favorites
	err := DB.Where("videoId = ?", videoId).Find(&favoriteActions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favoriteActions {
		result = append(result, v.UserId)
	}
	return result, nil
}

func GetFavoriterIdList(videoId int64) ([]int64, error) {
	if rdFavorite.CheckLiked(videoId) {
		return rdFavorite.GetLiked(videoId), nil
	}
	return getFavoriterIdList(videoId)
}

func GetFavoriteCount(videoId int64) (int64, error) {
	if rdFavorite.CheckLike(videoId) {
		return rdFavorite.CountLike(videoId)
	}

	likes, err := getFavoriterIdList(videoId)
	if err != nil {
		return 0, err
	}

	go func(users []int64, video int64) {
		for _, user := range users {
			rdFavorite.AddLiked(user, video)
		}
	}(likes, videoId)
	return int64(len(likes)), nil
}
