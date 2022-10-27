package impl

import (
	"github.com/xormplus/xorm"
	"swwgo/mframe/models"
	"swwgo/mframe/package/orm"
)

var UserDaoImpls = &UserDaoImpl{
	db : orm.GetDB(),
}

type UserDaoImpl struct {
	db *xorm.Engine
}

func (u *UserDaoImpl) List() ([]models.User, error) {
	ret := []models.User{}

	session := u.db.NewSession()
	session.Find(ret)
	return ret, nil
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

