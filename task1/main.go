package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	height int
	name   string
	age    int
}

func NewHuman(height int, name string, age int) *Human {
	return &Human{
		height: height,
		name:   name,
		age:    age,
	}
}

func (h *Human) Walk() {
	fmt.Println(h.name + ": is walking")
}

func (h *Human) ShowParametrs() string {
	return fmt.Sprint(h.name, ": is ", h.age, " y.o. and its height is ", h.height, "\n")
}

type Action struct {
	*Human
	actionType string
}

func NewAction(human *Human, actionType string) *Action {
	return &Action{
		Human:      human,
		actionType: actionType,
	}
}

func (a *Action) SayHello() {
	fmt.Printf("%s: %s", a.actionType, a.ShowParametrs())
}

func (a *Action) DoSomeAction() {
	a.Walk()
}

func main() {
	h := NewHuman(180, "John", 21)
	h.Walk()
	fmt.Println(h.ShowParametrs())
	h2 := NewHuman(190, "Nicolas", 44)
	a := NewAction(h2, "Walking")
	a.SayHello()
	a.DoSomeAction()
}
