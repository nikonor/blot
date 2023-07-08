package repo

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	StatusNotConfirmed = 0
	StatusConfirmed    = 1
)

type Repo struct {
	sync.Mutex
	db    *sql.DB
	cache map[string]string
}

type User struct {
	Login    string
	Token    *string
	Code     *string
	Status   int
	CreateAt time.Time
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
    status integer,
    code text,
    create_at text
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

func (r *Repo) GetToken(login string) (string, error) {
	upsert := `
insert into user (login, token, status, code, create_at)
values (?,?,?,?,?)
on conflict do
    update set token=?,status=?,code=?,create_at=?;
`
	code := genString(6)
	token := genToken()
	at := time.Now().String()

	if _, err := r.db.Exec(upsert, login, token, StatusNotConfirmed, code, at, token, StatusNotConfirmed,
		code, at); err != nil {
		return "", err
	}

	return code, nil
}

func (r *Repo) ConfirmUser(login, code string) (string, error) {
	upsert := `
update user
set status=?
where login=? and code=?;

`
	if _, err := r.db.Exec(upsert, StatusConfirmed, login, code); err != nil {
		return "", err
	}

	return code, nil
}

func (r *Repo) AddLink(login, link string) error {
	upsert := "insert into link (user_login, link, create_at) values (?,?,?)"
	at := time.Now().String()

	if _, err := r.db.Exec(upsert, login, link, at); err != nil {
		return err
	}

	return nil
}

func (r *Repo) AddNotify(login, duration, link string) error {
	upsert := "insert into link (user_login, link, create_at,notify_at) values (?,?,?,?)"

	now := time.Now()
	at := now.String()
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	notifyAt := now.Add(dur).String()

	if _, err := r.db.Exec(upsert, login, link, at, notifyAt); err != nil {
		return err
	}

	return nil
}
