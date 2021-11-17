package models

import "time"

type Order struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRef      int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRef"`
}
