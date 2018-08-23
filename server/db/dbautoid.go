package db

import (
	"fmt"
	"sync"
)

type increaseId struct {
	userId int32
	mailId int32
	taskId int32
	mutex sync.Mutex
}

var IdInstance *increaseId	// 全局共享实例,id分配

func NewIncrease() {
	userId := GetUserMaxId()
	fmt.Println("-----------max userid id = ", userId)
	mailId := GetMailMaxId()
	fmt.Println("-----------max mailid id = ", mailId)
	IdInstance = &increaseId{userId:userId, mailId:mailId}
}

func (i *increaseId) GetNewUserId() (userId int32) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	userId = i.userId
	i.userId += 1
	return
}

func (i *increaseId) GetNewMailId() (mailId int32) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	mailId = i.mailId
	i.mailId += 1
	return
}

