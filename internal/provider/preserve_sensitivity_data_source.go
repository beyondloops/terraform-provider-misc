// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &PreserveSensitivityDataSource{}

func NewPreserveSensitivityDataSource() datasource.DataSource {
	return &PreserveSensitivityDataSource{}
}

type PreserveSensitivityDataSource struct {
	client *http.Client
}

type PreserveSensitivityDataSourceModel struct {
	Input types.Dynamic `tfsdk:"input"`
}

func (ps *PreserveSensitivityDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_preserve_sensitivity"
}

func (ps *PreserveSensitivityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Process arbitrary data while preserving sensitive markings",

		Attributes: map[string]schema.Attribute{
			"input": schema.DynamicAttribute{
				MarkdownDescription: "Arbitrary input data that will be preserved.",
				Required:            true,
			},
		},
	}
}

func (ps *PreserveSensitivityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))

		return
	}

	ps.client = client
}

// Read simply echos the input back as output, preserving sensitivity.
func (ps *PreserveSensitivityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PreserveSensitivityDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
