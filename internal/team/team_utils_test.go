package team

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func TestBuildTeamDataForUtils(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"team_alias":              "test-team",
				"organization_id":         "org123",
				"budget_duration":         "monthly",
				"tpm_limit":               1000,
				"rpm_limit":               60,
				"max_budget":              100.50,
				"team_member_budget":      25.75,
				"blocked":                 true,
				"metadata":                map[string]interface{}{"env": "prod", "region": "us-east-1"},
				"models":                  []interface{}{"gpt-4", "gpt-3.5-turbo"},
				"team_member_permissions": []interface{}{"read", "write", "admin"},
			},
			expected: map[string]interface{}{
				"team_alias":              "test-team",
				"organization_id":         "org123",
				"budget_duration":         "monthly",
				"tpm_limit":               1000,
				"rpm_limit":               60,
				"max_budget":              100.50,
				"team_member_budget":      25.75,
				"blocked":                 true,
				"metadata":                map[string]interface{}{"env": "prod", "region": "us-east-1"},
				"models":                  []string{"gpt-4", "gpt-3.5-turbo"},
				"team_member_permissions": []string{"read", "write", "admin"},
			},
		},
		{
			name:  "empty input with default boolean values",
			input: map[string]interface{}{
				// Even with empty input, booleans should have their default values
			},
			expected: map[string]interface{}{
				"blocked": false, // Default boolean value
			},
		},
		{
			name: "partial fields",
			input: map[string]interface{}{
				"team_alias": "partial-team",
				"max_budget": 50.0,
				"tpm_limit":  500,
				"blocked":    false,
				"models":     []interface{}{"gpt-4"},
			},
			expected: map[string]interface{}{
				"team_alias": "partial-team",
				"max_budget": 50.0,
				"tpm_limit":  500,
				"blocked":    false,
				"models":     []string{"gpt-4"},
			},
		},
		{
			name: "only string lists",
			input: map[string]interface{}{
				"models":                  []interface{}{"claude-3", "gpt-4"},
				"team_member_permissions": []interface{}{"read", "write"},
			},
			expected: map[string]interface{}{
				"models":                  []string{"claude-3", "gpt-4"},
				"team_member_permissions": []string{"read", "write"},
				"blocked":                 false, // Default boolean value
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceTeam schema
			resource := ResourceTeam()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildTeamDataForUtils(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildTeamDataForUtils() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildTeamDataForUtils() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}

			// Check that no unexpected keys are present
			for key := range result {
				if _, expected := tt.expected[key]; !expected {
					t.Errorf("buildTeamDataForUtils() unexpected key %s with value %v", key, result[key])
				}
			}
		})
	}
}

func TestSetTeamResourceData(t *testing.T) {
	tests := []struct {
		name        string
		team        *TeamResponse
		expectError bool
	}{
		{
			name: "complete team data",
			team: &TeamResponse{
				TeamID:                "team123",
				TeamAlias:             "test-team",
				OrganizationID:        "org456",
				TPMLimit:              1000,
				RPMLimit:              60,
				MaxBudget:             100.0,
				BudgetDuration:        "monthly",
				TeamMemberBudget:      25.0,
				Blocked:               true,
				TeamMemberPermissions: []string{"read", "write"},
				Metadata:              map[string]interface{}{"env": "test"},
				Models:                []string{"gpt-4", "gpt-3.5-turbo"},
			},
			expectError: false,
		},
		{
			name: "minimal team data",
			team: &TeamResponse{
				TeamID:    "team-minimal",
				TeamAlias: "minimal-team",
			},
			expectError: false,
		},
		{
			name: "team with nil maps and empty slices",
			team: &TeamResponse{
				TeamID:                "team-empty",
				TeamAlias:             "empty-team",
				Models:                []string{},
				Metadata:              nil,
				TeamMemberPermissions: nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceTeam()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setTeamResourceData(d, tt.team)

			if tt.expectError && err == nil {
				t.Errorf("setTeamResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setTeamResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify some key fields were set correctly
				if d.Get("team_alias") != tt.team.TeamAlias {
					t.Errorf("Expected team_alias %s, got %s", tt.team.TeamAlias, d.Get("team_alias"))
				}
				if tt.team.MaxBudget != 0 && d.Get("max_budget") != tt.team.MaxBudget {
					t.Errorf("Expected max_budget %f, got %v", tt.team.MaxBudget, d.Get("max_budget"))
				}
				if tt.team.OrganizationID != "" && d.Get("organization_id") != tt.team.OrganizationID {
					t.Errorf("Expected organization_id %s, got %s", tt.team.OrganizationID, d.Get("organization_id"))
				}
			}
		})
	}
}

func TestParseTeamAPIResponse(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		expected    *TeamResponse
		expectError bool
	}{
		{
			name: "complete API response",
			input: map[string]interface{}{
				"team_id":                 "team123",
				"team_alias":              "test-team",
				"organization_id":         "org456",
				"budget_duration":         "monthly",
				"tpm_limit":               float64(1000),
				"rpm_limit":               float64(60),
				"max_budget":              100.0,
				"team_member_budget":      25.0,
				"blocked":                 true,
				"metadata":                map[string]interface{}{"env": "test"},
				"models":                  []interface{}{"gpt-4", "gpt-3.5-turbo"},
				"team_member_permissions": []interface{}{"read", "write"},
			},
			expected: &TeamResponse{
				TeamID:                "team123",
				TeamAlias:             "test-team",
				OrganizationID:        "org456",
				BudgetDuration:        "monthly",
				TPMLimit:              1000,
				RPMLimit:              60,
				MaxBudget:             100.0,
				TeamMemberBudget:      25.0,
				Blocked:               true,
				Metadata:              map[string]interface{}{"env": "test"},
				Models:                []string{"gpt-4", "gpt-3.5-turbo"},
				TeamMemberPermissions: []string{"read", "write"},
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
			expected:    &TeamResponse{},
			expectError: false,
		},
		{
			name: "partial response",
			input: map[string]interface{}{
				"team_id":    "team-partial",
				"team_alias": "partial-team",
				"max_budget": 50.0,
				"blocked":    false,
			},
			expected: &TeamResponse{
				TeamID:    "team-partial",
				TeamAlias: "partial-team",
				MaxBudget: 50.0,
				Blocked:   false,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTeamAPIResponse(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("parseTeamAPIResponse() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("parseTeamAPIResponse() unexpected error: %v", err)
			}

			if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseTeamAPIResponse() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildTeamForCreation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *Team
	}{
		{
			name: "complete creation data",
			input: map[string]interface{}{
				"team_alias":              "creation-team",
				"organization_id":         "org789",
				"budget_duration":         "weekly",
				"tpm_limit":               500,
				"rpm_limit":               30,
				"max_budget":              75.0,
				"team_member_budget":      15.0,
				"blocked":                 false,
				"metadata":                map[string]interface{}{"env": "staging"},
				"models":                  []string{"claude-3", "gpt-4"},
				"team_member_permissions": []string{"read", "write", "admin"},
			},
			expected: &Team{
				TeamAlias:             "creation-team",
				OrganizationID:        "org789",
				BudgetDuration:        "weekly",
				TPMLimit:              500,
				RPMLimit:              30,
				MaxBudget:             75.0,
				TeamMemberBudget:      15.0,
				Blocked:               false,
				Metadata:              map[string]interface{}{"env": "staging"},
				Models:                []string{"claude-3", "gpt-4"},
				TeamMemberPermissions: []string{"read", "write", "admin"},
			},
		},
		{
			name:     "empty creation data",
			input:    map[string]interface{}{},
			expected: &Team{},
		},
		{
			name: "partial creation data",
			input: map[string]interface{}{
				"team_alias": "partial-creation",
				"max_budget": 100.0,
				"blocked":    true,
			},
			expected: &Team{
				TeamAlias: "partial-creation",
				MaxBudget: 100.0,
				Blocked:   true,
			},
		},
		{
			name: "unknown fields ignored",
			input: map[string]interface{}{
				"team_alias":    "unknown-fields",
				"unknown_field": "should_be_ignored",
				"max_budget":    25.0,
			},
			expected: &Team{
				TeamAlias: "unknown-fields",
				MaxBudget: 25.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTeamForCreation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildTeamForCreation() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildTeamData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		teamID   string
		expected map[string]interface{}
	}{
		{
			name: "complete team data with ID",
			input: map[string]interface{}{
				"team_alias":              "legacy-team",
				"organization_id":         "org123",
				"metadata":                map[string]interface{}{"env": "prod"},
				"tpm_limit":               1000,
				"rpm_limit":               60,
				"max_budget":              100.0,
				"budget_duration":         "monthly",
				"models":                  []interface{}{"gpt-4"},
				"blocked":                 true,
				"team_member_permissions": []interface{}{"read", "write"},
				"team_member_budget":      25.0,
			},
			teamID: "team456",
			expected: map[string]interface{}{
				"team_id":                 "team456",
				"team_alias":              "legacy-team",
				"organization_id":         "org123",
				"metadata":                map[string]interface{}{"env": "prod"},
				"tpm_limit":               1000,
				"rpm_limit":               60,
				"max_budget":              100.0,
				"budget_duration":         "monthly",
				"models":                  []interface{}{"gpt-4"},
				"blocked":                 true,
				"team_member_permissions": []interface{}{"read", "write"},
				"team_member_budget":      25.0,
			},
		},
		{
			name: "minimal team data",
			input: map[string]interface{}{
				"team_alias": "minimal-legacy",
			},
			teamID: "team789",
			expected: map[string]interface{}{
				"team_id":    "team789",
				"team_alias": "minimal-legacy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceTeam()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildTeamData(d, tt.teamID)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildTeamData() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}
