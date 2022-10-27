package service

import (
	"swwgo/mframe/dao"
	"swwgo/mframe/dao/impl"
	"swwgo/mframe/models"
)

type UserService struct {
	UserDao dao.UserDao
}

func (u *UserService)AddUser(name string) (user *models.User) {
	u.UserDao.Save(name)
	return
}