package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	p := NewProxy(NewServer())
	p.HandlerRequest()
}
