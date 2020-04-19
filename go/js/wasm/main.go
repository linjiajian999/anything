package main

import (
	"strconv"
	"syscall/js"
)

func registerCallbacks() {
	global := js.Global()
	document := global.Get("document")

	getElementById := func(id string) js.Value {
		return document.Call("getElementById", id)
	}
	aValue := getElementById("aValue")
	bValue := getElementById("bValue")
	cValue := getElementById("cValue")
	sumValue := getElementById("sum")

	sumButton := getElementById("sumButton")
	runButton := getElementById("runButton")

	onRun := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("button on click")
		return nil
	})
	onSum := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		a, _ := strconv.Atoi(aValue.Get("value").String())
		b, _ := strconv.Atoi(bValue.Get("value").String())
		c, _ := strconv.Atoi(cValue.Get("value").String())
		sumValue.Set("value", a+b+c)
		return nil
	})

	sumButton.Call("addEventListener", "click", onSum)
	runButton.Call("addEventListener", "click", onRun)
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
