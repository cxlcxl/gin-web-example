package model

import (
	"gorm.io/gorm"
	"time"
)

var (
	// KeyUserById Redis Keys
	KeyUserById = "db:api:user"
)

type db struct {
	*gorm.DB
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
