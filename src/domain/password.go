package domain

type Password struct {
	id         int64
	name       string
	user       string
	password   string
	url        string
	tags       []Tag
	repository RepositoryPassword
}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) WithRepository(repository RepositoryPassword) *Password {
	p.repository = repository
	return p
}

func (p *Password) WithName(name string) *Password {
	p.name = name
	return p
}

func (p *Password) WithId(id int64) *Password {
	p.id = id
	return p
}

func (p *Password) WithUser(user string) *Password {
	p.user = user
	return p
}

func (p *Password) WithPassword(password string) *Password {
	p.password = password
	return p
}

func (p *Password) WithUrl(url string) *Password {
	p.url = url
	return p
}

func (p *Password) WithTags(tags []Tag) *Password {
	p.tags = tags
	return p
}

func (p *Password) Id() int64 {
	return p.id
}

func (p *Password) Name() string {
	return p.name
}

func (p *Password) User() string {
	return p.user
}

func (p *Password) Url() string {
	return p.url
}

func (p *Password) Password() string {
	return p.password
}

func (p *Password) Tags() []Tag {
	return p.tags
}

func (p *Password) AddTag(name string) {
	p.tags = append(p.tags, *NewTag(name))
}

func (p *Password) Exists() bool {
	return p.id > 0
}

func (p *Password) Save() error {
	password, err := p.repository.Save(*p)
	if err != nil {
		return err
	}
	p.WithId(password.Id())
	return nil
}

func (p *Password) Update() error {
	_, err := p.repository.Update(*p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Password) Delete() error {
	return p.repository.Delete(*p)
}

func PasswordList(repository RepositoryPassword) ([]Password, error) {
	return repository.All()
}

func FindPasswordById(id int64, repository RepositoryPassword) (Password, error) {
	return repository.FindById(id)
}
