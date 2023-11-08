package main

import (
	"fmt"
	"syscall/js"
	d "wasm-go/directive"
)

func main() {
	ch := make(chan chan struct{}, 0)
	fmt.Println("Hello WebAssembly from Go")

	directives := []d.Directive{&d.PhoneDirective{}, &d.CreditDirective{}}

	for _, dir := range directives {
		elements := js.Global().Get("document").Call("querySelectorAll", dir.Selector())
		fmt.Println(elements)
		for i := 0; i < elements.Length(); i++ {
			dir.SetElement(elements.Index(i))
			dir.Init()
		}
	}
	<-ch
}
