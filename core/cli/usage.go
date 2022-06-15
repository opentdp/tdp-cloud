package cli

import "flag"

var (
	Dsn     string
	Address string
)

func Flags() {

	flag.StringVar(&Dsn, "dsn", "cloud.db", "数据库存储路径")

	flag.StringVar(&Address, "address", ":7800", "服务器监听地址和端口")

	flag.Parse()
}
