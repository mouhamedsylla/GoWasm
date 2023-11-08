package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	ch := make(chan chan struct{}, 0)
	fmt.Println("Hello WebAssembly from Go")
	
	directives := []Directive{&PhoneDirective{}, &CreditDirective{} }

	for _, dir := range directives {
		elements := js.Global().Get("document").Call("querySelectorAll", dir.Selector())
		for i := 0; i < elements.Length(); i++ {
			dir.SetElement(elements.Index(i))
			dir.Init()
		}
	}
	<-ch
}
