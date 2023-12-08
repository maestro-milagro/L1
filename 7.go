package main

import (
	"log"
	"sync"
)

//Реализовать конкурентную запись данных в map.

//Чтобы избежать возникновения race condition из за одновременного доступа с ресурсу в виде map, воспользуемся мьютексом,
//а именно RWMutex т. к. он лучше подходит для частого обращения к ресурсу

var numbers = make(map[int]int, 100)

func main() {
	var m sync.RWMutex
	func() {
		wg := sync.WaitGroup{}
		// Запись.
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				//Блокируем доступ к ресурсу на время работы с ним горутины
				m.Lock()
				numbers[i] = i
				//Освобождаем ресурс
				m.Unlock()
				wg.Done()
			}(i)
		}
		// Чтение.
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				//Блокируем доступ к ресурсу на время работы с ним горутины
				m.Lock()
				log.Print(numbers[i])
				//Освобождаем ресурс
				m.Unlock()
				wg.Done()
			}(i)
		}
		wg.Wait()
	}()
}
