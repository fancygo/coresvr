package module

//游戏各模块的处理函数

var (
	UserApi  *UserMgr
	CraftApi *CraftMgr
	RankApi  *RankMgr
)

func Init() {
	UserApi = NewUserMgr()
	CraftApi = NewCraftMgr()
	RankApi = NewRankMgr()
}
