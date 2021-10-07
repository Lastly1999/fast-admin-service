package v1

import (
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleApi struct {
}

var roleService services.RoleService

// GetRoles
// @Tags Role
// @Summary 获取角色列表
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /role/role [get]
func (roleApi *RoleApi) GetRoles(c *gin.Context) {
	appRes := app.Gin{C: c}
	roles, err := roleService.GetRoles()
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"roles": roles,
	})
}

// DeleteRoleById
// @Tags Role
// @Summary 删除角色
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /role/role [delete]
func (roleApi *RoleApi) DeleteRoleById(c *gin.Context) {
	appRes := app.Gin{C: c}
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
	}
	err = roleService.DeleteRoleById(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// PutRole
// @Tags Role
// @Summary 新增角色
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /role/role [get]
func (roleApi *RoleApi) PutRole(c *gin.Context) {
	appRes := app.Gin{C: c}
	roleParams := request.SysRoleParams{}
	err := c.ShouldBindJSON(&roleParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	role := &model.SysRole{
		RoleName: roleParams.RoleName,
	}
	err = roleService.PutRole(role)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// UpdateRoleById
// @Tags Role
// @Summary 修改角色信息
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /role/role [patch]
func (roleApi *RoleApi) UpdateRoleById(c *gin.Context) {
	appRes := app.Gin{C: c}
	// 请求的结构体
	roleParams := request.SysRoleParams{}
	err := c.ShouldBindJSON(&roleParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	role := &model.SysRole{
		RoleId:   roleParams.RoleId,
		RoleName: roleParams.RoleName,
		Describe: roleParams.Describe,
	}
	err = roleService.UpdateRole(role)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// UpdateRoleBaseMenu
// @Tags Auth
// @Summary 修改角色权限菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /role/menu [patch]
func (roleApi *RoleApi) UpdateRoleBaseMenu(c *gin.Context) {
	appRes := app.Gin{C: c}
	// 解析token内的用户参数
	info, err := utils.ParseTokenRequest(c)
	// 解析请求结构体
	sysRoleMenuParams := request.SysRoleMenuParams{}
	if err := c.ShouldBindJSON(&sysRoleMenuParams); err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	err = roleService.UpdateRoleMenu(uint(info.RoleId), sysRoleMenuParams.PermissionId)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}
