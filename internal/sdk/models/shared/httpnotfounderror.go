// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HTTPNotFoundError struct {
	Detail string `json:"detail"`
}

func (o *HTTPNotFoundError) GetDetail() string {
	if o == nil {
		return ""
	}
	return o.Detail
}