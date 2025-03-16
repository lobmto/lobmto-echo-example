package words

type Tag struct {
	id ID
}

func NewTag(id ID) Tag {
	return Tag{id: id}
}

func (t Tag) ID() ID {
	return t.id
}

func (t Tag) Equals(other Tag) bool {
	return t.id.Equals(other.id)
}
