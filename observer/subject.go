// file: subject.go
// author: TaoZer
// Date: 2022/1/24
// Description: observer pattern, subject implement

package observer

import "errors"

type Subject interface {
	Register(o Observer) error
	Remove(o Observer)
	NotifyAll(msg string)
}

type Paper struct {
	follower map[Observer]bool
}

func (p *Paper) Register(o Observer) (err error) {
	if _, ok := p.follower[o]; ok {
		err = errors.New("observer already exist")
		return err
	}
	p.follower[o] = true
	return
}

func (p *Paper) Remove(o Observer) {
	delete(p.follower, o)
}

func (p *Paper) NotifyAll(msg string) {
	for k, _ := range p.follower {
		k.update(msg)
	}
}

func NewPaper() *Paper {
	return &Paper{
		follower: make(map[Observer]bool),
	}
}