package main

import "sync"

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	m := make(map[int]int)
	var mutex sync.Mutex
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}
	}()

	for i := 5; i < 20; i++ {
		mutex.Lock()
		m[i] = i + 5
		mutex.Unlock()
	}
}
