package validator
import (
"regexp"
)

type Validator struct {
	Errors map[string]string
}
