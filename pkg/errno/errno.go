package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode

	UserAlreadyExistErrCode
	UserIsNotExistErrCode

	FollowRelationAlreadyExistErrCode
	FollowRelationNotExistErrCode

	FavoriteRelationAlreadyExistErrCode
	FavoriteRelationNotExistErrCode
	FavoriteActionErrCode

	MessageAddFailedErrCode
	FriendListNoPermissionErrCode

	VideoIsNotExistErrCode
	CommentIsNotExistErrCode
)

const (
	SuccessMsg               = "Success"
	ServerErrMsg             = "Service is unable to start successfully"
	ParamErrMsg              = "Wrong Parameter has been given"
	UserIsNotExistErrMsg     = "user is not exist"
	PasswordIsNotVerifiedMsg = "username or password not verified"
	FavoriteActionErrMsg     = "favorite add failed"

	MessageAddFailedErrMsg    = "message add failed"
	FriendListNoPermissionMsg = "You can't query his friend list"
	VideoIsNotExistErrMsg     = "video is not exist"
	CommentIsNotExistErrMsg   = "comment is not exist"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

var (
	Success                         = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr                      = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr                        = NewErrNo(ParamErrCode, ParamErrMsg)
	UserAlreadyExistErr             = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr          = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserIsNotExistErr               = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerified           = NewErrNo(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)
	FollowRelationAlreadyExistErr   = NewErrNo(FollowRelationAlreadyExistErrCode, "Follow Relation already exist")
	FollowRelationNotExistErr       = NewErrNo(FollowRelationNotExistErrCode, "Follow Relation does not exist")
	FavoriteRelationAlreadyExistErr = NewErrNo(FavoriteRelationAlreadyExistErrCode, "Favorite Relation already exist")
	FavoriteRelationNotExistErr     = NewErrNo(FavoriteRelationNotExistErrCode, "FavoriteRelationNotExistErr")
	FavoriteActionErr               = NewErrNo(FavoriteActionErrCode, FavoriteActionErrMsg)

	MessageAddFailedErr       = NewErrNo(MessageAddFailedErrCode, MessageAddFailedErrMsg)
	FriendListNoPermissionErr = NewErrNo(FriendListNoPermissionErrCode, FriendListNoPermissionMsg)
	VideoIsNotExistErr        = NewErrNo(VideoIsNotExistErrCode, VideoIsNotExistErrMsg)
	CommentIsNotExistErr      = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	// err 是否为 Err 类型或者包含 Err 类型的错误，是则将信息传递到Err
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.WithMessage(err.Error())
	return s
}
