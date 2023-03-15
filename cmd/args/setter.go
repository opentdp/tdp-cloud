package args

import (
	"github.com/spf13/viper"

	"tdp-cloud/helper/strutil"
)

func init() {

	viper.SetDefault("dataset.dir", ".")
	viper.SetDefault("dataset.secret", strutil.Rand(32))

	viper.SetDefault("logger.dir", ".")
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.stdout", true)
	viper.SetDefault("logger.tofile", false)

	viper.SetDefault("server.jwtkey", strutil.Rand(32))

}

func Sync() {

	Debug = viper.GetBool("debug")

	Dataset.Dir = viper.GetString("dataset.dir")
	Dataset.Secret = viper.GetString("dataset.secret")

	Database.Type = viper.GetString("database.dir")
	Database.Host = viper.GetString("database.host")
	Database.User = viper.GetString("database.user")
	Database.Passwd = viper.GetString("database.passwd")
	Database.Name = viper.GetString("database.name")
	Database.Option = viper.GetString("database.option")

	Logger.Dir = viper.GetString("logger.dir")
	Logger.Level = viper.GetString("logger.level")
	Logger.Stdout = viper.GetBool("logger.stdout")
	Logger.ToFile = viper.GetBool("logger.tofile")

	Server.DSN = viper.GetString("server.dsn")
	Server.Listen = viper.GetString("server.listen")
	Server.JwtKey = viper.GetString("server.jwtkey")
	Server.Register = viper.GetBool("server.register")

	Worker.Remote = viper.GetString("worker.remote")

}
