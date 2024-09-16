package domain

import "time"


type Product struct {
	ID 				uint 		`json:"id" gorm:"primaryKey"`
	Name		 	string 		`json:"name" gorm:"index;"`
	Description 	string 		`json:"description"`
	Price 			float64 	`json:"price"`
	Stock 			uint 		`json:"stock"`
	ImageUrl 		string 		`json:"image_url"`
	CategoryId 		uint 		`json:"category_id"`
	UserId			uint 		`json:"user_id"`
	CreatedAt 		time.Time	`json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:current_timestamp"`
}
