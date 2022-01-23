// file: factory.go
// author: TaoZer
// Date: 2022/1/23
// Description: factory pattern

package factory

import (
	"fmt"
)

// SimpleFactoryI 简单工厂接口
type SimpleFactoryI interface {
	Create(tp string) Product
}

// SimpleFactory 简单工厂
type SimpleFactory struct {}

func (s *SimpleFactory) Create(tp string) Product {
	if tp == "A" {
		return newPizzaA()
	} else if tp == "B" {
		return newPizzaB()
	} else if tp == "C" {
		return newPizzaC()
	} else {
		_ = fmt.Errorf("invalid type")
		return nil
	}
}

// SimpleStore 使用简单工厂模式的对象
type SimpleStore struct {
	factory SimpleFactoryI
}

func  (s SimpleStore) Order(tp string) Product {
	p := s.factory.Create(tp)
	p.prepare()
	return p
}

func NewSimpleStore() SimpleStore {
	return SimpleStore{
		factory: &SimpleFactory{},
	}
}


// AbstractFactoryMethod 抽象工厂方法
type AbstractFactoryMethod struct {}

func (afm AbstractFactoryMethod) Order(tp string) Product{
	p := afm.create(tp)
	p.prepare()
	return p
}

func (afm AbstractFactoryMethod) create(tp string) Product {
	panic("implement me!")
}

// FactoryMethod 工厂方法
type FactoryMethod struct {
	AbstractFactoryMethod
}

func NewFactoryMethod() FactoryMethod {
	return FactoryMethod{}
}

func (fm FactoryMethod)Order(tp string) Product {
	p := fm.create(tp)
	p.prepare()
	return p
}

func (fm FactoryMethod) create(tp string) Product{
	if tp == "A" {
		return newPizzaA()
	} else if tp == "B" {
		return newPizzaB()
	} else if tp == "C" {
		return newPizzaC()
	} else {
		_ = fmt.Errorf("invalid type")
		return nil
	}
}

// IngredientFactory 抽象工厂接口
type IngredientFactoryI interface {
	createSalt() string
	createCheese() string
	createMeat() string
}


// IngredientFactory 抽象工厂
type IngredientFactoryA struct {}

func (i IngredientFactoryA) createSalt() string {
	return "salt A"
}

func (i IngredientFactoryA) createCheese() string {
	return "cheese A"
}

func (i IngredientFactoryA) createMeat() string {
	return "meat A"
}

type IngredientFactoryB struct {}

func (i IngredientFactoryB) createSalt() string {
	return "salt B"
}

func (i IngredientFactoryB) createCheese() string {
	return "cheese B"
}

func (i IngredientFactoryB) createMeat() string {
	return "meat B"
}

