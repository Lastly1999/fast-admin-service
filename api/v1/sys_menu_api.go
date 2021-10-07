package v1

import (
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BaseMenuApi struct {
}

var baseMenuService services.BaseMenuService

// PutBaseMenu
// @Tags Auth
// @Summary 添加系统菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /menu/menu [put]
func (baseMenuApi *BaseMenuApi) PutBaseMenu(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysBaseMenuParams := request.SysBaseMenuParams{}
	err := c.ShouldBindJSON(&sysBaseMenuParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	baseMenu := &model.SysBaseMenu{
		Name:     sysBaseMenuParams.MenuName,
		Icon:     sysBaseMenuParams.MenuIcon,
		Path:     sysBaseMenuParams.MenuPath,
		ParentId: uint(sysBaseMenuParams.MenuParentId),
	}
	err = baseMenuService.PutBaseMenu(baseMenu)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// DeleteBaseMenu
// @Tags Auth
// @Summary 删除系统菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /menu/menu [put]
func (baseMenuApi *BaseMenuApi) DeleteBaseMenu(c *gin.Context) {
	appRes := app.Gin{C: c}
	params := c.Param("id")
	// 参数转换
	id, err := strconv.Atoi(params)
	if err != nil {
		appRes.Response(http.StatusOK, enum.PARAMS_ERROR, nil)
		return
	}
	err = baseMenuService.DeleteBaseMenu(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// GetBaseMenu
// @Tags Auth
// @Summary 获取全部系统菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /menu/menu [put]
func (baseMenuApi *BaseMenuApi) GetBaseMenu(c *gin.Context) {
	appRes := app.Gin{C: c}
	menus, err := baseMenuService.GetBaseMenu()
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"menus": menus,
	})
}
