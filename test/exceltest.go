package main

import (
	"fmt"
	"os"
	"github.com/tealeg/xlsx"
	"strings"
)

func main() {
	excelFileName := "./doudizhu.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("config is wrong!!!")
		panic(err)
	}

	path := "../config/doudizhu.go"
	file, error := os.Create(path)
	defer file.Close()
	if error != nil {
		fmt.Println(error)
	}

	pkg := "config"
	obj := "Doudizhu"

	dataStruct := fmt.Sprintf("package %s\n\n", pkg)
	dataStruct += fmt.Sprintf("type %s struct{\n	", obj)

	row0 := xlFile.Sheets[0].Rows[0]
	row2 := xlFile.Sheets[0].Rows[2]

	var s string
	var dataType string
	for index, cell := range row0.Cells {
		fmt.Println(row2.Cells[index].String())
		fmt.Println("------------- ", row2.Cells[0].String())

		switch row2.Cells[index].String() {

		case "INT":
			dataType = " int"
		case "STRING":
			dataType = " string"
		case "TERM":
			dataType = " []int"

		}
		s = s + cell.String() + dataType + "\n	"
	}

	//file.WriteString(dataStruct + s + "\n}")

	key := "int"
	dataFunc := fmt.Sprintf("func Get_attr(key %s) *%s {\n", key, obj)
	dataFunc1 := dataStruct + s + "\n}\n\n" + dataFunc

	var msgdata string
	x := xlFile.Sheets[0]
	for k, v := range xlFile.Sheets[0].Rows[3:] {
		fmt.Println("-----------------k", k, v.Cells)
		var s2 string
		for index, cell := range v.Cells {
			array := cell.String()
			fmt.Println("----index", index)
			switch row2.Cells[index].String() {
			case "STRING":
				//c := "2"
				array = fmt.Sprintf("%q", array) //TODO
			case "TERM":
				array = strings.Replace(array, "[", "{", -1)
				array = strings.Replace(array, "]", "}", -1)
				array = fmt.Sprintf("[]int%s", array)

			}
			s2 += fmt.Sprintf("%s:%s, ", x.Rows[0].Cells[index].String(), array)
		}
		fmt.Println("----------------s2", s2)

		msg := fmt.Sprintf("		return &%s{%s}		\n", obj, s2)

		msgdata += fmt.Sprintf("	\n	if key == %s {\n %s \n	}", v.Cells[0].String(), msg)

		//data += s2
	}

	data := dataFunc1 + msgdata + "\n	return nil \n}	"

	file.WriteString(data)

}
