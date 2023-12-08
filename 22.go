package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

// Воспользуемся пакетом math/big для работы с большими числами
func main() {
	//Объявим переменные как рацианальные числа, чтобы охватить все наиболее часто используемые
	var a big.Rat
	var b big.Rat
	var c big.Rat
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = fmt.Scan(&b)
	fmt.Println("Умножение")
	fmt.Println(c.Mul(&a, &b))
	fmt.Println("Деление")
	fmt.Println(c.Quo(&a, &b))
	fmt.Println("Сложение")
	fmt.Println(c.Add(&a, &b))
	fmt.Println("Вычитание")
	fmt.Println(c.Add(&a, b.Neg(&b)))
}
