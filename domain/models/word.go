package models

type Word struct {
	ID        uint
	Word      string
	Meaning   string
	Example   string
	CreatedAt int64
	UpdatedAt int64
}
