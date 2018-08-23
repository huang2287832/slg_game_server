package db

import (
	"fmt"
	"database/sql"
	"slg_game_server/server/util"
	"slg_game_server/config/dataConfig"
	"slg_game_server/server/include"
)

type Attr include.PlayerAttr

func (a *Attr) InitData(name interface{}) interface{} {
	rows, err := DB.Query("select * from t_user where acct_name = ?", name)
	defer rows.Close()
	util.CheckErr(err)
	emptyData := true
	for rows.Next() {
		err = rows.Scan(&a.UserId, &a.AcctName, &a.LordName, &a.Sign, &a.X, &a.Y, &a.Country, &a.Level,
			&a.Exp, &a.Wood, &a.Iron, &a.Stone, &a.Forage, &a.Gold, &a.Diamond, &a.BindDiamond,
			&a.Decree, &a.ArmyOrder, &a.Power, &a.Domain)
		util.CheckErr(err)
		emptyData = false
	}

	if emptyData {
		userId := IdInstance.GetNewUserId()
		a.UserId = userId
		a.AcctName = name.(string)

		// 读取配置初始化值
		a.Level = dataConfig.GetCfgPlayerAttr(20).InitData
		a.Exp = dataConfig.GetCfgPlayerAttr(21).InitData
		a.Wood = dataConfig.GetCfgPlayerAttr(22).InitData
		a.Iron = dataConfig.GetCfgPlayerAttr(23).InitData
		a.Stone = dataConfig.GetCfgPlayerAttr(24).InitData
		a.Forage = dataConfig.GetCfgPlayerAttr(25).InitData
		a.Gold = dataConfig.GetCfgPlayerAttr(26).InitData
		a.Diamond = dataConfig.GetCfgPlayerAttr(27).InitData
		a.BindDiamond = dataConfig.GetCfgPlayerAttr(28).InitData
		a.Decree = dataConfig.GetCfgPlayerAttr(29).InitData
		a.ArmyOrder = dataConfig.GetCfgPlayerAttr(30).InitData
		a.Power = dataConfig.GetCfgPlayerAttr(31).InitData
		a.Domain = dataConfig.GetCfgPlayerAttr(32).InitData
		fmt.Println("create new userId = ", userId)
		a.SaveData()
	}

	return a
}

func (a *Attr) SaveData() interface{} {
	fmt.Println("---------------------------------- save a:", a.LordName, "--", a.Sign)
	_, err := DB.Exec("replace into t_user(user_id, acct_name, lord_name, sign, x, y, country, level, exp, wood, iron," +
		"stone, forage, gold, diamond, bind_diamond, decree, army_order, power, domain) " +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		a.UserId, a.AcctName, a.LordName, a.Sign, a.X, a.Y, a.Country, a.Level,
		a.Exp, a.Wood, a.Iron, a.Stone, a.Forage, a.Gold, a.Diamond, a.BindDiamond,
		a.Decree, a.ArmyOrder, a.Power, a.Domain)
	util.CheckErr(err)
	return nil
}

func GetUserMaxId() int32 {
	rows := DB.QueryRow("select max(user_id)+1 as id from t_user")
	var num sql.NullInt64
	err := rows.Scan(&num)
	util.CheckErr(err)
	// 因为这里有可能为空值
	if !num.Valid {
		return 100001

	}else {
		return int32(num.Int64)
	}

}