package main

import "syscall/js"


type CreditDirective struct{
	tp js.Value
}

// Constructeur d'un directive credit
func (d *CreditDirective) SetElement(el js.Value) {
	d.tp = el
}

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

func (d *CreditDirective) Selector() string {
	return "[credit-number]"
}