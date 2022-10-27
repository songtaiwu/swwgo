package dao

import "swwgo/mframe/models"

type UserDao interface {
	List() ([]models.User, error)

	Get(id int) (*models.User, error)

	Save(user *models.User) error

	Update(user *models.User, id int) error

	Delete(id string) error
}
