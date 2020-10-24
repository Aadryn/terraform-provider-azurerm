package validate

import (
	"fmt"
	"regexp"
)

func LogAnalyticsGenericName(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	if len(v) < 4 {
		errors = append(errors, fmt.Errorf("%q length should be greater than %d", k, 4))
		return
	}
	if len(v) > 63 {
		errors = append(errors, fmt.Errorf("%q length should be less than %d", k, 63))
		return
	}
	if !regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`).MatchString(v) {
		errors = append(errors, fmt.Errorf("the %q is invalid, the %q must begin with an alphanumeric character, end with an alphanumeric character and may only contain alphanumeric characters or hyphens, got %q", k, k, v))
		return
	}
	return
}
