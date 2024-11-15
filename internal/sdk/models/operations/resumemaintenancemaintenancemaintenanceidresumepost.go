// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

type ResumeMaintenanceMaintenanceMaintenanceIDResumePostRequest struct {
	MaintenanceID int64 `pathParam:"style=simple,explode=false,name=maintenance_id"`
}

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostRequest) GetMaintenanceID() int64 {
	if o == nil {
		return 0
	}
	return o.MaintenanceID
}

type ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse struct {
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

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *ResumeMaintenanceMaintenanceMaintenanceIDResumePostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
