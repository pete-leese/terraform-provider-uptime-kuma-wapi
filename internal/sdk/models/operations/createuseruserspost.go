// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type CreateUserUsersPostResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	User *shared.User
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
}

func (o *CreateUserUsersPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserUsersPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserUsersPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserUsersPostResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *CreateUserUsersPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
