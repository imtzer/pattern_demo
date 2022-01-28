// file: adapter.go
// author: TaoZer
// Date: 2022/1/28
// Description: adapter pattern

package adapter

import "fmt"

type Target interface {
	Request()
}

type Client struct {
	Target
}

func NewClient(t Target) Client {
	return Client{
		t,
	}
}

type Adaptee struct {}

func NewAdaptee() Adaptee {
	return Adaptee{}
}

func (a Adaptee) newRequest()  {
	fmt.Println("Adaptee new request")
}

type Adapter struct {
	Adaptee
}

func NewAdapter(aee Adaptee) Adapter {
	return Adapter{
		aee,
	}
}

func (a Adapter) Request() {
	a.Adaptee.newRequest()
}

