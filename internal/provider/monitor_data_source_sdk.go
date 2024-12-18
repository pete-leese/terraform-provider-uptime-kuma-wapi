// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *MonitorDataSourceModel) RefreshFromInterface(resp interface{}) {
	if resp == nil {
		r.Data = types.StringNull()
	} else {
		dataResult, _ := json.Marshal(resp)
		r.Data = types.StringValue(string(dataResult))
	}
}
