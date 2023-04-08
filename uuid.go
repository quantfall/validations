package validations

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/quantfall/rerror"
)

func ValidateUUID(value interface{}) (uuid.UUID, string, *rerror.Error) {
	v, ok := value.(string)
	if !ok {
		return uuid.UUID{}, "", rerror.ErrorF(http.StatusBadRequest, "value is not string: %v", value)
	}

	u, err := uuid.Parse(strings.ToLower(v))
	if err != nil {
		return uuid.UUID{}, "", rerror.ErrorF(http.StatusBadRequest, err.Error())
	}

	return u, GenerateSHA3256(u.String()), nil
}
