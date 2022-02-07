// file: template.go
// author: TaoZer
// Date: 2022/2/7
// Description: state pattern

package state

import "fmt"

type StateI interface {
	Turn2A(sm *StateMachine)
	Turn2B(sm *StateMachine)
	Turn2C(sm *StateMachine)
}

type StateMachine struct {
	present StateI
}

func (sm *StateMachine) TurnA()  {
	sm.present.Turn2A(sm)
}

func (sm *StateMachine) TurnB()  {
	sm.present.Turn2B(sm)
}

func (sm *StateMachine) TurnC()  {
	sm.present.Turn2C(sm)
}

func (sm *StateMachine) setState(s StateI)  {
	sm.present = s
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{}
	sm.present = GetStateA()
	return sm
}

type StateA struct {}

func (s StateA) Turn2A(sm *StateMachine) {
	fmt.Println("can not turn to self")
}

func (s StateA) Turn2B(sm *StateMachine) {
	fmt.Println("state A turn to state B")
	sm.setState(GetStateB())
}

func (s StateA) Turn2C(sm *StateMachine) {
	fmt.Println("state A can not turn to state C")
}

var sa StateA = StateA{}
func GetStateA() StateA {
	return sa
}

type StateB struct {}

func (s StateB) Turn2A(sm *StateMachine) {
	fmt.Println("state B can not turn to state A")
}

func (s StateB) Turn2B(sm *StateMachine) {
	fmt.Println("can not turn to self")
}

func (s StateB) Turn2C(sm *StateMachine) {
	fmt.Println("state B turn to state C")
	sm.setState(GetStateC())
}

var sb = StateB{}
func GetStateB() StateB {
	return sb
}

type StateC struct {}

func (s StateC) Turn2A(sm *StateMachine) {
	fmt.Println("state C turn to state A")
	sm.setState(GetStateA())
}

func (s StateC) Turn2B(sm *StateMachine) {
	fmt.Println("state C can not turn to state B")
}

func (s StateC) Turn2C(sm *StateMachine) {
	fmt.Println("can not turn to self")
}

var sc = StateC{}
func GetStateC() StateC {
	return sc
}