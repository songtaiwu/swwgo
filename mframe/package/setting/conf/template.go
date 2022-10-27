package conf

import (
	"time"
)

type PostgresDB struct {
	User string
	Password string
	Host string
	DbName string
	TablePrefix string
	IsShowSql bool
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
}

var ServiceSetting = &Service{}

// 日志相关
type Log struct {
	FilePath string
	Level string // debug、info、error、fatal
	MaxSize int // 每个日志文件保存的最大尺寸 单位：Mb
	MaxBackups int // 日志文件最多保存多少个备份
	MaxAge    int // 文件最多保存多少天
	Compress  bool // 是否压缩
}

var LogSetting = &Log{}