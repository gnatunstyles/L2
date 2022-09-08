package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	execute()
}

type client interface {
	eat(food string)
	pay(money int)
}

type payCommand struct {
	client client
	money  int
}

func newClientCommand(client client, money int) *payCommand {
	return &payCommand{
		client: client,
		money:  money,
	}
}
func (c *payCommand) execute() {
	c.client.pay(c.money)
}

type eatCommand struct {
	client client
	food   string
}

func newEatCommand(client client, food string) *eatCommand {
	return &eatCommand{
		client: client,
		food:   food,
	}
}

func (e *eatCommand) execute() {
	e.client.eat(e.food)
}

type Service struct {
	cmd Command
}

func newService(cmd Command) *Service {
	return &Service{
		cmd: cmd,
	}
}

func (s *Service) setCommand(cmd Command) {
	s.cmd = cmd
}

func (s *Service) start() {
	s.cmd.execute()
}

type Client struct {
	name  string
	money int
}

func (c *Client) eat(food string) {
	fmt.Printf("%s is eating %s. Bon appetit!", c.name, food)
}

func (c *Client) pay(m int) {
	c.money -= m
}
