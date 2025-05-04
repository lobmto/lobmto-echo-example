package repositories

import (
	"lobmto-echo-example/model/words"
)

type WordRepository interface {
	FindByID(id words.ID) (words.Word, error)
	Create(word words.Word) (words.Word, error)
	Delete(word words.Word) error
}
