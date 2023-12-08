package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

// Функция worker, которая принимает на вход контекст и данные из канала,
// осуществляет чтение из заданного канала chan int до тех пор пока не получит сигнал из контекста на завершение
func worker11(ctx context.Context, c chan int, wg *sync.WaitGroup) {
	for {
		//Проверяем на каждой итерации не пришел ли нам сигнал на завершение работы, если нет выводим данные из канала в консоль
		select {
		case <-ctx.Done():
			fmt.Println("Shot down")
			wg.Done()
			return
		case d := <-c:
			fmt.Println(d)
		}
	}
}

func main() {
	//Создаем контекст передающий сигнал на завершения процесса в то время когда пользователь нажмет Ctrl+C, чтобы прервать выполнение программы
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	wg := sync.WaitGroup{}
	var q int
	//Просим пользователя ввести количество воркеров
	_, err := fmt.Scan(&q)
	if err != nil {
		return
	}
	//Создаем канал
	ch := make(chan int)
	//Запускаем каждого воркера в горутине и передаем контекст чтобы завершить действие всех воркеров вместе с программой
	for i := 0; i < q; i++ {
		wg.Add(1)
		go worker11(ctx, ch, &wg)
	}
	func() {
		//Записываем в канал информацию в бесконечном цикле
		for {
			select {
			//Если пришел сигнал завершаем запись
			case <-ctx.Done():
				return
			default:
				ch <- 1
			}
		}
	}()
	//Ожидаем завершения всех горутин
	wg.Wait()
}