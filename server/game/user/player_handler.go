package user

import (
	"github.com/gogo/protobuf/proto"
	proto3 "slg_game_server/proto"
)

var Handler *MsgData

type Message struct {
	Cmd  proto3.ProtoCmd
	PbData proto.Message
}

type callBackFunc func(interface{}, *Player) interface{}

type MsgData struct {
	msgInfo map[proto3.ProtoCmd]*MsgInfo
}

type MsgInfo struct {
	pb proto.Message
	callback callBackFunc
}

func init()  {
	Handler = newRegister()
}

func newRegister() *MsgData {
	msgData := new(MsgData)
	msgData.msgInfo = make(map[proto3.ProtoCmd]*MsgInfo)
	return msgData
}

func (msgData *MsgData) RegistHandler(cmd proto3.ProtoCmd, pbData proto.Message, callback callBackFunc) {
	msgInfo := &MsgInfo{pb:pbData, callback:callback}
	msgData.msgInfo[cmd] = msgInfo
}

func (msgData *MsgData) GetPbData(cmd proto3.ProtoCmd) (pbData proto.Message) {
	msgInfo := msgData.msgInfo[cmd]
	pbData = msgInfo.pb
	return
}

// 可以实现不同的callback
func (msgData *MsgData) Callback(cmd proto3.ProtoCmd, pbData proto.Message, player *Player) {
	msgInfo := msgData.msgInfo[cmd]
	callBack := msgInfo.callback
	callBack(pbData, player)

}
