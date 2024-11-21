// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
)

func (r *PostIncidentResourceModel) ToSharedPostIncidentRequest() *shared.PostIncidentRequest {
	var title string
	title = r.Title.ValueString()

	var content string
	content = r.Content.ValueString()

	style := new(shared.IncidentStyle)
	if !r.Style.IsUnknown() && !r.Style.IsNull() {
		*style = shared.IncidentStyle(r.Style.ValueString())
	} else {
		style = nil
	}
	out := shared.PostIncidentRequest{
		Title:   title,
		Content: content,
		Style:   style,
	}
	return &out
}

func (r *PostIncidentResourceModel) RefreshFromSharedPostIncidentResponse(resp *shared.PostIncidentResponse) {
	if resp != nil {
		r.Content = types.StringValue(resp.Content)
		r.CreatedDate = types.StringValue(resp.CreatedDate)
		r.ID = types.Int64Value(resp.ID)
		r.Pin = types.BoolValue(resp.Pin)
		r.Style = types.StringValue(string(resp.Style))
		r.Title = types.StringValue(resp.Title)
	}
}