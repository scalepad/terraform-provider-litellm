package mcp

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

func ResourceLiteLLMMCPServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLiteLLMMCPServerCreate,
		ReadContext:   resourceLiteLLMMCPServerRead,
		UpdateContext: resourceLiteLLMMCPServerUpdate,
		DeleteContext: resourceLiteLLMMCPServerDelete,

		Schema: mcpServerSchema(),
	}
}

// mcpServerSchema returns the schema for the MCP server resource
func mcpServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the MCP server",
		},
		"alias": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Alias for the MCP server",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the MCP server",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "URL of the MCP server",
		},
		"transport": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"http",
				"sse",
				"stdio",
			}, false),
			Description: "Transport type for the MCP server (http, sse, stdio)",
		},
		"spec_version": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "2024-11-05",
			Description: "MCP specification version",
		},
		"auth_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "none",
			ValidateFunc: validation.StringInSlice([]string{
				"none",
				"bearer",
				"basic",
			}, false),
			Description: "Authentication type (none, bearer, basic)",
		},
		"mcp_access_groups": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "List of access groups for the MCP server",
		},
		"command": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Command to run for stdio transport",
		},
		"args": {
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Arguments for the command (stdio transport)",
		},
		"env": {
			Type:        schema.TypeMap,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Environment variables for the command (stdio transport)",
		},
		"mcp_info": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "MCP server information and configuration",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"server_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Server name in MCP info",
					},
					"description": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Description in MCP info",
					},
					"logo_url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Logo URL for the MCP server",
					},
					"mcp_server_cost_info": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "Cost information for MCP server tools",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"default_cost_per_query": {
									Type:        schema.TypeFloat,
									Optional:    true,
									Description: "Default cost per query",
								},
								"tool_name_to_cost_per_query": {
									Type:        schema.TypeMap,
									Optional:    true,
									Elem:        &schema.Schema{Type: schema.TypeFloat},
									Description: "Map of tool names to their cost per query",
								},
							},
						},
					},
				},
			},
		},
		// Read-only computed fields
		"server_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Unique identifier for the MCP server",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the server was created",
		},
		"created_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who created the server",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the server was last updated",
		},
		"updated_by": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User who last updated the server",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Current status of the MCP server",
		},
		"last_health_check": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of the last health check",
		},
		"health_check_error": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Error message from the last health check, if any",
		},
	}
}

// resourceLiteLLMMCPServerCreate creates a new MCP server
func resourceLiteLLMMCPServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	// Build MCP server data from schema
	mcpData := buildMCPServerData(d)
	mcpServer := buildMCPServerForCreation(mcpData)

	// Create the MCP server
	serverResp, err := createMCPServer(ctx, client, mcpServer)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create MCP server: %w", err))
	}

	// Set resource ID
	d.SetId(serverResp.ServerID)

	// Update resource data
	if err := setMCPServerResourceData(d, serverResp); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set MCP server resource data: %w", err))
	}

	return nil
}

// resourceLiteLLMMCPServerRead reads an existing MCP server
func resourceLiteLLMMCPServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	serverID := d.Id()

	// Get the MCP server
	serverResp, err := getMCPServer(ctx, client, serverID)
	if err != nil {
		// Handle not found case
		if err.Error() == "mcp_server_not_found" {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("failed to read MCP server: %w", err))
	}

	// Update resource data
	if err := setMCPServerResourceData(d, serverResp); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set MCP server resource data: %w", err))
	}

	return nil
}

// resourceLiteLLMMCPServerUpdate updates an existing MCP server
func resourceLiteLLMMCPServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	// Build MCP server data from schema
	mcpData := buildMCPServerData(d)
	mcpServer := buildMCPServerForCreation(mcpData)
	mcpServer.ServerID = d.Id() // Ensure we include the server ID for updates

	// Update the MCP server
	serverResp, err := updateMCPServer(ctx, client, mcpServer)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update MCP server: %w", err))
	}

	// Update resource data
	if err := setMCPServerResourceData(d, serverResp); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set MCP server resource data: %w", err))
	}

	return nil
}

// resourceLiteLLMMCPServerDelete deletes an existing MCP server
func resourceLiteLLMMCPServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	serverID := d.Id()

	// Delete the MCP server
	if err := deleteMCPServer(ctx, client, serverID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete MCP server: %w", err))
	}

	// Clear resource ID
	d.SetId("")

	return nil
}
