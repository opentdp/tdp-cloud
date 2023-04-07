package dborm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

func dialector() gorm.Dialector {

	switch args.Database.Type {
	case "sqlite":
		return useSqlite()
	case "mysql":
		return useMysql()
	default:
		logman.Fatal("Database type error", logman.String("type", args.Database.Type))
	}

	return nil

}

func useSqlite() gorm.Dialector {

	dir := args.Dataset.Dir
	name := args.Database.Name

	option := args.Database.Option
	if option == "" {
		option = "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}

	return sqlite.Open(dir + "/" + name + option)

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
