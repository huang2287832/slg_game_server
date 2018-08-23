package main

import (
	"fmt"
)
func main() {
	var a []int
	a = append(a, 333)
	for index, value := range a[:len(a)]{
		fmt.Println(index, value)
	}
	fmt.Println(a)

	//var v []int
	////v = make([]int, 1, 1)
	//v[0] = 2
	//fmt.Println(v)

	//l := list.New()
	//l.PushBack(4)
	//l.PushBack(5)
	//l.PushBack(6)
	//l.PushBack(7)
	//l.PushBack(17)
	//l.PushBack(27)
	////l.Remove(6)
	//fmt.Println(l.Len())

	b := []int{1,2,3,4}
	for k, v := range b[1:3] {

		fmt.Println("---------a[] ", k)
		fmt.Println("---------v", v)
	}

}
