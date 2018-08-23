package db

import (
	"fmt"
	"slg_game_server/server/util"
	"slg_game_server/server/include"
)

type PowerRank include.RankPower

func (p *PowerRank) getForceRank()  {
	rows, err := DB.Query("select user_id, lord_name, domain, force from t_user " +
		"order by force desc,domain desc limit 0, ?")
	defer rows.Close()
	util.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&p.UserId, &p.LordName, &p.Domain, &p.Force)
		util.CheckErr(err)
		fmt.Println("---------------rank:", p)
	}
}
