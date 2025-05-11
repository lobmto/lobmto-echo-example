package words

type Meaning struct {
	value string
}

func NewMeaning(meaning string) (Meaning, error) {
	if meaning == "" {
		return Meaning{}, ErrEmptyMeaning
	}
	return Meaning{value: meaning}, nil
}

func (m Meaning) String() string {
	return m.value
}

func (m Meaning) Equals(other Meaning) bool {
	return m.value == other.value
}
