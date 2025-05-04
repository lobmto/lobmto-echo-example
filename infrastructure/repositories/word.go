package repositories

import (
	"fmt"
	"lobmto-echo-example/domain/repositories"
	"lobmto-echo-example/infrastructure/models"
	"lobmto-echo-example/model/words"

	"gorm.io/gorm"
)

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) repositories.WordRepository {
	return wordRepository{db: db}
}

func (r wordRepository) FindByID(id words.ID) (words.Word, error) {
	var e models.Word
	if err := r.db.
		Preload("Meanings").
		First(&e, "id = ?", id.String()).Error; err != nil {
		return words.Word{}, err
	}

	id, err := words.NewIDFromString(e.ID)
	if err != nil {
		return words.Word{}, words.ErrInvalidWordID
	}

	meanings := make([]words.Meaning, len(e.Meanings))
	for i, meaning := range e.Meanings {
		meanings[i], err = words.NewMeaning(meaning.Meaning)
		if err != nil {
			return words.Word{}, err
		}
	}

	word, err := words.ReconstructWord(id, e.Word, meanings, []words.Tag{})
	if err != nil {
		fmt.Println(err)
		fmt.Println(word.MeaningList())
		return words.Word{}, err
	}
	return word, nil
}

func (r wordRepository) Create(word words.Word) (words.Word, error) {
	meanings := make([]models.Meaning, len(word.MeaningList()))
	for i, meaning := range word.MeaningList() {
		meanings[i] = models.Meaning{
			ID:      words.NewID().String(),
			Meaning: meaning.String(),
		}
	}
	e := models.Word{
		ID:       word.ID().String(),
		Word:     word.Word(),
		Meanings: meanings,
	}
	if err := r.db.Create(&e).Error; err != nil {
		return words.Word{}, err
	}
	return word, nil
}

func (r wordRepository) Delete(word words.Word) error {
	return r.db.
		Select("Meanings").
		Delete(&models.Word{
			ID: word.ID().String(),
		}).Error
}
