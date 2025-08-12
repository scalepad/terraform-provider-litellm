package member

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func TestBuildTeamMemberData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"team_id":            "team123",
				"user_id":            "user456",
				"user_email":         "user@example.com",
				"role":               "admin",
				"max_budget_in_team": 100.50,
			},
			expected: map[string]interface{}{
				"team_id":            "team123",
				"user_id":            "user456",
				"user_email":         "user@example.com",
				"role":               "admin",
				"max_budget_in_team": 100.50,
			},
		},
		{
			name: "minimal fields",
			input: map[string]interface{}{
				"team_id":    "team789",
				"user_id":    "user123",
				"user_email": "minimal@example.com",
				"role":       "user",
			},
			expected: map[string]interface{}{
				"team_id":    "team789",
				"user_id":    "user123",
				"user_email": "minimal@example.com",
				"role":       "user",
			},
		},
		{
			name: "with budget only",
			input: map[string]interface{}{
				"team_id":            "team999",
				"user_id":            "user999",
				"user_email":         "budget@example.com",
				"role":               "admin",
				"max_budget_in_team": 250.75,
			},
			expected: map[string]interface{}{
				"team_id":            "team999",
				"user_id":            "user999",
				"user_email":         "budget@example.com",
				"role":               "admin",
				"max_budget_in_team": 250.75,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceTeamMember schema
			resource := ResourceTeamMember()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildTeamMemberData(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildTeamMemberData() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildTeamMemberData() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}

			// Check that no unexpected keys are present
			for key := range result {
				if _, expected := tt.expected[key]; !expected {
					t.Errorf("buildTeamMemberData() unexpected key %s with value %v", key, result[key])
				}
			}
		})
	}
}

func TestSetTeamMemberResourceData(t *testing.T) {
	tests := []struct {
		name        string
		member      *TeamMemberResponse
		expectError bool
	}{
		{
			name: "complete team member data",
			member: &TeamMemberResponse{
				TeamID:          "team123",
				UserID:          "user456",
				UserEmail:       "user@example.com",
				Role:            "admin",
				MaxBudgetInTeam: 100.0,
				Status:          "active",
			},
			expectError: false,
		},
		{
			name: "minimal team member data",
			member: &TeamMemberResponse{
				TeamID:    "team-minimal",
				UserID:    "user-minimal",
				UserEmail: "minimal@example.com",
				Role:      "user",
			},
			expectError: false,
		},
		{
			name: "team member with zero budget",
			member: &TeamMemberResponse{
				TeamID:          "team-zero",
				UserID:          "user-zero",
				UserEmail:       "zero@example.com",
				Role:            "user",
				MaxBudgetInTeam: 0.0,
				Status:          "inactive",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceTeamMember()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setTeamMemberResourceData(d, tt.member)

			if tt.expectError && err == nil {
				t.Errorf("setTeamMemberResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setTeamMemberResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify some key fields were set correctly
				if d.Get("team_id") != tt.member.TeamID {
					t.Errorf("Expected team_id %s, got %s", tt.member.TeamID, d.Get("team_id"))
				}
				if d.Get("user_id") != tt.member.UserID {
					t.Errorf("Expected user_id %s, got %s", tt.member.UserID, d.Get("user_id"))
				}
				if d.Get("user_email") != tt.member.UserEmail {
					t.Errorf("Expected user_email %s, got %s", tt.member.UserEmail, d.Get("user_email"))
				}
				if d.Get("role") != tt.member.Role {
					t.Errorf("Expected role %s, got %s", tt.member.Role, d.Get("role"))
				}
			}
		})
	}
}

func TestBuildTeamMemberForCreation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *TeamMember
	}{
		{
			name: "complete creation data",
			input: map[string]interface{}{
				"team_id":            "team123",
				"user_id":            "user456",
				"user_email":         "user@example.com",
				"role":               "admin",
				"max_budget_in_team": 100.0,
			},
			expected: &TeamMember{
				TeamID:          "team123",
				UserID:          "user456",
				UserEmail:       "user@example.com",
				Role:            "admin",
				MaxBudgetInTeam: 100.0,
			},
		},
		{
			name:     "empty creation data",
			input:    map[string]interface{}{},
			expected: &TeamMember{},
		},
		{
			name: "partial creation data",
			input: map[string]interface{}{
				"team_id":    "team-partial",
				"user_email": "partial@example.com",
				"role":       "user",
			},
			expected: &TeamMember{
				TeamID:    "team-partial",
				UserEmail: "partial@example.com",
				Role:      "user",
			},
		},
		{
			name: "unknown fields ignored",
			input: map[string]interface{}{
				"team_id":       "team-unknown",
				"user_id":       "user-unknown",
				"unknown_field": "should_be_ignored",
				"role":          "admin",
			},
			expected: &TeamMember{
				TeamID: "team-unknown",
				UserID: "user-unknown",
				Role:   "admin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildTeamMemberForCreation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildTeamMemberForCreation() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestParseTeamMemberAPIResponse(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		expected    *TeamMemberResponse
		expectError bool
	}{
		{
			name: "complete API response",
			input: map[string]interface{}{
				"team_id":            "team123",
				"user_id":            "user456",
				"user_email":         "user@example.com",
				"role":               "admin",
				"max_budget_in_team": 100.0,
				"status":             "active",
			},
			expected: &TeamMemberResponse{
				TeamID:          "team123",
				UserID:          "user456",
				UserEmail:       "user@example.com",
				Role:            "admin",
				MaxBudgetInTeam: 100.0,
				Status:          "active",
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
			expected:    &TeamMemberResponse{},
			expectError: false,
		},
		{
			name: "partial response",
			input: map[string]interface{}{
				"team_id":    "team-partial",
				"user_email": "partial@example.com",
				"role":       "user",
			},
			expected: &TeamMemberResponse{
				TeamID:    "team-partial",
				UserEmail: "partial@example.com",
				Role:      "user",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTeamMemberAPIResponse(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("parseTeamMemberAPIResponse() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("parseTeamMemberAPIResponse() unexpected error: %v", err)
			}

			if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseTeamMemberAPIResponse() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildTeamMemberAddData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "complete team member add data",
			input: map[string]interface{}{
				"team_id": "team123",
				"member": []interface{}{
					map[string]interface{}{
						"user_id":    "user1",
						"user_email": "user1@example.com",
						"role":       "admin",
					},
					map[string]interface{}{
						"user_id":    "user2",
						"user_email": "user2@example.com",
						"role":       "user",
					},
				},
				"max_budget_in_team": 200.0,
			},
			expected: map[string]interface{}{
				"team_id": "team123",
				"member": []map[string]interface{}{
					{
						"user_id":    "user1",
						"user_email": "user1@example.com",
						"role":       "admin",
					},
					{
						"user_id":    "user2",
						"user_email": "user2@example.com",
						"role":       "user",
					},
				},
				"max_budget_in_team": 200.0,
			},
		},
		{
			name: "minimal team member add data",
			input: map[string]interface{}{
				"team_id": "team456",
				"member": []interface{}{
					map[string]interface{}{
						"user_email": "minimal@example.com",
						"role":       "user",
					},
				},
			},
			expected: map[string]interface{}{
				"team_id": "team456",
				"member": []map[string]interface{}{
					{
						"user_email": "minimal@example.com",
						"role":       "user",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceTeamMemberAdd schema
			resource := ResourceTeamMemberAdd()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildTeamMemberAddData(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildTeamMemberAddData() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildTeamMemberAddData() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestGetMemberKey(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected string
	}{
		{
			name: "user_id present",
			input: map[string]interface{}{
				"user_id":    "user123",
				"user_email": "user@example.com",
				"role":       "admin",
			},
			expected: "id:user123",
		},
		{
			name: "only user_email present",
			input: map[string]interface{}{
				"user_email": "user@example.com",
				"role":       "admin",
			},
			expected: "email:user@example.com",
		},
		{
			name: "empty user_id, use email",
			input: map[string]interface{}{
				"user_id":    "",
				"user_email": "user@example.com",
				"role":       "admin",
			},
			expected: "email:user@example.com",
		},
		{
			name: "no identifiers",
			input: map[string]interface{}{
				"role": "admin",
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMemberKey(tt.input)
			if result != tt.expected {
				t.Errorf("getMemberKey() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMemberAttributesChanged(t *testing.T) {
	tests := []struct {
		name      string
		oldMember map[string]interface{}
		newMember map[string]interface{}
		expected  bool
	}{
		{
			name: "role changed",
			oldMember: map[string]interface{}{
				"role": "user",
			},
			newMember: map[string]interface{}{
				"role": "admin",
			},
			expected: true,
		},
		{
			name: "role unchanged",
			oldMember: map[string]interface{}{
				"role": "admin",
			},
			newMember: map[string]interface{}{
				"role": "admin",
			},
			expected: false,
		},
		{
			name: "missing role in old",
			oldMember: map[string]interface{}{
				"user_id": "user123",
			},
			newMember: map[string]interface{}{
				"role": "admin",
			},
			expected: true,
		},
		{
			name: "missing role in new",
			oldMember: map[string]interface{}{
				"role": "admin",
			},
			newMember: map[string]interface{}{
				"user_id": "user123",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := memberAttributesChanged(tt.oldMember, tt.newMember)
			if result != tt.expected {
				t.Errorf("memberAttributesChanged() = %v, want %v", result, tt.expected)
			}
		})
	}
}
