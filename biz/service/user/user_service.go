package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-tiktok-new/biz/dal/db"
	"go-tiktok-new/biz/model/basic/user"
	"go-tiktok-new/biz/model/common"
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/errno"
	"go-tiktok-new/pkg/utils"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (s *UserService) UserRegister(req *user.DouyinUserRegisterRequest) (userId int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	passWord, err := utils.Crypt(req.Password)
	userId, err = db.CreateUser(&db.User{
		Username:        req.Username,
		Password:        passWord,
		Avatar:          constants.TestAva,
		BackgroundImage: constants.TestBackground,
		Signature:       constants.TestSign,
	})
	return userId, nil
}

func (s *UserService) UserInfo(req *user.DouyinUserRequest) (*common.User, error) {
	queryUserId := req.UserId
	currentUserId, exist := s.c.Get("current_user_id")
	if !exist {
		currentUserId = 0
	}
	return s.GetUserInfo(queryUserId, currentUserId.(int64))
}

func (s *UserService) GetUserInfo(queryUserId, userId int64) (*common.User, error) {
	u := &common.User{}

	dbUser, err := db.QueryUserById(queryUserId)
	if err != nil {
		return u, err
	}
	WorkCount, err := db.GetWorkCount(queryUserId)
	if err != nil {
		return u, err
	}
	FollowCount, err := db.GetFollowCount(queryUserId)
	if err != nil {
		return u, err
	}
	FollowerCount, err := db.GetFollowerCount(queryUserId)
	if err != nil {
		return u, err
	}

	var IsFollow bool
	if userId != 0 {
		IsFollow, err = db.QueryFollowExist(userId, queryUserId)
		if err != nil {
			return u, err
		} else {
			IsFollow = false
		}
	}
	FavoriteCount, err := db.GetFavoriteCountByUserId(queryUserId)
	if err != nil {
		return u, err
	}
	TotalFavorited, err := db.QueryTotalFavoritedByAuthorId(queryUserId)
	if err != nil {
		return u, err
	}

	u = &common.User{
		Id:              queryUserId,
		Name:            dbUser.Username,
		FollowCount:     FollowCount,
		FollowerCount:   FollowerCount,
		IsFollow:        IsFollow,
		Avatar:          utils.URLConvert(s.ctx, s.c, dbUser.Avatar),
		BackgroundImage: utils.URLConvert(s.ctx, s.c, dbUser.BackgroundImage),
		Signature:       dbUser.Signature,
		TotalFavorited:  TotalFavorited,
		WorkCount:       WorkCount,
		FavoriteCount:   FavoriteCount,
	}

	return u, nil
}
