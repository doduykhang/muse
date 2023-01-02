package repositories

import (
	"github.com/doduykhang/muse/pkg/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = config.GetDB()
}
