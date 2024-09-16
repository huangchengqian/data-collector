package dao

import (
	"data-collector/weibo/hot/vo"
	"gorm.io/gorm"
)

func SaveHot(data vo.Resp, conn *gorm.DB) {
	conn.Save(data)
}
