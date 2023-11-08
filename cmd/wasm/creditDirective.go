package main

import "syscall/js"


// The CreditDirective type represents a credit directive in the Go programming language.
// @property tp - The "tp" property is of type js.Value.
type CreditDirective struct{
	tp js.Value
}

// The `Selector()` function is defining a method for the `CreditDirective` struct. This method returns
// a CSS selector string that can be used to select elements in the HTML document that should be
// associated with the `CreditDirective` directive. In this case, the selector is `[credit-number]`,
// which means it will select elements that have the attribute `credit-number`.
func (d *CreditDirective) Selector() string {
	return "[credit-number]"
}

// The `SetElement` function is a method of the `CreditDirective` struct. It takes a `js.Value`
// parameter `el` and assigns it to the `tp` property of the `CreditDirective` instance `d`. This
// function is used to set the element associated with the `CreditDirective` directive.
func (d *CreditDirective) SetElement(el js.Value) {
	d.tp = el
}

// The `Init()` function is a method of the `CreditDirective` struct. It is responsible for
// initializing the credit directive by adding event listeners and applying some styling to the
// associated element.
func (d *CreditDirective) Init() {
	input := d.tp
	val := ""
	input.Get("style").Set("borderColor", "blue")

	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		val = input.Get("value").String()
		input.Set("value", FormatNumber(NumberVerify(val, 16), 4))
		return nil
	}))
}

