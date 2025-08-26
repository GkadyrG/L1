package main

import (
	"log"
	"os"
)

var stdLogger = log.New(os.Stdout, "[STD]", log.LstdFlags)

type Logger interface {
	Log(message string)
}

type StdLoggerAdapter struct {
	logger *log.Logger
}

func (a *StdLoggerAdapter) Log(message string) {
	a.logger.Println(message)
}

func main() {
	var l Logger
	l = &StdLoggerAdapter{logger: stdLogger}

	l.Log("Адаптер работает")
}
