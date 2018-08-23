package main

import (
	"sort"
	"fmt"
)

func main() {
	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	sort.Reverse(sort.IntSlice(intList))
	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

}