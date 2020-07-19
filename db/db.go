package db

import (
	"docApp/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConn() {

	cfg := config.GetConfig()
	connArgs := fmt.Sprintf(cfg.DbUser+"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbSchema)

	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection Successed")
	}

	defer db.Close()
}
