package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func startWorkers(n int, jobsCh <-chan int) {
	for i := 0; i < n; i++ {
		go func(id int) {
			for job := range jobsCh {
				fmt.Printf("Worker %d: %d\n", id, job)
			}
		}(i)
	}
}

func generateJobs(jobsCh chan<- int) {
	i := 0
	for {
		jobsCh <- i
		i++
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
	startWorkers(workerCount, jobsCh)
	generateJobs(jobsCh)
}
