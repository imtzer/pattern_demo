// file: decorate.go
// author: TaoZer
// Date: 2022/1/26
// Description: decoration pattern

package decorate

type Beverage interface {
	Cost() float64
	Description() string
}

type MilkTea struct {

}

func NewMilkTea() MilkTea {
	return MilkTea{}
}

func (mt MilkTea) Cost() float64 {
	return 6
}

func (mt MilkTea) Description() string {
	return "Milk Tea"
}

type MilkGreen struct {
}

func NewMilkGreen() MilkGreen {
	return MilkGreen{}
}

func (mg MilkGreen) Cost() float64 {
	return 5
}

func (mg MilkGreen) Description() string {
	return "Milk Green"
}

type Condiment struct {
	wrappedObj Beverage
}

func (c Condiment)Cost() float64 {
	return c.wrappedObj.Cost()
}

func (c Condiment)Description() string {
	return c.wrappedObj.Description()
}

type Bubble struct {
	Condiment
}

func AddBubble(b Beverage) Bubble {
	return Bubble{
		Condiment{
			wrappedObj: b,
		},
	}
}

func (b Bubble)Cost() float64 {
	return 1 + b.wrappedObj.Cost()
}

func (b Bubble)Description() string {
	return "Bubble "+b.wrappedObj.Description()
}

type MilkCover struct {
	Condiment
}

func AddMilkCover(b Beverage) MilkCover {
	return MilkCover{
		Condiment{
			wrappedObj: b,
		},
	}
}

func (mc MilkCover)Cost() float64 {
	return 1.5 + mc.wrappedObj.Cost()
}

func (mc MilkCover)Description() string {
	return "Milk Cover "+mc.wrappedObj.Description()
}

type Coconut struct {
	Condiment
}

func AddCoconut(b Beverage) Coconut {
	return Coconut{
		Condiment{
			wrappedObj: b,
		},
	}
}

func (c Coconut)Cost() float64 {
	return 1.5 + c.wrappedObj.Cost()
}

func (c Coconut)Description() string {
	return "Coconut "+c.wrappedObj.Description()
}

type GrassJelly struct {
	Condiment
}

func AddGrassJelly(b Beverage) MilkCover {
	return MilkCover{
		Condiment{
			wrappedObj: b,
		},
	}
}

func (gj GrassJelly)Cost() float64 {
	return 1.5 + gj.wrappedObj.Cost()
}

func (gj GrassJelly)Description() string {
	return "Bubble "+gj.wrappedObj.Description()
}