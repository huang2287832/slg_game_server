package rank

import (
	"fmt"
	"slg_game_server/server/db"
	proto3 "slg_game_server/proto"
)

const (
	powerNum = 300
)

type rankMsg struct {
	msg 	string
	attr 	*db.Attr
}

// 父类：用于通信和排序
type rankObj struct {
	rankChan  chan *rankMsg
	sortRank []*proto3.Rank
	//PowerRank *powerRank
}

// 存储所有子排行榜实例对象
type RkInstances struct {
	PowerRank *powerRank
}

var RkInstance *RkInstances		// 全局共享实例，排行榜

// 排序通用接口
func (r *rankObj) Len() int      { return len(r.sortRank) }
func (r *rankObj) Swap(i, j int) { r.sortRank[i], r.sortRank[j] = r.sortRank[j], r.sortRank[i] }

func NewRank() {
	rankChan := make(chan *rankMsg, 1024)
	rankObj := &rankObj{rankChan:rankChan}
	powerRank := NewPowerRank(rankObj)
	RkInstance = &RkInstances{PowerRank:powerRank}
	go func() {
		for {
			rankMsg, ok := <- rankObj.rankChan
			if !ok {
				fmt.Println("--------------- rankChan close = ", rankMsg)
				return
			}
			switch rankMsg.msg {
			case "SetPowerRank":
				//RkInstance.PowerRank.PowerRankSort(rankMsg.attr)
				powerRank.PowerRankSort(rankMsg.attr)
			default:

			}
		}
	}()
}
