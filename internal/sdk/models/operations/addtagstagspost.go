// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type AddTagsTagsPostResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	Any any
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
}

func (o *AddTagsTagsPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddTagsTagsPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddTagsTagsPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddTagsTagsPostResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *AddTagsTagsPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
