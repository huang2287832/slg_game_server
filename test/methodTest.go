package main

import "fmt"

//在编程语言中，方法和函数的概念需要搞清楚。函数指的是一个封装的代码块，我们可以直接调用它，并返回结果。
//而方法其实也是一种函数，只不过方法需要和某个对象绑定。Golang并没有类的概念，不过仍然有方法和接口这些概念。
type ICar interface {
	beep()
}

func main() {
	Method()
	fmt.Println("---------------- method = ", car1.id)
}