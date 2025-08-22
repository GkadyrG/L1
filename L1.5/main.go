package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int, interval time.Duration, timeout time.Duration) {
	defer close(out)
	i := 0
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	timer := time.After(timeout)

	for {
		select {
		case <-ticker.C:
			out <- i
			i++
		case <-timer:
			return
		}
	}
}

func consumer(in <-chan int, timeout time.Duration) {
	timer := time.After(timeout)
	for {
		select {
		case val, ok := <-in:
			if !ok {
				fmt.Println("канал закрыт")
				return
			}
			fmt.Println("Получено:", val)
		case <-timer:
			fmt.Println("время вышло")
			return
		}
	}

}

func main() {
	ch := make(chan int)
	go producer(ch, 500*time.Millisecond, 3*time.Second)
	consumer(ch, 4*time.Second)
}
