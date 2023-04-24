//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config codegen.yaml api.yaml
package ergast

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
)

// generic response type from a call to oapi-codegen's wrapper
type apiResponseType interface {
	StatusCode() int
}

// CallErgast is the lowest level api call helper
//
// T corresponds to the type expected from the JSON body as a result from the api call
// U is the returned type from the call to oapi codegen's wrapper
// currently supports 200 and 201 response calls for valid output structs
func CallErgast[T any, U apiResponseType](ctx context.Context, apiCall func() (res U, err error)) (*T, error) {
	var response *T

	res, resErr := apiCall()
	if resErr != nil {
		return nil, fmt.Errorf("error encountered on api call")
	}

	resBody := string(reflect.ValueOf(res).Elem().FieldByName("Body").Bytes())

	switch res.StatusCode() {
	case http.StatusOK:
		response = (*T)(reflect.ValueOf(res).Elem().FieldByName("JSON200").UnsafePointer())

		return response, nil
	case http.StatusCreated:
		response = (*T)(reflect.ValueOf(res).Elem().FieldByName("JSON201").UnsafePointer())

		return response, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("bad request: %s", resBody)
	case http.StatusNotFound:
		return nil, fmt.Errorf("not found: %s", resBody)
	case http.StatusConflict:
		return nil, fmt.Errorf("resource conflict: %s", resBody)
	default:
		return nil, fmt.Errorf("unknown error: %s", resBody)
	}
}
