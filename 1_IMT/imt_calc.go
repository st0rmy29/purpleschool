package main

import (
	"fmt"
	"math"
)

func calcIMT(height, weight float64) float64 {

	const IMTPower = 2

	IMT := weight / math.Pow(height/100, IMTPower)

	return IMT
}

func inputValue() (float64, float64) {
	var height float64
	var weight float64

	fmt.Println("___ Калькулятор индекса массы тела ___")

	fmt.Print("Рост (см): ")
	fmt.Scan(&height)

	fmt.Print("Вес (кг): ")
	fmt.Scan(&weight)

	return height, weight
}

func outPutResult(IMT float64) {

	fmt.Printf("Индекс массы тела: %.2f\n", IMT)

	switch {
	case IMT < 16:
		fmt.Println("Сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("Дефицит массы тела")
	case IMT < 25:
		fmt.Println("Нормальный вес")
	case IMT < 30:
		fmt.Println("Избыточный вес")
	default:
		fmt.Println("Ожирение")
	}
}

func main() {

	repeat := "н"

	for {

		height, weight := inputValue()

		IMT := calcIMT(height, weight)

		outPutResult(IMT)

		fmt.Print("Повторить? (д/н): ")
		fmt.Scan(&repeat)

		if repeat != "д" && repeat != "Д" {
			break
		}

	}
}
