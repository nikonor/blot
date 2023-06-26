package repo

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(fName string) (*Repo, error) {
	db, err := sql.Open("sqlite3", "file:"+fName+"?cache=shared")
	if err != nil {
		return nil, err
	}
	return &Repo{db: db}, nil
}

func (r Repo) AddNewUser(login, password string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repo) ConfirmUser(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repo) ResurrectToken(login string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repo) ConfirmResurrectToken(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r Repo) AddLink(login, link string) error {
	// TODO implement me
	panic("implement me")
}

func (r Repo) AddNotify(login, duration, link string) error {
	// TODO implement me
	panic("implement me")
}
