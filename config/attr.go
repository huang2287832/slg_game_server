package config


type attr struct {
	Userid int
	Name string
	Level int
	Info  []int
	Info1 attr1
}

type attr1 struct {
	a int
	b []int
	c string
}

func (a *attr)get_attr(key int) *attr {
	if key == 1 {
		return &attr{Userid:1, Name:"aa", Level:1, Info:[]int{12, 78, 50}, Info1:attr1{a:0,b:[]int{1,2,3},c:"a"}}
	}
	if key == 2 {
		return &attr{Userid:1, Name:"aa", Level:1, Info:[]int{12, 78, 50}}
	}
	if key == 3 {
		return &attr{Userid:1, Name:"aa", Level:1, Info:[]int{12, 78, 50}}
	}
	return nil
}

