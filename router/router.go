package router

import (
	v1 "fast-admin-service/api/v1"
	"fast-admin-service/middleware/jwt"
	"fast-admin-service/router/routes"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitRouter() (app *gin.Engine) {
	gin.ForceConsoleColor()
	app = gin.Default()
	api := app.Group("v1")
	authApi := v1.AuthApi{}
	// 记录到文件。
	f, _ := os.Create("./log/gin-example.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 授权登录
	api.POST("/auth/login", authApi.LoginAction)
	// 获取图形验证码
	api.GET("/auth/code", authApi.GetAuthCode)
	api.Use(jwt.JWT()) // jwt 中间件
	{
		// 授权模块
		routes.InitAuthRouter(api)
		// 用户模块
		routes.InitUserRouter(api)
		// 角色模块
		routes.InitRoleRouter(api)
		// 系统模块
		routes.InitSystemRouter(api)
		// 系统菜单模块
		routes.InitBaseMenuRouter(api)
	}
	return app
}
