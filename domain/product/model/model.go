package model

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Quantity    float64   `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestProduct struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" form:"name" valid:"required"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Quantity    float64   `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
