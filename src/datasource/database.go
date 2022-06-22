package datasource

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	postgres2 "rent/src/datasource/postgres"
)

type Database interface {
	AutoMigrate(...interface{}) error
	ProvideDb() *gorm.DB
	DropTables(...interface{}) error
}

func NewDatabase(dbConf DbConfig) (Database, error) {
	db, err := gorm.Open(postgres.Open(dbConf.GetPostgresDsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	fmt.Println(dbConf.GetPostgresDsn())

	if err != nil {
		return nil, err
	}
	return postgres2.PgDatabase{
		Db: db,
	}, nil
}
