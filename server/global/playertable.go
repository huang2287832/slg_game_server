package global

import (
	"sync"
	"fmt"
)

type globalList struct {
	playerList sync.Map
	len int32
}

var GloInstance *globalList

func NewGlobalPlayers() {
	GloInstance = &globalList{}
}

func (g *globalList) GetPlayerList() (listMap sync.Map) {
	return g.playerList
}

func (g *globalList) GetPlayersNum() (int, int32) {
	length := 0
	g.playerList.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length, g.len
}

func (g *globalList) GetPlayerIdList() []int32 {
	var listId []int32
	g.playerList.Range(func(userId, _ interface{}) bool {
		listId = append(listId, userId.(int32))
		return true
	})
	return listId
}

func (g *globalList) AddPlayer(userId int32, player interface{}) {
	g.playerList.Store(userId, player)
	g.len += 1
}

func (g *globalList) DelPlayer(userId int32) {
	g.playerList.Delete(userId)
	g.len -= 1
}

func (g *globalList) GetPlayer(userId int32) (player interface{}) {
	player, _ = g.playerList.Load(userId)
	return
}

func (g *globalList) BroadCast() {
	g.playerList.Range(func(userId, _ interface{}) bool {
		fmt.Printf("---------------- broadcast userid = %d\n", userId)
		return true
	})

}