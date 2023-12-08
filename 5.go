package main

import (
	"context"
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	var q time.Duration
	//Просим пользователя ввести время через которое программа закончится в консоль
	_, err := fmt.Scan(&q)
	if err != nil {
		return
	}
	//Для прекращения работы используем контекст WithTimeout
	ctx, cancel := context.WithTimeout(context.Background(), q*time.Second)
	defer cancel()
	//Создаем канал
	ch := make(chan int)
	//В отдельной горутине запускаем цикл который читает из канала
	go func() {
		for i := range ch {
			//Если пришел сигнал перестаём читать
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(i)
			}
		}
	}()
	//Пишем в канал
	for {
		//Если пришел сигнал завершаем запись
		select {
		case <-ctx.Done():
			return
		default:
			ch <- 1
		}
	}
}
