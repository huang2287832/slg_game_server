package db

import (
	"slg_game_server/server/util"
	"slg_game_server/server/include"
	proto3 "slg_game_server/proto"
)

func GetDBPowerRank(num int) (powerRankObj []*proto3.Rank) {
	rows, err := DB.Query("select user_id, lord_name, level, power, domain from t_user " +
		"order by power desc, level desc, domain desc limit 0, ?", num)
	defer rows.Close()
	util.CheckErr(err)
	//var rankNum int32
	for rows.Next() {
		//rankNum += 1
		player := &include.PlayerAttr{}
		err = rows.Scan(&player.UserId, &player.LordName, &player.Level, &player.Power, &player.Domain)
		util.CheckErr(err)

		playerAttr := &proto3.PlayerAttr{
			UserId: 	player.UserId,
			NickName:	player.LordName,
			Level: 		player.Level,
			Power: 		player.Power,
			Domain: 	player.Domain,
		}
		powerRankObj = append(powerRankObj, &proto3.Rank{PlayerAttr:playerAttr})
		util.CheckErr(err)
	}
	return

}