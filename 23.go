package main

import "fmt"

//Удалить i-ый элемент из слайса.

// Определим метод remove
func remove(slice []int, s int) []int {
	//Возвращаем слайс созданный из срезов всех элементов идущих до и после удалённого
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var s int
	sl := []int{1, 2, 3, 4, 5, 6}
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	fmt.Println(sl)
	fmt.Println(remove(sl, s))
}
