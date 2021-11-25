package enforce

import (
	"fast-admin-service/global"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ApiCheckRule Casbin的api鉴权认证中间件
func ApiCheckRule(e *casbin.Enforcer) gin.HandlerFunc {
	return func(context *gin.Context) {
		appRes := app.Gin{C: context}
		// 获取api路径
		obj := context.Request.URL.RequestURI()
		// 获取api请求方法
		act := context.Request.Method
		token := context.GetHeader("Authorization")
		parseToken, err := utils.ParseToken(token)
		if err != nil {
			appRes.Response(http.StatusUnauthorized, enum.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			return
		}
		// 获取角色id 为了跟数据库的policy规则对比 sub
		sub := strconv.Itoa(int(parseToken.RoleId))
		// 判断策略中是否存在放行
		enforce, _ := e.Enforce(sub, obj, act)
		if enforce {
			global.ZAP_LOG.Info("权限认证通过成功")
			context.Next()
		} else {
			global.ZAP_LOG.Info("权限认证失败，您没有此api权限")
			appRes.Response(http.StatusUnauthorized, enum.AUTH_FAIL, nil)
			context.Abort()
		}
	}
}
