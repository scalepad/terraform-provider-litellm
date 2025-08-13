package utils

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// expandStringList converts []interface{} to []string
func expandStringList(list []interface{}) []string {
	result := make([]string, len(list))
	for i, v := range list {
		if str, ok := v.(string); ok {
			result[i] = str
		}
	}
	return result
}

// GetValueDefault extracts a value from ResourceData with type assertion using the modern GetOk method
// For boolean types, it always includes the value even if false
func GetValueDefault[T any](d *schema.ResourceData, key string, keyData map[string]interface{}) {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, always get the value even if false
		keyData[key] = d.Get(key).(T)
	} else {
		// For other types, use GetOk to only include non-zero values
		if v, ok := d.GetOk(key); ok {
			keyData[key] = v.(T)
		}
	}
}

// GetStringListValue extracts a string list value from ResourceData using the modern GetOk method
func GetStringListValue(d *schema.ResourceData, key string, keyData map[string]interface{}) {
	if v, ok := d.GetOk(key); ok {
		keyData[key] = expandStringList(v.([]interface{}))
	}
}

// Helper functions to handle potential nil values from the API response with generics
// For boolean types, it returns the apiValue directly (no zero-value fallback)
func GetValueWithDefault[T comparable](apiValue, defaultValue T) T {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, return the actual value (including false)
		return apiValue
	}
	// For other types, use zero-value check
	if apiValue != zero {
		return apiValue
	}
	return defaultValue
}

// SetIfNotZero sets a value in ResourceData only if the API value is not zero,
// otherwise keeps the existing value from ResourceData
// For boolean types, it always sets the value (including false)
func SetIfNotZero[T comparable](d *schema.ResourceData, key string, apiValue T) {
	var zero T
	// Check if T is bool type
	if _, isBool := any(zero).(bool); isBool {
		// For booleans, always set the value (including false)
		d.Set(key, apiValue)
	} else {
		// For other types, only set if not zero
		if apiValue != zero {
			d.Set(key, apiValue)
		}
	}
	// If apiValue is zero for non-bool types, we don't set anything, keeping the existing value
}

// ShouldUseAPIValue determines if we should use the API value or preserve state
// This function helps decide whether to update a field based on the API response
func ShouldUseAPIValue(apiValue interface{}) bool {
	if apiValue == nil {
		return false
	}

	switch v := apiValue.(type) {
	case string:
		return v != ""
	case *string:
		return v != nil && *v != ""
	case []string:
		return len(v) > 0
	case *[]string:
		return v != nil && len(*v) > 0
	case map[string]interface{}:
		return len(v) > 0
	case *map[string]interface{}:
		return v != nil && len(*v) > 0
	case int:
		return true // Always use int values from API, including 0
	case *int:
		return v != nil
	case float64:
		return true // Always use float64 values from API, including 0.0
	case *float64:
		return v != nil
	case bool:
		return true // Always use bool values from API
	case *bool:
		return v != nil
	default:
		return apiValue != nil
	}
}

// Helper functions for creating pointers
// These functions return nil for zero values, which is useful for optional API fields

// StringPtr returns a pointer to the string value, or nil if the string is empty
func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// IntPtr returns a pointer to the int value, or nil if the int is zero
func IntPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

// FloatPtr returns a pointer to the float64 value, or nil if the float is zero
func FloatPtr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

// BoolPtr returns a pointer to the bool value (always returns a pointer, even for false)
func BoolPtr(b bool) *bool {
	return &b
}
