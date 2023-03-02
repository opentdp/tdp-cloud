package dborm

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

var Db *gorm.DB

func Connect() {

	config := &gorm.Config{
		Logger: NewLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if db, err := gorm.Open(dialector(), config); err != nil {
		logman.Fatal("Connect to databse error:", err)
	} else {
		Db = db
	}

}

func dialector() gorm.Dialector {

	switch args.Database.Type {
	case "sqlite":
		return dsn_sqlite()
	case "mysql":
		return dsn_mysql()
	default:
		return dsn_cli()
	}

}

func dsn_sqlite() gorm.Dialector {

	dir := args.Dataset.Dir
	name := args.Database.Name
	option := args.Database.Option

	dsn := dir + "/" + name + option

	if !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}

	return sqlite.Open(dsn)

}

func dsn_mysql() gorm.Dialector {

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

func dsn_cli() gorm.Dialector {

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
