package main

import (
	"fmt"
	"time"
)

func openhandler(ch chan int) {
	for  {
		fmt.Println("1")
		a, ok := <- ch
		fmt.Println("ok = ", ok)
		select {
		case <-ch:
			fmt.Println("2222222222222222")
			return
		default:
			fmt.Println("3333")
		}

		fmt.Println("3444", a)
	}
}

func main() {
	c := make(chan int, 1024)
	go openhandler(c)
	time.Sleep(1e9)
	//close(c)
	//fmt.Println(IsClosed(c)) // true
}