package db

import (
	"go-tiktok-new/pkg/constants"
	"go-tiktok-new/pkg/errno"
)

type User struct {
	ID              int64  `json:"id"`
	Username        string `json:"user_name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func CreateUser(user *User) (int64, error) {
	if err := DB.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func QueryUser(userName string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserById(userId int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}
	if user == (User{}) {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return &user, nil
}

func VerifyUser(userName, password string) (int64, error) {
	var user User
	if err := DB.Where("user_name = ? AND password = ?", userName, password).Find(&user).Error; err != nil {
		return 0, err
	}
	if user.ID == 0 {
		err := errno.PasswordIsNotVerified
		return user.ID, err
	}
	return user.ID, nil
}

func CheckUserExistById(userId int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
