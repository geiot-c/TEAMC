package controller

import (
	"TEAMC/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CandidateShop struct {
	Id   string `json:"id"`
	Name string `json:"Name"`
}

type SelectedShop struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetCandidates(c *gin.Context) {
	LunchAndDinner := []CandidateShop{}
	LongerThanThirty := []CandidateShop{}
	ShorterThanThirty := []CandidateShop{}
	ids := []string{"test"}
	res := model.FindShop(ids)
	fmt.Println(res)

	for _, shop := range res {
		LunchAndDinner = append(LunchAndDinner, CandidateShop{strconv.Itoa(int(shop.ID)), shop.ShopName})
	}
	// LunchAndDinner = append(LunchAndDinner, CandidateShop{"1", "栗"})
	// LunchAndDinner = append(LunchAndDinner, CandidateShop{"2", "おんどり"})
	// LunchAndDinner = append(LunchAndDinner, CandidateShop{"3", "La Terrasse “irisée”"})
	// LunchAndDinner = append(LunchAndDinner, CandidateShop{"4", "カフェ エトランジェ ナラッド "})
	// LunchAndDinner = append(LunchAndDinner, CandidateShop{"5", "ALL DAY DINING"})
	// LongerThanThirty = append(LongerThanThirty, CandidateShop{"6", "中川政七商店　奈良の工芸に触れる体験"})
	// LongerThanThirty = append(LongerThanThirty, CandidateShop{"7", "今西清兵衛商店"})
	// LongerThanThirty = append(LongerThanThirty, CandidateShop{"8", "ならまち格子の家"})
	// LongerThanThirty = append(LongerThanThirty, CandidateShop{"9", "瑜伽山園地"})
	// LongerThanThirty = append(LongerThanThirty, CandidateShop{"10", "寧楽美術館"})
	// ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"11", "日本酒とおつまみ　chuin"})
	// ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"12", "なら泉勇斎"})
	// ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"13", "樫舎"})
	// ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"14", "SUNNY and MORE"})
	// ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"15", "ボリクコーヒー"})

	c.JSON(http.StatusOK, gin.H{"lnd": LunchAndDinner, "lthan30": LongerThanThirty, "sthan30": ShorterThanThirty})
}

func GetResult(c *gin.Context) {
	SelectedShops := []SelectedShop{}

	fmt.Println(c.PostFormArray("ids[]"))

	SelectedShops = append(SelectedShops, SelectedShop{"1", "栗"})
	SelectedShops = append(SelectedShops, SelectedShop{"2", "おんどり"})
	SelectedShops = append(SelectedShops, SelectedShop{"3", "La Terrasse “irisée”"})
	SelectedShops = append(SelectedShops, SelectedShop{"4", "カフェ エトランジェ ナラッド "})
	SelectedShops = append(SelectedShops, SelectedShop{"5", "ALL DAY DINING"})
	SelectedShops = append(SelectedShops, SelectedShop{"6", "中川政七商店　奈良の工芸に触れる体験"})
	SelectedShops = append(SelectedShops, SelectedShop{"7", "今西清兵衛商店"})
	SelectedShops = append(SelectedShops, SelectedShop{"8", "ならまち格子の家"})
	SelectedShops = append(SelectedShops, SelectedShop{"9", "瑜伽山園地"})
	SelectedShops = append(SelectedShops, SelectedShop{"10", "寧楽美術館"})
	SelectedShops = append(SelectedShops, SelectedShop{"11", "日本酒とおつまみ　chuin"})
	SelectedShops = append(SelectedShops, SelectedShop{"12", "なら泉勇斎"})
	SelectedShops = append(SelectedShops, SelectedShop{"13", "樫舎"})
	SelectedShops = append(SelectedShops, SelectedShop{"14", "SUNNY and MORE"})
	SelectedShops = append(SelectedShops, SelectedShop{"15", "ボリクコーヒー"})

	reqTime := 230

	c.JSON(http.StatusOK, gin.H{"result": SelectedShops, "time": reqTime})
}
