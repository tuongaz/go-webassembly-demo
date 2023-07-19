package main

import (
	"fmt"
	"syscall/js" // work hand-in-hand with wasm_exec.js
)

func main() {
	// how to call go func from browser
	demo1()

	// how to update DOM from Go
	demo2()

	// Keep go running
	select {}
}

// demo1 Call Go from browser
func demo1() {
	// js.Global() ~ window in JS
	js.Global().Set("hello", js.FuncOf(hello))
}

func hello(this js.Value, inputs []js.Value) any {
	name := inputs[0].String()
	return fmt.Sprintf("Hello %s!", name)
}

// demo2 Call JS function from Go
func demo2() {
	js.Global().Set("changeToRed", js.FuncOf(changeToRed))
}

func changeToRed(_ js.Value, _ []js.Value) any {
	document := js.Global().Get("document")
	button := document.Call("getElementById", "demo2")
	button.Get("style").Set("color", "red")
	return nil
}
