// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PostIncidentResponse struct {
	Content     string `json:"content"`
	CreatedDate string `json:"createdDate"`
	ID          int64  `json:"id"`
	Pin         bool   `json:"pin"`
	Style       string `json:"style"`
	Title       string `json:"title"`
}

func (o *PostIncidentResponse) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *PostIncidentResponse) GetCreatedDate() string {
	if o == nil {
		return ""
	}
	return o.CreatedDate
}

func (o *PostIncidentResponse) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PostIncidentResponse) GetPin() bool {
	if o == nil {
		return false
	}
	return o.Pin
}

func (o *PostIncidentResponse) GetStyle() string {
	if o == nil {
		return ""
	}
	return o.Style
}

func (o *PostIncidentResponse) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}
