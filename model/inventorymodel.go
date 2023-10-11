package model

import "time"

type Inventory struct {
	ID         uint      `json:"id"; gorm:"primary_key"`
	Firearm    string    `json:"firearm"`    //ปืน
	Bullet     string    `json:"bullet"`     //ขนาดกระสุน
	Barrel     string    `json:"barrel"`     //ลำกล้อง
	Brand      string    `json:"brand"`      //ยี่ห้อ
	Models     string    `json:"models"`     //โมเดิล
	Ammunition string    `json:"ammunition"` //กระสุน
	Stock      int64     `json:"stock"`
	Price      float64   `json:"price"`
	Image      string    `json:"image"`
	CreatedAt  time.Time `json:"time"`
}
