package infra

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQL(config *AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQL.Usr, config.MySQL.Pw, config.MySQL.Host, config.MySQL.Port, config.MySQL.Db)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panic(err)
	}

	return db
}
