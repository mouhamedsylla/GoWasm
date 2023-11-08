package services

import (
	"regexp"
	"strings"
)

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
func FormatNumber(s string, group int, sep string) string {
	r := ""
	cpt := 0
	for i, v := range s {
		r += string(v)
		if (len(r)-cpt)%group == 0 && i != len(s)-1 {
			r += sep
			cpt++
		}
	}
	return r
}
