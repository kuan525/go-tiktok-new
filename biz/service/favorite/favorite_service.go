package favorite

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/basic/feed"
	"go-tiktok-new/biz/model/common"
	"go-tiktok-new/biz/model/interact/favorite"
	feed_service "go-tiktok-new/biz/service/feed"
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFavoriteService(ctx context.Context, c *app.RequestContext) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
		c:   c,
	}
}

func (r *FavoriteService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) (flag bool, err error) {
	_, err = db.CheckVideoExistById(req.VideoId)
	if err != nil {
		return false, err
	}
	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return false, errno.ParamErr
	}
	currentUserId, _ := r.c.Get("current_user_id")

	newFavoriteRelation := &db.Favorites{
		UserId:  currentUserId.(int64),
		VideoId: req.VideoId,
	}
	favoriteExist, _ := db.QueryFavoriteExist(newFavoriteRelation.UserId, newFavoriteRelation.VideoId)
	if req.ActionType == constants.FavoriteActionType {
		if favoriteExist {
			return false, errno.FavoriteRelationAlreadyExistErr
		}
		flag, err = db.AddNewFavorite(newFavoriteRelation)
	} else {
		if !favoriteExist {
			return false, errno.FavoriteRelationNotExistErr
		}
		flag, err = db.DeleteFavorite(newFavoriteRelation)
	}
	return flag, err
}

func (r *FavoriteService) GetGavoriteList(req *favorite.DouyinFavoriteListRequest) (favoritelist []*common.Video, err error) {
	queryUserId := req.UserId
	_, err = db.CheckUserExistById(queryUserId)

	if err != nil {
		return nil, err
	}
	currentUserId, _ := r.c.Get("current_user_id")

	videoIdList, err := db.GetFavoriteIdList(queryUserId)

	dbVideos, err := db.GetVideoListByVideoIdList(videoIdList)
	var videos []*feed.Video
	f := feed_service.NewFeedService(r.ctx, r.c)
	err = f.CopyVideos(&videos, &dbVideos, currentUserId.(int64))
	for _, item := range videos {
		video := &common.Video{
			Id: item.Id,
			Author: &common.User{
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		favoritelist = append(favoritelist, video)
	}
	return favoritelist, err
}
