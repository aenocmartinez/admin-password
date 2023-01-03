package domain

type Tag struct {
	id        int64
	name      string
	passwords []Password
}

func NewTag(name string) *Tag {
	return &Tag{
		name: name,
	}
}

func (t *Tag) WithId(id int64) *Tag {
	t.id = id
	return t
}

func (t *Tag) WithName(name string) *Tag {
	t.name = name
	return t
}

func (t *Tag) WithPasswords(passwords []Password) *Tag {
	t.passwords = passwords
	return t
}

func (t *Tag) Id() int64 {
	return t.id
}

func (t *Tag) Name() string {
	return t.name
}

func (t *Tag) Passwords() []Password {
	return t.passwords
}
