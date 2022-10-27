package models

type User struct {
	Id         string `xorm:"not null pk UUID"`
	Name      string `xorm:"not null VARCHAR(200)"`
	Created    int64  `xorm:"default 0 BIGINT"`
}