// file: strategy.go
// author: TaoZer
// Date: 2022/1/21
// Description: strategy pattern

package strategy

import "fmt"

type swimming interface {
	Swim()
}

type quacking interface {
	Quack()
}

type canSwim struct {}

func (c *canSwim) Swim() {
	fmt.Println("is swimming")
}

type notSwim struct {}

func (n *notSwim) Swim() {
	fmt.Println("can not swim")
}

type canQuack struct {}

func (c *canQuack) Quack() {
	fmt.Println("is quacking")
}

type notQuack struct {}

func (n *notQuack) Quack() {
	fmt.Println("can not quack")
}

type Duck struct {
	sm swimming
	qk quacking
	name string
}

func (d *Duck) Quack() {
	fmt.Print(d.name + " ")
	d.qk.Quack()
}

func (d *Duck) Swim() {
	fmt.Print(d.name + " ")
	d.sm.Swim()
}

// 由于go没有java的构造器，所以省略了继承抽象类重写构造方法并在构造方法中定义策略的步骤
// 取而代之的是利用New方法去定义

func NewRedDuck() *Duck {
	return &Duck{
		sm: &canSwim{},
		qk: &canQuack{},
		name: "red duck",
	}
}

func NewFakeDuck() *Duck {
	return &Duck{
		sm: &notSwim{},
		qk: &canQuack{},
		name: "fake duck",
	}
}
