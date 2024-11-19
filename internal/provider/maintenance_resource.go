// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/sdk"
	"github.com/pete-leese/terraform-provider-ukumawapi/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MaintenanceResource{}
var _ resource.ResourceWithImportState = &MaintenanceResource{}

func NewMaintenanceResource() resource.Resource {
	return &MaintenanceResource{}
}

// MaintenanceResource defines the resource implementation.
type MaintenanceResource struct {
	client *sdk.TerraformProviderUkumawapi
}

// MaintenanceResourceModel describes the resource data model.
type MaintenanceResourceModel struct {
	Active      types.Bool     `tfsdk:"active"`
	Data        types.String   `tfsdk:"data"`
	DateRange   []types.String `tfsdk:"date_range"`
	DaysOfMonth []types.String `tfsdk:"days_of_month"`
	Description types.String   `tfsdk:"description"`
	IntervalDay types.Int64    `tfsdk:"interval_day"`
	Strategy    types.String   `tfsdk:"strategy"`
	TimeRange   []types.String `tfsdk:"time_range"`
	Title       types.String   `tfsdk:"title"`
	Weekdays    []types.String `tfsdk:"weekdays"`
}

func (r *MaintenanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_maintenance"
}

func (r *MaintenanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Maintenance Resource",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(true),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Default: true; Requires replacement if changed.`,
			},
			"data": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"date_range": schema.ListAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				ElementType: types.StringType,
				Description: `Requires replacement if changed.`,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"days_of_month": schema.ListAttribute{
				Computed: true,
				Optional: true,
				Default:  listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				ElementType: types.StringType,
				Description: `Requires replacement if changed.`,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Default:  stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Default: ""; Requires replacement if changed.`,
			},
			"interval_day": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Default:  int64default.StaticInt64(1),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Default: 1; Requires replacement if changed.`,
			},
			"strategy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Enumerate maintenance strategies. must be one of ["manual", "single", "recurring-interval", "recurring-weekday", "recurring-day-of-month", "cron"]; Requires replacement if changed.`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"manual",
						"single",
						"recurring-interval",
						"recurring-weekday",
						"recurring-day-of-month",
						"cron",
					),
				},
			},
			"time_range": schema.ListAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				ElementType: types.StringType,
				Description: `Requires replacement if changed.`,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"title": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Requires replacement if changed.`,
			},
			"weekdays": schema.ListAttribute{
				Computed: true,
				Optional: true,
				Default:  listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				ElementType: types.StringType,
				Description: `Requires replacement if changed.`,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
		},
	}
}

func (r *MaintenanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.TerraformProviderUkumawapi)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.TerraformProviderUkumawapi, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MaintenanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MaintenanceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedMaintenance()
	res, err := r.client.Maintenance.CreateMaintenanceMaintenancePost(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Any != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromInterface(res.Any)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MaintenanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MaintenanceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MaintenanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MaintenanceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MaintenanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MaintenanceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *MaintenanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource maintenance.")
}