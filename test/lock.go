package main

import (
	"fmt"
	"time"
)

//func main() {
//	var mutex sync.Mutex
//	fmt.Println("Lock the lock")
//	mutex.Lock()
//	fmt.Println("The lock is locked")
//	channels := make([]chan int, 4)
//	for i := 0; i < 4; i++ {
//		channels[i] = make(chan int)
//		go func(i int, c chan int) {
//			fmt.Println("Not lock: ", i)
//			mutex.Lock()
//			fmt.Println("Locked: ", i)
//			time.Sleep(time.Second)
//			fmt.Println("Unlock the lock0: ", i)
//			mutex.Unlock()
//			c <- i
//		}(i, channels[i])
//	}
//	time.Sleep(time.Second)
//	fmt.Println("Unlock the lock")
//	mutex.Unlock()
//	time.Sleep(time.Second)
//
//	for _, c := range channels {
//		<-c
//	}
//}

//func main() {
//	var mutex sync.Mutex
//	fmt.Println("Lock the lock. (G0)")
//	mutex.Lock()
//	fmt.Println("The lock is locked. (G0)")
//	for i := 1; i <= 3; i++ {
//		go func(i int) {
//			fmt.Printf("Lock the lock. (G%d)\n", i)
//			mutex.Lock()
//			fmt.Printf("The lock is locked. (G%d)\n", i)
//			mutex.Unlock()
//		}(i)
//	}
//	time.Sleep(time.Second)
//	fmt.Println("Unlock the lock. (G0)")
//	mutex.Unlock()
//	fmt.Println("The lock is unlocked. (G0)")
//	time.Sleep(time.Second)
//}


func print(i int) {
	fmt.Println("Hello world ", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go print(i)
	}
	time.Sleep(time.Second)
}

//func main() {
//	for i := 1; i <= 100; i++ {
//		go func(i int) {
//			fmt.Printf("Lock the lock. (G%d)\n", i)
//		}(i)
//	}
//	time.Sleep(time.Second)
//	fmt.Println("Unlock the lock. (G0)")
//	time.Sleep(time.Second)
//}