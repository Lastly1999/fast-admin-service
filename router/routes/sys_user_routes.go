package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	userApi := v1.UserApi{}
	{
		// 获取系统用户
		userRouter.POST("user", userApi.GetSystemUsers)
	}
}
