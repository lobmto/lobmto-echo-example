package repositories

import (
	"lobmto-echo-example/domain/models"
	"lobmto-echo-example/model/words"
)

type WordRepository interface {
	FindByID(id words.ID) (models.Word, error)
}
