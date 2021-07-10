package controller

import (
	"net/http"

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

// "昼食、夕食	栗	34.68096923516614, 135.83505440564
// 昼食、夕食	おんどり	34.697889069751675, 135.8452012810097
// 昼食、夕食	La Terrasse “irisée” 	34.70355272581798, 135.76559570736123
// 昼食、夕食	カフェ エトランジェ ナラッド 	34.682309921526944, 135.82655925771624
// 昼食、夕食	ALL DAY DINING	34.683420861642915, 135.82780645642984
// 30分以上	中川政七商店　奈良の工芸に触れる体験	34.66844085304311, 135.83071231303978
// 30分以上	今西清兵衛商店	34.6771167799442, 135.8346802547372
// 30分以上	ならまち格子の家	34.67519626978004, 135.83075745458194
// 30分以上	瑜伽山園地	34.6795394700259, 135.83924008867956
// 30分以上	寧楽美術館	34.68621397298806, 135.83719518341763
// 30分以下	日本酒とおつまみ　chuin	34.67879867636942, 135.83066651225275
// 30分以下	なら泉勇斎	34.679852836545734, 135.82983086566745
// 30分以下	樫舎	34.67875598474761, 135.8312995000453
// 30分以下	SUNNY and MORE	34.679749291117396, 135.8312133815698
// 30分以下	ボリクコーヒー	34.67665865074141, 135.83004641964288"

func GetCandidates(c *gin.Context) {
	LunchAndDinner := []CandidateShop{}
	LongerThanThirty := []CandidateShop{}
	ShorterThanThirty := []CandidateShop{}

	LunchAndDinner = append(LunchAndDinner, CandidateShop{"1", "栗"})
	LunchAndDinner = append(LunchAndDinner, CandidateShop{"2", "おんどり"})
	LunchAndDinner = append(LunchAndDinner, CandidateShop{"3", "La Terrasse “irisée”"})
	LunchAndDinner = append(LunchAndDinner, CandidateShop{"4", "カフェ エトランジェ ナラッド "})
	LunchAndDinner = append(LunchAndDinner, CandidateShop{"5", "ALL DAY DINING"})
	LongerThanThirty = append(LongerThanThirty, CandidateShop{"6", "中川政七商店　奈良の工芸に触れる体験"})
	LongerThanThirty = append(LongerThanThirty, CandidateShop{"7", "今西清兵衛商店"})
	LongerThanThirty = append(LongerThanThirty, CandidateShop{"8", "ならまち格子の家"})
	LongerThanThirty = append(LongerThanThirty, CandidateShop{"9", "瑜伽山園地"})
	LongerThanThirty = append(LongerThanThirty, CandidateShop{"10", "寧楽美術館"})
	ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"11", "日本酒とおつまみ　chuin"})
	ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"12", "なら泉勇斎"})
	ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"13", "樫舎"})
	ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"14", "SUNNY and MORE"})
	ShorterThanThirty = append(ShorterThanThirty, CandidateShop{"15", "ボリクコーヒー"})

	c.JSON(http.StatusOK, gin.H{"lnd": LunchAndDinner, "lthan30": LongerThanThirty, "sthan30": ShorterThanThirty})
}

func GetResult(c *gin.Context) {
	SelectedShops := []SelectedShop{}

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
