package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/model/request"
)

type UserService struct {
}

type IUserService interface {
	GetUsers(pageInfo request.PageInfo) (user []*model.SysUser, total int64, err error)
	CreateUser(user model.SysUser) (err error)
}

// GetUsers 获取系统用户
func (userService *UserService) GetUsers(pageInfo request.PageInfo) (users []*model.SysUser, total int64, err error) {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	userModel := global.GLOBAL_DB.Model(&model.SysUser{})
	err = userModel.Count(&total).Error
	if err != nil {
		return users, total, err
	}
	err = userModel.First(&users).Limit(pageInfo.PageSize).Offset(offset).Error
	if err != nil {
		return users, total, err
	}
	return users, total, nil
}

// CreateUser 新增系统用户
func (userService *UserService) CreateUser(user model.SysUser) (err error) {
	if err = global.GLOBAL_DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
