// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// tests for the helper functions to convert to JSONAPI Errors
//
// Command:
// $ goagen
// --design=github.com/golang-starters/golang-rest-http/design
// --out=$(GOPATH)/src/github.com/golang-starters/golang-rest-http/app
// --version=v1.3.0

package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"github.com/fabric8-services/fabric8-common/errors"
	"github.com/fabric8-services/fabric8-common/resource"
	errs "github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)


func TestErrorToJSONAPIError(t *testing.T) {
	t.Parallel()
	resource.Require(t, resource.UnitTest)

	var jerr JSONAPIError
	var httpStatus int

	// test not found error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewNotFoundError("foo", "bar"))
	require.Equal(t, http.StatusNotFound, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeNotFound, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test not found error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewConversionError("foo"))
	require.Equal(t, http.StatusBadRequest, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeConversionError, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test bad parameter error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewBadParameterError("foo", "bar"))
	require.Equal(t, http.StatusBadRequest, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeBadParameter, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test internal server error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewInternalError(context.Background(), errs.New("foo")))
	require.Equal(t, http.StatusInternalServerError, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeInternalError, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test unauthorized error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewUnauthorizedError("foo"))
	require.Equal(t, http.StatusUnauthorized, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeUnauthorizedError, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test forbidden error
	jerr, httpStatus = ErrorToJSONAPIError(nil, errors.NewForbiddenError("foo"))
	require.Equal(t, http.StatusForbidden, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeForbiddenError, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)

	// test unspecified error
	jerr, httpStatus = ErrorToJSONAPIError(nil, fmt.Errorf("foobar"))
	require.Equal(t, http.StatusInternalServerError, httpStatus)
	require.NotNil(t, jerr.Code)
	require.NotNil(t, jerr.Status)
	require.Equal(t, ErrorCodeUnknownError, *jerr.Code)
	require.Equal(t, strconv.Itoa(httpStatus), *jerr.Status)
}
