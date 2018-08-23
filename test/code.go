package main

import (
	"fmt"
)

type abc struct {
	a int
}

func main() {
	a := make(map[interface{}]interface{})
	a[1] = 2
	//fmt.Println(a[1])

	fmt.Println(new(abc))
	fmt.Println(abc{})
}
