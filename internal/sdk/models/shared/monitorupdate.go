// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-ukumawapi/internal/sdk/internal/utils"
)

// MonitorUpdateAuthMethod - Enumerate authentication methods for monitors.
type MonitorUpdateAuthMethod string

const (
	MonitorUpdateAuthMethodUnknown MonitorUpdateAuthMethod = ""
	MonitorUpdateAuthMethodBasic   MonitorUpdateAuthMethod = "basic"
	MonitorUpdateAuthMethodNtlm    MonitorUpdateAuthMethod = "ntlm"
	MonitorUpdateAuthMethodMtls    MonitorUpdateAuthMethod = "mtls"
)

func (e MonitorUpdateAuthMethod) ToPointer() *MonitorUpdateAuthMethod {
	return &e
}
func (e *MonitorUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "basic":
		fallthrough
	case "ntlm":
		fallthrough
	case "mtls":
		*e = MonitorUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MonitorUpdateAuthMethod: %v", v)
	}
}

type MonitorUpdate struct {
	// Enumerate monitor types.
	Type                     *MonitorType             `json:"type,omitempty"`
	Name                     *string                  `json:"name,omitempty"`
	Interval                 *int64                   `default:"60" json:"interval"`
	RetryInterval            *int64                   `default:"60" json:"retryInterval"`
	ResendInterval           *int64                   `default:"0" json:"resendInterval"`
	Maxretries               *int64                   `default:"0" json:"maxretries"`
	UpsideDown               *bool                    `default:"false" json:"upsideDown"`
	NotificationIDList       []any                    `json:"notificationIDList,omitempty"`
	URL                      *string                  `json:"url,omitempty"`
	ExpiryNotification       *bool                    `default:"false" json:"expiryNotification"`
	IgnoreTLS                *bool                    `default:"false" json:"ignoreTls"`
	Maxredirects             *int64                   `default:"10" json:"maxredirects"`
	AcceptedStatuscodes      []any                    `json:"accepted_statuscodes,omitempty"`
	ProxyID                  *int64                   `json:"proxyId,omitempty"`
	Method                   *string                  `default:"GET" json:"method"`
	Body                     *string                  `json:"body,omitempty"`
	Headers                  *string                  `json:"headers,omitempty"`
	AuthMethod               *MonitorUpdateAuthMethod `default:"" json:"authMethod"`
	BasicAuthUser            *string                  `json:"basic_auth_user,omitempty"`
	BasicAuthPass            *string                  `json:"basic_auth_pass,omitempty"`
	AuthDomain               *string                  `json:"authDomain,omitempty"`
	AuthWorkstation          *string                  `json:"authWorkstation,omitempty"`
	Keyword                  *string                  `json:"keyword,omitempty"`
	Hostname                 *string                  `json:"hostname,omitempty"`
	Port                     *int64                   `default:"53" json:"port"`
	DNSResolveServer         *string                  `default:"1.1.1.1" json:"dns_resolve_server"`
	DNSResolveType           *string                  `default:"A" json:"dns_resolve_type"`
	MqttUsername             *string                  `json:"mqttUsername,omitempty"`
	MqttPassword             *string                  `json:"mqttPassword,omitempty"`
	MqttTopic                *string                  `json:"mqttTopic,omitempty"`
	MqttSuccessMessage       *string                  `json:"mqttSuccessMessage,omitempty"`
	DatabaseConnectionString *string                  `json:"databaseConnectionString,omitempty"`
	DatabaseQuery            *string                  `json:"databaseQuery,omitempty"`
	DockerContainer          *string                  `default:"" json:"docker_container"`
	DockerHost               *int64                   `json:"docker_host,omitempty"`
	RadiusUsername           *string                  `json:"radiusUsername,omitempty"`
	RadiusPassword           *string                  `json:"radiusPassword,omitempty"`
	RadiusSecret             *string                  `json:"radiusSecret,omitempty"`
	RadiusCalledStationID    *string                  `json:"radiusCalledStationId,omitempty"`
	RadiusCallingStationID   *string                  `json:"radiusCallingStationId,omitempty"`
}

func (m MonitorUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MonitorUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MonitorUpdate) GetType() *MonitorType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *MonitorUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MonitorUpdate) GetInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *MonitorUpdate) GetRetryInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryInterval
}

func (o *MonitorUpdate) GetResendInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.ResendInterval
}

func (o *MonitorUpdate) GetMaxretries() *int64 {
	if o == nil {
		return nil
	}
	return o.Maxretries
}

func (o *MonitorUpdate) GetUpsideDown() *bool {
	if o == nil {
		return nil
	}
	return o.UpsideDown
}

func (o *MonitorUpdate) GetNotificationIDList() []any {
	if o == nil {
		return nil
	}
	return o.NotificationIDList
}

func (o *MonitorUpdate) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *MonitorUpdate) GetExpiryNotification() *bool {
	if o == nil {
		return nil
	}
	return o.ExpiryNotification
}

func (o *MonitorUpdate) GetIgnoreTLS() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreTLS
}

func (o *MonitorUpdate) GetMaxredirects() *int64 {
	if o == nil {
		return nil
	}
	return o.Maxredirects
}

func (o *MonitorUpdate) GetAcceptedStatuscodes() []any {
	if o == nil {
		return nil
	}
	return o.AcceptedStatuscodes
}

func (o *MonitorUpdate) GetProxyID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProxyID
}

func (o *MonitorUpdate) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *MonitorUpdate) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *MonitorUpdate) GetHeaders() *string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *MonitorUpdate) GetAuthMethod() *MonitorUpdateAuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *MonitorUpdate) GetBasicAuthUser() *string {
	if o == nil {
		return nil
	}
	return o.BasicAuthUser
}

func (o *MonitorUpdate) GetBasicAuthPass() *string {
	if o == nil {
		return nil
	}
	return o.BasicAuthPass
}

func (o *MonitorUpdate) GetAuthDomain() *string {
	if o == nil {
		return nil
	}
	return o.AuthDomain
}

func (o *MonitorUpdate) GetAuthWorkstation() *string {
	if o == nil {
		return nil
	}
	return o.AuthWorkstation
}

func (o *MonitorUpdate) GetKeyword() *string {
	if o == nil {
		return nil
	}
	return o.Keyword
}

func (o *MonitorUpdate) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *MonitorUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *MonitorUpdate) GetDNSResolveServer() *string {
	if o == nil {
		return nil
	}
	return o.DNSResolveServer
}

func (o *MonitorUpdate) GetDNSResolveType() *string {
	if o == nil {
		return nil
	}
	return o.DNSResolveType
}

func (o *MonitorUpdate) GetMqttUsername() *string {
	if o == nil {
		return nil
	}
	return o.MqttUsername
}

func (o *MonitorUpdate) GetMqttPassword() *string {
	if o == nil {
		return nil
	}
	return o.MqttPassword
}

func (o *MonitorUpdate) GetMqttTopic() *string {
	if o == nil {
		return nil
	}
	return o.MqttTopic
}

func (o *MonitorUpdate) GetMqttSuccessMessage() *string {
	if o == nil {
		return nil
	}
	return o.MqttSuccessMessage
}

func (o *MonitorUpdate) GetDatabaseConnectionString() *string {
	if o == nil {
		return nil
	}
	return o.DatabaseConnectionString
}

func (o *MonitorUpdate) GetDatabaseQuery() *string {
	if o == nil {
		return nil
	}
	return o.DatabaseQuery
}

func (o *MonitorUpdate) GetDockerContainer() *string {
	if o == nil {
		return nil
	}
	return o.DockerContainer
}

func (o *MonitorUpdate) GetDockerHost() *int64 {
	if o == nil {
		return nil
	}
	return o.DockerHost
}

func (o *MonitorUpdate) GetRadiusUsername() *string {
	if o == nil {
		return nil
	}
	return o.RadiusUsername
}

func (o *MonitorUpdate) GetRadiusPassword() *string {
	if o == nil {
		return nil
	}
	return o.RadiusPassword
}

func (o *MonitorUpdate) GetRadiusSecret() *string {
	if o == nil {
		return nil
	}
	return o.RadiusSecret
}

func (o *MonitorUpdate) GetRadiusCalledStationID() *string {
	if o == nil {
		return nil
	}
	return o.RadiusCalledStationID
}

func (o *MonitorUpdate) GetRadiusCallingStationID() *string {
	if o == nil {
		return nil
	}
	return o.RadiusCallingStationID
}
