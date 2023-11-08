package directive

import "syscall/js"


// The CreditDirective type represents a credit directive in the Go programming language.
// @property tp - The "tp" property is of type js.Value.
type CreditDirective struct{
	tp js.Value
}

func (d *CreditDirective) Selector() string {
	return "[credit-number]"
}

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

