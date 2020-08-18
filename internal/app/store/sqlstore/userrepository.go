package sqlstore

import (
	"database/sql"

	"github.com/Orynik/Backend_Blog/internal/app/model"
	"github.com/Orynik/Backend_Blog/internal/app/store"
)

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	v, err := r.store.db.Exec("INSERT INTO user (`email`, `encryptedPassword`) VALUES(?,?)",
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		return err
	}
	resultID, err := v.LastInsertId()

	u.ID = int(resultID)

	if err != nil {
		return err
	}

	return nil

}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT `id`, `email`, `encryptedPassword` FROM user WHERE email = ?",
		email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT `id`, `email`, `encryptedPassword` FROM user WHERE `id` = ?",
		id).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
