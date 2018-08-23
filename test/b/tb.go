package main

import (
	"fmt"
	"slg_game_server/test/i"
	"slg_game_server/test/a"
)

type B struct {
}

func (b B) PrintA() {
	fmt.Println("bbbbbbbb")
}

func NewB() *B {
	b := new(B)
	return b
}

func RequireA(o i.Interfa) {
	o.PrintA()
}

func main()  {
	var obj i.Interfa
	obj = a.NewA()
	RequireA(obj)

}
