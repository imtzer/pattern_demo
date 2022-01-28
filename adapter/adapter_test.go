package adapter

import "testing"

func TestAdapter(t *testing.T) {
	ae := NewAdaptee()
	adp := NewAdapter(ae)
	c := NewClient(adp)
	c.Request()
}
