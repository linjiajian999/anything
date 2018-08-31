package main

import (
	"syscall/js"
)

func sum(args []js.Value) {
	var sum int
	for val := range args {
		sum += args[val].Int()
	}
	println(sum)
}
func registerCallbacks() {
	js.Global().Set("sum", js.NewCallback(sum))
	run := js.NewCallback(func(args []js.Value) {
		println("button on click")
	})
	js.Global().Get("document").Call("getElementById", "runButton").Call("addEventListener", "click", run)
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
