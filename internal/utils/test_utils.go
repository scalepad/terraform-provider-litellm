package utils

import (
	"reflect"
	"strconv"
)

// CompareMapValues compares two maps, handling type conversions that Terraform state performs
func CompareMapValues(actual, expected map[string]interface{}) bool {
	if len(actual) != len(expected) {
		return false
	}

	for key, expectedValue := range expected {
		actualValue, exists := actual[key]
		if !exists {
			return false
		}

		// Handle numeric type conversions (int, int64, float64)
		if !CompareNumericValues(actualValue, expectedValue) {
			return false
		}
	}

	return true
}

// CompareNumericValues handles comparison of numeric values that might have different types
func CompareNumericValues(actual, expected interface{}) bool {
	// Convert both values to float64 for comparison if they're numeric or string representations of numbers
	actualFloat, actualIsNumeric := ConvertToFloat64(actual)
	expectedFloat, expectedIsNumeric := ConvertToFloat64(expected)

	if actualIsNumeric && expectedIsNumeric {
		return actualFloat == expectedFloat
	}

	// If not both numeric, use regular comparison
	return reflect.DeepEqual(actual, expected)
}

// ConvertToFloat64 converts various numeric types and string representations to float64
func ConvertToFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case string:
		// Try to parse string as number
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// CompareValues compares two values, using special logic for maps
func CompareValues(actual, expected interface{}) bool {
	// Special handling for maps
	if actualMap, ok := actual.(map[string]interface{}); ok {
		if expectedMap, ok := expected.(map[string]interface{}); ok {
			return CompareMapValues(actualMap, expectedMap)
		}
	}

	// For non-map types, use regular DeepEqual
	return reflect.DeepEqual(actual, expected)
}
