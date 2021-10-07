package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
)

type UserService struct {
}

type IUserService interface {
	GetUsers() (user model.SysUser, err error)
	CreateUser(user model.SysUser) (err error)
}

func (userService *UserService) GetUsers() (user model.SysUser, err error) {
	if err := global.GLOBAL_DB.First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (userService *UserService) CreateUser(user model.SysUser) (err error) {
	if err = global.GLOBAL_DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
