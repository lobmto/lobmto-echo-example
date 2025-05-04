package entity

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	Word    string
	Meaning string
	Example string
}
