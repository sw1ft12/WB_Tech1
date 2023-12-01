package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	ch := make(chan any)
	done := make(chan int)
	termChan := make(chan os.Signal)                         // Создаём канал, который будет принимать сигналы
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM) // Связываем канал с сигналами SIGINT и SIGTERM
	var n int
	fmt.Scan(&n)
	var wg sync.WaitGroup
	wg.Add(n + 1)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case msg := <-ch:
					fmt.Printf("%v ", msg)
				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case ch <- "123":
			}
		}
	}()
	<-termChan // При поступлении сигнала закрываем канал, который означает завершение горутины
	close(done)
	wg.Wait()
}
