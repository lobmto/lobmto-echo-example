package tags

import (
	"github.com/google/uuid"
)

// TODO: ドメイン間で共通化する
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
		return ID{}, ErrInvalidTagID
	}
	return ID{value: id}, nil
}

func (id ID) Value() uuid.UUID {
	return uuid.MustParse(id.value)
}

func (id ID) String() string {
	return id.value
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
} 