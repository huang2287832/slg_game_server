package handler

import (
	"fmt"
	"slg_game_server/server/game/user"
	"slg_game_server/server/game/rank"
	proto3 "slg_game_server/proto"
)

func init() {
	user.Handler.RegistHandler(proto3.ProtoCmd_CMD_ShowRankReq, &proto3.ShowRankReq{}, handleRank)
}

func handleRank(req interface{},  player *user.Player) interface{} {
	msg := req.(*proto3.ShowRankReq)
	var rankObj []*proto3.Rank
	var myRank int32
	switch msg.Key {
	case proto3.RankEnum_Rank_Power:
		rankObj, myRank = rank.RkInstance.PowerRank.GetPowerRank(player.Attr.UserId)
	default:
		fmt.Println("rank client send err = ", msg)
		panic(msg)
		return nil
	}
	pbData := &proto3.ShowRankResp{MyRank:myRank, NickName:player.Attr.LordName, RankObj:rankObj}
	player.SendMessage(&user.Message{Cmd:proto3.ProtoCmd_CMD_ShowRankResp, PbData:pbData})
	return nil
}

