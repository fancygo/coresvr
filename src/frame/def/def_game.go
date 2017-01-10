package def

//游戏内用到的定义值, 以及消息结构, 消息号

const (
	CraftGood       = 1
	CraftNormal     = 2
	CraftPopular    = 3
	GetCraftNum     = 30
	GetPraiseNum    = 50
	GoodBaseId      = 700000
	NormalBaseId    = 200000
	RankNum         = 7
	RankShowNum     = 3
	RankPraiseLimit = 1

	OK                     = 1
	Ret_UnMarshalERR       = 1000 //解析错误
	Ret_CraftTypeErr       = 1001 //上传craft类型错误
	Ret_CraftRectErr       = 1002 //上传craft rect错误
	Ret_CraftIdxErr        = 1003 //获取列表idx超出范围
	Ret_CraftIdNotMatch    = 1004 //点暂时id不匹配
	Ret_LoginAccidNameNUll = 1005 //登录accid或者name为空
	Ret_LoginAccidErr      = 1006 //登录accid不存在
	Ret_GoldNotEnough      = 1007 //金币不足
	Ret_UserNotLogin       = 1008 //用户未登录

	MsgId_UserLogin      = 101
	MsgId_UserLoginRet   = 102
	MsgId_UploadCraft    = 103
	MsgId_UploadCraftRet = 104
	MsgId_GetCraft       = 105
	MsgId_GetCraftRet    = 106
	MsgId_Praise         = 107
	MsgId_PraiseRet      = 108
	MsgId_GetPraise      = 109
	MsgId_GetPraiseRet   = 110
	MsgId_SyncInfo       = 111
	MsgId_SyncInfoRet    = 112
	MsgId_GetPopular     = 113
	MsgId_GetPopularRet  = 114

	MsgId_UploadCraftNoti = 201
)

//用户信息
type UserInfo struct {
	Gold  int   `json:"gold" codec:"gold"`
	Color []int `json:"color" codec:"color"`
	Skin  []int `json:"skin" codec:"skin"`
}

//用户登录
type UserLoginA struct {
	Accid string `json:"accid" codec:"accid"`
	Name  string `json:"name" codec:"name"`
}
type UserLoginR struct {
	Ret  int       `json:"ret" codec:"ret"`
	Info *UserInfo `json:"info" codec:"info"`
}

//关卡
type CraftEntry struct {
	Id     int    `json:"id" codec:"id"`
	Author string `json:"author" codec:"author"`
	Rect   int    `json:"rect" codec:"rect"`
	Data   string `json:"data" codec:"data"`
	Praise int    `json:"praise" codec:"praise"`
}

//上传新关卡
type UploadCraftA struct {
	Typ     int        `json:"typ" codec:"typ"`
	Craft   CraftEntry `json:"craft" codec:"craft"`
	Author  string     `json:"author" codec:"author"`
	Ifbroad int        `json:"ifbroad" codec:"ifbroad"`
}
type UploadCraftR struct {
	Ret     int `json:"ret" codec:"ret"`
	Craftid int `json:"craftid" codec:"craftid"`
}

//拉取关卡
type GetCraftA struct {
	Typ int `json:"typ" codec:"typ"`
	Idx int `json:"idx" codec:"idx"`
}
type GetCraftR struct {
	Ret    int           `json:"ret" codec:"ret"`
	Crafts []*CraftEntry `json:"crafts" codec:"crafts"`
}

//点赞结构
type PraiseEntry struct {
	Id     int `json:"id" codec:"id"`
	Praise int `json:"praise" codec:"praise"`
}

//点赞
type PraiseCraftA struct {
	Typ int `json:"typ" codec:"typ"`
	Idx int `json:"idx" codec:"idx"`
	Id  int `json:"id" codec:"id"`
}
type PraiseCraftR struct {
	Ret int `json:"ret" codec:"ret"`
}

//拉取点赞
type GetPraiseA struct {
	Typ int `json:"typ" codec:"typ"`
	Idx int `json:"idx" codec:"idx"`
}
type GetPraiseR struct {
	Ret     int            `json:"ret" codec:"ret"`
	Praises []*PraiseEntry `json:"praises" codec:"praises"`
}

//拉取人气
type GetPopularA struct {
	Idx int `json:"idx" codec:"idx"`
}
type GetPopularR struct {
	Ret    int           `json:"ret" codec:"ret"`
	Crafts []*CraftEntry `json:"crafts" codec:"crafts"`
}

//同步玩家信息entey
type SyncInfoA struct {
	Accid string    `json:"id" codec:"id"`
	Info  *UserInfo `json:"info" codec:"info"`
}

type SyncInfoR struct {
	Ret int `json:"ret" codec:"ret"`
}

///////通知////
type UploadCraftNoti struct {
	Author string `json:"author" codec:"author"`
}
