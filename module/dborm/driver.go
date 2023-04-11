package dborm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tdp-cloud/cmd/args"
)

func useSqlite() gorm.Dialector {

	name := args.Database.Name

	option := args.Database.Option
	if option == "" {
		option = "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}

	return sqlite.Open(name + option)

}

func useMysql() gorm.Dialector {

	host := args.Database.Host
	user := args.Database.User
	passwd := args.Database.Passwd
	name := args.Database.Name

	option := args.Database.Option
	if option == "" {
		option = "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return mysql.Open(user + ":" + passwd + "@tcp(" + host + ")/" + name + option)

}
