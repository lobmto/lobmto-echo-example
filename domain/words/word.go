package words

type Word struct {
	id          ID
	word        string
	meaningList []Meaning
	tagList     []Tag
}

func NewWord(word string, meaningList []Meaning, tagList []Tag) (Word, error) {
	if err := validateWord(word); err != nil {
		return Word{}, err
	}
	if err := validateMeaningList(meaningList); err != nil {
		return Word{}, err
	}
	if err := validateTagList(tagList); err != nil {
		return Word{}, err
	}

	return Word{
		id:          NewID(),
		word:        word,
		meaningList: meaningList,
		tagList:     tagList,
	}, nil
}

func ReconstructWord(id ID, word string, meaningList []Meaning, tagList []Tag) (Word, error) {
	if err := validateWord(word); err != nil {
		return Word{}, err
	}
	if err := validateMeaningList(meaningList); err != nil {
		return Word{}, err
	}
	if err := validateTagList(tagList); err != nil {
		return Word{}, err
	}

	return Word{
		id:          id,
		word:        word,
		meaningList: meaningList,
		tagList:     tagList,
	}, nil
}

func (w Word) ID() ID {
	return w.id
}

func (w Word) Word() string {
	return w.word
}

func (w Word) MeaningList() []Meaning {
	return w.meaningList
}

func (w Word) TagList() []Tag {
	return w.tagList
}

func validateWord(word string) error {
	if word == "" {
		return ErrEmptyWord
	}
	return nil
}

func validateMeaningList(meaningList []Meaning) error {
	if len(meaningList) == 0 {
		return ErrEmptyMeaning
	}
	return nil
}

func validateTagList(tagList []Tag) error {
	if len(tagList) > 5 {
		return ErrTooManyTags
	}
	return nil
}
