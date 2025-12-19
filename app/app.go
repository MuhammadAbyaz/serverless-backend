package app

import (
	"serverless/api"
	"serverless/interfaces"

	"gorm.io/gorm"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp(driver gorm.Dialector, dsn string) *App {
	db, err := interfaces.NewDbInstance(dsn, driver)
	if err != nil {
		panic("Error Getting Database Instance!!")
	}
	return &App{
		ApiHandler: *api.NewApiHandler(&db),
	}
}
