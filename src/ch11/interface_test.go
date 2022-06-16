package test

import "testing"

type Bus struct {
	Speed int
	Price int
	Color string
}

func (b *Bus) GetSpeed() int {
	return b.Speed
}

func (b *Bus) GetPrice() int {
	return b.Price
}

func (b *Bus) GetColor() string {
	return b.Color
}

type Car struct {
	Speed int
	Price int
	Seats int
}

func (c *Car) GetSpeed() int {
	return c.Speed
}

func (c *Car) GetPrice() int {
	return c.Price
}

func (c *Car) GetSeats() int {
	return c.Seats
}

type CommomCar interface {
	GetColor() int
	GetPrice() int
}

func TestInterface(t *testing.T) {

	bus := new(Bus)
	bus.Speed = 100
	bus.Price = 20

	c1 := Car{Speed: 10, Price: 20}

	c1.GetPrice()

}
