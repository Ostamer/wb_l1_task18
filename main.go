package main

import (
	"fmt"
	"sync"
)

// Структура счетчика
type Counter struct {
	counterMutex sync.Mutex
	count        int
}

// Метод для инкрементации счетчика
func (counter *Counter) Increment() {
	counter.counterMutex.Lock()         // Блокируем доступ к счетчику
	defer counter.counterMutex.Unlock() // Разблокируем мьютекс после завершения работы

	counter.count++ // Инкрементируем счетчик
}

// Метод для получения текущего значения счетчика
func (counter *Counter) Get() int {
	counter.counterMutex.Lock()         // Блокируем доступ к счетчику
	defer counter.counterMutex.Unlock() // Разблокируем мьютекс после завершения работы

	return counter.count // Возвращаем значение счетчика
}

func main() {
	var counterGroup sync.WaitGroup
	counter := &Counter{} // Создаем счетчик

	// Запускаем 100 горутин для инкрементирования счетчика
	for i := 0; i < 100; i++ {
		counterGroup.Add(1)
		go func() {
			defer counterGroup.Done()
			counter.Increment() // Увеличиваем счетчик
		}()
	}

	counterGroup.Wait()

	fmt.Println("Итоговое значение счетчика:", counter.Get())
}
