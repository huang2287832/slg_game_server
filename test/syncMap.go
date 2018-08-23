package main

import (
	"fmt"
	"sync"
)

type userInfo struct {
	Name string
	Age  int
}

var m sync.Map

func main() {

	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) // oneone true

	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println("--------0", k, v)
		return true
	})

	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println("--------1", k, v)
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChamPly"
	user1.Age = 24
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = 18
	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")

	for k, v := range mapValue.(interface{}).(map[string]userInfo) {
		fmt.Println(k, v)
		fmt.Println("name:", v.Name)
	}
}