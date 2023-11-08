package main

import (
	"regexp"
	"strings"
	"syscall/js"
)

type Directive interface {
	Init()
	Selector() string
	SetElement(js.Value)
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
	cpt := 0
	for i, v := range s {
		r += string(v)
		if (len(r)-cpt)%group == 0 && i != len(s)-1 {
			r += " "
			cpt++
		}
	}
	return r
}
