// file: command.go
// author: TaoZer
// Date: 2022/1/27
// Description: command pattern

package command

import "fmt"

type Command interface {
	Do()
	Undo()
}

type ZipCommand struct {
	ZipReceiver
}

func NewZipCommand() ZipCommand {
	return ZipCommand{
		ZipReceiver{},
	}
}

func (z ZipCommand) Do() {
	z.ZipReceiver.compress()
}

func (z ZipCommand) Undo() {
	z.ZipReceiver.decompress()
}

type ZipReceiver struct {}

func (z ZipReceiver) compress() {
	fmt.Println("Zip compress")
}

func (z ZipReceiver) decompress()  {
	fmt.Println("Zip decompress")
}

// Invoker 定义某些命令触发的方式（这里没有定义）
type Invoker struct {
	Command
}

func (ivk *Invoker)SetCommand(c Command)  {
	ivk.Command = c
}

func NewInvoker(c Command) Invoker {
	return Invoker{
		Command: c,
	}
}