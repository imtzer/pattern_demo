package proxy

import (
	"reflect"
	"testing"
)

func TestDynamic(t *testing.T)  {
	s1 := NewServer1()
	//s2 := NewServer2()
	dproxy := NewDProxy(s1)
	dproxy.Call("Register1", []reflect.Value{})
}