package words

type WordsRepository interface {
	FindByID(id ID) (Word, error)
	Create(word Word) (Word, error)
	Delete(word Word) error
}
