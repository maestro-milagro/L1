package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

// Горутины могут попытаться одновременно изменить переменную из за чего часть вычислений может пропасть.
// Чтобы избежать этого мы поставим тип переменной int64 и будем увеличивать её значение с помощью atomic.AddInt64
// Таким образом обнспечив попеременный доступ к объекту
func main() {
	var count int64
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}