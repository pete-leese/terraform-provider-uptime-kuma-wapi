// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MonitorMaintenance struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (o *MonitorMaintenance) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *MonitorMaintenance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
