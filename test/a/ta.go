package a

import (
	"fmt"
)

type A struct {
}

func (a A) PrintA() {
	fmt.Println("aaaa0")
}

func NewA() *A {
	a := new(A)
	return a
}