package main

import (
	"context"
	"sync"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	var wg sync.WaitGroup
	{
		done := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				}
			}
		}()
		close(done)
		wg.Wait()
	}

	{
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-done:
					return
				}
			}
		}()
		done <- struct{}{}
	}

	{
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				}
			}
		}()
		cancel()
		wg.Wait()
	}
}
