package words

type WordRepository interface {
	FindByID(id ID) (Word, error)
	Create(word Word) (Word, error)
	Delete(word Word) error
}
