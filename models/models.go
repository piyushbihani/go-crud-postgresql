package models

type Stock struct {
	StockId int `json:"stockid"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Company string `json:"company"`
}