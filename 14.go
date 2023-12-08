package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

// С помощью метода Sprintf из пакета fmt
func what(i interface{}) {
	fmt.Println(fmt.Sprintf("%T", i))
}

// С помощью type switch
func what2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan<- int:
		fmt.Println("chan:", v)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	r := true
	what(r)
	what2(r)
	//С помощью TypeOf из пакета reflect
	fmt.Println(reflect.TypeOf(r))
}
