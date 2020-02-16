package common

import "time"

type SoftDeletableEntity struct {
	DeletedAt *time.Time `json:"-" gorm:"column:updated_at"`
}

type TimestampEntity struct {
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
