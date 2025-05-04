package repositories

import (
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
	if err := r.db.First(&e, id).Error; err != nil {
		return words.Word{}, err
	}

	id, err := words.NewIDFromString(e.ID)
	if err != nil {
		return words.Word{}, words.ErrInvalidWordID
	}

	word, err := words.ReconstructWord(id, e.Word, []words.Meaning{}, []words.Tag{})
	if err != nil {
		return words.Word{}, err
	}
	return word, nil
}

func (r wordRepository) Create(word words.Word) (words.Word, error) {
	e := models.Word{
		ID:   word.ID().String(),
		Word: word.Word(),
	}
	if err := r.db.Create(&e).Error; err != nil {
		return words.Word{}, err
	}
	return word, nil
}

func (r wordRepository) Delete(word words.Word) error {
	return r.db.Delete(&models.Word{}, word.ID().String()).Error
}
