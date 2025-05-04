package repositories

import (
	"lobmto-echo-example/model/words"
)

type WordRepository interface {
	FindByID(id words.ID) (words.Word, error)
}
