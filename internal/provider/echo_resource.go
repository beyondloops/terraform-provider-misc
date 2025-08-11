// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &EchoResource{}

func NewEchoResource() resource.Resource {
	return &EchoResource{}
}

type EchoResource struct{}

type EchoResourceModel struct {
	ID     types.String  `tfsdk:"id"`
	Input  types.Dynamic `tfsdk:"input"`
	Output types.Dynamic `tfsdk:"output"`
}

func (r *EchoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_echo"
}

func (r *EchoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Echo resource that maintains input data in managed state",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource identifier",
			},
			"input": schema.DynamicAttribute{
				Required:            true,
				MarkdownDescription: "Input data to echo",
			},
			"output": schema.DynamicAttribute{
				Computed:            true,
				MarkdownDescription: "Echoed output data",
			},
		},
	}
}

func (r *EchoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EchoResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Output = data.Input
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EchoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EchoResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EchoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data EchoResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Output = data.Input
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EchoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
