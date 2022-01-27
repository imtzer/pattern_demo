package command

import "testing"

func TestCommand(t *testing.T) {
	c := NewZipCommand()
	ivk := NewInvoker(c)
	ivk.Command.Do()
	ivk.Command.Undo()
}