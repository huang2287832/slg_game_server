package main

import "fmt"

//golang的指针receiver和非指针receiver的区别?
//
//最大的区别应该是指针传递的是对像的引用，这样在方法里操作的时候可以动态修改对像的属性值。
//
//非指针传递的是对像的拷贝。

//eg:
//-record(Person, {
//	Name,
//	Age
//})
//eg:
// #r_user{Name = "hb", Age = 18}

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHi1() {
	p.Name = "leon1"

}

func (p Person) SayHi2() {
	p.Name = "leon2"

}

func main() {
	p1 := &Person{Name: "test", Age: 10}
	fmt.Printf("p1t = %T\n", &Person{Name: "test", Age: 10})
	fmt.Printf("p2t = %T\n", Person{Name: "test", Age: 10})
	fmt.Println("name1 : " + p1.Name)
	p1.SayHi1()
	fmt.Println("name2 : " + p1.Name)

	p2 := Person{Name: "test1", Age: 11}
	fmt.Println("name3: " + p2.Name)
	p2.SayHi2()
	fmt.Println("name4 : " + p2.Name)




}
