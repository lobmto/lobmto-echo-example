package models

type Word struct {
	ID   string `gorm:"primaryKey"`
	Word string `gorm:"not null"`
}
