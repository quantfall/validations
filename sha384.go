package validations

import (
	"net/http"
	"strings"

	"github.com/quantfall/rerror"
)

func ValidateSHA384(value interface{}) (string, string, *rerror.Error) {
	v, ok := value.(string)
	if !ok {
		return "", "", rerror.ErrorF(false, http.StatusBadRequest, "TYPE_VALIDATION", "value is not string: %v", value)
	}
	v = strings.ToLower(v)
	e := ValidateRegEx(`^[0-9a-f]{96}$`, v)
	if e != nil {
		return "", "", e
	}

	return v, GenerateSHA3256(v), nil
}
