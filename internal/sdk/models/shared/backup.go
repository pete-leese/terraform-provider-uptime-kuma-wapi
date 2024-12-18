// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Backup struct {
	Version          *string                `json:"version,omitempty"`
	NotificationList []NotificationListItem `json:"notificationList,omitempty"`
	MonitorList      []MonitorListItem      `json:"monitorList,omitempty"`
	ProxyList        []any                  `json:"proxyList,omitempty"`
}

func (o *Backup) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *Backup) GetNotificationList() []NotificationListItem {
	if o == nil {
		return nil
	}
	return o.NotificationList
}

func (o *Backup) GetMonitorList() []MonitorListItem {
	if o == nil {
		return nil
	}
	return o.MonitorList
}

func (o *Backup) GetProxyList() []any {
	if o == nil {
		return nil
	}
	return o.ProxyList
}
