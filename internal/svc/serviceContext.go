package svc

import (
	"system-service/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	DB     sqlx.SqlConn // 注意是 SqlConn 接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
