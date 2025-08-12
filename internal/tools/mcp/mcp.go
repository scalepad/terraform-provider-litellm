package mcp

// MCPServerCostInfo represents cost information for MCP server tools.
type MCPServerCostInfo struct {
	DefaultCostPerQuery    float64            `json:"default_cost_per_query,omitempty"`
	ToolNameToCostPerQuery map[string]float64 `json:"tool_name_to_cost_per_query,omitempty"`
}

// MCPInfo represents MCP server information and configuration.
type MCPInfo struct {
	ServerName        string             `json:"server_name,omitempty"`
	Description       string             `json:"description,omitempty"`
	LogoURL           string             `json:"logo_url,omitempty"`
	MCPServerCostInfo *MCPServerCostInfo `json:"mcp_server_cost_info,omitempty"`
}

// MCPServer represents an MCP server for input operations
type MCPServer struct {
	ServerID        string            `json:"server_id,omitempty"`
	ServerName      string            `json:"server_name"`
	Alias           string            `json:"alias,omitempty"`
	Description     string            `json:"description,omitempty"`
	Transport       string            `json:"transport"`
	SpecVersion     string            `json:"spec_version,omitempty"`
	AuthType        string            `json:"auth_type,omitempty"`
	URL             string            `json:"url"`
	MCPInfo         *MCPInfo          `json:"mcp_info,omitempty"`
	MCPAccessGroups []string          `json:"mcp_access_groups,omitempty"`
	Command         string            `json:"command,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
}

// MCPServerRequest represents a request to create or update an MCP server.
type MCPServerRequest struct {
	ServerID        string            `json:"server_id,omitempty"`
	ServerName      string            `json:"server_name"`
	Alias           string            `json:"alias,omitempty"`
	Description     string            `json:"description,omitempty"`
	Transport       string            `json:"transport"`
	SpecVersion     string            `json:"spec_version,omitempty"`
	AuthType        string            `json:"auth_type,omitempty"`
	URL             string            `json:"url"`
	MCPInfo         *MCPInfo          `json:"mcp_info,omitempty"`
	MCPAccessGroups []string          `json:"mcp_access_groups,omitempty"`
	Command         string            `json:"command,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
}

// MCPServerResponse represents a response from the API containing MCP server information.
type MCPServerResponse struct {
	ServerID         string              `json:"server_id"`
	ServerName       string              `json:"server_name"`
	Alias            string              `json:"alias,omitempty"`
	Description      string              `json:"description,omitempty"`
	URL              string              `json:"url"`
	Transport        string              `json:"transport"`
	SpecVersion      string              `json:"spec_version,omitempty"`
	AuthType         string              `json:"auth_type,omitempty"`
	CreatedAt        string              `json:"created_at,omitempty"`
	CreatedBy        string              `json:"created_by,omitempty"`
	UpdatedAt        string              `json:"updated_at,omitempty"`
	UpdatedBy        string              `json:"updated_by,omitempty"`
	Teams            []map[string]string `json:"teams,omitempty"`
	MCPAccessGroups  []string            `json:"mcp_access_groups,omitempty"`
	MCPInfo          *MCPInfo            `json:"mcp_info,omitempty"`
	Status           string              `json:"status,omitempty"`
	LastHealthCheck  string              `json:"last_health_check,omitempty"`
	HealthCheckError string              `json:"health_check_error,omitempty"`
	Command          string              `json:"command,omitempty"`
	Args             []string            `json:"args,omitempty"`
	Env              map[string]string   `json:"env,omitempty"`
}
