package validation

import (
	"regexp"
	"strings"
)

type Rule3 struct {
	Name     string
	Pattern  *regexp.Regexp
	Required bool
}

func NewRule3(name, pattern string, required bool) (Rule3, error) {
	p, err := regexp.Compile(pattern)
	if err != nil {
		return Rule3{}, err
	}
	return Rule3{Name: strings.TrimSpace(name), Pattern: p, Required: required}, nil
}

func (r Rule3) Check(value string) bool {
	if r.Required && strings.TrimSpace(value) == "" {
		return false
	}
	if r.Pattern == nil {
		return !r.Required || value != ""
	}
	return r.Pattern.MatchString(value)
}
