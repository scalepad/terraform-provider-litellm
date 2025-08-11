package litellm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetValueWithDefault(t *testing.T) {
	tests := []struct {
		name         string
		apiValue     interface{}
		defaultValue interface{}
		expected     interface{}
	}{
		// String tests
		{
			name:         "string with non-empty api value",
			apiValue:     "api_value",
			defaultValue: "default_value",
			expected:     "api_value",
		},
		{
			name:         "string with empty api value",
			apiValue:     "",
			defaultValue: "default_value",
			expected:     "default_value",
		},
		// Int tests
		{
			name:         "int with non-zero api value",
			apiValue:     42,
			defaultValue: 10,
			expected:     42,
		},
		{
			name:         "int with zero api value",
			apiValue:     0,
			defaultValue: 10,
			expected:     10,
		},
		// Float64 tests
		{
			name:         "float64 with non-zero api value",
			apiValue:     3.14,
			defaultValue: 2.71,
			expected:     3.14,
		},
		{
			name:         "float64 with zero api value",
			apiValue:     0.0,
			defaultValue: 2.71,
			expected:     2.71,
		},
		// Bool tests
		{
			name:         "bool with true api value",
			apiValue:     true,
			defaultValue: false,
			expected:     true,
		},
		{
			name:         "bool with false api value",
			apiValue:     false,
			defaultValue: true,
			expected:     true, // Note: for bools, zero value (false) falls back to default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.apiValue.(type) {
			case string:
				result := GetValueWithDefault(v, tt.defaultValue.(string))
				if result != tt.expected {
					t.Errorf("GetValueWithDefault() = %v, want %v", result, tt.expected)
				}
			case int:
				result := GetValueWithDefault(v, tt.defaultValue.(int))
				if result != tt.expected {
					t.Errorf("GetValueWithDefault() = %v, want %v", result, tt.expected)
				}
			case float64:
				result := GetValueWithDefault(v, tt.defaultValue.(float64))
				if result != tt.expected {
					t.Errorf("GetValueWithDefault() = %v, want %v", result, tt.expected)
				}
			case bool:
				// For bool, we test GetValueWithDefault
				result := GetValueWithDefault(v, tt.defaultValue.(bool))
				if result != tt.expected {
					t.Errorf("GetValueWithDefault() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestSetIfNotZero(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		apiValue     interface{}
		initialValue interface{}
		expectedSet  bool
	}{
		// String tests
		{
			name:         "string with non-empty api value",
			key:          "test_string",
			apiValue:     "new_value",
			initialValue: "old_value",
			expectedSet:  true,
		},
		{
			name:         "string with empty api value",
			key:          "test_string",
			apiValue:     "",
			initialValue: "old_value",
			expectedSet:  false,
		},
		// Int tests
		{
			name:         "int with non-zero api value",
			key:          "test_int",
			apiValue:     42,
			initialValue: 10,
			expectedSet:  true,
		},
		{
			name:         "int with zero api value",
			key:          "test_int",
			apiValue:     0,
			initialValue: 10,
			expectedSet:  false,
		},
		// Float64 tests
		{
			name:         "float64 with non-zero api value",
			key:          "test_float",
			apiValue:     3.14,
			initialValue: 2.71,
			expectedSet:  true,
		},
		{
			name:         "float64 with zero api value",
			key:          "test_float",
			apiValue:     0.0,
			initialValue: 2.71,
			expectedSet:  false,
		},
		// Bool tests
		{
			name:         "bool with true api value",
			key:          "test_bool",
			apiValue:     true,
			initialValue: false,
			expectedSet:  true,
		},
		{
			name:         "bool with false api value",
			key:          "test_bool",
			apiValue:     false,
			initialValue: true,
			expectedSet:  false, // false is zero value for bool, so it shouldn't set
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock ResourceData
			resourceSchema := map[string]*schema.Schema{
				tt.key: {
					Type: getSchemaType(tt.apiValue),
				},
			}

			d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
				tt.key: tt.initialValue,
			})

			// Call SetIfNotZero with the appropriate type
			switch v := tt.apiValue.(type) {
			case string:
				SetIfNotZero(d, tt.key, v)
			case int:
				SetIfNotZero(d, tt.key, v)
			case float64:
				SetIfNotZero(d, tt.key, v)
			case bool:
				SetIfNotZero(d, tt.key, v)
			}

			// Check the result
			actualValue := d.Get(tt.key)
			if tt.expectedSet {
				if actualValue != tt.apiValue {
					t.Errorf("SetIfNotZero() should have set %v, but got %v", tt.apiValue, actualValue)
				}
			} else {
				if actualValue != tt.initialValue {
					t.Errorf("SetIfNotZero() should have kept initial value %v, but got %v", tt.initialValue, actualValue)
				}
			}
		})
	}
}

// Helper function to get the appropriate schema type for testing
func getSchemaType(value interface{}) schema.ValueType {
	switch value.(type) {
	case string:
		return schema.TypeString
	case int:
		return schema.TypeInt
	case float64:
		return schema.TypeFloat
	case bool:
		return schema.TypeBool
	default:
		return schema.TypeString
	}
}
