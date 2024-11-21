// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type DeleteMonitorMonitorsMonitorIDDeleteRequest struct {
	MonitorID int64 `pathParam:"style=simple,explode=false,name=monitor_id"`
}

func (o *DeleteMonitorMonitorsMonitorIDDeleteRequest) GetMonitorID() int64 {
	if o == nil {
		return 0
	}
	return o.MonitorID
}

type DeleteMonitorMonitorsMonitorIDDeleteResponse struct {
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

func (o *DeleteMonitorMonitorsMonitorIDDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMonitorMonitorsMonitorIDDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMonitorMonitorsMonitorIDDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMonitorMonitorsMonitorIDDeleteResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *DeleteMonitorMonitorsMonitorIDDeleteResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
