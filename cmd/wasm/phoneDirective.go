package main

import (
	"syscall/js"
)


type PhoneDirective struct {
	tp js.Value
}

func (d *PhoneDirective) Selector() string {
	return "[phone-number]"
}

func (d *PhoneDirective) SetElement(el js.Value){
	d.tp = el
}

// Définition des directives
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