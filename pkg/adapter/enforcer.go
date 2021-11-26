package adapter

import (
	"fast-admin-service/global"
	"fast-admin-service/utils"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// EnforcerTools Casbin执法者认证
func EnforcerTools() *casbin.Enforcer {
	// 创建适配器 使用 MySQL 数据库初始化一个 Gorm 适配器
	adapter, err := gormadapter.NewAdapter("mysql", "root:1234@tcp(127.0.0.1:3306)/fav", true)
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		return nil
	}
	global.GLOBAL_Enforcer, err = casbin.NewEnforcer("config/casbin.conf", adapter)
	global.GLOBAL_Enforcer.AddFunction("ParamsMatch", utils.RegexMatchFunc)
	global.GLOBAL_Enforcer.AddFunction("regexMatch", utils.RegexMatchFunc)
	//global.GLOBAL_Enforcer = enforcer
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		return nil
	}
	// 从数据库加载casbin策略
	err = global.GLOBAL_Enforcer.LoadPolicy()
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		return nil
	}
	return global.GLOBAL_Enforcer
}
