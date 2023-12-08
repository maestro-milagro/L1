package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	//Сохдадим множество и
	set := make(map[string]bool)
	for _, v := range str {
		set[v] = true
	}
	fmt.Print(set)
}
