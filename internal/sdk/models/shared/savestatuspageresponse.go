// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SaveStatusPageResponse struct {
	Detail any `json:"detail,omitempty"`
}

func (o *SaveStatusPageResponse) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}
