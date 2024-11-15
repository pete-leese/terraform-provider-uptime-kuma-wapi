// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type StatusPage struct {
	CustomCSS         *string       `json:"customCSS,omitempty"`
	Description       *string       `json:"description,omitempty"`
	DomainNameList    []string      `json:"domainNameList"`
	FooterText        *string       `json:"footerText,omitempty"`
	GoogleAnalyticsID *string       `json:"googleAnalyticsId,omitempty"`
	Icon              string        `json:"icon"`
	ID                int64         `json:"id"`
	Incident          *Incident     `json:"incident,omitempty"`
	MaintenanceList   []any         `json:"maintenanceList,omitempty"`
	Published         bool          `json:"published"`
	ShowPoweredBy     bool          `json:"showPoweredBy"`
	ShowTags          bool          `json:"showTags"`
	Slug              string        `json:"slug"`
	Theme             string        `json:"theme"`
	Title             string        `json:"title"`
	PublicGroupList   []PublicGroup `json:"publicGroupList,omitempty"`
}

func (o *StatusPage) GetCustomCSS() *string {
	if o == nil {
		return nil
	}
	return o.CustomCSS
}

func (o *StatusPage) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *StatusPage) GetDomainNameList() []string {
	if o == nil {
		return []string{}
	}
	return o.DomainNameList
}

func (o *StatusPage) GetFooterText() *string {
	if o == nil {
		return nil
	}
	return o.FooterText
}

func (o *StatusPage) GetGoogleAnalyticsID() *string {
	if o == nil {
		return nil
	}
	return o.GoogleAnalyticsID
}

func (o *StatusPage) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *StatusPage) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *StatusPage) GetIncident() *Incident {
	if o == nil {
		return nil
	}
	return o.Incident
}

func (o *StatusPage) GetMaintenanceList() []any {
	if o == nil {
		return nil
	}
	return o.MaintenanceList
}

func (o *StatusPage) GetPublished() bool {
	if o == nil {
		return false
	}
	return o.Published
}

func (o *StatusPage) GetShowPoweredBy() bool {
	if o == nil {
		return false
	}
	return o.ShowPoweredBy
}

func (o *StatusPage) GetShowTags() bool {
	if o == nil {
		return false
	}
	return o.ShowTags
}

func (o *StatusPage) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *StatusPage) GetTheme() string {
	if o == nil {
		return ""
	}
	return o.Theme
}

func (o *StatusPage) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *StatusPage) GetPublicGroupList() []PublicGroup {
	if o == nil {
		return nil
	}
	return o.PublicGroupList
}