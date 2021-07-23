package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Shop struct {
	// gorm.Model
	ID        uint   `gorm:"primary_key" json:"id"`
	ShopName  string `gorm:"unique;not null" json:"name"`
	SelfIntro string `json:"self_intro"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Recommend struct {
	Intro string `json:"intro"`
	Shop
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

func ShowShops() []*Shop {
	db := dbConnect()
	defer db.Close()

	result := []*Shop{}
	err := db.Find(&result).Error
	if err != nil {
		panic(err.Error())
	}
	return result
}

func FindShop(id string) Shop {
	db := dbConnect()
	defer db.Close()

	result := Shop{}
	// err := db.Find(&result).Error
	err := db.Where("id = ?", id).Find(&result).Error
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

func GetRecommendations(id string) []Recommend {
	db := dbConnect()
	defer db.Close()

	result := []Recommend{}
	err := db.Table("recommends").Select("shops.id, shop_name, intro, self_intro, latitude, longitude").Joins("inner join shops on recommended_shop_id = shops.id").Where("recommender_id = ?", id).Scan(&result).Error
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
	return result
}

func GetRecommendationsByOthers(id string) []Recommend {
	db := dbConnect()
	defer db.Close()

	result := []Recommend{}
	err := db.Table("recommends").Select("shops.id, shop_name, intro, self_intro, latitude, longitude").Joins("inner join shops on recommender_id = shops.id").Where("recommended_shop_id = ?", id).Scan(&result).Error
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
	return result
}
