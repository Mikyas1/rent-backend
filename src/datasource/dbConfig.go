package datasource

import "fmt"

type DbConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	dbTimeZone string
}

func (c DbConfig) GetPostgresDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=%s",
		c.dbHost,
		c.dbUser,
		c.dbPassword,
		c.dbName,
		c.dbPort,
		c.dbTimeZone,
	)
}

func NewDbConfig(dbUser, dbPassword, dbName, dbHost, dbPort, dbTimeZone string) DbConfig {
	return DbConfig{
		dbUser:     dbUser,
		dbPassword: dbPassword,
		dbName:     dbName,
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbTimeZone: dbTimeZone,
	}
}
