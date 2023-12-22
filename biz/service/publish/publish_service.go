package publish

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/basic/feed"
	"go-tiktok-new/biz/model/basic/publish"
	"go-tiktok-new/biz/model/common"
	"go-tiktok-new/biz/mw/ffmpeg"
	"go-tiktok-new/biz/mw/minio"
	feed_service "go-tiktok-new/biz/service/feed"
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/utils"
	"strconv"
	"time"
)

type PublishService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewPublishService(ctx context.Context, c *app.RequestContext) *PublishService {
	return &PublishService{
		ctx: ctx,
		c:   c,
	}
}

func (s *PublishService) PublishAction(req *publish.DouyinPublishActionRequest, fileName string) (err error) {
	// 初始化信息
	v, _ := s.c.Get("current_user_id")
	title := s.c.PostForm("title")
	userId := v.(int64)
	nowTime := time.Now()
	fileName = utils.NewFileName(userId, nowTime.Unix()) + fileName

	// 上传视频
	uploadInfo, err := minio.PutToBucketByBuf(s.ctx, constants.MinioVideoBucketName, fileName, bytes.NewBuffer(req.Data))
	hlog.CtxInfof(s.ctx, "video upload size:"+strconv.FormatInt(uploadInfo.Size, 10))
	PlayURL := constants.MinioVideoBucketName + "/" + fileName

	// 上传封面
	buf, err := ffmpeg.GetFirstFrame(utils.URLConvert(s.ctx, s.c, PlayURL))
	uploadInfo, err = minio.PutToBucketByBuf(s.ctx, constants.MinioImgBucketName, fileName+".png", buf)
	hlog.CtxInfof(s.ctx, "image upload size:"+strconv.FormatInt(uploadInfo.Size, 10))
	if err != nil {
		hlog.CtxInfof(s.ctx, "err:"+err.Error())
	}

	// 插入数据库
	_, err = db.CreateVideo(&db.Video{
		AuthorId:    userId,
		PlayURL:     PlayURL,
		CoverURL:    constants.MinioImgBucketName + "/" + fileName + ".png",
		PublishTime: nowTime,
		Title:       title,
	})
	return err
}

func (s *PublishService) PublishList(req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp = &publish.DouyinPublishListResponse{}
	queryUserId := req.UserId
	currentUserId, exist := s.c.Get("current_user_id")
	if !exist {
		currentUserId = int64(0)
	}
	dbVideos, err := db.GetVideoByUserId(queryUserId)
	if err != nil {
		return resp, err
	}
	var videos []*feed.Video

	f := feed_service.NewFeedService(s.ctx, s.c)
	err = f.CopyVideos(&videos, &dbVideos, currentUserId.(int64))
	if err != nil {
		return resp, err
	}

	for _, item := range videos {
		video := common.Video{
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
		resp.VideoList = append(resp.VideoList, &video)
	}
	return resp, nil
}
