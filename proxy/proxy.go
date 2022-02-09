// file: proxy.go
// author: TaoZer
// Date: 2022/2/9
// Description: static proxy pattern

package proxy

import "fmt"

type ServerI interface {
	HandlerRequest()
}

type Server struct {}

func (s Server) HandlerRequest() {
	fmt.Println("server handler request")
}

func NewServer() Server {
	return Server{}
}

type Proxy struct {
	s Server
}

func NewProxy(s Server) Proxy {
	return Proxy{s: s}
}

func (p Proxy) HandlerRequest() {
	fmt.Println("proxy receive request, send to server")
	p.s.HandlerRequest()
}

