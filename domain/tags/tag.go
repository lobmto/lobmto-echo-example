package tags

type Tag struct {
	id   ID
	name string
}

func NewTag(name string) (Tag, error) {
	if err := validateTagName(name); err != nil {
		return Tag{}, err
	}
	
	return Tag{
		id:   NewID(),
		name: name,
	}, nil
}

func ReconstructTag(id ID, name string) (Tag, error) {
	if err := validateTagName(name); err != nil {
		return Tag{}, err
	}
	
	return Tag{
		id:   id,
		name: name,
	}, nil
}

func (t Tag) ID() ID {
	return t.id
}

func (t Tag) Name() string {
	return t.name
}

func (t Tag) Equals(other Tag) bool {
	return t.id.Equals(other.id)
}

func validateTagName(name string) error {
	if name == "" {
		return ErrEmptyTagName
	}
	return nil
} 