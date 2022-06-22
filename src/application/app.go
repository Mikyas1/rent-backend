package application

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"os"
	"rent/src/datasource"
	"rent/src/middleware"
	"rent/src/storage"
	"rent/src/storage/minioAPI"
)

var (
	Db      datasource.Database
	Storage storage.Storage
	//ListCache cache.ListCache
	//MapApi    mapApi.MapApi
	//Ws        messaging.MessagingAPI
)

func setUpDb() error {
	var err error

	Db, err = datasource.NewDatabase(DbConf)

	if err != nil {
		return err
	}

	//SetUpTable(Db)
	//
	//err = DropTables(Db)
	//if err != nil {
	//	return err
	//}
	//
	//err = CreateTables(Db)
	//if err != nil {
	//	return err
	//}

	return nil
}

func setUpStorage() error {
	var err error
	Storage, err = minioAPI.NewMinIOStorage()
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = Storage.CreateBucket(ctx, os.Getenv("MINIO_BUCKET_NAME"))
	if err != nil {
		return err
	}

	return nil
}

func setUpFiberApp() error {
	app := fiber.New()
	app.Use(middleware.GetBearerToken)
	RegisterApi(app)
	if err := app.Listen(":" + Port); err != nil {
		return err
	}
	return nil
}

func StartApplication() {
	// setup db
	if err := setUpDb(); err != nil {
		panic(err)
	}
	// setup storage
	if err := setUpStorage(); err != nil {
		panic(err)
	}
	// setup fiber app
	if err := setUpFiberApp(); err != nil {
		panic(err)
	}
}
