package def

const (
	CORE_SVR_ID = 1
	LOG_SVR_ID  = 5
	GAME_SVR_ID = 11
	DB_SVR_ID   = 12
)

//gorm数据库的定义

type User struct {
	Accid string `gorm:"not null;unique_index"`
	Name  string
	Gold  int
	Info  string
}

type NormalCraft struct {
	Id     int
	Author string
	Rect   int
	Data   string
	Praise int
}

type GoodCraft struct {
	Id     int
	Author string
	Rect   int
	Data   string
	Praise int
}
