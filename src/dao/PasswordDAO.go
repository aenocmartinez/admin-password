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

func (p *PasswordDAO) Update(password domain.Password) (domain.Password, error) {
	var sql bytes.Buffer
	sql.WriteString("UPDATE passwords SET name=?, user=?, password=?, url=?, updatedOn=NOW() WHERE id=?")

	stmt, err := p.db.Source().Conn().Prepare(sql.String())
	if err != nil {
		log.Println("PassworDAO / Update / Prepare: ", err)
		return domain.Password{}, err
	}
	_, err = stmt.Exec(password.Name(), password.User(), password.Password(), password.Url(), password.Id())
	if err != nil {
		log.Println("PassworDAO / Update / stmt.Exec: ", err)
		return domain.Password{}, err
	}

	return password, nil
}

func (p *PasswordDAO) Delete(password domain.Password) error {
	var sql bytes.Buffer
	sql.WriteString("DELETE FROM passwords WHERE id=?")

	stmt, err := p.db.Source().Conn().Prepare(sql.String())
	if err != nil {
		log.Println("PassworDAO / Delete / Prepare: ", err)
		return err
	}
	_, err = stmt.Exec(password.Id())
	if err != nil {
		log.Println("PassworDAO / Delete / stmt.Exec: ", err)
		return err
	}

	return nil
}

func (p *PasswordDAO) FindById(id int64) (password domain.Password, err error) {
	password = domain.Password{}
	var sql bytes.Buffer
	sql.WriteString("SELECT id, name, user, password, url, createdOn, updatedOn FROM passwords WHERE id = ?")
	stmt, err := p.db.Source().Conn().Prepare(sql.String())
	if err != nil {
		log.Println("PassworDAO / FindById / Prepare: ", err)
		return password, err
	}
	row := stmt.QueryRow(id)
	id = 0
	var name, user, pass, url, createdOn, updatedOn string
	row.Scan(&id, &name, &user, &pass, &url, &createdOn, &updatedOn)
	password.WithId(id).WithName(name).WithUser(user).WithUrl(url).WithPassword(pass)

	return password, err
}

func (p *PasswordDAO) All() (passwordList []domain.Password, err error) {
	passwordList = []domain.Password{}
	var sql bytes.Buffer
	sql.WriteString("SELECT id, name, user, password, url, createdOn, updatedOn FROM passwords")

	stmt, err := p.db.Source().Conn().Prepare(sql.String())
	if err != nil {
		log.Println("PassworDAO / All / Prepare: ", err)
		return passwordList, err
	}
	rows, _ := stmt.Query()
	for rows.Next() {
		var id int64
		var name, user, password, url, createdOn, updatedOn string
		rows.Scan(&id, &name, &user, &password, &url, &createdOn, &updatedOn)
		pass := domain.NewPassword()
		pass.WithId(id).WithName(name).WithUser(user).WithPassword(password).WithUrl(url)
		passwordList = append(passwordList, *pass)
	}

	return passwordList, nil
}
