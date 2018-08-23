package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"slg_game_server/server/util"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8")
	util.CheckErr(err)
	err = DB.Ping()
	util.CheckErr(err)
	//defer DB.Close()	why close?
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(50)

}