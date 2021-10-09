package v1

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApi struct {
}

var userService services.UserService

// GetSystemUsers
// @Tags User
// @Summary 获取系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [get]
func (userApi *UserApi) GetSystemUsers(c *gin.Context) {
	appRes := app.Gin{C: c}
	infoParams := request.PageInfo{}
	users, total, err := userService.GetUsers(infoParams)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"users": users,
		"total": total,
	})
}

// CreateSystemUser
// @Tags User
// @Summary 新增系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [put]
func (userApi *UserApi) CreateSystemUser(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysUserParams := request.SystemUserParams{}
	err := c.ShouldBindJSON(&sysUserParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	user := &model.SysUser{
		UserName: sysUserParams.UserName,
		PassWord: sysUserParams.PassWord,
		NikeName: sysUserParams.NikeName,
	}
	err = userService.CreateUser(user)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// UpdateSystemUserById
// @Tags User
// @Summary 修改系统用户信息
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [patch]
func (userApi *UserApi) UpdateSystemUserById(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysUserParams := request.SystemUserParams{}
	err := c.ShouldBindJSON(&sysUserParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	user := &model.SysUser{
		Model: global.Model{
			ID: sysUserParams.Id,
		},
		UserName: sysUserParams.UserName,
		PassWord: sysUserParams.PassWord,
		NikeName: sysUserParams.NikeName,
	}
	err = userService.UpdateUserById(user)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// DeleteSystemUserById
// @Tags User
// @Summary 删除系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [DELETE]
func (userApi *UserApi) DeleteSystemUserById(c *gin.Context) {
	appRes := app.Gin{C: c}
	param := c.Param("id")
	// 参数转换
	id, err := strconv.Atoi(param)
	if err != nil {
		appRes.Response(http.StatusOK, enum.PARAMS_ERROR, nil)
		return
	}
	err = userService.DeleteUserById(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}
