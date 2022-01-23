package main

import (
	"pattern/factory"
)

func main() {
	ss := factory.NewSimpleStore()
	pizza := ss.Order("A")
	pizza.Show()

	fm := factory.NewFactoryMethod()
	pizzaB := fm.Order("B")
	pizzaB.Show()
}