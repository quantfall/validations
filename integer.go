package validations

import (
	"fmt"
	"net/http"

	"github.com/quantfall/rerror"
	"google.golang.org/grpc/codes"
)

func ValidateInteger(value interface{}) (int64, string, *rerror.Error) {
	v, ok := value.(int64)
	if !ok {
		return 0, "", rerror.ErrorF(http.StatusBadRequest, codes.InvalidArgument, "value is not integer: %v", value)
	}

	return v, GenerateSHA3256(fmt.Sprint(v)), nil
}
