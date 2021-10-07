package gorm

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/pkg/setting"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() {
	//	构造连接池链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DataBaseOptions.User,
		setting.DataBaseOptions.Password,
		setting.DataBaseOptions.Host,
		setting.DataBaseOptions.Db,
	)
	var err error
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	global.GLOBAL_DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		zap.L().Error("database content error..")
		return
	}
	zap.L().Info("database content success...")
}

// InitDataBase 表迁移
func InitDataBase() {
	var err error
	err = global.GLOBAL_DB.AutoMigrate(
		&model.SysUser{},
		&model.SysRole{},
		&model.SysBaseMenu{},
		&model.SysIcon{},
	)
	if err != nil {
		zap.S().Error("automiGrate error...")
		return
	}
}
