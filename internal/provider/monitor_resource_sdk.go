// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
)

func (r *MonitorResourceModel) ToSharedSchemasMonitorMonitor() *shared.SchemasMonitorMonitor {
	out := shared.SchemasMonitorMonitor{}
	return &out
}

func (r *MonitorResourceModel) RefreshFromInterface(resp interface{}) {
	if resp == nil {
		r.Data = types.StringNull()
	} else {
		dataResult, _ := json.Marshal(resp)
		r.Data = types.StringValue(string(dataResult))
	}
}

func (r *MonitorResourceModel) ToSharedMonitorUpdate() *shared.MonitorUpdate {
	typeVar := new(shared.MonitorType)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = shared.MonitorType(r.Type.ValueString())
	} else {
		typeVar = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	interval := new(int64)
	if !r.Interval.IsUnknown() && !r.Interval.IsNull() {
		*interval = r.Interval.ValueInt64()
	} else {
		interval = nil
	}
	retryInterval := new(int64)
	if !r.RetryInterval.IsUnknown() && !r.RetryInterval.IsNull() {
		*retryInterval = r.RetryInterval.ValueInt64()
	} else {
		retryInterval = nil
	}
	resendInterval := new(int64)
	if !r.ResendInterval.IsUnknown() && !r.ResendInterval.IsNull() {
		*resendInterval = r.ResendInterval.ValueInt64()
	} else {
		resendInterval = nil
	}
	maxretries := new(int64)
	if !r.Maxretries.IsUnknown() && !r.Maxretries.IsNull() {
		*maxretries = r.Maxretries.ValueInt64()
	} else {
		maxretries = nil
	}
	upsideDown := new(bool)
	if !r.UpsideDown.IsUnknown() && !r.UpsideDown.IsNull() {
		*upsideDown = r.UpsideDown.ValueBool()
	} else {
		upsideDown = nil
	}
	var notificationIDList []interface{} = []interface{}{}
	for _, notificationIDListItem := range r.NotificationIDList {
		var notificationIDListTmp interface{}
		_ = json.Unmarshal([]byte(notificationIDListItem.ValueString()), &notificationIDListTmp)
		notificationIDList = append(notificationIDList, notificationIDListTmp)
	}
	url := new(string)
	if !r.URL.IsUnknown() && !r.URL.IsNull() {
		*url = r.URL.ValueString()
	} else {
		url = nil
	}
	expiryNotification := new(bool)
	if !r.ExpiryNotification.IsUnknown() && !r.ExpiryNotification.IsNull() {
		*expiryNotification = r.ExpiryNotification.ValueBool()
	} else {
		expiryNotification = nil
	}
	ignoreTLS := new(bool)
	if !r.IgnoreTLS.IsUnknown() && !r.IgnoreTLS.IsNull() {
		*ignoreTLS = r.IgnoreTLS.ValueBool()
	} else {
		ignoreTLS = nil
	}
	maxredirects := new(int64)
	if !r.Maxredirects.IsUnknown() && !r.Maxredirects.IsNull() {
		*maxredirects = r.Maxredirects.ValueInt64()
	} else {
		maxredirects = nil
	}
	var acceptedStatuscodes []interface{} = []interface{}{}
	for _, acceptedStatuscodesItem := range r.AcceptedStatuscodes {
		var acceptedStatuscodesTmp interface{}
		_ = json.Unmarshal([]byte(acceptedStatuscodesItem.ValueString()), &acceptedStatuscodesTmp)
		acceptedStatuscodes = append(acceptedStatuscodes, acceptedStatuscodesTmp)
	}
	proxyID := new(int64)
	if !r.ProxyID.IsUnknown() && !r.ProxyID.IsNull() {
		*proxyID = r.ProxyID.ValueInt64()
	} else {
		proxyID = nil
	}
	method := new(string)
	if !r.Method.IsUnknown() && !r.Method.IsNull() {
		*method = r.Method.ValueString()
	} else {
		method = nil
	}
	body := new(string)
	if !r.Body.IsUnknown() && !r.Body.IsNull() {
		*body = r.Body.ValueString()
	} else {
		body = nil
	}
	headers := new(string)
	if !r.Headers.IsUnknown() && !r.Headers.IsNull() {
		*headers = r.Headers.ValueString()
	} else {
		headers = nil
	}
	authMethod := new(shared.MonitorUpdateAuthMethod)
	if !r.AuthMethod.IsUnknown() && !r.AuthMethod.IsNull() {
		*authMethod = shared.MonitorUpdateAuthMethod(r.AuthMethod.ValueString())
	} else {
		authMethod = nil
	}
	basicAuthUser := new(string)
	if !r.BasicAuthUser.IsUnknown() && !r.BasicAuthUser.IsNull() {
		*basicAuthUser = r.BasicAuthUser.ValueString()
	} else {
		basicAuthUser = nil
	}
	basicAuthPass := new(string)
	if !r.BasicAuthPass.IsUnknown() && !r.BasicAuthPass.IsNull() {
		*basicAuthPass = r.BasicAuthPass.ValueString()
	} else {
		basicAuthPass = nil
	}
	authDomain := new(string)
	if !r.AuthDomain.IsUnknown() && !r.AuthDomain.IsNull() {
		*authDomain = r.AuthDomain.ValueString()
	} else {
		authDomain = nil
	}
	authWorkstation := new(string)
	if !r.AuthWorkstation.IsUnknown() && !r.AuthWorkstation.IsNull() {
		*authWorkstation = r.AuthWorkstation.ValueString()
	} else {
		authWorkstation = nil
	}
	keyword := new(string)
	if !r.Keyword.IsUnknown() && !r.Keyword.IsNull() {
		*keyword = r.Keyword.ValueString()
	} else {
		keyword = nil
	}
	hostname := new(string)
	if !r.Hostname.IsUnknown() && !r.Hostname.IsNull() {
		*hostname = r.Hostname.ValueString()
	} else {
		hostname = nil
	}
	port := new(int64)
	if !r.Port.IsUnknown() && !r.Port.IsNull() {
		*port = r.Port.ValueInt64()
	} else {
		port = nil
	}
	dnsResolveServer := new(string)
	if !r.DNSResolveServer.IsUnknown() && !r.DNSResolveServer.IsNull() {
		*dnsResolveServer = r.DNSResolveServer.ValueString()
	} else {
		dnsResolveServer = nil
	}
	dnsResolveType := new(string)
	if !r.DNSResolveType.IsUnknown() && !r.DNSResolveType.IsNull() {
		*dnsResolveType = r.DNSResolveType.ValueString()
	} else {
		dnsResolveType = nil
	}
	mqttUsername := new(string)
	if !r.MqttUsername.IsUnknown() && !r.MqttUsername.IsNull() {
		*mqttUsername = r.MqttUsername.ValueString()
	} else {
		mqttUsername = nil
	}
	mqttPassword := new(string)
	if !r.MqttPassword.IsUnknown() && !r.MqttPassword.IsNull() {
		*mqttPassword = r.MqttPassword.ValueString()
	} else {
		mqttPassword = nil
	}
	mqttTopic := new(string)
	if !r.MqttTopic.IsUnknown() && !r.MqttTopic.IsNull() {
		*mqttTopic = r.MqttTopic.ValueString()
	} else {
		mqttTopic = nil
	}
	mqttSuccessMessage := new(string)
	if !r.MqttSuccessMessage.IsUnknown() && !r.MqttSuccessMessage.IsNull() {
		*mqttSuccessMessage = r.MqttSuccessMessage.ValueString()
	} else {
		mqttSuccessMessage = nil
	}
	databaseConnectionString := new(string)
	if !r.DatabaseConnectionString.IsUnknown() && !r.DatabaseConnectionString.IsNull() {
		*databaseConnectionString = r.DatabaseConnectionString.ValueString()
	} else {
		databaseConnectionString = nil
	}
	databaseQuery := new(string)
	if !r.DatabaseQuery.IsUnknown() && !r.DatabaseQuery.IsNull() {
		*databaseQuery = r.DatabaseQuery.ValueString()
	} else {
		databaseQuery = nil
	}
	dockerContainer := new(string)
	if !r.DockerContainer.IsUnknown() && !r.DockerContainer.IsNull() {
		*dockerContainer = r.DockerContainer.ValueString()
	} else {
		dockerContainer = nil
	}
	dockerHost := new(int64)
	if !r.DockerHost.IsUnknown() && !r.DockerHost.IsNull() {
		*dockerHost = r.DockerHost.ValueInt64()
	} else {
		dockerHost = nil
	}
	radiusUsername := new(string)
	if !r.RadiusUsername.IsUnknown() && !r.RadiusUsername.IsNull() {
		*radiusUsername = r.RadiusUsername.ValueString()
	} else {
		radiusUsername = nil
	}
	radiusPassword := new(string)
	if !r.RadiusPassword.IsUnknown() && !r.RadiusPassword.IsNull() {
		*radiusPassword = r.RadiusPassword.ValueString()
	} else {
		radiusPassword = nil
	}
	radiusSecret := new(string)
	if !r.RadiusSecret.IsUnknown() && !r.RadiusSecret.IsNull() {
		*radiusSecret = r.RadiusSecret.ValueString()
	} else {
		radiusSecret = nil
	}
	radiusCalledStationID := new(string)
	if !r.RadiusCalledStationID.IsUnknown() && !r.RadiusCalledStationID.IsNull() {
		*radiusCalledStationID = r.RadiusCalledStationID.ValueString()
	} else {
		radiusCalledStationID = nil
	}
	radiusCallingStationID := new(string)
	if !r.RadiusCallingStationID.IsUnknown() && !r.RadiusCallingStationID.IsNull() {
		*radiusCallingStationID = r.RadiusCallingStationID.ValueString()
	} else {
		radiusCallingStationID = nil
	}
	out := shared.MonitorUpdate{
		Type:                     typeVar,
		Name:                     name,
		Interval:                 interval,
		RetryInterval:            retryInterval,
		ResendInterval:           resendInterval,
		Maxretries:               maxretries,
		UpsideDown:               upsideDown,
		NotificationIDList:       notificationIDList,
		URL:                      url,
		ExpiryNotification:       expiryNotification,
		IgnoreTLS:                ignoreTLS,
		Maxredirects:             maxredirects,
		AcceptedStatuscodes:      acceptedStatuscodes,
		ProxyID:                  proxyID,
		Method:                   method,
		Body:                     body,
		Headers:                  headers,
		AuthMethod:               authMethod,
		BasicAuthUser:            basicAuthUser,
		BasicAuthPass:            basicAuthPass,
		AuthDomain:               authDomain,
		AuthWorkstation:          authWorkstation,
		Keyword:                  keyword,
		Hostname:                 hostname,
		Port:                     port,
		DNSResolveServer:         dnsResolveServer,
		DNSResolveType:           dnsResolveType,
		MqttUsername:             mqttUsername,
		MqttPassword:             mqttPassword,
		MqttTopic:                mqttTopic,
		MqttSuccessMessage:       mqttSuccessMessage,
		DatabaseConnectionString: databaseConnectionString,
		DatabaseQuery:            databaseQuery,
		DockerContainer:          dockerContainer,
		DockerHost:               dockerHost,
		RadiusUsername:           radiusUsername,
		RadiusPassword:           radiusPassword,
		RadiusSecret:             radiusSecret,
		RadiusCalledStationID:    radiusCalledStationID,
		RadiusCallingStationID:   radiusCallingStationID,
	}
	return &out
}
