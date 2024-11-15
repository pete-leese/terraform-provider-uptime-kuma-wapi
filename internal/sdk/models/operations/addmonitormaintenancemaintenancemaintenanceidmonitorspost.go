// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostRequest struct {
	MaintenanceID int64                       `pathParam:"style=simple,explode=false,name=maintenance_id"`
	RequestBody   []shared.MonitorMaintenance `request:"mediaType=application/json"`
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostRequest) GetMaintenanceID() int64 {
	if o == nil {
		return 0
	}
	return o.MaintenanceID
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostRequest) GetRequestBody() []shared.MonitorMaintenance {
	if o == nil {
		return []shared.MonitorMaintenance{}
	}
	return o.RequestBody
}

type AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse struct {
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

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *AddMonitorMaintenanceMaintenanceMaintenanceIDMonitorsPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
