package main

import (
	"fmt"
	"reflect"
)

func main() {
	isType(0)
	isType("123")
	isType(false)
	isType(make(chan struct{}))

}

func isType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(i).Kind() == reflect.Chan {
			fmt.Println("chan")
		}
	}
}
