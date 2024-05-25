package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(s string) *Logger {
	logger = NewLogger(s)
	return logger
}
