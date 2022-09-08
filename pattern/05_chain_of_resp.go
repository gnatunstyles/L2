package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type floor int

const (
	null floor = iota
	first
	second
	roof
)

type currentFloor struct {
	button bool
}

type handler interface {
	check(*currentFloor)
	setNextFloor(handler)
}

type nilFloor struct {
	next handler
}

func (n *nilFloor) check(c *currentFloor) {
	if !c.button {
		fmt.Println("It's not your floor yet. Wait")
		n.next.setNextFloor(n)
	} else {
		fmt.Println("It's your floor. Good luck!")
		return
	}
}

func (n *nilFloor) setNextFloor(next handler) {
	n.next = next
}

type firstFloor struct {
	next handler
}

func (n *firstFloor) check(c *currentFloor) {
	if !c.button {
		fmt.Println("It's not your floor yet. Wait")
		n.next.setNextFloor(n)
	} else {
		fmt.Println("It's your floor. Good luck!")
		return
	}
}

func (n *firstFloor) setNextFloor(next handler) {
	n.next = next
}

type secondFloor struct {
	next handler
}

func (n *secondFloor) check(c *currentFloor) {
	if !c.button {
		fmt.Println("It's not your floor yet. Wait")
		n.next.setNextFloor(n)
	} else {
		fmt.Println("It's your floor. Good luck!")
		return
	}
}

func (n *secondFloor) setNextFloor(next handler) {
	n.next = next
}

type Roof struct{}

func (n *Roof) check(c *currentFloor) {
	fmt.Println("You are on the roof! Use stairs to climb down^)")
	return
}
