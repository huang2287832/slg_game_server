package db

import (
	"fmt"
	"database/sql"
	"slg_game_server/server/util"
	"slg_game_server/server/include"
)

type Mail include.PlayerMail

func (m *Mail) InitData(name interface{}) interface{} {
	return m
}

func (m *Mail) SaveData() interface{} {
	return nil
}

func GetMailMaxId() int32 {
	rows := DB.QueryRow("select max(mail_id)+1 as id from t_user_mail")
	var num sql.NullInt64
	err := rows.Scan(&num)
	util.CheckErr(err)
	fmt.Println("------------null", num.Int64)
	if !num.Valid {
		return 100001

	}else {
		return int32(num.Int64)
	}
}