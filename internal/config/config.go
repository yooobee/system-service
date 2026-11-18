package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string `json:",env=MYSQL_DSN"` // 支持从环境变量覆盖
	}
}
