package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

// Определим метод для поиска заданного в сигнатуре элемента в множестве и вернём true в случае если такой элемент есть
func Contains(set map[interface{}]bool, i interface{}) bool {
	if _, ok := set[i]; !ok {
		return false
	}
	return true
}

// Определим метод для нахождения пересечения между двумя множествами
// и в ответ вернём множество хронящее элементы которые встречаются в обоих множествах
func Intersect(set1, set2 map[interface{}]bool) map[interface{}]bool {
	intersect := make(map[interface{}]bool)
	//Итерируемся по множество с меньшим числом элементов
	if len(set1) < len(set2) {
		for v := range set1 {
			if Contains(set2, v) {
				intersect[v] = true
			}
		}
	} else {
		for v := range set2 {
			if Contains(set1, v) {
				intersect[v] = true
			}
		}
	}
	return intersect
}

func main() {
	set1 := make(map[interface{}]bool)
	for i := 0; i < 10; i++ {
		set1[i] = true
	}
	set2 := make(map[interface{}]bool)
	for i := 0; i < 5; i++ {
		set2[i] = true
	}
	fmt.Print(Intersect(set1, set2))
}
