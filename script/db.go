package script

import (
	"docApp/config"
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


	//db.DropTableIfExists(
	//	&models.User{},
	//	&models.Admin{},
	//	&models.Services{},
	//	&models.Reservation{},
	//	&models.Variety{},
	//	&models.Doctor{},
	//	&models.Schedule{},
	//	&models.Product{},
	//	&models.Category{},
	//	&models.OrderDetail{},
	//	&models.Order{},
	//)

	//create
	//db.AutoMigrate(
	//	&models.User{},
	//	&models.Admin{},
	//	&models.Services{},
	//	&models.Reservation{},
	//	&models.Variety{},
	//	&models.Doctor{},
	//	&models.Schedule{},
	//	&models.Product{},
	//	&models.Category{},
	//	&models.OrderDetail{},
	//	&models.Order{},
	//)

	//constraint
	//CreateConstraint(db)



	//defer db.Close()
	return db, nil
}
