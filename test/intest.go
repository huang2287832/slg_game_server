package main

import "fmt"

type CallBack interface {
	getName() string
	BaseCall
}

type BaseCall interface {
	doSomething()
}

func (user User) doSomething() {
	fmt.Println("do something")
}

type User struct {
	name string
	age int
}

func (user User) getName() string {
	return user.name
}

func main()  {
	user := User{name:"hb"}
	user.doSomething()
	a := CallBack(user)
	fmt.Println(a)
}
