package dao

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type PasswordDAO struct {
	db *database.ConnectDB
}

func NewPasswordDAO() *PasswordDAO {
	return &PasswordDAO{
		db: database.Instance(),
	}
}

func (p *PasswordDAO) Save(password domain.Password) (domain.Password, error) {
	var sql bytes.Buffer
	sql.WriteString("INSERT INTO passwords(name, user, password, url) VALUES (?, ?, ?, ?)")

	stmt, err := p.db.Source().Conn().Prepare(sql.String())
	if err != nil {
		log.Println("PassworDAO / Save / Prepare: ", err)
		return domain.Password{}, err
	}
	result, err := stmt.Exec(password.Name(), password.User(), password.Password(), password.Url())
	if err != nil {
		log.Println("PassworDAO / Save / stmt.Exec: ", err)
		return domain.Password{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("PassworDAO / Save / LastInsertId(): ", err)
		return domain.Password{}, err
	}
	password.WithId(id)
	return password, nil
}

func (p *PasswordDAO) Update(password domain.Password) error {
	return nil
}

func (p *PasswordDAO) Delete(password domain.Password) error {
	return nil
}

func (p *PasswordDAO) FindById(id int64) domain.Password {
	return domain.Password{}
}
