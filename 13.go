package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

// Первый способ поменять местами указатели на объекты
func swap1(first, last *int) {
	*first, *last = *last, *first
}

func main() {
	first, last := 1, 2
	fmt.Println(1)
	fmt.Println("до: ", first, last)
	swap1(&first, &last)
	fmt.Println("после: ", first, last)
	//Второй споссоб
	fmt.Println(2)
	first2, last2 := 1, 2
	fmt.Println("до: ", first2, last2)
	first2, last2 = last2, first2
	fmt.Println("после: ", first2, last2)
	//Третий способ с помощью операции xor
	fmt.Println(3)
	a, b := 1, 2
	fmt.Println("до: ", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("после: ", a, b)
	//Четвертый способ с помощью сложения вычитания
	fmt.Println(4)
	a1, b1 := 1, 2
	fmt.Println("до: ", a1, b1)
	a1 = a1 + b1
	b1 = a1 - b1
	a1 = a1 - b1
	fmt.Println("после: ", a1, b1)

}
