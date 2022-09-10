package deferstate

import (
	"fmt"
	"time"
)

func A() {
	fmt.Println("test")
}
func C() {
	fmt.Println("test C")
}

func DeferState() {
	fmt.Println("Runnn")
	time.Sleep(2 * time.Second)
	defer A()
	defer C()
}
