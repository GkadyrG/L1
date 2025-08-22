package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// 1️ Выход по условию
func workerCondition(stopAfter int) {
	i := 0
	for {
		if i >= stopAfter {
			fmt.Println("[Condition] Горутина остановлена по условию")
			return
		}
		fmt.Println("[Condition] Работаем:", i)
		i++
		time.Sleep(300 * time.Millisecond)
	}
}

// 2️ Через канал уведомления
func workerChannel(stopChan <-chan struct{}) {

	for _ = range stopChan {
		fmt.Println("[Channel] Работаем")
		time.Sleep(400 * time.Millisecond)
	}
	fmt.Println("[Channel] Горутина остановлена через канал")

}

// 3️ Через контекст
func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Context] Горутина остановлена через контекст:", ctx.Err())
			return
		default:
			fmt.Println("[Context] Работаем")
			time.Sleep(350 * time.Millisecond)
		}
	}
}

// 4️ Использование runtime.Goexit()
func workerGoexit() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("[Goexit] Немедленно выходим из горутины на i=2")
			runtime.Goexit() // прерываем горутину прямо здесь
		}
		fmt.Println("[Goexit] Работаем:", i)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("[Goexit] Этот код никогда не выполнится")
}

// 5️ Через таймер / time.After
func workerTimer(timeout time.Duration) {
	timer := time.After(timeout)
	for {
		select {
		case <-timer:
			fmt.Println("[Timer] Горутина остановлена через таймер")
			return
		default:
			fmt.Println("[Timer] Работаем")
			time.Sleep(450 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("=== Демонстрация остановки горутин ===")

	// 1️ Условие
	workerCondition(5)

	// 2️ Канал
	stopChan := make(chan struct{})
	go workerChannel(stopChan)
	// Остановка канала после 2 секунд
	time.Sleep(2 * time.Second)
	close(stopChan)

	// 3️ Контекст
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	workerContext(ctx)

	// 4️ Goexit
	go workerGoexit()
	time.Sleep(1 * time.Second)
	// 5️ Таймер
	workerTimer(3 * time.Second)

	// Ждём, чтобы все горутины завершились
	time.Sleep(6 * time.Second)

	fmt.Println("=== Все горутины завершены ===")
}
