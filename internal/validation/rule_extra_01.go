package validation

import (
	"regexp"
	"strings"
)

type Rule1 struct {
	Name     string
	Pattern  *regexp.Regexp
	Required bool
}

func NewRule1(name, pattern string, required bool) (Rule1, error) {
	p, err := regexp.Compile(pattern)
	if err != nil {
		return Rule1{}, err
	}
	return Rule1{Name: strings.TrimSpace(name), Pattern: p, Required: required}, nil
}

func (r Rule1) Check(value string) bool {
	if r.Required && strings.TrimSpace(value) == "" {
		return false
	}
	if r.Pattern == nil {
		return !r.Required || value != ""
	}
	return r.Pattern.MatchString(value)
}
