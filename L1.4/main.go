package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func readWorkerCount() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число воркеров: ")
	if !scanner.Scan() {
		return 0, fmt.Errorf("ошибка ввода")
	}
	return strconv.Atoi(scanner.Text())
}

func startWorkers(ctx context.Context, n int, jobsCh <-chan int, wg *sync.WaitGroup) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for job := range jobsCh {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Printf("Worker %d: %d\n", id, job)
				}

			}

		}(i)

	}
}

func generateJobs(ctx context.Context, jobsCh chan<- int) {
	defer close(jobsCh)
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case jobsCh <- i:
			i++
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	workerCount, err := readWorkerCount()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	jobsCh := make(chan int)
	wg := &sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	startWorkers(ctx, workerCount, jobsCh, wg)
	generateJobs(ctx, jobsCh)

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
