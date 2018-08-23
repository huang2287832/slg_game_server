package user

import (
	"io"
	"fmt"
	"slg_game_server/server/db"
	"slg_game_server/server/global"
	"github.com/gogo/protobuf/proto"
	"slg_game_server/server/game/rank"
	proto3 "slg_game_server/proto"
)

type Player struct {
	Attr 			*db.Attr
	Mail 			*db.Mail
	Task 			*db.Task
	Writer			io.Writer
	writeChan 		chan interface{}	// 等待接受的数据
	PidChan			chan interface{}

}

func NewPlayer() *Player {
	player := new(Player)
	player.PidChan = make(chan interface{}, 1024)
	go func() {
		for {
			msgData, ok := <- player.PidChan
			if !ok {
				fmt.Println("--------------- pidchan close = ", msgData)
				return
			}
			switch msg := msgData.(type) {
			case *Message:
				fmt.Printf("--------------msgid = %v, pbData = %v\n", msg.Cmd, msg.PbData)
				Handler.Callback(msg.Cmd, msg.PbData, player)
			default:

			}
		}
	}()

	return player
}

func (player *Player) CreatePlayer(name interface{}){

	var playerObj global.PlayerObj		// 定义一个接口类型

	playerObj = new(db.Attr)
	player.Attr = playerObj.InitData(name).(*db.Attr)
	userId := player.Attr.UserId

	// 重复登录的玩家,需要先摧毁之前的旧对象
	OldPlayer := global.GloInstance.GetPlayer(userId)
	fmt.Println("------------ oldplayer", OldPlayer)
	if OldPlayer != nil {
		OldPlayer.(*Player).DestoryPlayer()
	}

	playerObj = new(db.Mail)
	player.Mail = playerObj.InitData(name).(*db.Mail)

	playerObj = new(db.Task)
	player.Task = playerObj.InitData(name).(*db.Task)

	global.GloInstance.AddPlayer(userId, player)

}

func (player *Player) DestoryPlayer() {
	fmt.Println("--------------- DestoryPlayer = ")
	close(player.PidChan)
	close(player.writeChan)
	global.PlayerObj.SaveData(player.Attr)
	global.GloInstance.DelPlayer(player.Attr.UserId)

}

func (player *Player) SetWriteChan(sendChan chan interface{}) {
	player.writeChan = sendChan
}

func (player *Player) SendMessage(msg interface{}) {
	fmt.Println("SendMessage", msg)
	player.writeChan <- msg
}

func (player *Player) EventHandler(cmd proto3.ProtoCmd, pbData proto.Message) {
	player.PidChan <- &Message{Cmd:cmd, PbData:pbData}
}

func (player *Player) getPlayerAttr(attrId int32) int32 {
	return 0
}

func (player *Player) GetLevel() int32 {
	return player.Attr.Level
}

func (player *Player) SetLevel(level int32)  {
	player.Attr.Level = level
}

func (player *Player) SetPower(power int32)  {
	player.Attr.Power = power
	rank.RkInstance.PowerRank.SetPowerRank(player.Attr)
}