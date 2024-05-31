package handler

import (
	"github.com/wendellast/Maybe/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}
