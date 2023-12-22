package relation

import (
	"context"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/common"
	"go-tiktok-new/biz/model/social/relation"
	user_service "go-tiktok-new/biz/service/user"
	"go-tiktok-new/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	FOLLOW   = 1
	UNFOLLOW = 2
)

type RelationService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewRelationService(ctx context.Context, c *app.RequestContext) *RelationService {
	return &RelationService{
		ctx: ctx,
		c:   c,
	}
}

func (r *RelationService) FollowAction(req *relation.DouyinRelationActionRequest) (flag bool, err error) {
	_, err = db.CheckUserExistById(req.ToUserId)
	if err != nil {
		return false, err
	}
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return false, errno.ParamErr
	}
	currentUserid, _ := r.c.Get("current_user_id")
	if req.ToUserId == currentUserid.(int64) {
		return false, errno.ParamErr
	}
	newFollowRelation := &db.Follows{
		UserId:     req.ToUserId,
		FollowerId: currentUserid.(int64),
	}
	followExist, _ := db.QueryFollowExist(newFollowRelation.UserId, newFollowRelation.FollowerId)
	if req.ActionType == FOLLOW {
		if followExist {
			return false, errno.FollowRelationAlreadyExistErr
		}
		flag, err = db.AddNewFollow(newFollowRelation)
	} else {
		if !followExist {
			return false, errno.FollowRelationNotExistErr
		}
		flag, err = db.DeleteFollow(newFollowRelation)
	}
	return flag, err
}

func (r *RelationService) GetFollowList(req *relation.DouyinRelationFollowListRequest) (followerList []*common.User, err error) {
	_, err = db.CheckUserExistById(req.UserId)
	if err != nil {
		return nil, err
	}

	var followList []*common.User
	currentUserId, exist := r.c.Get("current_user_id")
	if !exist {
		currentUserId = int64(0)
	}
	followIdList, err := db.GetFollowIdList(req.UserId)
	if err != nil {
		return followList, err
	}

	for _, follow := range followIdList {
		userInfo, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(follow, currentUserId.(int64))
		if err != nil {
			continue
		}
		user := common.User{
			Id:              userInfo.Id,
			Name:            userInfo.Name,
			FollowCount:     userInfo.FollowCount,
			FollowerCount:   userInfo.FollowerCount,
			IsFollow:        userInfo.IsFollow,
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			TotalFavorited:  userInfo.TotalFavorited,
			WorkCount:       userInfo.WorkCount,
			FavoriteCount:   userInfo.FavoriteCount,
		}
		followList = append(followList, &user)
	}
	return followList, nil
}

func (r *RelationService) GetFollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*common.User, error) {
	userId := req.UserId
	var followerList []*common.User
	currentUserId, exist := r.c.Get("current_user_id")
	if !exist {
		currentUserId = int64(0)
	}

	followerIdList, err := db.GetFollowerIdList(userId)
	if err != nil {
		return followerList, err
	}

	for _, follower := range followerIdList {
		userInfo, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(follower, currentUserId.(int64))
		if err != nil {
			hlog.Error("func error: GetFollowerList -> GetUserInfo")
		}
		user := &common.User{
			Id:              userInfo.Id,
			Name:            userInfo.Name,
			FollowCount:     userInfo.FollowCount,
			FollowerCount:   userInfo.FollowerCount,
			IsFollow:        userInfo.IsFollow,
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			TotalFavorited:  userInfo.TotalFavorited,
			WorkCount:       userInfo.WorkCount,
			FavoriteCount:   userInfo.FavoriteCount,
		}
		followerList = append(followerList, user)
	}
	return followerList, nil
}

func (r *RelationService) GetFriendList(req *relation.DouyinRelationFriendListRequest) ([]*relation.FriendUser, error) {
	userId := req.UserId
	currentUserId, _ := r.c.Get("current_user_id")

	if currentUserId.(int64) != userId {
		return nil, errno.FriendListNoPermissionErr
	}

	var friendList []*relation.FriendUser

	friendIdList, err := db.GetFriendIdList(userId)
	if err != nil {
		return friendList, err
	}

	for _, id := range friendIdList {
		userInfo, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(id, userId)
		if err != nil {
			hlog.CtxInfof(context.Background(), "func error: GetFriendList -> GetUserinfo")
		}
		message, err := db.GetLatestMessageByIdPair(userId, id)
		if err != nil {
			hlog.CtxInfof(context.Background(), "func error: GetFriendList -> GetLatestMessageByIdPair")
		}
		var msgType int64
		if message == nil {
			msgType = 2
			message = &db.Messages{}
		} else if userId == message.FromUserId {
			msgType = 1
		} else {
			msgType = 0
		}

		friendList = append(friendList, &relation.FriendUser{
			User: &common.User{
				Id:              userInfo.Id,
				Name:            userInfo.Name,
				FollowCount:     userInfo.FollowCount,
				FollowerCount:   userInfo.FollowerCount,
				IsFollow:        userInfo.IsFollow,
				Avatar:          userInfo.Avatar,
				BackgroundImage: userInfo.BackgroundImage,
				Signature:       userInfo.Signature,
				TotalFavorited:  userInfo.TotalFavorited,
				WorkCount:       userInfo.WorkCount,
				FavoriteCount:   userInfo.FavoriteCount,
			},
			Message: message.Content,
			MsgType: msgType,
		})
	}

	return friendList, nil
}
