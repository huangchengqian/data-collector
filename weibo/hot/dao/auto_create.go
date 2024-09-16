package dao

import (
	"data-collector/weibo/hot/dao/do"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitPartition(conn *do.DBConn) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conn.Host, conn.User, conn.Pass, conn.DBName, conn.Port, conn.SslMode, conn.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	dateNow := time.Now()
	db.Exec("CREATE TABLE \"weibo_hot\" (\"id\" bigserial NOT NULL,\"period\" int8,\"create_time\" timestamp,\"note\" varchar(255),\"word\" varchar(255),\"num\" bigint,\"url\" varchar(255), \"rank\" int8) PARTITION BY list(period)")
	for i := 0; i < 100; i++ {
		dateCycle := dateNow.Format("20060102")
		db.Exec(fmt.Sprintf("CREATE TABLE weibo_hot_%s PARTITION OF weibo_hot FOR VALUES in (%s)", dateCycle, dateCycle))
		dateNow = dateNow.AddDate(0, 0, 1)
	}

}
