// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type GetAllStatusPagesStatuspagesGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	StatusPageList *shared.StatusPageList
}

func (o *GetAllStatusPagesStatuspagesGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllStatusPagesStatuspagesGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllStatusPagesStatuspagesGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllStatusPagesStatuspagesGetResponse) GetStatusPageList() *shared.StatusPageList {
	if o == nil {
		return nil
	}
	return o.StatusPageList
}
