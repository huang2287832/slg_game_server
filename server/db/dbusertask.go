package db

import "slg_game_server/server/include"

type Task include.PlayerTask

func (t *Task) InitData(name interface{}) interface{} {
	return t
}

func (t *Task) SaveData() interface{} {
	return nil
}