// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AddStatusPageRequest struct {
	Slug  *string `json:"slug,omitempty"`
	Title *string `json:"title,omitempty"`
	Msg   *string `json:"msg,omitempty"`
}

func (o *AddStatusPageRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *AddStatusPageRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AddStatusPageRequest) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}