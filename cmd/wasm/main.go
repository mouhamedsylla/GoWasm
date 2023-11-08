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
	// element := js.Global().Get("document").Call("querySelector", B.Selector())

	// fmt.Println(B.Selector())
	// B.el = element

	// B.Init()

	// element1 := js.Global().Get("document").Call("querySelector", C.Selector())

	// fmt.Println(C.Selector())
	// C.el = element1

	// C.Init()

	directives := []Directive{&PhoneDirective{}, &CreditDirective{} }

	for _, dir := range directives {
		elements := js.Global().Get("document").Call("querySelectorAll", dir.Selector())
		for i := 0; i < elements.Length(); i++ {
			dir.SetElement(elements.Index(i))
			dir.Init()
		}
	}

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

	// 	phoneNumber := &PhoneDirective{}
	// 	creditCard := &CreditDirective{}
	// directives := []Directive{phoneNumber, creditCard}

	// for _, directive := range directives {
	// 	elements := js.Global().Call("querySelectorAll", directive.Selector())
	// 	for i := 0; i < elements.Length(); i++ {
	// 		element := elements.Index(i)
	// 		directive.NewDirective(element)
	// 		directive.Init()
	// 	}
	// }

	// document := js.Global().Get("document")
	// phones := document.Call("querySelectorAll", "[phone-number]")

	// for i := 0; i < phones.Length(); i++ {
	// 	element := phones.Index(i)
	// 	if !element.IsNull() {
	// 		directiv := PhoneNumberDirective(element)
	// 		directiv.Init()
	// 	}
	// }

	// credits := document.Call("querySelectorAll", "[credit-number]")

	// for i := 0; i < credits.Length(); i++ {
	// 	element := credits.Index(i)
	// 	if !element.IsNull() {
	// 		directiv := CreditCardDirective(element)
	// 		directiv.Init()
	// 	}
	// }

	<-ch
}
