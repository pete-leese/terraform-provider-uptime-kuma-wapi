// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type GetUserUsersUsernameGetRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserUsersUsernameGetRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserUsersUsernameGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	User *shared.User
	// Not Found
	HTTPNotFoundError *shared.HTTPNotFoundError
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
}

func (o *GetUserUsersUsernameGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserUsersUsernameGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserUsersUsernameGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserUsersUsernameGetResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *GetUserUsersUsernameGetResponse) GetHTTPNotFoundError() *shared.HTTPNotFoundError {
	if o == nil {
		return nil
	}
	return o.HTTPNotFoundError
}

func (o *GetUserUsersUsernameGetResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
