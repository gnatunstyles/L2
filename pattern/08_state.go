package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	pressButton()
}

type faucet struct {
	currState State
}

func newFaucet() *faucet {
	faucet := faucet{}

	faucet.changeState(&TurnedOff{faucet: &faucet})

	return &faucet
}

func (c *faucet) pressButton() {
	c.currState.pressButton()
}

func (c *faucet) changeState(state State) {
	c.currState = state
}

type TurnedOff struct {
	faucet *faucet
}

func (t *TurnedOff) pressButton() {
	fmt.Println("Water goes")
	t.faucet.changeState(&TurnedOn{faucet: t.faucet})
}

type TurnedOn struct {
	faucet *faucet
}

func (t *TurnedOn) pressButton() {
	fmt.Println("Water stops")
	t.faucet.changeState(&TurnedOff{faucet: t.faucet})
}
