package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TouristComments struct {
	ID        uint   `gorm:"primary_key" json:"tourist_comments_id"`
	ShopId    uint   `json:"shop_id"`
	Commenter string `json:"commenter"`
	Comment   string `json:"comment"`
}

type Shop struct {
	// gorm.Model
	ID              uint              `gorm:"primary_key" json:"id"`
	ShopName        string            `gorm:"unique;not null" json:"name"`
	SelfIntro       string            `json:"self_intro"`
	Category1       string            `json:"category_1"`
	Category2       string            `json:"category_2"`
	Category3       string            `json:"category_3"`
	TouristComments []TouristComments `json:"tourist_comments"`
	Latitude        string            `json:"latitude"`
	Longitude       string            `json:"longitude"`
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
	// err := db.Where("id = ?", id).Find(&result).Error
	err := db.Table("shops").Select("shops.id, shop_name, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude").
		Where("shops.id = ?", id).
		Joins("inner join categories as cat1 on category_1_id = cat1.id").
		Joins("inner join categories as cat2 on category_2_id = cat2.id").
		Joins("inner join categories as cat3 on category_3_id = cat3.id").
		Scan(&result).Error
	// rows, err := db.Table("shops").Where("shop_name = ?", ids).Select("id, shop_name").Rows()
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}

	comments := []TouristComments{}
	err = db.Where("shop_id = ?", id).Find(&comments).Error
	if err != nil {
		panic(err.Error())
	}

	result.TouristComments = comments
	// for rows.Next() {
	// 	rows.Scan(&shop.ID, &shop.ShopName)
	// }
	return result
}

func GetRecommendations(id string) []Recommend {
	db := dbConnect()
	defer db.Close()

	result := []Recommend{}
	err := db.Table("recommends").
		Select("shops.id, shop_name, intro, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude").
		Joins("inner join shops on recommended_shop_id = shops.id").
		Where("recommender_id = ?", id).
		Joins("inner join categories as cat1 on category_1_id = cat1.id").
		Joins("inner join categories as cat2 on category_2_id = cat2.id").
		Joins("inner join categories as cat3 on category_3_id = cat3.id").
		Order("shops.id").
		Scan(&result).Error
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}

	comments := []TouristComments{}
	err = db.Table("recommends").
		Select("tourist_comments.id, shop_id, commenter, comment").
		Joins("inner join tourist_comments on recommended_shop_id = shop_id").
		Where("recommender_id = ?", id).
		Order("shop_id").
		Scan(&comments).Error
	if err != nil {
		panic(err.Error())
	}

	ind := 0
	for i := range result {
		result[i].TouristComments = comments[ind : ind+3]
		ind += 3
	}

	fmt.Println(result)
	return result
}

func GetRecommendationsByOthers(id string) []Recommend {
	db := dbConnect()
	defer db.Close()

	result := []Recommend{}
	err := db.Table("recommends").
		Select("shops.id, shop_name, intro, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude").
		Joins("inner join shops on recommender_id = shops.id").
		Where("recommended_shop_id = ?", id).
		Joins("inner join categories as cat1 on category_1_id = cat1.id").
		Joins("inner join categories as cat2 on category_2_id = cat2.id").
		Joins("inner join categories as cat3 on category_3_id = cat3.id").
		Order("shops.id").
		Scan(&result).Error
	// rows, err := db.Table("shops").Where("shop_name in (?)", ids).Select("id, shop_name").Rows()
	if err != nil {
		panic(err.Error())
	}

	comments := []TouristComments{}
	err = db.Table("recommends").
		Select("tourist_comments.id, shop_id, commenter, comment").
		Joins("inner join tourist_comments on recommender_id = shop_id").
		Where("recommended_shop_id = ?", id).
		Order("shop_id").
		Scan(&comments).Error
	if err != nil {
		panic(err.Error())
	}

	ind := 0
	for i := range result {
		result[i].TouristComments = comments[ind : ind+3]
		ind += 3
	}

	fmt.Println(result)
	return result
}
