package global

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	GLOBAL_DB       *gorm.DB
	ZAP_LOG         *zap.Logger
	GLOBAL_Enforcer *casbin.Enforcer
)

type Model struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
