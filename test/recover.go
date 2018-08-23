package main

import (
	"fmt"
)

func main() {

	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err!=nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			//debug.PrintStack()
		}
	}()




	//index := 4
	//ia := [3]int{1,2,3}
	//_ = ia[index]


	panic("12345")
}
