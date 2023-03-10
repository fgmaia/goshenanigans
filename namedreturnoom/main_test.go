package main_test

import (
	"fmt"
	"testing"
)

func TestMainWillOutOfMemory(t *testing.T) {
	fmt.Println("start")
	print()()
}

func print() (f func()) { // <- will be called recursivle
	total := 0
	f = func() { // <- will be ignored by f from return
		fmt.Println("total")
		total += 1
	}
	return func() {
		fmt.Println("ini!")
		f() // calling recursivle, will ignore f in line 15
		fmt.Println("done!")
	}
}

func TestMainFixed(t *testing.T) {
	fmt.Println("start")
	printFixed()()
}

func printFixed() func() { // <- removed named return
	total := 0
	f := func() { // <- will be called
		fmt.Println("total")
		total += 1
	}
	return func() {
		fmt.Println("ini!")
		f() // calling f from line 33
		fmt.Println("done!")
	}
}
