// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/internal/hooks"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/internal/utils"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://localhost:8000/openapi.json",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type TerraformProviderUkumawapi struct {
	Users             *Users
	Settings          *Settings
	DataBase          *DataBase
	Monitor           *Monitor
	StatusPage        *StatusPage
	Maintenance       *Maintenance
	Tags              *Tags
	CertificationInfo *CertificationInfo
	Informations      *Informations
	PingAverage       *PingAverage
	Uptime            *Uptime
	Authentication    *Authentication

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*TerraformProviderUkumawapi)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *TerraformProviderUkumawapi) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *TerraformProviderUkumawapi {
	sdk := &TerraformProviderUkumawapi{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.1.0",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.460.1",
			UserAgent:         "speakeasy-sdk/go 0.0.1 2.460.1 0.1.0 github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Settings = newSettings(sdk.sdkConfiguration)

	sdk.DataBase = newDataBase(sdk.sdkConfiguration)

	sdk.Monitor = newMonitor(sdk.sdkConfiguration)

	sdk.StatusPage = newStatusPage(sdk.sdkConfiguration)

	sdk.Maintenance = newMaintenance(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.CertificationInfo = newCertificationInfo(sdk.sdkConfiguration)

	sdk.Informations = newInformations(sdk.sdkConfiguration)

	sdk.PingAverage = newPingAverage(sdk.sdkConfiguration)

	sdk.Uptime = newUptime(sdk.sdkConfiguration)

	sdk.Authentication = newAuthentication(sdk.sdkConfiguration)

	return sdk
}
