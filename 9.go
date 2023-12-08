package main

import "fmt"

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	//Ресурс из которого будем брать числа
	xr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	norm := make(chan int)
	normx2 := make(chan int)
	//Запись в первый канал
	go func() {
		for x := 0; x < len(xr); x++ {
			go func(x int) {
				norm <- xr[x]
			}(x)
		}
	}()
	//Запись во второй канал
	go func() {
		for x := 0; x < len(xr); x++ {
			normx2 <- xr[x] * 2
		}
		close(normx2)
	}()
	//Чтение
	for i := range normx2 {
		fmt.Println(i)
	}
}
