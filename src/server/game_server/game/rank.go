package module

import (
    "frame/def"
    "frame/logger"
    "sort"
    "time"
    "math/rand"
)

type RankObj struct {
    Id int
    Typ int
    Praise int
}

func NewRankObj(id int, praise int, typ int) *RankObj {
    return &RankObj{
        Id: id,
        Typ: typ,
        Praise: praise,
    }
}

type RankMgr struct {
    Objs    []*RankObj
    IdMap  map[int]int
}

func NewRankMgr() *RankMgr {
    return &RankMgr{
        Objs: make([]*RankObj, 0, def.RankNum),
        IdMap: make(map[int]int),
    }
}

func (this *RankMgr) DoSort() {
    sort.Stable(*this)
    this.IdMap = make(map[int]int)
    for k, v := range this.Objs {
        this.IdMap[v.Id] = k
    }
    this.DumpData()
}

func (this *RankMgr) DumpData() {
    for _, v := range this.Objs {
        logger.Info("Rank: id = %d, praise = %d", v.Id, v.Praise)
    }
}

func (this *RankMgr) AddRankData(id int, praise int, typ int) bool {
    if praise < def.RankPraiseLimit {
        return false
    }
    rankLen := len(this.Objs)
    obj := NewRankObj(id, praise, typ)
    if rankLen < def.RankNum {
        this.Objs = append(this.Objs, obj)
        return true
    } else {
        if obj.Praise > this.Objs[rankLen-1].Praise {
            this.Objs[rankLen-1] = obj
            return true
        }
    }
    return false
}

func (this *RankMgr) UpdateRankData(id int, praise int, typ int) bool {
    if praise < def.RankPraiseLimit {
        return false
    }
    if idx, ok := this.IdMap[id]; ok {
        this.Objs[idx].Praise = praise
        return true
    }
    return this.AddRankData(id, praise, typ)
}

func (this *RankMgr) GetRandomPopular() []*def.CraftEntry {
    rankLen := len(this.Objs)
    showLen := def.RankShowNum
    if showLen > rankLen {
        showLen = rankLen
    }
    rand.Seed(time.Now().UnixNano())
    randIdx := rand.Perm(rankLen)

    crafts := make([]*def.CraftEntry, 0, rankLen)
    for k, idx := range randIdx {
        if k >= showLen {
            break
        }
        obj := this.Objs[idx]
        craftEntry := CraftApi.GetCraftInfo(obj.Typ, obj.Id)
        crafts = append(crafts, craftEntry)
    }
    return crafts
}

////////sort 接口///////
func (this RankMgr) Len() int {
    return len(this.Objs)
}

func (this RankMgr) Less(i, j int) bool {
    return this.Objs[i].Praise > this.Objs[j].Praise
}

func (this RankMgr) Swap(i, j int) {
    this.Objs[i], this.Objs[j] = this.Objs[j], this.Objs[i]
}
