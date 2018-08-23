
package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func function(i int) {
	fmt.Println(i)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
	for i := 0; i < 10; i++ {
		//每创建一个goroutine，就把任务队列中任务的数量+1
		waitgroup.Add(1)
		go function(i)
	}
	//这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	waitgroup.Wait()
	fmt.Println("over")
}