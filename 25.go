package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

// Воспользуемся функцией After, которая ожидает истечения продолжительности
// duration, а затем отправляет текущее время по возвращаемому каналу(все это время канал ждет ответа).
func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("some")
	Sleep(2)
	fmt.Println("...sleep")
}
