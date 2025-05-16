package main

import "fmt"

func main() {

	defer fmt.Print("завершена")
	defer fmt.Print("Работа программы ")

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)

}
