package orm

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/names"
	"swwgo/mframe/package/setting/conf"
	"time"
)

var db *xorm.Engine
var err error

func Init() {
	dbname := conf.PostgresDBSetting.DbName
	user := conf.PostgresDBSetting.User
	password := conf.PostgresDBSetting.Password
	host := conf.PostgresDBSetting.Host
	tablePrefix := conf.PostgresDBSetting.TablePrefix
	isShowSql := conf.PostgresDBSetting.IsShowSql

	str := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err = xorm.NewPostgreSQL(str)
	if err != nil {
		panic(err)
	}

	// xorm内置了三种Mapper实现，默认为SnakeMapper，即struct为驼峰命名，表结构为下划线命名
	mapper := names.NewPrefixMapper(names.SnakeMapper{}, tablePrefix)
	db.SetTableMapper(mapper)

	// 显示sql语句
	db.ShowSQL(isShowSql)

	// 设置连接池的空闲数大小
	db.SetMaxIdleConns(2)

	// 设置最大打开连接数
	db.SetMaxOpenConns(1000)

	// 设置连接的最大生存时间
	db.SetConnMaxLifetime(time.Hour)

}