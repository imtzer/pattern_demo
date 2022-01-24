// file: observer.go
// author: TaoZer
// Date: 2022/1/24
// Description: observer pattern, observer implement

package observer

import "fmt"

type Observer interface {
	update(msg string)
}

type User struct {
	Name string
	msg string
}

func (u *User) update(msg string) {
	u.msg = msg
}

func (u *User)Show() {
	fmt.Println(fmt.Sprintf("%s, msg: %s", u.Name, u.msg))
}