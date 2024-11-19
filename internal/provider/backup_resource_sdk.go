// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
)

func (r *BackupResourceModel) ToSharedBackup() *shared.Backup {
	out := shared.Backup{}
	return &out
}

func (r *BackupResourceModel) RefreshFromInterface(resp interface{}) {
	if resp == nil {
		r.Data = types.StringNull()
	} else {
		dataResult, _ := json.Marshal(resp)
		r.Data = types.StringValue(string(dataResult))
	}
}
