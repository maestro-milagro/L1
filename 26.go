package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет,
//что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

// Определим функцию для проверки возвращающую true если все символы уникальные и false если нет
func checkUniq(s string) bool {
	//Сразу переведём заданную стрроку в верхний регистр чтобы сделать проверку регистронезависимой
	newS := strings.ToUpper(s)
	//Пройдемся по строке вызывая метод Count для каждого элемента и если число больше 1 возвращаем false
	for _, i := range newS {
		if strings.Count(newS, string(i)) > 1 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkUniq("abcdfghjkl;'va"))
}
