package words

import "errors"

var (
	ErrEmptyWord     = errors.New("word cannot be empty")
	ErrEmptyMeaning  = errors.New("meaning cannot be empty")
	ErrTooManyTags   = errors.New("too many tags")
	ErrInvalidWordID = errors.New("invalid word id")
)
