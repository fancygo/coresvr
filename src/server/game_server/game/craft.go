package module

import (
	"frame/def"
	"frame/logger"
	"server/game_server/dbdata"
	"sync"
)

type Craft struct {
	Id     int
	Author string
	Rect   int
	Data   string
	Praise int
}

type CraftMgr struct {
	NormalIdMax int
	GoodIdMax   int
	NormalCraft []*Craft
	GoodCraft   []*Craft
	NormalMap   map[int]int
	GoodMap     map[int]int

	lock *sync.RWMutex
}

func NewCraftMgr() *CraftMgr {
	return &CraftMgr{
		NormalIdMax: def.NormalBaseId,
		GoodIdMax:   def.GoodBaseId,
		NormalCraft: make([]*Craft, 0),
		GoodCraft:   make([]*Craft, 0),
		NormalMap:   make(map[int]int),
		GoodMap:     make(map[int]int),

		lock: new(sync.RWMutex),
	}
}

func NewCraft(id int, author string, rect int, data string, praise int) *Craft {
	return &Craft{
		Id:     id,
		Author: author,
		Rect:   rect,
		Data:   data,
		Praise: praise,
	}
}

func (this *CraftMgr) InitNormalCraft(id int, author string, rect int, data string, praise int) {
	if this.NormalIdMax <= id {
		this.NormalIdMax = id + 1
	}
	craft := NewCraft(id, author, rect, data, praise)
	this.NormalCraft = append(this.NormalCraft, craft)
	this.NormalMap[id] = len(this.NormalCraft) - 1
}

func (this *CraftMgr) AddNormalCraft(author string, rect int, data string) int {
	this.lock.Lock()

	id := this.NormalIdMax
	this.NormalIdMax += 1
	craft := NewCraft(id, author, rect, data, 0)
	this.NormalCraft = append(this.NormalCraft, craft)
	this.NormalMap[id] = len(this.NormalCraft) - 1

	this.lock.Unlock()

	dbdata.SaveNormalCraft(id, author, rect, data, 0)
	return id
}

func (this *CraftMgr) GetNormalCraft() []*Craft {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.NormalCraft
}

func (this *CraftMgr) PraisePopular(id int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	sortFlag := false
	/*
	   if idx, ok := this.NormalMap[id]; !ok {
	       if idx, ok := this.GoodMap[id]; !ok {
	           return def.Ret_CraftIdNotMatch
	       } else {
	           craft := this.GoodCraft[idx]
	           craft.Praise += 1
	           sortFlag = RankApi.UpdateRankData(id, craft.Praise, def.CraftGood)
	           dbdata.SaveGoodCraft(craft.Id, craft.Author, craft.Rect, craft.Data, craft.Praise)
	       }
	   } else {
	       craft := this.NormalCraft[idx]
	       craft.Praise += 1
	       sortFlag = RankApi.UpdateRankData(id, craft.Praise, def.CraftNormal)
	       dbdata.SaveNormalCraft(craft.Id, craft.Author, craft.Rect, craft.Data, craft.Praise)
	   }
	   if sortFlag {
	       RankApi.DoSort()
	   }
	*/
	if idx, ok := this.NormalMap[id]; ok {
		craft := this.NormalCraft[idx]
		craft.Praise += 1
		sortFlag = RankApi.UpdateRankData(id, craft.Praise, def.CraftNormal)
		dbdata.SaveNormalCraft(craft.Id, craft.Author, craft.Rect, craft.Data, craft.Praise)
		if sortFlag {
			RankApi.DoSort()
		}
		return def.OK
	}
	if idx, ok := this.GoodMap[id]; ok {
		craft := this.GoodCraft[idx]
		craft.Praise += 1
		sortFlag = RankApi.UpdateRankData(id, craft.Praise, def.CraftGood)
		dbdata.SaveGoodCraft(craft.Id, craft.Author, craft.Rect, craft.Data, craft.Praise)
		if sortFlag {
			RankApi.DoSort()
		}
		return def.OK
	}
	return def.OK
}

////////////////////////精选/////////////////////

func (this *CraftMgr) InitGoodCraft(id int, author string, rect int, data string, praise int) {
	if this.GoodIdMax <= id {
		this.GoodIdMax = id + 1
	}
	craft := NewCraft(id, author, rect, data, praise)
	this.GoodCraft = append(this.GoodCraft, craft)
	this.GoodMap[id] = len(this.GoodCraft) - 1
}

func (this *CraftMgr) AddGoodCraft(author string, rect int, data string) int {
	this.lock.Lock()

	id := this.GoodIdMax
	this.GoodIdMax += 1
	craft := NewCraft(id, author, rect, data, 0)
	this.GoodCraft = append(this.GoodCraft, craft)
	this.GoodMap[id] = len(this.GoodCraft) - 1

	this.lock.Unlock()

	dbdata.SaveGoodCraft(id, author, rect, data, 0)
	return id
}

func (this *CraftMgr) GetGoodCraft() []*Craft {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.GoodCraft
}

////////////////////通用////////////////////////////
func (this *CraftMgr) GetCraftInfo(typ, id int) *def.CraftEntry {
	craftEntry := &def.CraftEntry{}
	craft := &Craft{}
	if typ == def.CraftGood {
		idx, ok := this.GoodMap[id]
		if !ok {
			return nil
		}
		if idx >= len(this.GoodCraft) {
			return nil
		}
		craft = this.GoodCraft[idx]
	} else if typ == def.CraftNormal {
		idx, ok := this.NormalMap[id]
		if !ok {
			return nil
		}
		if idx >= len(this.NormalCraft) {
			return nil
		}
		craft = this.NormalCraft[idx]
	} else {
		return nil
	}
	craftEntry.Id = craft.Id
	craftEntry.Author = craft.Author
	craftEntry.Rect = craft.Rect
	craftEntry.Data = craft.Data
	craftEntry.Praise = craft.Praise
	return craftEntry
}

func (this *CraftMgr) DumpCraft() {
	logger.Info("Dump NormalCraft")
	for _, v := range this.NormalCraft {
		logger.Info("craft = %+v", v)
	}
	logger.Info("Dump GoodCraft")
	for _, v := range this.GoodCraft {
		logger.Info("craft = %+v", v)
	}
}
