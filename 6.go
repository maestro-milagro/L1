package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// 1-й Вариант через канал, который отпраляет сигнал о том что горутина завешила свою работу
func worker1(done chan bool) {
	// Выполняем задачу
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Посылаем сигнал о завершении работы в канал
	done <- true
}

// 2-й Вариант с помощью контекста, мы ждем когда нам придет сигнал из контекста на завершение программы
func worker2(ctx context.Context) {
	// Выполняем задачу
	// Проверяем, не требуется ли завершение
	for i := 10; i < 200000; i++ {
		select {
		case <-ctx.Done():
			// Завершаем работу, если контекст отменен
			fmt.Println("Get back")
			return
		default:
			fmt.Println(i)
			// Продолжаем работу
		}
	}
}

// 3-й Вариант с помощью sync.WaitGroup, мы ждем когда завершиться работа в горутине и вызываем wg.Done() чтобы дать понять вейтгруппе что горутина выполнилась
func worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	s := "some calc"
	for _, i := range s {
		fmt.Println(string(i))
	}
}
func main() {
	wg := sync.WaitGroup{}
	done := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
	defer cancel()
	go worker1(done)
	go worker2(ctx)
	wg.Add(1)
	go worker3(&wg)
	//Ждем сигнала из канала done
	<-done
	time.Sleep(2 * time.Second)
	//Ждем, пока горутина завершится
	wg.Wait()
}
