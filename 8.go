package main

import "fmt"

func main() {
	var i int64 = 2
	var z int64 = 12
	fmt.Printf("%064b", uint64(i))
	//инвертируем z-й бит в переменной i
	i ^= 1 << z
	fmt.Printf("%064b", uint64(i))
}
