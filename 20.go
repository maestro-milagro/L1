package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

// Создадим из исходной строки массив строк разделив её при помощи strings.Fields
func ReversWords(inp string) (output string) {
	words := strings.Fields(inp)
	//Так же как и в случае с символами будем добавлять элементы массива в итоговую строку добавляя пробел в начало
	for _, i := range words {
		output = " " + i + output
	}
	//Избавимся от первого пробела в строке
	return strings.TrimSpace(output)
}
func main() {
	fmt.Println(ReversWords("The quick brown 狐 jumped over the lazy 犬"))
}
