package db

import (
	"docApp/config"
	"docApp/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConn() (*gorm.DB, error){

	cfg := config.GetConfig()
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbSchema)

	db, err := gorm.Open(cfg.DbName, connArgs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Connection Successed")


	//migration
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})


	//defer db.Close()
	return db, nil
}
