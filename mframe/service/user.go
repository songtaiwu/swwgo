package service

import (
	"swwgo/mframe/dao"
	"swwgo/mframe/dao/impl"
	"swwgo/mframe/models"
	"time"
)

var UserServices = UserService{
	UserDao: impl.UserDaoImpls,
}

type UserService struct {
	UserDao dao.UserDao
}

func (u *UserService)AddUser(name string) (*models.User, error) {
	user := &models.User{
		Name:    name,
		Created: time.Now().Unix(),
	}
	err := u.UserDao.Save(user)
	if err != nil {
		return user, nil
	}
	return nil, err
}