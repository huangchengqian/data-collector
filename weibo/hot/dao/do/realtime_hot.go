package do

type Realtime struct {
	Id         int64  `gorm:"column:id;not null;type:int primary key auto_increment;comment:'id'"`
	Period     string `gorm:"column:period;type:varchar(255);comment:'Period'"`
	CreateTime int64  `gorm:"column:crete_time;type:int;comment:'timestamp'"`
	Note       string `gorm:"column:note;type:varchar(255);comment:'note'"`
	Word       string `gorm:"column:word;type:varchar(255);comment:'word'"`
	Num        int64  `gorm:"column:num;type:int;comment:'num'"`
	Url        string `gorm:"column:url;type:varchar(255);comment:'url'"`
}

func (u *Realtime) TableName() string {
	return "weibo_hot"
}
