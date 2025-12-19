package interfaces

import (
	"gorm.io/gorm"
)

type IDbInstance struct {
	DataStore *gorm.DB
}

func NewDbInstance(connectionString string, driver gorm.Dialector) (IDbInstance, error) {
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return IDbInstance{}, err
	}
	return IDbInstance{
		DataStore: db,
	}, nil
}
