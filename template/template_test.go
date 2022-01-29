package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T)  {

	var brewTea FuncBrew = func() {
		fmt.Println("brewing tea!")
	}
	var addLemon FuncCondiment = func() {
		fmt.Println("add lemon")
	}
	tea := NewBeverage(brewTea, addLemon)
	tea.Prepare()

	var brewCoffee FuncBrew = func() {
		fmt.Println("brewing coffee!")
	}
	var addSurge FuncCondiment = func() {
		fmt.Println("add surge")
	}
	coffee := NewBeverage(brewCoffee, addSurge)
	coffee.Prepare()
}