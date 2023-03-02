package dborm

import (
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

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

	switch viper.GetString("database.type") {
	case "sqlite":
		return dsn_sqlite()
	case "mysql":
		return dsn_mysql()
	default:
		return dsn_cli()
	}

}

func dsn_sqlite() gorm.Dialector {

	dir := viper.GetString("dataset.dir")
	name := viper.GetString("database.name")
	pragma := viper.GetString("database.param")

	dsn := dir + "/" + name + pragma

	if !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}

	return sqlite.Open(dsn)

}

func dsn_mysql() gorm.Dialector {

	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	passwd := viper.GetString("database.passwd")
	name := viper.GetString("database.name")
	pragma := viper.GetString("database.param")

	dsn := user + ":" + passwd + "@tcp(" + host + ")/" + name + pragma

	if !strings.Contains(dsn, "?") {
		pragma = "?charset=utf8mb4&parseTime=True&loc=Local"
	}

	return mysql.Open(dsn)

}

func dsn_cli() gorm.Dialector {

	dsn := viper.GetString("server.dsn")

	// mysql

	if strings.Contains(dsn, "@tcp") {
		if !strings.Contains(dsn, "?") {
			dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
		}
		return mysql.Open(dsn)
	}

	// sqlite

	if !strings.HasPrefix(dsn, "/") {
		dsn = viper.GetString("dataset.dir") + "/" + dsn
	}
	if !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
	}
	return sqlite.Open(dsn)

}
