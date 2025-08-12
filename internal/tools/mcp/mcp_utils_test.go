package mcp

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func TestBuildMCPServerData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"server_name":       "test-server",
				"alias":             "test-alias",
				"description":       "Test MCP server",
				"url":               "https://api.example.com/mcp",
				"transport":         "http",
				"spec_version":      "2024-11-05",
				"auth_type":         "bearer",
				"command":           "python3",
				"mcp_access_groups": []interface{}{"group1", "group2"},
				"args":              []interface{}{"--port", "8080"},
				"env": map[string]interface{}{
					"DEBUG":     "true",
					"LOG_LEVEL": "info",
				},
				"mcp_info": []interface{}{
					map[string]interface{}{
						"server_name": "Test Server Info",
						"description": "Server info description",
						"logo_url":    "https://example.com/logo.png",
						"mcp_server_cost_info": []interface{}{
							map[string]interface{}{
								"default_cost_per_query": 0.01,
								"tool_name_to_cost_per_query": map[string]interface{}{
									"search": 0.02,
									"write":  0.05,
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"server_name":       "test-server",
				"alias":             "test-alias",
				"description":       "Test MCP server",
				"url":               "https://api.example.com/mcp",
				"transport":         "http",
				"spec_version":      "2024-11-05",
				"auth_type":         "bearer",
				"command":           "python3",
				"mcp_access_groups": []string{"group1", "group2"},
				"args":              []string{"--port", "8080"},
				"env": map[string]interface{}{
					"DEBUG":     "true",
					"LOG_LEVEL": "info",
				},
				"mcp_info": map[string]interface{}{
					"server_name": "Test Server Info",
					"description": "Server info description",
					"logo_url":    "https://example.com/logo.png",
					"mcp_server_cost_info": map[string]interface{}{
						"default_cost_per_query": 0.01,
						"tool_name_to_cost_per_query": map[string]interface{}{
							"search": 0.02,
							"write":  0.05,
						},
					},
				},
			},
		},
		{
			name: "minimal fields",
			input: map[string]interface{}{
				"server_name": "minimal-server",
				"url":         "https://minimal.example.com",
				"transport":   "stdio",
			},
			expected: map[string]interface{}{
				"server_name":  "minimal-server",
				"url":          "https://minimal.example.com",
				"transport":    "stdio",
				"spec_version": "2024-11-05", // Default value from schema
				"auth_type":    "none",       // Default value from schema
			},
		},
		{
			name: "with empty lists and maps",
			input: map[string]interface{}{
				"server_name":       "empty-server",
				"url":               "https://empty.example.com",
				"transport":         "sse",
				"mcp_access_groups": []interface{}{},
				"args":              []interface{}{},
				"env":               map[string]interface{}{},
			},
			expected: map[string]interface{}{
				"server_name":  "empty-server",
				"url":          "https://empty.example.com",
				"transport":    "sse",
				"spec_version": "2024-11-05", // Default value from schema
				"auth_type":    "none",       // Default value from schema
				// Empty lists and maps are not included by utils functions
			},
		},
		{
			name: "only string lists",
			input: map[string]interface{}{
				"server_name":       "list-server",
				"url":               "https://list.example.com",
				"transport":         "http",
				"mcp_access_groups": []interface{}{"admin", "user"},
				"args":              []interface{}{"--verbose", "--debug"},
			},
			expected: map[string]interface{}{
				"server_name":       "list-server",
				"url":               "https://list.example.com",
				"transport":         "http",
				"spec_version":      "2024-11-05", // Default value from schema
				"auth_type":         "none",       // Default value from schema
				"mcp_access_groups": []string{"admin", "user"},
				"args":              []string{"--verbose", "--debug"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceLiteLLMMCPServer schema
			resource := ResourceLiteLLMMCPServer()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildMCPServerData(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildMCPServerData() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildMCPServerData() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}

			// Check that no unexpected keys are present
			for key := range result {
				if _, expected := tt.expected[key]; !expected {
					t.Errorf("buildMCPServerData() unexpected key %s with value %v", key, result[key])
				}
			}
		})
	}
}

func TestSetMCPServerResourceData(t *testing.T) {
	tests := []struct {
		name        string
		server      *MCPServerResponse
		expectError bool
	}{
		{
			name: "complete server data",
			server: &MCPServerResponse{
				ServerID:         "server123",
				ServerName:       "test-server",
				Alias:            "test-alias",
				Description:      "Test MCP server",
				URL:              "https://api.example.com/mcp",
				Transport:        "http",
				SpecVersion:      "2024-11-05",
				AuthType:         "bearer",
				CreatedAt:        "2024-01-01T00:00:00Z",
				CreatedBy:        "user123",
				UpdatedAt:        "2024-01-02T00:00:00Z",
				UpdatedBy:        "user456",
				Status:           "active",
				LastHealthCheck:  "2024-01-02T12:00:00Z",
				HealthCheckError: "",
				Command:          "python3",
				MCPAccessGroups:  []string{"group1", "group2"},
				Args:             []string{"--port", "8080"},
				Env:              map[string]string{"DEBUG": "true"},
				MCPInfo: &MCPInfo{
					ServerName:  "Test Server Info",
					Description: "Server info description",
					LogoURL:     "https://example.com/logo.png",
					MCPServerCostInfo: &MCPServerCostInfo{
						DefaultCostPerQuery: 0.01,
						ToolNameToCostPerQuery: map[string]float64{
							"search": 0.02,
							"write":  0.05,
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "minimal server data",
			server: &MCPServerResponse{
				ServerID:   "server-minimal",
				ServerName: "minimal-server",
				URL:        "https://minimal.example.com",
				Transport:  "stdio",
			},
			expectError: false,
		},
		{
			name: "server with nil maps and empty slices",
			server: &MCPServerResponse{
				ServerID:        "server-empty",
				ServerName:      "empty-server",
				URL:             "https://empty.example.com",
				Transport:       "sse",
				MCPAccessGroups: []string{},
				Args:            []string{},
				Env:             nil,
				MCPInfo:         nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceLiteLLMMCPServer()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setMCPServerResourceData(d, tt.server)

			if tt.expectError && err == nil {
				t.Errorf("setMCPServerResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setMCPServerResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify some key fields were set correctly
				if d.Get("server_name") != tt.server.ServerName {
					t.Errorf("Expected server_name %s, got %s", tt.server.ServerName, d.Get("server_name"))
				}
				if d.Get("url") != tt.server.URL {
					t.Errorf("Expected url %s, got %s", tt.server.URL, d.Get("url"))
				}
				if d.Get("transport") != tt.server.Transport {
					t.Errorf("Expected transport %s, got %s", tt.server.Transport, d.Get("transport"))
				}
				if tt.server.ServerID != "" && d.Get("server_id") != tt.server.ServerID {
					t.Errorf("Expected server_id %s, got %s", tt.server.ServerID, d.Get("server_id"))
				}
			}
		})
	}
}

func TestParseMCPServerAPIResponse(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		expected    *MCPServerResponse
		expectError bool
	}{
		{
			name: "complete API response",
			input: map[string]interface{}{
				"server_id":          "server123",
				"server_name":        "test-server",
				"alias":              "test-alias",
				"description":        "Test MCP server",
				"url":                "https://api.example.com/mcp",
				"transport":          "http",
				"spec_version":       "2024-11-05",
				"auth_type":          "bearer",
				"created_at":         "2024-01-01T00:00:00Z",
				"created_by":         "user123",
				"updated_at":         "2024-01-02T00:00:00Z",
				"updated_by":         "user456",
				"status":             "active",
				"last_health_check":  "2024-01-02T12:00:00Z",
				"health_check_error": "",
				"command":            "python3",
				"mcp_access_groups":  []interface{}{"group1", "group2"},
				"args":               []interface{}{"--port", "8080"},
				"env": map[string]interface{}{
					"DEBUG":     "true",
					"LOG_LEVEL": "info",
				},
				"teams": []interface{}{
					map[string]interface{}{
						"team_id":   "team1",
						"team_name": "Team One",
					},
				},
				"mcp_info": map[string]interface{}{
					"server_name": "Test Server Info",
					"description": "Server info description",
					"logo_url":    "https://example.com/logo.png",
					"mcp_server_cost_info": map[string]interface{}{
						"default_cost_per_query": 0.01,
						"tool_name_to_cost_per_query": map[string]interface{}{
							"search": 0.02,
							"write":  0.05,
						},
					},
				},
			},
			expected: &MCPServerResponse{
				ServerID:         "server123",
				ServerName:       "test-server",
				Alias:            "test-alias",
				Description:      "Test MCP server",
				URL:              "https://api.example.com/mcp",
				Transport:        "http",
				SpecVersion:      "2024-11-05",
				AuthType:         "bearer",
				CreatedAt:        "2024-01-01T00:00:00Z",
				CreatedBy:        "user123",
				UpdatedAt:        "2024-01-02T00:00:00Z",
				UpdatedBy:        "user456",
				Status:           "active",
				LastHealthCheck:  "2024-01-02T12:00:00Z",
				HealthCheckError: "",
				Command:          "python3",
				MCPAccessGroups:  []string{"group1", "group2"},
				Args:             []string{"--port", "8080"},
				Env: map[string]string{
					"DEBUG":     "true",
					"LOG_LEVEL": "info",
				},
				Teams: []map[string]string{
					{
						"team_id":   "team1",
						"team_name": "Team One",
					},
				},
				MCPInfo: &MCPInfo{
					ServerName:  "Test Server Info",
					Description: "Server info description",
					LogoURL:     "https://example.com/logo.png",
					MCPServerCostInfo: &MCPServerCostInfo{
						DefaultCostPerQuery: 0.01,
						ToolNameToCostPerQuery: map[string]float64{
							"search": 0.02,
							"write":  0.05,
						},
					},
				},
			},
			expectError: false,
		},
		{
			name:        "nil response",
			input:       nil,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty response",
			input:       map[string]interface{}{},
			expected:    &MCPServerResponse{},
			expectError: false,
		},
		{
			name: "partial response",
			input: map[string]interface{}{
				"server_id":   "server-partial",
				"server_name": "partial-server",
				"url":         "https://partial.example.com",
				"transport":   "stdio",
			},
			expected: &MCPServerResponse{
				ServerID:   "server-partial",
				ServerName: "partial-server",
				URL:        "https://partial.example.com",
				Transport:  "stdio",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseMCPServerAPIResponse(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("parseMCPServerAPIResponse() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("parseMCPServerAPIResponse() unexpected error: %v", err)
			}

			if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseMCPServerAPIResponse() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildMCPServerForCreation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *MCPServer
	}{
		{
			name: "complete creation data",
			input: map[string]interface{}{
				"server_name":       "creation-server",
				"alias":             "creation-alias",
				"description":       "Creation test server",
				"url":               "https://creation.example.com",
				"transport":         "http",
				"spec_version":      "2024-11-05",
				"auth_type":         "bearer",
				"command":           "node",
				"mcp_access_groups": []string{"admin", "user"},
				"args":              []string{"server.js", "--port", "3000"},
				"env": map[string]string{
					"NODE_ENV": "production",
					"DEBUG":    "false",
				},
				"mcp_info": map[string]interface{}{
					"server_name": "Creation Server Info",
					"description": "Creation server description",
					"logo_url":    "https://creation.example.com/logo.png",
					"mcp_server_cost_info": map[string]interface{}{
						"default_cost_per_query": 0.02,
						"tool_name_to_cost_per_query": map[string]float64{
							"analyze": 0.03,
							"process": 0.04,
						},
					},
				},
			},
			expected: &MCPServer{
				ServerName:      "creation-server",
				Alias:           "creation-alias",
				Description:     "Creation test server",
				URL:             "https://creation.example.com",
				Transport:       "http",
				SpecVersion:     "2024-11-05",
				AuthType:        "bearer",
				Command:         "node",
				MCPAccessGroups: []string{"admin", "user"},
				Args:            []string{"server.js", "--port", "3000"},
				Env: map[string]string{
					"NODE_ENV": "production",
					"DEBUG":    "false",
				},
				MCPInfo: &MCPInfo{
					ServerName:  "Creation Server Info",
					Description: "Creation server description",
					LogoURL:     "https://creation.example.com/logo.png",
					MCPServerCostInfo: &MCPServerCostInfo{
						DefaultCostPerQuery: 0.02,
						ToolNameToCostPerQuery: map[string]float64{
							"analyze": 0.03,
							"process": 0.04,
						},
					},
				},
			},
		},
		{
			name:     "empty creation data",
			input:    map[string]interface{}{},
			expected: &MCPServer{},
		},
		{
			name: "partial creation data",
			input: map[string]interface{}{
				"server_name": "partial-creation",
				"url":         "https://partial.example.com",
				"transport":   "stdio",
				"command":     "python3",
			},
			expected: &MCPServer{
				ServerName: "partial-creation",
				URL:        "https://partial.example.com",
				Transport:  "stdio",
				Command:    "python3",
			},
		},
		{
			name: "unknown fields ignored",
			input: map[string]interface{}{
				"server_name":   "unknown-fields",
				"url":           "https://unknown.example.com",
				"transport":     "http",
				"unknown_field": "should_be_ignored",
			},
			expected: &MCPServer{
				ServerName: "unknown-fields",
				URL:        "https://unknown.example.com",
				Transport:  "http",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildMCPServerForCreation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildMCPServerForCreation() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildMCPServerRequestFromStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    *MCPServer
		expected *MCPServerRequest
	}{
		{
			name: "complete server struct",
			input: &MCPServer{
				ServerID:        "server123",
				ServerName:      "request-server",
				Alias:           "request-alias",
				Description:     "Request test server",
				URL:             "https://request.example.com",
				Transport:       "sse",
				SpecVersion:     "2024-11-05",
				AuthType:        "basic",
				Command:         "java",
				MCPAccessGroups: []string{"dev", "test"},
				Args:            []string{"-jar", "server.jar"},
				Env: map[string]string{
					"JAVA_OPTS": "-Xmx1g",
				},
				MCPInfo: &MCPInfo{
					ServerName:  "Request Server Info",
					Description: "Request server description",
					LogoURL:     "https://request.example.com/logo.png",
				},
			},
			expected: &MCPServerRequest{
				ServerID:        "server123",
				ServerName:      "request-server",
				Alias:           "request-alias",
				Description:     "Request test server",
				URL:             "https://request.example.com",
				Transport:       "sse",
				SpecVersion:     "2024-11-05",
				AuthType:        "basic",
				Command:         "java",
				MCPAccessGroups: []string{"dev", "test"},
				Args:            []string{"-jar", "server.jar"},
				Env: map[string]string{
					"JAVA_OPTS": "-Xmx1g",
				},
				MCPInfo: &MCPInfo{
					ServerName:  "Request Server Info",
					Description: "Request server description",
					LogoURL:     "https://request.example.com/logo.png",
				},
			},
		},
		{
			name:     "empty server struct",
			input:    &MCPServer{},
			expected: &MCPServerRequest{},
		},
		{
			name: "minimal server struct",
			input: &MCPServer{
				ServerName: "minimal-request",
				URL:        "https://minimal-request.example.com",
				Transport:  "stdio",
			},
			expected: &MCPServerRequest{
				ServerName: "minimal-request",
				URL:        "https://minimal-request.example.com",
				Transport:  "stdio",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildMCPServerRequestFromStruct(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildMCPServerRequestFromStruct() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}
