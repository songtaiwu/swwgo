package impl

import (
	"swwgo/mframe/models"
)

type UserDaoImpl struct {
	
}

func (u *UserDaoImpl) List() ([]models.User, error) {
	panic("implement me")
}

func (u *UserDaoImpl) Get(id int) (*models.User, error) {
	panic("implement me")
}

func (u *UserDaoImpl) Save(user *models.User) error {
	panic("implement me")
}

func (u *UserDaoImpl) Update(user *models.User, id int) error {
	panic("implement me")
}

func (u *UserDaoImpl) Delete(id string) error {
	panic("implement me")
}

