// file: template.go
// author: TaoZer
// Date: 2022/1/29
// Description: template pattern

package template

import "fmt"

type FuncBrew func()
type FuncCondiment func()

type Beverage struct {
	brew         FuncBrew
	addCondiment FuncCondiment
}

func NewBeverage(brew FuncBrew, addCondiment FuncCondiment) Beverage {
	return Beverage{
		brew: brew,
		addCondiment: addCondiment,
	}
}

func (b Beverage) Prepare()  {
	b.boilWater()
	b._brew()
	b.pourIntoCup()
	b._addCondiment()
}

func (b Beverage) boilWater()  {
	fmt.Println("boiling water!")
}

func (b Beverage) _brew()  {
	if b.brew == nil {
		panic("invalid brew")
	} else {
		b.brew()
	}
}

func (b Beverage) pourIntoCup()  {
	fmt.Println("pour into cup!")
}

func (b Beverage) _addCondiment()  {
	if b.addCondiment == nil {
		panic("invalid addCondiment")
	} else {
		b.addCondiment()
	}
}