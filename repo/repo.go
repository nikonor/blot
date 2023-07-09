package repo

import (
	"database/sql"
	"errors"
	"strconv"
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
	Link     string
	Note     string
	CreateAt time.Time
	NotifyAt *time.Time
}

const schema = `create table if not exists user (
    id text primary key,
    login text unique,
    token text,
    status integer not null default 0,
    code text,
    create_at text
);

create index if not exists idx_user_login on user (login);

create table if not exists link (
    user_login int not null references user(login),
    link text not null,
    archive int not null default 0,
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
	var (
		upsert = `
update user
set status=?
where login=? and code=? and status=?
returning token
`
		token string
	)

	if err := r.db.QueryRow(upsert, StatusConfirmed, login, code, StatusNotConfirmed).Scan(&token); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("invalid code or confirmation has already been confirmed")
		}
		return "", err
	}

	return token, nil
}

func (r *Repo) AddLink(token, link string) error {
	upsert := "insert into link (user_login, link, create_at) values (?,?,?)"
	at := time.Now().String()

	if _, err := r.db.Exec(upsert, token, link, at); err != nil {
		return err
	}

	return nil
}

func (r *Repo) AddNotify(token, duration, link string) error {
	upsert := "insert into link (user_login, link, create_at,notify_at) values (?,?,?,?)"

	now := time.Now()
	at := now.String()
	notifyAt, err := getEventDate(now, duration)
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(upsert, token, link, at, notifyAt); err != nil {
		return err
	}

	return nil
}

// getEventDate - разбираем период
//
//	duration = "<число>[m|h|d|y]"
func getEventDate(now time.Time, duration string) (time.Time, error) {
	timeDim := time.Second
	switch duration[len(duration)-1] {
	case 'm':
		timeDim = time.Minute
	case 'h':
		timeDim = time.Hour
	case 'd':
		timeDim = 24 * time.Hour
	case 'y':
		dur, err := strconv.ParseInt(duration[:len(duration)-1], 10, 64)
		if err != nil {
			return now, err
		}
		return time.Date(now.Year()+int(dur), now.Month(), now.Day(), now.Hour(), now.Minute(),
			now.Second(), 0, now.Location()), nil
	default:
		return now, errors.New("invalid dimension")
	}

	dur, err := strconv.ParseInt(duration[:len(duration)-1], 10, 64)
	if err != nil {
		return now, err
	}

	return now.Add(time.Duration(dur) * timeDim), nil
}
