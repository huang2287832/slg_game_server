package main

import "fmt"

// 定义接口类型 PeopleGetter 包含获取基本信息的方法
type PeopleGetter interface {
	GetName() string
	GetAge() int
}

// 定义接口类型 EmployeeGetter 包含获取薪水的方法
// EmployeeGetter 接口中嵌入了 PeopleGetter 接口，前者将获取后者的所有方法
type EmployeeGetter interface {
	PeopleGetter
	GetSalary() int
	Help()
}

// 定义结构 Employee
type Employee struct {
	name   string
	age    int
	salary int
	gender string
}

// 定义结构 Employee 的方法
func (self *Employee) GetName() string {
	return self.name
}

func (self *Employee) GetAge() int {
	return self.age
}

func (self *Employee) GetSalary() int {
	return self.salary
}

func (self *Employee) Help() {
	fmt.Println("This is help info.", self.age)
}

// 匿名接口可以被用作变量或者结构属性类型
type Man struct {
	gender interface {
		GetGender() string
	}
}

func (self *Employee) GetGender() string {
	return self.gender
}

// 定义执行回调函数的接口
type Callbacker interface {
	Execute()
}

// 定义函数类型 func() 的新类型 CallbackFunc
type CallbackFunc func()

// 实现 CallbackFunc 的 Execute() 方法
func (self CallbackFunc) Execute() { self() }

func main() {
	// 空接口的使用，空接口类型的变量可以保存任何类型的值
	// 空格口类型的变量非常类似于弱类型语言中的变量
	var varEmptyInterface interface{}
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)
	varEmptyInterface = 100
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)
	varEmptyInterface = "Golang"
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	// Employee 实现了 PeopleGetter 和 EmployeeGetter 两个接口
	varEmployee := Employee{
		name:   "Jack Ma",
		age:    50,
		salary: 100000000,
		gender: "Male",
	}
	fmt.Println("varEmployee is: ", varEmployee)
	varEmployee.Help()
	fmt.Println("varEmployee.name = ", varEmployee.GetName())
	fmt.Println("varEmployee.age = ", varEmployee.GetAge())
	fmt.Println("varEmployee.salary = ", varEmployee.GetSalary())

	// 匿名接口对象的使用
	varMan := Man{&Employee{
		name:   "Nobody",
		age:    20,
		salary: 10000,
		gender: "Unknown",
	}}
	fmt.Println("The gender of Nobody is: ", varMan.gender.GetGender())

	// 接口类型转换，从超集到子集的转换是可以的
	// 从方法集的子集到超集的转换会导致编译错误
	// 这种情况下 switch 不支持 fallthrough
	var varEmpInter EmployeeGetter = &varEmployee
	switch varEmpInter.(type) {
	case nil:
		fmt.Println("nil")
	case PeopleGetter:
		fmt.Println("PeopleGetter")
	default:
		fmt.Println("Unknown")
	}

	// 使用 “执行回调函数的接口对象” 执行回调函数
	// 这种做法的优势是函数显式地 “实现” 特定接口
	varCallbacker := CallbackFunc(func() { println("I am a callback function.") })
	varCallbacker.Execute()

}