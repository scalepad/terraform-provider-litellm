package key

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func TestBuildKeyData(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"models":                 []interface{}{"gpt-4", "gpt-3.5-turbo"},
				"allowed_cache_controls": []interface{}{"no-cache", "max-age=3600"},
				"guardrails":             []interface{}{"content-filter", "safety-check"},
				"tags":                   []interface{}{"production", "api-key"},
				"max_budget":             100.50,
				"soft_budget":            80.25,
				"user_id":                "user123",
				"team_id":                "team456",
				"budget_duration":        "monthly",
				"key_alias":              "my-api-key",
				"duration":               "30d",
				"max_parallel_requests":  10,
				"tpm_limit":              1000,
				"rpm_limit":              60,
				"blocked":                true,
				"send_invite_email":      false,
				"metadata":               map[string]interface{}{"env": "prod"},
				"aliases":                map[string]interface{}{"alias1": "value1"},
				"config":                 map[string]interface{}{"timeout": "30s"},
				"permissions":            map[string]interface{}{"read": true, "write": false},
				"model_max_budget":       map[string]interface{}{"gpt-4": 50.0},
				"model_rpm_limit":        map[string]interface{}{"gpt-4": 30},
				"model_tpm_limit":        map[string]interface{}{"gpt-4": 500},
			},
			expected: map[string]interface{}{
				"models":                 []string{"gpt-4", "gpt-3.5-turbo"},
				"allowed_cache_controls": []string{"no-cache", "max-age=3600"},
				"guardrails":             []string{"content-filter", "safety-check"},
				"tags":                   []string{"production", "api-key"},
				"max_budget":             100.50,
				"soft_budget":            80.25,
				"user_id":                "user123",
				"team_id":                "team456",
				"budget_duration":        "monthly",
				"key_alias":              "my-api-key",
				"duration":               "30d",
				"max_parallel_requests":  10,
				"tpm_limit":              1000,
				"rpm_limit":              60,
				"blocked":                true,
				"send_invite_email":      false,
				"metadata":               map[string]interface{}{"env": "prod"},
				"aliases":                map[string]interface{}{"alias1": "value1"},
				"config":                 map[string]interface{}{"timeout": "30s"},
				"permissions":            map[string]interface{}{"read": 1, "write": 0},
				"model_max_budget":       map[string]interface{}{"gpt-4": 50.0},
				"model_rpm_limit":        map[string]interface{}{"gpt-4": 30},
				"model_tpm_limit":        map[string]interface{}{"gpt-4": 500},
			},
		},
		{
			name:  "empty input with default boolean values",
			input: map[string]interface{}{
				// Even with empty input, booleans should have their default values
			},
			expected: map[string]interface{}{
				"blocked":           false, // Default boolean value
				"send_invite_email": false, // Default boolean value
			},
		},
		{
			name: "partial fields",
			input: map[string]interface{}{
				"models":     []interface{}{"gpt-4"},
				"max_budget": 50.0,
				"user_id":    "user789",
				"blocked":    false,
			},
			expected: map[string]interface{}{
				"models":            []string{"gpt-4"},
				"max_budget":        50.0,
				"user_id":           "user789",
				"blocked":           false,
				"send_invite_email": false, // Boolean fields are always included
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceKey schema
			resource := ResourceKey()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildKeyData(d)

			// Check that all expected keys are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("buildKeyData() missing key %s", key)
					continue
				}
				if !utils.CompareValues(actualValue, expectedValue) {
					t.Errorf("buildKeyData() key %s = %v, want %v", key, actualValue, expectedValue)
				}
			}

			// Check that no unexpected keys are present
			for key := range result {
				if _, expected := tt.expected[key]; !expected {
					t.Errorf("buildKeyData() unexpected key %s with value %v", key, result[key])
				}
			}
		})
	}
}

func TestSetKeyResourceData(t *testing.T) {
	tests := []struct {
		name        string
		key         *Key
		expectError bool
	}{
		{
			name: "complete key data",
			key: &Key{
				Key:                  "sk-test123",
				Models:               []string{"gpt-4", "gpt-3.5-turbo"},
				Spend:                25.75,
				MaxBudget:            100.0,
				UserID:               "user123",
				TeamID:               "team456",
				MaxParallelRequests:  5,
				Metadata:             map[string]interface{}{"env": "test"},
				TPMLimit:             1000,
				RPMLimit:             60,
				BudgetDuration:       "monthly",
				AllowedCacheControls: []string{"no-cache"},
				SoftBudget:           80.0,
				KeyAlias:             "test-key",
				Duration:             "30d",
				Aliases:              map[string]interface{}{"alias1": "value1"},
				Config:               map[string]interface{}{"timeout": "30s"},
				Permissions:          map[string]interface{}{"read": "true"},
				ModelMaxBudget:       map[string]interface{}{"gpt-4": 50.0},
				ModelRPMLimit:        map[string]interface{}{"gpt-4": 30},
				ModelTPMLimit:        map[string]interface{}{"gpt-4": 500},
				Guardrails:           []string{"content-filter"},
				Blocked:              true,
				Tags:                 []string{"production"},
				SendInviteEmail:      false,
			},
			expectError: false,
		},
		{
			name: "minimal key data",
			key: &Key{
				Key: "sk-minimal",
			},
			expectError: false,
		},
		{
			name: "key with nil maps and empty slices",
			key: &Key{
				Key:                  "sk-empty",
				Models:               []string{},
				Metadata:             nil,
				AllowedCacheControls: nil,
				Guardrails:           []string{},
				Tags:                 nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceKey()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setKeyResourceData(d, tt.key)

			if tt.expectError && err == nil {
				t.Errorf("setKeyResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setKeyResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify some key fields were set correctly
				if d.Get("key") != tt.key.Key {
					t.Errorf("Expected key %s, got %s", tt.key.Key, d.Get("key"))
				}
				if tt.key.MaxBudget != 0 && d.Get("max_budget") != tt.key.MaxBudget {
					t.Errorf("Expected max_budget %f, got %v", tt.key.MaxBudget, d.Get("max_budget"))
				}
				if tt.key.UserID != "" && d.Get("user_id") != tt.key.UserID {
					t.Errorf("Expected user_id %s, got %s", tt.key.UserID, d.Get("user_id"))
				}
			}
		})
	}
}

func TestMapToKey(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *Key
	}{
		{
			name: "complete data map",
			input: map[string]interface{}{
				"key":                    "sk-test123",
				"models":                 []string{"gpt-4", "gpt-3.5-turbo"},
				"max_budget":             100.0,
				"user_id":                "user123",
				"team_id":                "team456",
				"max_parallel_requests":  5,
				"metadata":               map[string]interface{}{"env": "test"},
				"tpm_limit":              1000,
				"rpm_limit":              60,
				"budget_duration":        "monthly",
				"allowed_cache_controls": []string{"no-cache"},
				"soft_budget":            80.0,
				"key_alias":              "test-key",
				"duration":               "30d",
				"aliases":                map[string]interface{}{"alias1": "value1"},
				"config":                 map[string]interface{}{"timeout": "30s"},
				"permissions":            map[string]interface{}{"read": true},
				"model_max_budget":       map[string]interface{}{"gpt-4": 50.0},
				"model_rpm_limit":        map[string]interface{}{"gpt-4": 30},
				"model_tpm_limit":        map[string]interface{}{"gpt-4": 500},
				"guardrails":             []string{"content-filter"},
				"blocked":                true,
				"tags":                   []string{"production"},
				"send_invite_email":      false,
			},
			expected: &Key{
				Key:                  "sk-test123",
				Models:               []string{"gpt-4", "gpt-3.5-turbo"},
				MaxBudget:            100.0,
				UserID:               "user123",
				TeamID:               "team456",
				MaxParallelRequests:  5,
				Metadata:             map[string]interface{}{"env": "test"},
				TPMLimit:             1000,
				RPMLimit:             60,
				BudgetDuration:       "monthly",
				AllowedCacheControls: []string{"no-cache"},
				SoftBudget:           80.0,
				KeyAlias:             "test-key",
				Duration:             "30d",
				Aliases:              map[string]interface{}{"alias1": "value1"},
				Config:               map[string]interface{}{"timeout": "30s"},
				Permissions:          map[string]interface{}{"read": true},
				ModelMaxBudget:       map[string]interface{}{"gpt-4": 50.0},
				ModelRPMLimit:        map[string]interface{}{"gpt-4": 30},
				ModelTPMLimit:        map[string]interface{}{"gpt-4": 500},
				Guardrails:           []string{"content-filter"},
				Blocked:              true,
				Tags:                 []string{"production"},
				SendInviteEmail:      false,
			},
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: &Key{},
		},
		{
			name: "partial data",
			input: map[string]interface{}{
				"key":        "sk-partial",
				"models":     []string{"gpt-4"},
				"max_budget": 50.0,
				"blocked":    false,
			},
			expected: &Key{
				Key:       "sk-partial",
				Models:    []string{"gpt-4"},
				MaxBudget: 50.0,
				Blocked:   false,
			},
		},
		{
			name: "unknown fields ignored",
			input: map[string]interface{}{
				"key":           "sk-unknown",
				"unknown_field": "should_be_ignored",
				"max_budget":    25.0,
			},
			expected: &Key{
				Key:       "sk-unknown",
				MaxBudget: 25.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapToKey(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mapToKey() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildKeyForCreation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *Key
	}{
		{
			name: "creation data",
			input: map[string]interface{}{
				"models":     []string{"gpt-4"},
				"max_budget": 100.0,
				"user_id":    "user123",
				"blocked":    false,
			},
			expected: &Key{
				Models:    []string{"gpt-4"},
				MaxBudget: 100.0,
				UserID:    "user123",
				Blocked:   false,
			},
		},
		{
			name:     "empty creation data",
			input:    map[string]interface{}{},
			expected: &Key{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildKeyForCreation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildKeyForCreation() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}
