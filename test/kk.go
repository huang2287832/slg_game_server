package main

func main()  {
	var a int
	a = change(&a)
	println(a)
}

func change(a *int) int {
	return *a + 3
}
