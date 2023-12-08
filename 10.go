package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

// Создадим структуру в которую поместим по одному множеству для хранения значений температур с заданным шагом
type Set struct {
	subSet10    map[float64]bool
	subSet20    map[float64]bool
	subSet30    map[float64]bool
	subSetNeg20 map[float64]bool
}

func main() {
	ran := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	s := Set{subSet10: make(map[float64]bool), subSet20: make(map[float64]bool), subSet30: make(map[float64]bool), subSetNeg20: make(map[float64]bool)}
	// Пройдемся по массиву с значениями температур и с помощью switch будем добавлять значение в подходящую для него мапу
	for _, i := range ran {
		switch int(i / 10) {
		case 1:
			s.subSet10[i] = true
		case 2:
			s.subSet20[i] = true
		case 3:
			s.subSet30[i] = true
		case -2:
			s.subSetNeg20[i] = true
		default:
			fmt.Println("non")
		}
	}
	fmt.Print("10:")
	fmt.Println(s.subSet10)
	fmt.Print("20:")
	fmt.Println(s.subSet20)
	fmt.Print("30:")
	fmt.Println(s.subSet30)
	fmt.Print("-20:")
	fmt.Println(s.subSetNeg20)
}
