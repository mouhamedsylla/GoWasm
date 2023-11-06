package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `<h4>Hello, I am HTML snippet from GO!</h4>`

// func GetHtml() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		return htmlString
// 	})
// }

func main() {
	ch := make(chan chan struct{}, 0)
	fmt.Println("Hello WebAssembly from Go")

	// selectors := map[interface{}]string{
	// 	PhoneNumberDirective: "[phone-number]",
	// 	CreditCardDirective: "[credit-number]",
	// }

	// directives := []interface{}{
	// 	PhoneNumberDirective,
	// 	CreditCardDirective,
	// }

	// for _, obj := range directives {
	// 	var elements []js.Value
	// 	//elements = js.Global().Get("document").Call("querySelectorAll", )
	// 	switch v := obj.(type) {
	// 	case PhoneDirective:
	// 		elements = js.Global().Get("document").Call("querySelectorAll", v.selector)
	// 		for i := 0; i < elements.Length(); i++ {
	// 			element := elements.Index(i)
	// 			if !element.IsNull() {
	// 				v = v
	// 			}
	// 		}
	// 	case CreditDirective:
	// 		elements = js.Global().Get("document").Call("querySelectorAll", v.selector)
	// 	}

	// 	// elements :=  js.Global().Get("document").Call("querySelectorAll", )
	// 	// for i := 0; i < elements.Length(); i++ {
	// 	// 	element := elements.Index(i)
	// 	// 	if !element.IsNull() {
	// 	// }
	// 	// }
	// }

	// for _, v := range directives{
	// 	elements :=  js.Global().Get("document").Call("querySelectorAll", selector)
	// 	for i := 0; i < elements.Length(); i++ {
	// 		element := elements.Index(i)

	// 	}
	// }

	document := js.Global().Get("document")
	phones := document.Call("querySelectorAll", "[phone-number]")

	for i := 0; i < phones.Length(); i++ {
		element := phones.Index(i)
		if !element.IsNull() {
			directiv := PhoneNumberDirective(element)
			directiv.Init()
		}
	}

	credits := document.Call("querySelectorAll", "[credit-number]")

	for i := 0; i < credits.Length(); i++ {
		element := credits.Index(i)
		if !element.IsNull() {
			directiv := CreditCardDirective(element)
			directiv.Init()
		}
	}

	<-ch
}
