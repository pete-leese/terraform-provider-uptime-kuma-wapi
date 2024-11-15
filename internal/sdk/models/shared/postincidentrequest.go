// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform-provider-ukumawapi/internal/sdk/internal/utils"
)

// IncidentStyle - Enumerate incident styles.
type IncidentStyle string

const (
	IncidentStyleInfo    IncidentStyle = "info"
	IncidentStyleWarning IncidentStyle = "warning"
	IncidentStyleDanger  IncidentStyle = "danger"
	IncidentStylePrimary IncidentStyle = "primary"
	IncidentStyleLight   IncidentStyle = "light"
	IncidentStyleDark    IncidentStyle = "dark"
)

func (e IncidentStyle) ToPointer() *IncidentStyle {
	return &e
}
func (e *IncidentStyle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warning":
		fallthrough
	case "danger":
		fallthrough
	case "primary":
		fallthrough
	case "light":
		fallthrough
	case "dark":
		*e = IncidentStyle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncidentStyle: %v", v)
	}
}

type PostIncidentRequest struct {
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Style   *IncidentStyle `default:"primary" json:"style"`
}

func (p PostIncidentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostIncidentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostIncidentRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PostIncidentRequest) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *PostIncidentRequest) GetStyle() *IncidentStyle {
	if o == nil {
		return nil
	}
	return o.Style
}