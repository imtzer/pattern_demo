// file: product.go
// author: TaoZer
// Date: 2022/1/23
// Description: factory pattern product

package factory

import "fmt"

// Product 产品
type Product interface {
	prepare()
	Show()
}

// Pizza 抽象类
type Pizza struct {
	Name string
	salt string
	cheese string
	meat string
	factory IngredientFactoryI
}

func (p *Pizza) prepare() {
	panic("implement me")
}

func (p *Pizza) Show() {
	fmt.Println(p.Name + " ingredient: "+fmt.Sprintf("%s, %s, %s", p.salt, p.cheese, p.meat))
}

// 三种具体的pizza
type pizzaA struct {
	Pizza
}

func (p *pizzaA)prepare()  {
	p.meat = p.factory.createMeat()
	p.salt = p.factory.createSalt()
	p.cheese = p.factory.createCheese()
}

func newPizzaA() *pizzaA {
	return &pizzaA{
		Pizza{
			Name: "pizza A",
			factory: IngredientFactoryA{},
		},
	}
}

type pizzaB struct {
	Pizza
}

func (p *pizzaB)prepare()  {
	p.meat = p.factory.createMeat()
	p.salt = p.factory.createSalt()
	p.cheese = p.factory.createCheese()
}

func newPizzaB() *pizzaB {
	return &pizzaB{
		Pizza{
			Name: "pizza B",
			factory: IngredientFactoryB{},
		},
	}
}

type pizzaC struct {
	Pizza
}

func (p *pizzaC)prepare()  {

}

func newPizzaC() *pizzaC {
	return &pizzaC{

	}
}