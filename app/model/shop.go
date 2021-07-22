package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Shop struct {
	// gorm.Model
	ID        uint   `gorm:"primary_key"`
	ShopName  string `gorm:"unique;not null"`
	Latitude  string
	Longitude string
}

func dbConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "test"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "TEAMC"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func FindShop(ids []string) []*Shop {
	db := dbConnect()
	defer db.Close()

	result := []*Shop{}
	err := db.Find(&result).Error
	// db.Where("name = ?", ids).Find(&result)
	// rows, err := db.Table("shops").Where("shop_name = ?", ids).Select("id, shop_name").Rows()
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}
	// for rows.Next() {
	// 	rows.Scan(&shop.ID, &shop.ShopName)
	// }
	return result
}
