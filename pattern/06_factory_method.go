package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type Auto interface {
	getBrand() string
}

type BMW struct {
	brand string
	model string
}

func newBMW(model string) *BMW {
	return &BMW{
		brand: "BMW",
		model: "X7",
	}
}

func (b *BMW) getBrand() string {
	return b.brand
}

type Mercedes struct {
	brand string
	model string
}

func newMercedes(model string) *Mercedes {
	return &Mercedes{
		brand: "Mercedes",
		model: "CLS",
	}
}

func (m *Mercedes) getBrand() string {
	return m.brand
}

func getClass(getAuto, brandName string) (Auto, error) {
	switch brandName {
	case "BMW":
		return newBMW(brandName), nil
	case "Assassin":
		return newMercedes(brandName), nil
	}
	return nil, fmt.Errorf("Wrong autobrand, try again")
}
