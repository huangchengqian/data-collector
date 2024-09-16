package dao

import (
	"data-collector/weibo/hot/dao/do"
	"testing"
)

func TestInitPartition(t *testing.T) {

	conn := &do.DBConn{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		DBName:   "postgres",
		Pass:     "hcq10086",
		SslMode:  "disable",
		Timezone: "Asia/Shanghai",
	}

	InitPartition(conn)

}
