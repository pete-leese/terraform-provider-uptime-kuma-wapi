// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-ukumawapi/internal/sdk/internal/utils"
)

type SaveStatusPageRequest struct {
	Title             *string  `json:"title,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Theme             *string  `default:"light" json:"theme"`
	Published         *bool    `default:"true" json:"published"`
	ShowTags          *bool    `default:"false" json:"showTags"`
	DomainNameList    []string `json:"domainNameList,omitempty"`
	GoogleAnalyticsID *string  `json:"googleAnalyticsId,omitempty"`
	CustomCSS         *string  `default:"" json:"customCSS"`
	FooterText        *string  `json:"footerText,omitempty"`
	ShowPoweredBy     *bool    `default:"true" json:"showPoweredBy"`
	Icon              *string  `default:"/icon.svg" json:"icon"`
	PublicGroupList   []any    `json:"publicGroupList,omitempty"`
}

func (s SaveStatusPageRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SaveStatusPageRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SaveStatusPageRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *SaveStatusPageRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SaveStatusPageRequest) GetTheme() *string {
	if o == nil {
		return nil
	}
	return o.Theme
}

func (o *SaveStatusPageRequest) GetPublished() *bool {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *SaveStatusPageRequest) GetShowTags() *bool {
	if o == nil {
		return nil
	}
	return o.ShowTags
}

func (o *SaveStatusPageRequest) GetDomainNameList() []string {
	if o == nil {
		return nil
	}
	return o.DomainNameList
}

func (o *SaveStatusPageRequest) GetGoogleAnalyticsID() *string {
	if o == nil {
		return nil
	}
	return o.GoogleAnalyticsID
}

func (o *SaveStatusPageRequest) GetCustomCSS() *string {
	if o == nil {
		return nil
	}
	return o.CustomCSS
}

func (o *SaveStatusPageRequest) GetFooterText() *string {
	if o == nil {
		return nil
	}
	return o.FooterText
}

func (o *SaveStatusPageRequest) GetShowPoweredBy() *bool {
	if o == nil {
		return nil
	}
	return o.ShowPoweredBy
}

func (o *SaveStatusPageRequest) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *SaveStatusPageRequest) GetPublicGroupList() []any {
	if o == nil {
		return nil
	}
	return o.PublicGroupList
}
