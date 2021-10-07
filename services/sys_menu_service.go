package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
)

type BaseMenuService struct {
}

type IBaseMenuService interface {
	PutBaseMenu(baseMenu *model.SysBaseMenu) (err error)
	DeleteBaseMenu(id int) (err error)
	GetBaseMenu() (baseMenus []*model.SysBaseMenu, err error)
}

// PutBaseMenu  添加系统菜单
func (baseMenuService *BaseMenuService) PutBaseMenu(baseMenu *model.SysBaseMenu) (err error) {
	if err = global.GLOBAL_DB.Create(&baseMenu).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBaseMenu 删除系统菜单
func (baseMenuService *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	if err = global.GLOBAL_DB.Delete(&model.SysBaseMenu{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetBaseMenu 获取全部系统菜单
func (baseMenuService *BaseMenuService) GetBaseMenu() (baseMenus []*model.SysBaseMenu, err error) {
	if err = global.GLOBAL_DB.Find(&baseMenus).Error; err != nil {
		return nil, err
	}
	return baseMenus, nil
}
