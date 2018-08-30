package main

import (
	"syscall/js"
)

func add(args []js.Value) {
	var sum int
	for val := range i {
		sum += args[val].Int()
	}
	println(sum)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
