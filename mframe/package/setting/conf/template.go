package conf

import "time"

type PostgresDB struct {
	User string
	Password string
	Host string
	DbName string
	TablePrefix string
}

var PostgresDBSetting = &PostgresDB{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

type Service struct {
	HttpPort int
	LogPath string
}

var ServiceSetting = &Service{}