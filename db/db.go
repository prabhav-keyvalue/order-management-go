package db

import (
	"fmt"

	"github.com/prabhav-keyvalue/order-management-go/config"
	loggerr "github.com/prabhav-keyvalue/order-management-go/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDb() error {
	d, err := newDb()

	if err != nil {
		return err
	}

	db = d
	return nil
}

func newDb() (d *gorm.DB, err error) {
	databaseConfig := config.GetDbConfig()

	dbUrl := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode = disable password= %s",
		databaseConfig.DbHost,
		databaseConfig.DbPort,
		databaseConfig.DbUser,
		databaseConfig.DbName,
		databaseConfig.DbPassword,
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseConfig.DbSchema + ".",
			SingularTable: true,
		},
	}

	if config.GetEnv() != string(config.Production) {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	d, err = gorm.Open(postgres.Open(dbUrl), gormConfig)

	if err != nil {
		return
	}

	// d.NamingStrategy.SchemaName("test")

	loggerr.Info("DB Connection successful")
	return
}

func GetDB() *gorm.DB {
	return db
}
