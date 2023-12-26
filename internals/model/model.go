package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model            // Adds some metadata fields to the table
	ID          uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
}
