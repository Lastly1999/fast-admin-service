package v1

import (
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/captcha"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"fast-admin-service/services"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthApi struct {
}

var authService services.AuthService

// LoginAction
// @Tags Auth
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /auth/login [post]
func (authApi *AuthApi) LoginAction(c *gin.Context) {
	appRes := app.Gin{C: c}
	loginParam := request.Login{}
	valid := validation.Validation{}
	err := c.ShouldBind(&loginParam)
	// 绑定json错误
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, "绑定json失败，结构体参数绑定失败")
		return
	}
	// 字段验证
	ok, err := valid.Valid(&loginParam)
	// 字段验证错误
	if !ok {
		appRes.Response(http.StatusOK, enum.ERROR, "参数解析异常")
		return
	}
	userReqBody := &model.SysUser{
		UserName: loginParam.UserName,
		PassWord: loginParam.PassWord,
	}
	// 验证图形验证码
	res := captcha.Verify(loginParam.CodeAuth, loginParam.Code)
	if res {
		auth, err := authService.CheckAuth(userReqBody)
		if err != nil {
			appRes.Response(http.StatusOK, enum.AUTH_ERROR, nil)
			return
		}
		// 成功 派发token 用户的权限 默认选择第一个作为默认角色
		token, err := utils.GenerateToken(auth.UserName, auth.PassWord, int(auth.ID), int(auth.Role[0].RoleId))
		if err != nil {
			appRes.Response(http.StatusOK, enum.ERROR_AUTH, "token派发错误")
			return
		}
		// 返回接口
		appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
			"token": token,
		})
	} else {
		appRes.Response(http.StatusOK, enum.CODE_ERROR, nil)
	}
}

// GetAuthCode
// @Tags Auth
// @Summary 获取图片验证码
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/code [post]
func (authApi *AuthApi) GetAuthCode(c *gin.Context) {
	appRes := app.Gin{C: c}
	code, base, err := authService.GenerateVerificode()
	if err != nil {
		appRes.Response(http.StatusBadGateway, enum.ERROR, err.Error())
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"code":     code,
		"codeBase": base,
	})
}

// GetBaseMenus
// @Tags Auth
// @Summary 获取基础权限菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/menu [get]
func (authApi *AuthApi) GetBaseMenus(c *gin.Context) {
	appRes := app.Gin{C: c}
	info, err := utils.ParseTokenRequest(c)
	if err != nil {
		appRes.Response(http.StatusOK, enum.INVALID_TOKEN_PARAMS_ERROR, nil)
		return
	}
	menus, err := authService.GetSystemPermissionsMenu(info.RoleId)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"menus": menus,
	})
}

// GetBaseMenusIds
// @Tags Auth
// @Summary 获取用户权限菜单id组
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/menuids/:id [get]
func (authApi *AuthApi) GetBaseMenusIds(c *gin.Context) {
	appRes := app.Gin{C: c}
	// 解析token内的用户参数
	id := c.Param("id")
	// 参数转换
	uid, err := strconv.Atoi(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.PARAMS_ERROR, nil)
		return
	}
	ids, err := authService.GetSystemPermissionsMenuIds(uid)
	if err != nil {
		appRes.Response(http.StatusOK, enum.INVALID_TOKEN_PARAMS_ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"roleIds": ids,
	})
}
