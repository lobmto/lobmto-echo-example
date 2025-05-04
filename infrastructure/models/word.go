package models

type Word struct {
	ID       string    `gorm:"primaryKey"`
	Word     string    `gorm:"not null"`
	Meanings []Meaning `gorm:"foreignKey:WordID;constraint:OnDelete:RESTRICT"`
}

type Meaning struct {
	ID      string `gorm:"primaryKey"`
	Meaning string `gorm:"not null"`
	WordID  string `gorm:"not null"`
}
