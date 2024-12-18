// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk/models/shared"
	"net/http"
)

var _ provider.Provider = &UkumawapiProvider{}

type UkumawapiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// UkumawapiProviderModel describes the provider data model.
type UkumawapiProviderModel struct {
	ServerURL            types.String `tfsdk:"server_url"`
	OAuth2PasswordBearer types.String `tfsdk:"o_auth2_password_bearer"`
}

func (p *UkumawapiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ukumawapi"
	resp.Version = p.version
}

func (p *UkumawapiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to http://localhost:8000/openapi.json)",
				Optional:            true,
				Required:            false,
			},
			"o_auth2_password_bearer": schema.StringAttribute{
				Sensitive: true,
				Optional:  true,
			},
		},
	}
}

func (p *UkumawapiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data UkumawapiProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "http://localhost:8000/openapi.json"
	}

	oAuth2PasswordBearer := new(string)
	if !data.OAuth2PasswordBearer.IsUnknown() && !data.OAuth2PasswordBearer.IsNull() {
		*oAuth2PasswordBearer = data.OAuth2PasswordBearer.ValueString()
	} else {
		oAuth2PasswordBearer = nil
	}
	security := shared.Security{
		OAuth2PasswordBearer: oAuth2PasswordBearer,
	}

	httpClient := http.DefaultClient
	httpClient.Transport = NewLoggingHTTPTransport(http.DefaultTransport)

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(httpClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *UkumawapiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewMaintenanceResource,
		NewMonitorResource,
		NewMonitorTagResource,
		NewPauseResource,
		NewPostIncidentResource,
		NewResumeResource,
		NewStatusPageResource,
		NewUserResource,
	}
}

func (p *UkumawapiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewMonitorDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UkumawapiProvider{
			version: version,
		}
	}
}
