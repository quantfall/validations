package validations

import (
	"net/http"
	"strings"

	"github.com/quantfall/rerror"
)

func ValidateMime(value interface{}) (string, string, *rerror.Error) {
	v, ok := value.(string)
	if !ok {
		return "", "", rerror.ErrorF(http.StatusBadRequest, "value is not string: %v", value)
	}

	v = strings.ToLower(v)

	e := ValidateRegEx(`^([a-z]+)[/]([a-z]+[a-z+-.][a-z]+)+$`, v)
	if e != nil {
		return "", "", e
	}

	return v, GenerateSHA3256(v), nil
}
