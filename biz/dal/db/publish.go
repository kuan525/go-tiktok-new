package db

import (
	"go-tiktok-new/pkg/constants"
	"time"
)

type Video struct {
	ID          int64
	AuthorId    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func (Video) TableName() string {
	return constants.VideosTableName
}

func CreateVideo(video *Video) (VideoId int64, err error) {
	err = DB.Create(video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, err
}

func GetVideosByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, constants.VideoFeedCount)
	err := DB.Where("publish_time < ?", lastTime).
		Order("publish_time desc").
		Limit(constants.VideoFeedCount).
		Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func GetVideoByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	if err := DB.Where("author_id = ?", userId).Find(&videos).Error; err != nil {
		return videos, err
	}
	return videos, nil
}

func GetVideoListByVideoIdList(videoIdList []int64) ([]*Video, error) {
	var videoList []*Video
	var err error
	for _, item := range videoIdList {
		var video *Video
		err = DB.Where("id = ?", item).Find(&video).Error
		if err != nil {
			return videoList, err
		}
		videoList = append(videoList, video)
	}
	return videoList, err
}

func GetWorkCount(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Video{}).Where("author_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CheckVideoExistById(videoId int64) (bool, error) {
	var video Video
	if err := DB.Where("id = ?", videoId).Find(&video).Error; err != nil {
		return false, err
	}
	if video == (Video{}) {
		return false, nil
	}
	return true, nil
}
