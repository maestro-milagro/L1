package main

import (
	"fmt"
	"strconv"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(arr []int, low, high int) {
	//Пока массив не отсортируется рекурсивно вызываем quickSort для левой и правой части массива
	//т. к. посути первый вызов partition делит массив на 2 части
	if low < high {
		var pivot = partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot+1, high)
	}
}

// Функция для частичной сортировки и нахождения индекса элемента
// вокруг которого будет проводиться сортировка в рекурсивных вызовах в функции quickSort
func partition(arr []int, low, high int) int {
	// Выбираем основное значение
	var pivot = arr[low]
	//Выделяем сортируемую область выбирая нижнюю и верхнюю границу
	var i = low
	var j = high
	//Запускаем цикл выполнение которого будет значить что мы прошлись по всем элементам
	for i < j {
		//Если элемент слева меньше или равен основному значению пропускаем его
		for arr[i] <= pivot && i < high {
			i++
		}
		//Если элемент справа больше основного значения пропускаем его
		for arr[j] > pivot && j > low {
			j--
		}
		//Если границы еще не соеденились на одном элементе
		//меняем левое и правое значение полученное в результате циклов местами
		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot
	//Возвращаем значение pivot для следующего рекурсивного вызова
	return j
}

// Функция для вывода массива в консоль
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int{15, 3, 12, 6, -9, 9, 0}

	fmt.Print("Before Sorting: ")
	printArray(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Print("After Sorting: ")
	printArray(arr)

}
