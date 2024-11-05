package connector

import (
	"errors"
	"fmt"
	"go-todo/config"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	m "go-todo/pkg/db/models"
)

type MysqlConnector struct {
	DB *gorm.DB
	Config *config.Config
}

var SingletonMysqlInstance *MysqlConnector

var OnceMysqlInstance sync.Once

func GetDBConnectorInstance(config *config.Config) *MysqlConnector {
	if SingletonMysqlInstance == nil {
		OnceMysqlInstance.Do(func ()  {
			if db, err := setupMysql(config); err == nil {
				SingletonMysqlInstance = &MysqlConnector{
					DB: db,
					Config: config,
				}
			}
		})
	}
	return SingletonMysqlInstance
}

func setupMysql(config *config.Config) (*gorm.DB, error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	 	config.DB.DbUserName, config.DB.DbPassword, config.DB.DbHost, config.DB.DbPort, config.DB.DbName)
	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		log.Printf("can't connect to db :%v", err)
		return nil, errors.New(fmt.Sprintf("DB err occured %s", err))
	}
	// auto migrate
	db.AutoMigrate(&m.Todo{})
	return db, nil
}
