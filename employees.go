package main

import "gorm.io/gorm"

type Weapons struct {
	gorm.Model
	// WId   int     `json:"id"`
	WName string  `json:"wname"`
	Wtype string  `json:"type"`
	Price float64 `json:"price"`
}

//start
