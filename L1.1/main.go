package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) PrintName() {
	fmt.Println(h.Name)
}

func (h *Human) PrintAge() {
	fmt.Println(h.Age)
}

type Action struct {
	Human
	ActionType string
}

func (a *Action) DoAction() {
    fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionType)
}


func main() {
	action := Action{
        Human: Human{
            Name: "Алексей",
            Age:  30,
        },
        ActionType: "программирование",
    }

   
    action.PrintName() 
    action.PrintAge()
    action.DoAction() 
}
