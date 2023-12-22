package feed

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/basic/feed"
	"go-tiktok-new/biz/model/common"
	user_service "go-tiktok-new/biz/service/user"
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/utils"
	"sync"
	"time"
)

type FeedService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFeedService(ctx context.Context, c *app.RequestContext) *FeedService {
	return &FeedService{
		ctx: ctx,
		c:   c,
	}
}

func (s *FeedService) Feed(req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	resp := &feed.DouyinFeedResponse{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime/1000, 0)
	}

	currentUserId, exists := s.c.Get("current_user_id")
	if !exists {
		currentUserId = int64(0)
	}

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return resp, err
	}

	videos := make([]*feed.Video, 0, constants.VideoFeedCount)
	err = s.CopyVideos(&videos, &dbVideos, currentUserId.(int64))
	if err != nil {
		return resp, nil
	}
	resp.VideoList = videos
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

func (s *FeedService) CopyVideos(result *[]*feed.Video, data *[]*db.Video, userId int64) error {
	for _, item := range *data {
		video := s.createVideo(item, userId)
		*result = append(*result, video)
	}
	return nil
}

func (s *FeedService) createVideo(data *db.Video, userId int64) *feed.Video {
	video := &feed.Video{
		Id:       data.ID,
		PlayUrl:  utils.URLConvert(s.ctx, s.c, data.PlayURL),
		CoverUrl: utils.URLConvert(s.ctx, s.c, data.CoverURL),
		Title:    data.Title,
	}

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		author, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(data.AuthorId, userId)
		if err != nil {
			hlog.CtxInfof(s.ctx, "GetUserInfo func error:"+err.Error())
		}
		video.Author = &common.User{
			Id:              author.Id,
			Name:            author.Name,
			FollowCount:     author.FollowCount,
			FollowerCount:   author.FollowerCount,
			IsFollow:        author.IsFollow,
			Avatar:          author.Avatar,
			BackgroundImage: author.BackgroundImage,
			Signature:       author.Signature,
			TotalFavorited:  author.TotalFavorited,
			WorkCount:       author.WorkCount,
			FavoriteCount:   author.FavoriteCount,
		}
		wg.Done()
	}()

	go func() {
		var err error
		video.FavoriteCount, err = db.GetFavoriteCount(data.ID)
		if err != nil {
			hlog.CtxInfof(s.ctx, "GetFavoriteCount func error: "+err.Error())
		}
		wg.Done()
	}()

	go func() {
		var err error
		video.CommentCount, err = db.GetCommentCountByVideoId(data.ID)
		if err != nil {
			hlog.CtxInfof(s.ctx, "GetCommentCountByVideoId func error:"+err.Error())
		}
		wg.Done()
	}()

	go func() {
		var err error
		video.IsFavorite, err = db.QueryFavoriteExist(userId, data.ID)
		if err != nil {
			hlog.CtxInfof(s.ctx, "QueryFavoriteExist func error:"+err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
	return video
}
