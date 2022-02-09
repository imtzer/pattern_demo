// file: proxy.go
// author: TaoZer
// Date: 2022/2/9
// Description: dynamic proxy pattern
// Todo: 1.两种接口，接口大于一种方法，一个代理对所有类的方法调用进行计时
// Todo: 2.为计时代理添加其他用户设定的职责

package proxy

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type ServerI1 interface {
	Register1()
	Login1()
}

type ServerI2 interface {
	Register2()
	Login2()
}

type Server1 struct {

}

func NewServer1() Server1 {
	return Server1{}
}

func (s Server1) Register1() {
	fmt.Println("register Server1")
}

func (s Server1) Login1() {
	fmt.Println("login Server1")
}

type Server2 struct {

}

func NewServer2() Server2 {
	return Server2{}
}

func (s Server2) Register2() {
	fmt.Println("register Server2")
}

func (s Server2) Login2() {
	fmt.Println("login Server2")
}

// DProxy 动态代理，接收任何类型的对象
type DProxy struct {
	ih InvocationHandlerI
	subject interface{}
	methods map[string]reflect.Value
}

func (d DProxy) Call(name string, in []reflect.Value) {
	fmt.Println("dynamic proxy call")
	d.ih.invoke(d.methods[name], in)
}

func NewDProxy(subject interface{}) DProxy {
	v := reflect.ValueOf(subject)
	t := reflect.TypeOf(subject)
	dproxy := DProxy{
		subject: subject,
		ih: NewTimeInvocationHandler(subject),
		methods: make(map[string]reflect.Value),
	}
	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		mt := t.Method(i)
		fmt.Println("method name: "+mt.Name)
		dproxy.methods[mt.Name] = m
	}
	return dproxy
}


// InvocationHandlerI
type InvocationHandlerI interface {
	invoke(method reflect.Value, in []reflect.Value)
}

// TimeInvocationHandler 和所有函数运行时间相关的逻辑
type TimeInvocationHandler struct {
	i interface{}
}

func NewTimeInvocationHandler(i interface{}) TimeInvocationHandler {
	return TimeInvocationHandler{i: i}
}

func (tih TimeInvocationHandler) invoke(method reflect.Value, in []reflect.Value) {
	t := time.Now().Unix()
	method.Call(in)
	fmt.Println("used time: "+strconv.Itoa(int(time.Now().Unix()-t))+"ms")
}

