package validation

import (
	"regexp"
	"strings"
)

type Rule2 struct {
	Name     string
	Pattern  *regexp.Regexp
	Required bool
}

func NewRule2(name, pattern string, required bool) (Rule2, error) {
	p, err := regexp.Compile(pattern)
	if err != nil {
		return Rule2{}, err
	}
	return Rule2{Name: strings.TrimSpace(name), Pattern: p, Required: required}, nil
}

func (r Rule2) Check(value string) bool {
	if r.Required && strings.TrimSpace(value) == "" {
		return false
	}
	if r.Pattern == nil {
		return !r.Required || value != ""
	}
	return r.Pattern.MatchString(value)
}
