package teststore

import (
	"github.com/Kentabr5427/http-rest-api/internal/app/model"
	"github.com/Kentabr5427/http-rest-api/internal/app/store"
)

//UserRepository ...
type UserRepository struct {
	//не указывать *teststore.Store, иначе начинается цикличный импорт пакетов
	//т.к пока мы находимся в данном пакете, он видит конкретно teststore.Store,
	//а не внешний sqlstore.Store
	store *Store
	users map[int]*model.User
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	for _, itemMap := range r.users {
		if itemMap.Email == email {
			return itemMap, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {

	if _, ok := r.users[id]; ok {
		return r.users[id], nil
	}

	return nil, store.ErrRecordNotFound

}
