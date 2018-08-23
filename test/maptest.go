package main

import (
	"fmt"
	"sync"
)

func main()  {

	m := map[string]int{"a":1, "b":2, "c":3, "d":4, "e":5, "f":6}

	fmt.Println(m)
	fmt.Println("-------")

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(p map[string]int) {

		defer wg.Done()
		p["a"] = 10

	}(m) //因为go语言中的map为引用类型， 虽然go语言函数以传值方式调用，即函数持有的为参数副本，但因是引用类型， 故依然指向映射m ， 类似c的指针.
	go func(p map[string]int) {

		defer wg.Done()
		p["a"] = 14

	}(m)
	go func(p map[string]int) {

		defer wg.Done()
		p["a"] = 12

	}(m)

	wg.Wait()

	fmt.Println(m)
	}