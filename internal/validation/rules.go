package validation

import (
	"regexp"
	"strings"
)

var sampleID = regexp.MustCompile(`^[A-Z]{2,5}-[0-9]{4,8}$`)

func ValidID(id string) bool { return sampleID.MatchString(strings.TrimSpace(id)) }
func Required(values ...string) bool {
	for _, v := range values {
		if strings.TrimSpace(v) == "" {
			return false
		}
	}
	return true
}
