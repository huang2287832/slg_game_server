package main
import "fmt"
import "time"

func f1(ch chan int) {
	time.Sleep(time.Second * 1)
	ch <- 1
}
func f2(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1
}
func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go f1(ch1)
	go f2(ch2)

	//for  {
	//	select {
	//	case a := <-ch1:
	//		fmt.Println("The first case is selected.", a)
	//	case b := <-ch2:
	//		fmt.Println("The second case is selected.", b)
	//
	//	default:
	//		//fmt.Println("The 3 case is selected.")
	//	}
	//}

	//select {
	//case <-ch1:
	//	fmt.Println("The first case is selected.")
	//case <-ch2:
	//	fmt.Println("The second case is selected.")
	//}

	//go func() {
	//	for  {
	//		fmt.Println("looping")
	//	}
	//}()

	select {
	case a := <-ch1:
		fmt.Println("The first case is selected.", a)
	case b := <-ch2:
		fmt.Println("The second case is selected.", b)

	//default:
	//	fmt.Println("The 3 case is selected.")
	}
	fmt.Println("The 4 case is selected.")
}