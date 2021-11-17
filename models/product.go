package models

import "time"

type Product struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	Name        string `json:"name"`
	SerialNumer string `json:"serialNumer"`
}
