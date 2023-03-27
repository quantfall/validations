package validations

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/quantfall/rerror"
	"google.golang.org/grpc/codes"
)

func ValidateURL(value interface{}) (string, string, *rerror.Error) {
	v, ok := value.(string)
	if !ok {
		return "", "", rerror.ErrorF(http.StatusBadRequest, codes.InvalidArgument, "value is not string: %v", value)
	}

	tmp, err := url.ParseRequestURI(v)
	if err != nil {
		return "", "", rerror.ErrorF(http.StatusBadRequest, codes.InvalidArgument, err.Error())
	}

	surl := tmp.String()
	tmp.Host = strings.ToLower(tmp.Host)
	tmp.Scheme = strings.ToLower(tmp.Scheme)

	return surl, GenerateSHA3256(surl), nil
}
