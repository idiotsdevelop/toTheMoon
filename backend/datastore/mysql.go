package datastore

import (
	"gorm.io/gorm"
	"toTheMoon/backend/config"
)

var DB *gorm.DB

func ConnectDB() {
	//var err error
	if config.IsLocal() {
		//dsn := fmt.Sprintf(
		//	"%s:%s@tcp(%s)/%s?parseTime=True&loc=%s",
		//		config.Local,
		//		config.Dev,
		//		config.Prod,
		//	)
	}
}