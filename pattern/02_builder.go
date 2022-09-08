package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Laptop struct {
	motherBoard string
	cpu         string
	gpu         string
	battery     string
	ram         string
	fan         string
	display     string
}
type LaptopBuilder interface {
	setMotherBoard()
	setCPU()
	setGPU()
	setBattery()
	setRAM()
	setFAN()
	setDisplay()
	getLaptop() *Laptop
}
type asusBuilder struct {
	motherBoard string
	cpu         string
	gpu         string
	battery     string
	ram         string
	fan         string
	display     string
}
type appleBuilder struct {
	motherBoard string
	cpu         string
	gpu         string
	battery     string
	ram         string
	fan         string
	display     string
}

func getBuilder(brand string) LaptopBuilder {
	if brand == "apple" {
		return newAppleBuilder()
	}
	if brand == "asus" {
		return newAsusBuilder()
	}
	return nil
}

func newAsusBuilder() *asusBuilder {
	return &asusBuilder{}
}
func newAppleBuilder() *appleBuilder {
	return &appleBuilder{}
}

func (b *asusBuilder) setMotherBoard() {
	b.motherBoard = "asus motherboard"
}
func (b *asusBuilder) setCPU() {
	b.cpu = "amd"
}
func (b *asusBuilder) setGPU() {
	b.gpu = "radeon"
}
func (b *asusBuilder) setBattery() {
	b.battery = "228mAh"
}
func (b *asusBuilder) setRAM() {
	b.ram = "8GB"
}
func (b *asusBuilder) setFAN() {
	b.fan = "deepcool mini"
}
func (b *asusBuilder) setDisplay() {
	b.display = "some cheap chinese display"
}
func (b *asusBuilder) getLaptop() *Laptop {
	return &Laptop{
		motherBoard: b.motherBoard,
		cpu:         b.cpu,
		gpu:         b.gpu,
		battery:     b.battery,
		ram:         b.ram,
		fan:         b.fan,
		display:     b.display,
	}
}

func (b *appleBuilder) setMotherBoard() {
	b.motherBoard = "apple motherboard"
}
func (b *appleBuilder) setCPU() {
	b.cpu = "m1"
}
func (b *appleBuilder) setGPU() {
	b.gpu = "nvidia"
}
func (b *appleBuilder) setBattery() {
	b.battery = "100mAh"
}
func (b *appleBuilder) setRAM() {
	b.ram = "4GB"
}
func (b *appleBuilder) setFAN() {
	b.fan = "apple fan"
}
func (b *appleBuilder) setDisplay() {
	b.display = "retina"
}
func (b *appleBuilder) getLaptop() *Laptop {
	return &Laptop{
		motherBoard: b.motherBoard,
		cpu:         b.cpu,
		gpu:         b.gpu,
		battery:     b.battery,
		ram:         b.ram,
		fan:         b.fan,
		display:     b.display,
	}
}

type director struct {
	laptopBuilder LaptopBuilder
}

func newDirector(c LaptopBuilder) *director {
	return &director{
		laptopBuilder: c,
	}
}

func (d *director) setBuilder(b LaptopBuilder) {
	d.laptopBuilder = b
}

func (d *director) buildComputer() Laptop {
	d.laptopBuilder.setMotherBoard()
	d.laptopBuilder.setBattery()
	d.laptopBuilder.setCPU()
	d.laptopBuilder.setGPU()
	d.laptopBuilder.setDisplay()
	d.laptopBuilder.setRAM()
	d.laptopBuilder.setFAN()
	return *d.laptopBuilder.getLaptop()
}
