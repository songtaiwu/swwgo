package tools

import (
	"swwgo/mframe/package/setting/conf"
)

const (
	envPgDbHost = "PG_DB_HOST"
	envPgDbName = "PG_DB_NAME"
	envPgDbUser = "PG_DB_USER"
	envPgDbPASS = "PG_DB_PASS"
	envPgDbTablePrefix = "PG_DB_TABLE_PREFIX"

	envHttpPort = "SERVICE_HTTP_PORT"
	envLogPath  = "SERVICE_LOG_PATH"

	defPgDbHost = "127.0.0.1"
	defPgDbName = "test"
	defPgDbUser = "root"
	defPgDbPASS = "123456"
	defPgDbTablePrefix = ""

	defHttpPort = 80
	defLogPath = "./log"
)



func EnvInit() {
	conf.PostgresDBSetting = &conf.PostgresDB{
		User:        Env(envPgDbUser, defPgDbUser),
		Password:    Env(envPgDbPASS, defPgDbPASS),
		Host:        Env(envPgDbHost, defPgDbHost),
		DbName:      Env(envPgDbName, defPgDbName),
		TablePrefix: Env(envPgDbTablePrefix, defPgDbTablePrefix),
	}

	conf.ServiceSetting = &conf.Service{
		HttpPort: EnvInt(envHttpPort, defHttpPort),
	}
}

