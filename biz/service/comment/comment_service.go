package comment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/common"
	"go-tiktok-new/biz/model/interact/comment"
	user_service "go-tiktok-new/biz/service/user"
	"go-tiktok-new/pkg/errno"
)

type CommentService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewCommentService(ctx context.Context, c *app.RequestContext) *CommentService {
	return &CommentService{
		ctx: ctx,
		c:   c,
	}
}

func (c *CommentService) AddNewComment(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	currentUserId, _ := c.c.Get("current_user_id")
	videoId := req.VideoId
	actionType := req.ActionType
	CommentText := req.CommentText
	CommentId := req.CommentId
	comment := &comment.Comment{}

	if actionType == 1 {
		dbComment := &db.Comment{
			UserId:      currentUserId.(int64),
			VideoId:     videoId,
			CommentText: CommentText,
		}
		err := db.AddNewComment(dbComment)
		if err != nil {
			return comment, err
		}
		comment.Id = dbComment.ID
		comment.CreateDate = dbComment.CreateAt.Format("01-02")
		comment.Content = dbComment.CommentText
		comment.User, err = c.GetUserInfoById(currentUserId.(int64), currentUserId.(int64))
		if err != nil {
			return comment, nil
		}
		return comment, nil
	} else {
		err := db.DeleteCommentById(CommentId)
		if err != nil {
			return comment, err
		}
		return comment, nil
	}
}

func (c *CommentService) GetUserInfoById(currentUserId, userId int64) (*common.User, error) {
	u, err := user_service.NewUserService(c.ctx, c.c).GetUserInfo(userId, currentUserId)
	var commentUser *common.User
	if err != nil {
		return commentUser, err
	}
	commentUser = &common.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
	return commentUser, nil
}

func (c *CommentService) CommentList(req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	resp := &comment.DouyinCommentListResponse{}
	videoId := req.VideoId

	currentUserId, _ := c.c.Get("current_user_id")

	dbComments, err := db.GetCommentListByVideoId(videoId)
	if err != nil {
		return resp, err
	}
	var comments []*comment.Comment
	err = c.copyComment(&comments, &dbComments, currentUserId.(int64))
	if err != nil {
		return resp, err
	}
	resp.CommentList = comments
	resp.StatusMsg = errno.SuccessMsg
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

func (c *CommentService) copyComment(result *[]*comment.Comment, data *[]*db.Comment, currentUserId int64) error {
	for _, item := range *data {
		comment := c.createComment(item, currentUserId)
		*result = append(*result, comment)
	}
	return nil
}

func (c *CommentService) createComment(data *db.Comment, userId int64) *comment.Comment {
	comment := &comment.Comment{
		Id:         data.ID,
		Content:    data.CommentText,
		CreateDate: data.CreateAt.Format("01-02"),
	}

	userInfo, err := c.GetUserInfoById(userId, data.UserId)
	if err != nil {
		hlog.CtxInfof(context.Background(), "func error")
	}
	comment.User = userInfo
	return comment
}
