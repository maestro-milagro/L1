package main

import (
	"fmt"
	"math/rand"
)

// Предыдщая реализация была плоха тем что при создании среза HugeString мы все еще хранили
// ссылку да длинную строку в переменной justString хотя использовали только 100 из них
// В данной реализации в justString = []rune(v[:100]) мы создаем новый слайс рун и храним ссылку на него а не на длинную строку
var justString []rune

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Заполняем стрингу определённым в сигнатуре числом символов
func createHugeString(i int) (ans string) {
	b := make([]byte, i)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)
	//Создаем новый слайс рун
	justString = []rune(v[:100])
}

func main() {
	someFunc()
	fmt.Print(justString)
}
