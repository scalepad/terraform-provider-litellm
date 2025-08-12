package mcp

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// buildMCPServerData builds MCP server data using utils methods
func buildMCPServerData(d *schema.ResourceData) map[string]interface{} {
	mcpData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "server_name", mcpData)
	utils.GetValueDefault[string](d, "alias", mcpData)
	utils.GetValueDefault[string](d, "description", mcpData)
	utils.GetValueDefault[string](d, "url", mcpData)
	utils.GetValueDefault[string](d, "transport", mcpData)
	utils.GetValueDefault[string](d, "spec_version", mcpData)
	utils.GetValueDefault[string](d, "auth_type", mcpData)
	utils.GetValueDefault[string](d, "command", mcpData)

	// String list fields
	utils.GetStringListValue(d, "mcp_access_groups", mcpData)
	utils.GetStringListValue(d, "args", mcpData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "env", mcpData)

	// Handle mcp_info nested structure
	if v, ok := d.GetOk("mcp_info"); ok {
		mcpInfoList := v.([]interface{})
		if len(mcpInfoList) > 0 {
			mcpInfoMap := mcpInfoList[0].(map[string]interface{})
			mcpInfoData := make(map[string]interface{})

			if serverName, exists := mcpInfoMap["server_name"]; exists {
				mcpInfoData["server_name"] = serverName
			}
			if description, exists := mcpInfoMap["description"]; exists {
				mcpInfoData["description"] = description
			}
			if logoURL, exists := mcpInfoMap["logo_url"]; exists {
				mcpInfoData["logo_url"] = logoURL
			}

			// Handle cost info
			if costInfoList, exists := mcpInfoMap["mcp_server_cost_info"]; exists {
				costInfos := costInfoList.([]interface{})
				if len(costInfos) > 0 {
					costInfoMap := costInfos[0].(map[string]interface{})
					costInfoData := make(map[string]interface{})

					if defaultCost, exists := costInfoMap["default_cost_per_query"]; exists {
						costInfoData["default_cost_per_query"] = defaultCost
					}
					if toolCosts, exists := costInfoMap["tool_name_to_cost_per_query"]; exists {
						costInfoData["tool_name_to_cost_per_query"] = toolCosts
					}

					mcpInfoData["mcp_server_cost_info"] = costInfoData
				}
			}

			mcpData["mcp_info"] = mcpInfoData
		}
	}

	return mcpData
}

// setMCPServerResourceData sets MCP server resource data using utils methods
func setMCPServerResourceData(d *schema.ResourceData, server *MCPServerResponse) error {
	fields := map[string]interface{}{
		"server_id":          server.ServerID,
		"server_name":        server.ServerName,
		"alias":              server.Alias,
		"description":        server.Description,
		"url":                server.URL,
		"transport":          server.Transport,
		"spec_version":       server.SpecVersion,
		"auth_type":          server.AuthType,
		"created_at":         server.CreatedAt,
		"created_by":         server.CreatedBy,
		"updated_at":         server.UpdatedAt,
		"updated_by":         server.UpdatedBy,
		"status":             server.Status,
		"last_health_check":  server.LastHealthCheck,
		"health_check_error": server.HealthCheckError,
		"command":            server.Command,
	}

	for field, value := range fields {
		utils.SetIfNotZero(d, field, value)
	}

	// Handle string lists
	if server.MCPAccessGroups != nil {
		d.Set("mcp_access_groups", server.MCPAccessGroups)
	}
	if server.Args != nil {
		d.Set("args", server.Args)
	}

	// Handle env map
	if server.Env != nil {
		d.Set("env", server.Env)
	}

	// Handle mcp_info nested structure
	if server.MCPInfo != nil {
		mcpInfoList := make([]map[string]interface{}, 1)
		mcpInfoMap := make(map[string]interface{})

		mcpInfoMap["server_name"] = server.MCPInfo.ServerName
		mcpInfoMap["description"] = server.MCPInfo.Description
		mcpInfoMap["logo_url"] = server.MCPInfo.LogoURL

		if server.MCPInfo.MCPServerCostInfo != nil {
			costInfoList := make([]map[string]interface{}, 1)
			costInfoMap := make(map[string]interface{})

			costInfoMap["default_cost_per_query"] = server.MCPInfo.MCPServerCostInfo.DefaultCostPerQuery
			if server.MCPInfo.MCPServerCostInfo.ToolNameToCostPerQuery != nil {
				costInfoMap["tool_name_to_cost_per_query"] = server.MCPInfo.MCPServerCostInfo.ToolNameToCostPerQuery
			}

			costInfoList[0] = costInfoMap
			mcpInfoMap["mcp_server_cost_info"] = costInfoList
		}

		mcpInfoList[0] = mcpInfoMap
		d.Set("mcp_info", mcpInfoList)
	}

	return nil
}

// buildMCPServerForCreation converts map data to MCPServer struct
func buildMCPServerForCreation(data map[string]interface{}) *MCPServer {
	server := &MCPServer{}

	if v, ok := data["server_name"].(string); ok {
		server.ServerName = v
	}
	if v, ok := data["alias"].(string); ok {
		server.Alias = v
	}
	if v, ok := data["description"].(string); ok {
		server.Description = v
	}
	if v, ok := data["url"].(string); ok {
		server.URL = v
	}
	if v, ok := data["transport"].(string); ok {
		server.Transport = v
	}
	if v, ok := data["spec_version"].(string); ok {
		server.SpecVersion = v
	}
	if v, ok := data["auth_type"].(string); ok {
		server.AuthType = v
	}
	if v, ok := data["command"].(string); ok {
		server.Command = v
	}

	if v, ok := data["mcp_access_groups"].([]string); ok {
		server.MCPAccessGroups = v
	}
	if v, ok := data["args"].([]string); ok {
		server.Args = v
	}
	if v, ok := data["env"].(map[string]string); ok {
		server.Env = v
	}

	// Handle mcp_info
	if v, ok := data["mcp_info"].(map[string]interface{}); ok {
		server.MCPInfo = &MCPInfo{}

		if serverName, exists := v["server_name"].(string); exists {
			server.MCPInfo.ServerName = serverName
		}
		if description, exists := v["description"].(string); exists {
			server.MCPInfo.Description = description
		}
		if logoURL, exists := v["logo_url"].(string); exists {
			server.MCPInfo.LogoURL = logoURL
		}

		if costInfo, exists := v["mcp_server_cost_info"].(map[string]interface{}); exists {
			server.MCPInfo.MCPServerCostInfo = &MCPServerCostInfo{}

			if defaultCost, exists := costInfo["default_cost_per_query"].(float64); exists {
				server.MCPInfo.MCPServerCostInfo.DefaultCostPerQuery = defaultCost
			}
			if toolCosts, exists := costInfo["tool_name_to_cost_per_query"].(map[string]float64); exists {
				server.MCPInfo.MCPServerCostInfo.ToolNameToCostPerQuery = toolCosts
			}
		}
	}

	return server
}

// parseMCPServerAPIResponse parses API response into MCPServerResponse
func parseMCPServerAPIResponse(resp map[string]interface{}) (*MCPServerResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	serverResp := &MCPServerResponse{}

	// Parse string fields
	if v, ok := resp["server_id"].(string); ok {
		serverResp.ServerID = v
	}
	if v, ok := resp["server_name"].(string); ok {
		serverResp.ServerName = v
	}
	if v, ok := resp["alias"].(string); ok {
		serverResp.Alias = v
	}
	if v, ok := resp["description"].(string); ok {
		serverResp.Description = v
	}
	if v, ok := resp["url"].(string); ok {
		serverResp.URL = v
	}
	if v, ok := resp["transport"].(string); ok {
		serverResp.Transport = v
	}
	if v, ok := resp["spec_version"].(string); ok {
		serverResp.SpecVersion = v
	}
	if v, ok := resp["auth_type"].(string); ok {
		serverResp.AuthType = v
	}
	if v, ok := resp["created_at"].(string); ok {
		serverResp.CreatedAt = v
	}
	if v, ok := resp["created_by"].(string); ok {
		serverResp.CreatedBy = v
	}
	if v, ok := resp["updated_at"].(string); ok {
		serverResp.UpdatedAt = v
	}
	if v, ok := resp["updated_by"].(string); ok {
		serverResp.UpdatedBy = v
	}
	if v, ok := resp["status"].(string); ok {
		serverResp.Status = v
	}
	if v, ok := resp["last_health_check"].(string); ok {
		serverResp.LastHealthCheck = v
	}
	if v, ok := resp["health_check_error"].(string); ok {
		serverResp.HealthCheckError = v
	}
	if v, ok := resp["command"].(string); ok {
		serverResp.Command = v
	}

	// Parse array fields
	if accessGroups, ok := resp["mcp_access_groups"].([]interface{}); ok {
		serverResp.MCPAccessGroups = make([]string, len(accessGroups))
		for i, group := range accessGroups {
			if s, ok := group.(string); ok {
				serverResp.MCPAccessGroups[i] = s
			}
		}
	}

	if args, ok := resp["args"].([]interface{}); ok {
		serverResp.Args = make([]string, len(args))
		for i, arg := range args {
			if s, ok := arg.(string); ok {
				serverResp.Args[i] = s
			}
		}
	}

	// Parse env map
	if env, ok := resp["env"].(map[string]interface{}); ok {
		serverResp.Env = make(map[string]string)
		for k, v := range env {
			if s, ok := v.(string); ok {
				serverResp.Env[k] = s
			}
		}
	}

	// Parse teams
	if teams, ok := resp["teams"].([]interface{}); ok {
		serverResp.Teams = make([]map[string]string, len(teams))
		for i, team := range teams {
			if teamMap, ok := team.(map[string]interface{}); ok {
				serverResp.Teams[i] = make(map[string]string)
				for k, v := range teamMap {
					if s, ok := v.(string); ok {
						serverResp.Teams[i][k] = s
					}
				}
			}
		}
	}

	// Parse mcp_info
	if mcpInfo, ok := resp["mcp_info"].(map[string]interface{}); ok {
		serverResp.MCPInfo = &MCPInfo{}

		if serverName, exists := mcpInfo["server_name"].(string); exists {
			serverResp.MCPInfo.ServerName = serverName
		}
		if description, exists := mcpInfo["description"].(string); exists {
			serverResp.MCPInfo.Description = description
		}
		if logoURL, exists := mcpInfo["logo_url"].(string); exists {
			serverResp.MCPInfo.LogoURL = logoURL
		}

		if costInfo, exists := mcpInfo["mcp_server_cost_info"].(map[string]interface{}); exists {
			serverResp.MCPInfo.MCPServerCostInfo = &MCPServerCostInfo{}

			if defaultCost, exists := costInfo["default_cost_per_query"].(float64); exists {
				serverResp.MCPInfo.MCPServerCostInfo.DefaultCostPerQuery = defaultCost
			}
			if toolCosts, exists := costInfo["tool_name_to_cost_per_query"].(map[string]interface{}); exists {
				serverResp.MCPInfo.MCPServerCostInfo.ToolNameToCostPerQuery = make(map[string]float64)
				for k, v := range toolCosts {
					if f, ok := v.(float64); ok {
						serverResp.MCPInfo.MCPServerCostInfo.ToolNameToCostPerQuery[k] = f
					}
				}
			}
		}
	}

	return serverResp, nil
}

// buildMCPServerRequestFromStruct converts MCPServer to MCPServerRequest
func buildMCPServerRequestFromStruct(server *MCPServer) *MCPServerRequest {
	return &MCPServerRequest{
		ServerID:        server.ServerID,
		ServerName:      server.ServerName,
		Alias:           server.Alias,
		Description:     server.Description,
		Transport:       server.Transport,
		SpecVersion:     server.SpecVersion,
		AuthType:        server.AuthType,
		URL:             server.URL,
		MCPInfo:         server.MCPInfo,
		MCPAccessGroups: server.MCPAccessGroups,
		Command:         server.Command,
		Args:            server.Args,
		Env:             server.Env,
	}
}
