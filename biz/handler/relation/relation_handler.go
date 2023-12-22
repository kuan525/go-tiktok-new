package relation

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-tiktok-new/biz/model/social/relation"
	service "go-tiktok-new/biz/service/relation"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

// RelationAction users follow or unfollow other users
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_, err = service.NewRelationService(ctx, c).FollowAction(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// RelationFollowList
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
			StatusMsg:  resp.StatusMsg,
			StatusCode: resp.StatusCode,
			UserList:   nil,
		})
		return
	}

	FollowInfo, err := service.NewRelationService(ctx, c).GetFollowList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
		StatusMsg:  errno.SuccessMsg,
		StatusCode: errno.SuccessCode,
		UserList:   FollowInfo,
	})
}

// RelationFollowerList
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	followerList, err := service.NewRelationService(ctx, c).GetFollowerList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   followerList,
		})
	}
}

// RelationFriendList
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusMsg:  resp.StatusMsg,
			StatusCode: resp.StatusCode,
			UserList:   nil,
		})
		return
	}

	friendList, err := service.NewRelationService(ctx, c).GetFriendList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusMsg:  resp.StatusMsg,
			StatusCode: resp.StatusCode,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusMsg:  errno.SuccessMsg,
			StatusCode: errno.SuccessCode,
			UserList:   friendList,
		})
	}
}
