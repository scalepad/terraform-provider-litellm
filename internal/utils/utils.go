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
