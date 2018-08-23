package main
import (
	"fmt"
	"sort"
)
type Tom struct {
	Name string
	Age int
}
type Persons []Tom

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (p Persons) Len() int { return len(p) }

// 交换数据
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// 嵌套结构体 将继承 Person 的所有属性和方法
// 所以相当于SortByName 也实现了 Len() 和 Swap() 方法
type SortByName struct{ Persons }

// 根据元素的姓名长度降序排序 （此处按照自己的业务逻辑写）
func (p SortByName) Less(i, j int) bool {
	return len(p.Persons[i].Name) > len(p.Persons[j].Name)
}

type SortByAge struct{ Persons }
// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）

func (p SortByAge) Less(i, j int) bool {
	return p.Persons[i].Age > p.Persons[j].Age
}

func main() {
	persons := Persons{
		{
			Name: "test123",
			Age: 20,
		},
		{
			Name: "test1",
			Age: 22,
		},
		{
			Name: "test12",
			Age: 21,
		},
	}
	fmt.Println("排序前")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
	sort.Sort(SortByName{persons})
	fmt.Println("排序后")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
}