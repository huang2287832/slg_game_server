package main

import (
	"fmt"
	"os"
)

func main() {

	type person struct {
		Id      int
		Name    string
		Country string
	}

	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "China"}

	fmt.Println("liumiaocn = ", liumiaocn)

	tmpl := template.New("tmpl1")
	tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	tmpl.Execute(os.Stdout, liumiaocn)

}
