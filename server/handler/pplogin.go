package handler

import (
	"fmt"
	"time"
	"slg_game_server/server/game/user"
	proto3 "slg_game_server/proto"
)

func init() {
	user.Handler.RegistHandler(proto3.ProtoCmd_CMD_HeartBeatReq, &proto3.HeartBeatReq{}, handleHeartbeat)
	user.Handler.RegistHandler(proto3.ProtoCmd_CMD_LoginReq, &proto3.LoginReq{}, handleLogin)
	user.Handler.RegistHandler(proto3.ProtoCmd_CMD_LogoutReq, &proto3.LogOutReq{}, handleLogout)
}

func handleHeartbeat(req interface{},  player *user.Player) interface{} {
	cmd := proto3.ProtoCmd_CMD_HeartBeatResp
	unixTime := time.Now().Unix()
	pbData := &proto3.HeartBeatResp{ServerSec:unixTime}
	player.SendMessage(&user.Message{Cmd:cmd, PbData:pbData})
	return nil
}

func handleLogin(req interface{},  player *user.Player) interface{} {
	msg := req.(*proto3.LoginReq)
	acctName := msg.AcctName

	player.CreatePlayer(acctName)
	fmt.Println("login success = ", player.Attr.UserId)

	cmd := proto3.ProtoCmd_CMD_LoginResp
	attr := player.Attr
	playerAttr := &proto3.PlayerAttr{
		UserId:attr.UserId, NickName:attr.LordName, Sign:attr.Sign, X:attr.X, Y:attr.Y, Country:attr.Country, Level:attr.Level,
		Exp:attr.Exp, Wood:attr.Wood, Iron:attr.Iron, Stone:attr.Stone, Forage:attr.Forage, Gold:attr.Gold, Diamond:attr.Diamond,
		BindDiamond:attr.BindDiamond, Decree:attr.Decree, ArmyOrder:attr.ArmyOrder, Power:attr.Power, Domain:attr.Domain,
	}
	pbData := &proto3.LoginResp{PlayerAttr:playerAttr}	// TODO
	player.SendMessage(&user.Message{Cmd:cmd, PbData:pbData})

	//pbData1 := &proto3.ErrResp{Cmd:proto3.ProtoCmd_CMD_LoginReq, ErrCode:cmd, ErrMsg:"fuck"}	// TODO delete
	//player.SendMessage(&user.Message{Cmd:cmd, PbData:pbData1})

	return nil
}



func handleLogout(req interface{},  player *user.Player) interface{} {
	fmt.Println("------------handle_logOut = ", player.GetLevel())
	player.DestoryPlayer()
	return nil
}