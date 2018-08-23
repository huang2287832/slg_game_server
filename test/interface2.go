package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct {

}

func (c Cat) Speak() string {
	return "cat"
}

type Dog struct {

}

func (d Dog) Speak() string {
	return "dog"
}

func Test(params interface{})  {
	fmt.Println(params)
}

func main()  {
	animals := []Animal{Cat{}, Dog{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	b := &Cat{}
	a := b.Speak()
	fmt.Println("-------------a = ", a)
	Test("string")
	Test(123)
	Test(true)
}