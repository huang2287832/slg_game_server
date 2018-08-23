package main

//type gameObject struct {
//	name string
//}
//
//type player struct {
//	gameObject
//}
//
//func (p *player) Name() string {
//	return p.name
//}

type challenger interface {
	Name() string
	Attack() interface{}
}

func main() {
	a := challenger.Name
	println("=----------------a = ", a)
}
