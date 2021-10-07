package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	GLOBAL_DB *gorm.DB
	ZAP_LOG   = zap.S()
)

type Model struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
