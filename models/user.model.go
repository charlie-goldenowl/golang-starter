package models

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string    `gorm:"type:varchar(255);not null"`
	Email string    `gorm:"uniqueIndex;not null"`
}
