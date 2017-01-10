package def

//gorm数据库的定义

type User struct {
	Accid string `gorm:"not null;unique_index"`
	Name  string
	Gold  int
	Info  string
}
