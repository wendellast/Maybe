package config

import (
	"os"

	"github.com/wendellast/Maybe/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializerSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {

		logger.Infof("Database not found, creating")

		err = os.MkdirAll("db", os.ModePerm)
		if err != nil {
			logger.Errof("Error creating directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errof("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()

		logger.Infof("Database file created at %s", dbPath)
	} else {
		logger.Infof("Database file already exists at %s", dbPath)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errof("Error opening SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errof("Error performing auto migration: %v", err)
		return nil, err
	}

	logger.Infof("SQLite database initialized successfully")

	return db, nil
}
