// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type LoginAccessTokenLoginAccessTokenPostResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	JWToken *shared.JWToken
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
}

func (o *LoginAccessTokenLoginAccessTokenPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LoginAccessTokenLoginAccessTokenPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LoginAccessTokenLoginAccessTokenPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LoginAccessTokenLoginAccessTokenPostResponse) GetJWToken() *shared.JWToken {
	if o == nil {
		return nil
	}
	return o.JWToken
}

func (o *LoginAccessTokenLoginAccessTokenPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
