package directive

import (
	"syscall/js"
	"wasm-go/services"
)

// The PhoneDirective type represents a phone directive in Go.
// @property tp - The "tp" property is a JavaScript value that represents the type of phone directive.
type PhoneDirective struct {
	tp js.Value
}

func (d *PhoneDirective) Selector() string {
	return "[phone-number]"
}

func (d *PhoneDirective) SetElement(el js.Value) {
	d.tp = el
}

func (d *PhoneDirective) Init() {
	input := d.tp
	val := ""
	input.Get("style").Set("borderColor", SetBorder(input))
	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		val = input.Get("value").String()
		input.Set("value", services.FormatNumber(services.NumberVerify(val, 10), 2, WithSpace(input)))
		return nil
	}))
}

func SetBorder(el js.Value) string {
	if el.Call("hasAttribute", "border-color").Bool() {
		val := el.Call("getAttribute", "border-color").String()
		return val
	}
	return ""
}

func WithSpace(el js.Value) string {
	if el.Call("hasAttribute", "with-space").Bool() {
		val := el.Call("getAttribute", "with-space").String()
		if val == "false" {
			return ""
		}
	}
	return " "
}
