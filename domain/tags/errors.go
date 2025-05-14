package tags

import "errors"

var (
	ErrEmptyTagName = errors.New("tag name cannot be empty")
	ErrInvalidTagID = errors.New("invalid tag id")
) 