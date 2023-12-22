package db

import (
	"go-tiktok-new/biz/mw/redis"
	"go-tiktok-new/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Follows struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreateAt   time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `grom:"index" json:"deleted_at"`
}

var rdFollows redis.Follows

func (Follows) TableName() string {
	return constants.FollowsTableName
}

func AddNewFollow(follow *Follows) (bool, error) {
	if err := DB.Create(follow).Error; err != nil {
		return false, err
	}

	// data to redis
	if rdFollows.CheckFollow(follow.FollowerId) {
		rdFollows.AddFollow(follow.UserId, follow.FollowerId)
	}
	if rdFollows.CheckFollower(follow.UserId) {
		rdFollows.AddFollower(follow.UserId, follow.FollowerId)
	}

	return true, nil
}

func DeleteFollow(follow *Follows) (bool, error) {
	if err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).
		Delete(follow).Error; err != nil {
		return false, err
	}

	// if redis hit del
	if rdFollows.CheckFollow(follow.FollowerId) {
		rdFollows.DelFollow(follow.UserId, follow.FollowerId)
	}
	if rdFollows.CheckFollow(follow.UserId) {
		rdFollows.DelFollower(follow.UserId, follow.FollowerId)
	}
	return true, nil
}

func QueryFollowExist(userId, followerId int64) (bool, error) {
	if rdFollows.CheckFollow(followerId) {
		return rdFollows.ExistFollow(userId, followerId), nil
	}
	if rdFollows.CheckFollow(userId) {
		return rdFollows.ExistFollower(userId, followerId), nil
	}
	follow := Follows{
		UserId:     userId,
		FollowerId: followerId,
	}
	if err := DB.Where("user_id = ? AND follower_id = ?", userId, followerId).
		First(&follow).Error; err != nil {
		return false, err
	}
	if follow.ID == 0 {
		return false, nil
	}
	return true, nil
}

func GetFollowCount(followerId int64) (int64, error) {
	if rdFollows.CheckFollow(followerId) {
		return rdFollows.CountFollow(followerId)
	}

	followings, err := getFollowIdList(followerId)
	if err != nil {
		return 0, err
	}

	// async update redis
	go addFollowRelationToRedis(followerId, followings)
	return int64(len(followings)), nil
}

func GetFollowerCount(userId int64) (int64, error) {
	if rdFollows.CheckFollower(userId) {
		return rdFollows.CountFollower(userId)
	}

	followers, err := getFollowerIdList(userId)
	if err != nil {
		return 0, err
	}
	go addFollowerRelationToRedis(userId, followers)
	return int64(len(followers)), nil
}

func GetFollowIdList(followerId int64) ([]int64, error) {
	if rdFollows.CheckFollow(followerId) {
		return rdFollows.GetFollow(followerId), nil
	}
	return getFollowIdList(followerId)
}

func GetFollowerIdList(userId int64) ([]int64, error) {
	if rdFollows.CheckFollower(userId) {
		return rdFollows.GetFollower(userId), nil
	}
	return getFollowerIdList(userId)
}

func getFollowIdList(followerId int64) ([]int64, error) {
	var followActions []Follows
	if err := DB.Where("follower_id = ?", followerId).Find(&followActions).Error; err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range followActions {
		result = append(result, v.UserId)
	}
	return result, nil
}

func getFollowerIdList(userId int64) ([]int64, error) {
	var followActions []Follows
	if err := DB.Where("user_id = ?", userId).Find(&followActions).Error; err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range followActions {
		result = append(result, v.FollowerId)
	}
	return result, nil
}

func addFollowRelationToRedis(followerId int64, followings []int64) {
	for _, following := range followings {
		rdFollows.AddFollow(following, followerId)
	}
}

func addFollowerRelationToRedis(userId int64, followers []int64) {
	for _, follower := range followers {
		rdFollows.AddFollower(userId, follower)
	}
}
