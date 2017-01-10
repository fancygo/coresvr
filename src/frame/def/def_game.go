package def

//游戏内用到的定义值, 以及消息结构, 消息号

const (
	OK = 1
)

//用户信息
type UserInfo struct {
	Gold  int   `json:"gold" codec:"gold"`
	Color []int `json:"color" codec:"color"`
	Skin  []int `json:"skin" codec:"skin"`
}
