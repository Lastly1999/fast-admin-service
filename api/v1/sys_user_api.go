package v1

import (
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

var userService services.UserService

// GetSystemUsers
// @Tags User
// @Summary 获取系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [put]
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
