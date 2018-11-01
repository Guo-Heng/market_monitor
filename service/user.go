package service

import (
	"github.com/JiangInk/market_monitor/helpers"
	"github.com/JiangInk/market_monitor/models"
)

// UserService 用户服务层逻辑
type UserService struct{}

// StoreUser 添加用户
func (us UserService) StoreUser(email string, pass string) (userID uint, err error) {
	// log.Info().Msg("enter signup service.")

	user := &models.User{
		Email:    email,
		UserName: email,
		Password: pass,
		Status:   true,
	}
	user.Password = helpers.Md5(user.Email + user.Password)
	userID, err = user.Insert()
	return
}

// UpdateUser 更新用户
func (us UserService) UpdateUser(userID uint) {
	return
}

// DestroyUser 删除用户
func (us UserService) DestroyUser(userID uint) error {
	// log.Info().Msg("enter removeUser service.")
	return nil
}
