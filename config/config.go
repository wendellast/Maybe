package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializerSqlite()

	if err != nil {
		return fmt.Errorf("erro initializer sqlite: %v ", err)
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db

}

func GetLogger(s string) *Logger {
	logger = NewLogger(s)
	return logger
}
