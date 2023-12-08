package main

import (
	"fmt"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// 1. Пройдемся циклом по строке поочереди добавляя каждый символ к новой строке
func Reverse1(input string) (output string) {
	for _, r := range input {
		output = string(r) + output
	}
	return
}

// 2. Идем циклом с двух сторон меняя по две руны местами и преобразуем слайс рун в строку
func Reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	fmt.Println(Reverse1("The quick brown 狐 jumped over the lazy 犬"))
	fmt.Println(Reverse2("The quick brown 狐 jumped over the lazy 犬"))
}
