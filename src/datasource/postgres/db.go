package postgres

import (
	"gorm.io/gorm"
)

type PgDatabase struct {
	Db *gorm.DB
}

func (db PgDatabase) AutoMigrate(models ...interface{}) error {
	err := db.Db.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func (db PgDatabase) ProvideDb() *gorm.DB {
	return db.Db
}

func (db PgDatabase) DropTables(models ...interface{}) error {
	err := db.Db.Migrator().DropTable(models...)
	if err != nil {
		return err
	}
	return nil
}
