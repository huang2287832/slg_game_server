package rank

import (
	"fmt"
	"sort"
	"slg_game_server/server/db"
	proto3 "slg_game_server/proto"
)

// 实时排名榜-势力值power

type powerRank struct {
	*rankObj
	powerList []*proto3.Rank
}

func (p *powerRank) Less(i, j int) bool {
	return p.powerList[i].PlayerAttr.Power > p.powerList[j].PlayerAttr.Power
}

func NewPowerRank(obj *rankObj) *powerRank {
	p := db.GetDBPowerRank(powerNum)
	return &powerRank{rankObj:obj, powerList:p}
}

func (p *powerRank) GetPowerRank(userId int32) (powerList []*proto3.Rank, myRank int32) {
	//r.mutex.RLock()
	//defer r.mutex.RUnlock()
	powerList = p.powerList
	var rankNum int32
	for _, value := range powerList {
		rankNum += 1
		value.RankNum = rankNum
		fmt.Printf("uid = %v, RankNum = %v, power = %v\n", value.PlayerAttr.UserId, value.RankNum, value.PlayerAttr.Power)
		if userId == value.PlayerAttr.UserId {
			myRank = rankNum
		}
	}
	fmt.Println("myrank = ", myRank)
	return powerList, myRank
}

func (p *powerRank) SetPowerRank(attr *db.Attr) {
	rankLen := len(p.powerList)
	last := p.powerList[rankLen - 1]
	_, myRank := p.GetPowerRank(attr.UserId)
	fmt.Printf("---------------- power : %d, attrpow : %d, last : %d\n", attr.Power, last.PlayerAttr.Power, last.PlayerAttr.Power)
	if attr.Power > last.PlayerAttr.Power || myRank > 0 {
		msg := &rankMsg{msg:"SetPowerRank", attr:attr}
		p.rankObj.rankChan <- msg
		//p.rankObj.rankChan <- msg
		fmt.Println("---------- rankchan send")
	}
}

func (p *powerRank) PowerRankSort(attr *db.Attr) interface{} {
	powerList, myRank := p.GetPowerRank(attr.UserId)
	playerAttr := &proto3.PlayerAttr{
		UserId: attr.UserId,
		NickName:attr.LordName,
		Level: attr.Level,
		Power: attr.Power,
		Domain: attr.Domain,
	}
	if myRank > 0 {		// 前300名里面有玩家变化
		index := myRank - 1
		powerList[index].PlayerAttr = playerAttr		// 首先更新这个玩家的信息
		if myRank == 1 && attr.Power >= powerList[index + 1].PlayerAttr.Power {	// 排名第一，而且最新值仍然大于第二名，排名不需要变动，只需要更新值
			return nil
		}
		// 最后一名，而且最新值仍然低于比倒数第二名，则不需要变动
		if myRank == int32(len(powerList)) && powerList[index - 1].PlayerAttr.Power >= attr.Power {
			return nil
		}
		if myRank > 1 && (powerList[index - 1].PlayerAttr.Power >= attr.Power && attr.Power >= powerList[index + 1].PlayerAttr.Power) {
			return nil
		}
		p.rankObj.sortRank = powerList
		sort.Sort(p)
	}else {				// 最新进入前300名，需要把最后一个玩家t掉
		lastIndex := len(powerList) - 1
		powerList[lastIndex].PlayerAttr = playerAttr
		p.rankObj.sortRank = powerList
		sort.Sort(p)
	}
	return nil

}
