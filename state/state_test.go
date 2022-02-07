package state

import "testing"

func TestState(t *testing.T) {
	sm := NewStateMachine()
	sm.TurnB()
	sm.TurnC()
	sm.TurnA()
	sm.TurnA()
}
