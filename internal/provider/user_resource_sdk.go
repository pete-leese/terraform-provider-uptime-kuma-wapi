// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"time"
)

func (r *UserResourceModel) ToSharedRegisterUser() *shared.RegisterUser {
	var username string
	username = r.Username.ValueString()

	var password string
	password = r.Password.ValueString()

	out := shared.RegisterUser{
		Username: username,
		Password: password,
	}
	return &out
}

func (r *UserResourceModel) RefreshFromSharedUser(resp *shared.User) {
	if resp != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.Int64Value(resp.ID)
		r.LastVisit = types.StringValue(resp.LastVisit.Format(time.RFC3339Nano))
		r.Username = types.StringValue(resp.Username)
	}
}
