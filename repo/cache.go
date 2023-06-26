package repo

func (r *Repo) fillCache() error {
	r.Lock()
	defer r.Unlock()

	rows, err := r.db.Query("select token,login from user where length(token) > 0")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var t, l string
		if err = rows.Scan(&t, &l); err != nil {
			return err
		}
		r.cache[t] = l
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) addUser(u User) error {
	// TODO: валидация
	// TODO: запись в БД
	// TODO: запись в кэш
	return nil
}

func (r *Repo) updateUser(u User) error {
	// TODO: валидация
	// TODO: запись в БД
	// TODO: запись в кэш
	return nil
}
