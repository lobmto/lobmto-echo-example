package repositories

import (
	"lobmto-echo-example/domain/models"
	"lobmto-echo-example/domain/repositories"
	"lobmto-echo-example/infrastructure/entity"
	"lobmto-echo-example/model/words"

	"gorm.io/gorm"
)

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) repositories.WordRepository {
	return wordRepository{db: db}
}

func (r wordRepository) FindByID(id words.ID) (models.Word, error) {
	var e entity.Word
	if err := r.db.First(&e, id).Error; err != nil {
		return models.Word{}, err
	}
	return models.Word{
		ID:        e.ID,
		Word:      e.Word,
		Meaning:   e.Meaning,
		Example:   e.Example,
		CreatedAt: e.CreatedAt.Unix(),
		UpdatedAt: e.UpdatedAt.Unix(),
	}, nil
}
