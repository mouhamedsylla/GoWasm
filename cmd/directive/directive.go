package directive

import (
	"regexp"
	"strings"
	"syscall/js"
)

// The above type defines an interface for directives in Go.
// @property Init - The Init method is used to initialize the directive. It can be used to set up any
// necessary state or perform any initializations that are required for the directive to function
// properly.
// @property {string} Selector - The `Selector` method returns a string that represents the CSS
// selector used to select the element(s) that the directive will be applied to.
// @property SetElement - SetElement is a method that sets the element on which the directive is
// applied. It takes a js.Value parameter, which represents the JavaScript object that corresponds to
// the element.
type Directive interface {
	Init()
	Selector() string
	SetElement(js.Value)
}


// The NumberVerify function takes a string and a limit as input, removes any spaces from the string,
// truncates it if it exceeds the limit, and returns the string if it consists only of digits,
// otherwise it returns the string without the last character.
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

// The FormatNumber function takes a string and a group size as input, and returns the string with
// spaces inserted every group size characters.
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
