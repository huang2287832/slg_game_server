package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	a := new(MyStruct)
	a.name = "yejianfeng"
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("-------------------", reflect.ValueOf(a).Elem().FieldByName("name"))
	a1 := reflect.TypeOf(a).Elem().Name()
	fmt.Printf("-------------------%T \n", a1)
	var b MyStruct
	b.name = "abc"
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b).FieldByName("name"))
}