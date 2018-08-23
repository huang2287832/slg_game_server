package global

type PlayerObj interface {
	InitData(interface{}) interface{}
	SaveData() interface{}
}