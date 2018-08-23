package config
//
//import (
//	"time"
//	"fmt"
//)
//
//type attr struct {
//	Userid int
//	Name string
//	Level int
//
//}
//
//func main()  {
//	a := attr{}
//	for i := 0; i < 100000; i++ {
//		b := a.get_attr(1)
//		fmt.Println(b.Level)
//	}
//
//	time.Sleep(10e9)
//}
//
//
//func (a *attr)get_attr(key int) attr {
//	if key == 1 {
//		return attr{Userid:1, Name:"aa", Level:1}
//	}
//	return attr{}
//}
//
