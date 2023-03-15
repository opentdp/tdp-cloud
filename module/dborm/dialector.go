package dborm

import (
	"strings"
	"tdp-cloud/cmd/args"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dialector() gorm.Dialector {

	switch args.Database.Type {
	case "sqlite":
		return useSqlite()
	case "mysql":
		return useMysql()
	default:
		return useCli()
	}

}

func useSqlite() gorm.Dialector {

	dir := args.Dataset.Dir
	name := args.Database.Name
	option := args.Database.Option

	dsn := dir + "/" + name + option

	if !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}

	return sqlite.Open(dsn)

}

func useMysql() gorm.Dialector {

	host := args.Database.Host
	user := args.Database.User
	passwd := args.Database.Passwd
	name := args.Database.Name
	option := args.Database.Option

	dsn := user + ":" + passwd + "@tcp(" + host + ")/" + name + option

	if !strings.Contains(dsn, "?") {
		dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return mysql.Open(dsn)

}

func useCli() gorm.Dialector {

	dsn := args.Server.DSN

	// mysql

	if strings.Contains(dsn, "@tcp") {
		if !strings.Contains(dsn, "?") {
			dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
		}
		return mysql.Open(dsn)
	}

	// sqlite

	if !strings.HasPrefix(dsn, "/") {
		dsn = args.Dataset.Dir + "/" + dsn
	}
	if !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}
	return sqlite.Open(dsn)

}
