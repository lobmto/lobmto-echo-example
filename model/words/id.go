package words

import (
	"github.com/google/uuid"
)

type ID struct {
	value string
}

func NewID() ID {
	return ID{
		value: uuid.New().String(),
	}
}

func NewIDFromString(id string) (ID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return ID{}, ErrInvalidWordID
	}
	return ID{value: id}, nil
}

func (id ID) Value() uuid.UUID {
	// バリデーション済みなのでエラーは想定しない
	return uuid.MustParse(id.value)
}

func (id ID) String() string {
	return id.value
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
}
