package main

import (
	"fmt"
	"pattern/observer"
)

func main() {
	u1 := &observer.User{
		Name: "u1",
	}
	u2 := &observer.User{
		Name: "u2",
	}
	paperPublisher := observer.NewPaper()

	_ = paperPublisher.Register(u1)
	_ = paperPublisher.Register(u2)
	fmt.Println("users msg:")
	fmt.Print("  ")
	u1.Show()
	fmt.Print("  ")
	u2.Show()
	fmt.Println("publish new msg...")
	paperPublisher.NotifyAll("new paper")
	fmt.Println("users msg:")
	fmt.Print("  ")
	u1.Show()
	fmt.Print("  ")
	u2.Show()
}