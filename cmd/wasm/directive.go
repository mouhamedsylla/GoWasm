package main

import (
	"regexp"
	"strings"
	"syscall/js"
)


// Phone directive
type PhoneDirective struct {
	tp js.Value
	selector string
}

// Credit directive
type CreditDirective struct{
	tp js.Value
	selector string
}


// Constructeur d'un directive phone
func PhoneNumberDirective(el js.Value) PhoneDirective {
	n := PhoneDirective{
		tp: el,
		selector: "[phone-number]",
	}
	
	return n
}

// Constructeur d'un directive credit
func CreditCardDirective(el js.Value) CreditDirective {
	n := CreditDirective{
		tp: el,
		selector: "[credit-number]",
	}
	
	return n
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


func NumberVerify(s string, limit int) string {
	s = strings.Join(strings.Split(s, " "), "")
	if len(s) > limit {
		s = s[:limit]
	}
	regex := regexp.MustCompile(`^\d+\s*$`)
	match := regex.MatchString(s)
	if match {
		return s
	}
	return s[:len(s)-1]
}

func FormatNumber(s string, group int) string {
	r := ""
	for i, v := range s {
		if i != 0 && i%group == 0 && i != len(s)-1 {
			r += string(v) + " "
		} else {
			r += string(v)
		}
	}
	return r
}

