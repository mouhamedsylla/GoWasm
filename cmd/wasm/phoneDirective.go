package main

import (
	"syscall/js"
)


// The PhoneDirective type represents a phone directive in Go.
// @property tp - The "tp" property is a JavaScript value that represents the type of phone directive.
type PhoneDirective struct {
	tp js.Value
}

// The `Selector()` function is a method of the `PhoneDirective` struct. It returns a CSS selector
// string that is used to select the phone number input element associated with the `PhoneDirective`
// directive. In this case, the selector is `[phone-number]`, which means it will select any element
// with the `phone-number` attribute.
func (d *PhoneDirective) Selector() string {
	return "[phone-number]"
}

// The `SetElement` function is a method of the `PhoneDirective` struct. It is used to set the value of
// the `tp` field of the `PhoneDirective` struct to the provided `el` value. This allows the
// `PhoneDirective` struct to have access to the phone number input element that it is associated with.
func (d *PhoneDirective) SetElement(el js.Value){
	d.tp = el
}

// The `Init()` function is a method of the `PhoneDirective` struct. It is responsible for initializing
// the phone number input element.
func (d *PhoneDirective) Init() {
	input := d.tp
	val := ""
	input.Get("style").Set("borderColor", "red")

	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		val = input.Get("value").String()
		input.Set("value", FormatNumber(NumberVerify(val, 10), 2))
		return nil
	}))
}