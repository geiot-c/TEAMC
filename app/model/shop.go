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

type ShopRet struct {
	// gorm.Model
	ID              uint              `gorm:"primary_key" json:"id"`
	ShopName        string            `gorm:"unique;not null" json:"name"`
	SelfIntro       string            `json:"self_intro"`
	Category1       string            `json:"category_1"`
	Category2       string            `json:"category_2"`
	Category3       string            `json:"category_3"`
	TouristComments []TouristComments `json:"tourist_comments"`
	Latitude        float64           `gorm:"type:decimal(20,17)" json:"latitude"`
	Longitude       float64           `gorm:"type:decimal(20,17)" json:"longitude"`
	IsHot           *bool             `json:"is_hot"`
}

type Shop struct {
	// gorm.Model
	ID              uint              `gorm:"primary_key" json:"id" binding:"required"` //binding : ginのvalidationに利用
	ShopName        string            `gorm:"unique;not null" json:"name" binding:"required"`
	SelfIntro       string            `json:"self_intro" binding:"required"`
	Category1Id     uint              `json:"category_1" binding:"required"`
	Category2Id     uint              `json:"category_2" binding:"required"`
	Category3Id     uint              `json:"category_3" binding:"required"`
	TouristComments []TouristComments `json:"tourist_comments" binding:"required"`
	Latitude        float64           `gorm:"type:decimal(20,17)" json:"latitude" binding:"required"`
	Longitude       float64           `gorm:"type:decimal(20,17)" json:"longitude" binding:"required"`
	IsHot           *bool             `json:"is_hot" binding:"required"`
}

type Recommend struct {
	Intro string `json:"intro"`
	ShopRet
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

func FindShop(id string) ShopRet {
	db := dbConnect()
	defer db.Close()

	result := ShopRet{}
	// err := db.Find(&result).Error
	// err := db.Where("id = ?", id).Find(&result).Error
	err := db.Table("shops").Select("shops.id, shop_name, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude, is_hot").
		Where("shops.id = ?", id).
		Joins("inner join categories as cat1 on category1_id = cat1.id").
		Joins("inner join categories as cat2 on category2_id = cat2.id").
		Joins("inner join categories as cat3 on category3_id = cat3.id").
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
		Select("shops.id, shop_name, intro, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude, is_hot").
		Joins("inner join shops on recommended_shop_id = shops.id").
		Where("recommender_id = ?", id).
		Joins("inner join categories as cat1 on category1_id = cat1.id").
		Joins("inner join categories as cat2 on category2_id = cat2.id").
		Joins("inner join categories as cat3 on category3_id = cat3.id").
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
		Select("shops.id, shop_name, intro, self_intro, cat1.category_name as category1, cat2.category_name as category2, cat3.category_name as category3, latitude, longitude, is_hot").
		Joins("inner join shops on recommender_id = shops.id").
		Where("recommended_shop_id = ?", id).
		Joins("inner join categories as cat1 on category1_id = cat1.id").
		Joins("inner join categories as cat2 on category2_id = cat2.id").
		Joins("inner join categories as cat3 on category3_id = cat3.id").
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

func GetHotShops(center_shop ShopRet) []ShopRet {
	result := []ShopRet{}
	db := dbConnect()
	defer db.Close()

	err := db.Table("shops").
		Select("shops.id, shop_name, latitude, longitude, is_hot").
		Where("latitude between ? and ? and longitude between ? and ? and is_hot = ?", center_shop.Latitude-0.00500, center_shop.Latitude+0.00500, center_shop.Longitude-0.00500, center_shop.Longitude+0.00500, true).
		Order("shops.id").
		Scan(&result).Error

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
	return result
}

func UpdateShopStatus(shop *Shop) error {
	db := dbConnect()
	defer db.Close()

	err := db.Save(shop)
	if err != nil {
		panic(err.Error)
	}
	return err.Error
}

func SwitchShopHotness(id string, is_hot bool) error {
	db := dbConnect()
	defer db.Close()

	err := db.Table("shops").
		Where("shops.id = ?", id).
		Updates(map[string]interface{}{"is_hot": is_hot})

	if err != nil {
		panic(err.Error)
	}

	return err.Error
}
