package repo

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	sync.Mutex
	db    *sql.DB
	cache map[string]string
}

type User struct {
	Login    string
	Token    *string
	Password string
}

type Link struct {
	URL      string
	Note     string
	CreateAt time.Time
	NotifyAt *time.Time
}

const schema = `create table if not exists user (
    id text primary key not null,
    login text unique,
    token text,
    password text
);

create index if not exists idx_user_login on user (login);

create table if not exists link (
    id text primary key not null,
    user_id int not null references user(id),
    link text not null,
    create_at text,
    notify_at text
);
`

func open(fname string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(schema); err != nil {
		return nil, err
	}
	return db, nil
}

func NewRepo(fName string) (*Repo, error) {
	db, err := open(fName)
	if err != nil {
		return nil, err
	}

	r := Repo{
		db:    db,
		cache: make(map[string]string),
	}

	if err = r.fillCache(); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *Repo) AddNewUser(login, password string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repo) ConfirmUser(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repo) ResurrectToken(login string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repo) ConfirmResurrectToken(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repo) AddLink(login, link string) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repo) AddNotify(login, duration, link string) error {
	// TODO implement me
	panic("implement me")
}
