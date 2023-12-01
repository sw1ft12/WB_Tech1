package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {

	var n int
	fmt.Scan(&n)

	ch := make(chan any)

	go func() {
		for {
			ch <- "one"
		}
	}()

	go func() {
		for {
			fmt.Printf("%v\n", <-ch)
		}
	}()

	time.Sleep(time.Duration(n) * time.Second) // Программа завершается через n секунд
}
