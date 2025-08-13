package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// Mock HTTP server for testing
func createMockServer(t *testing.T, responses map[string]interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set content type
		w.Header().Set("Content-Type", "application/json")

		// Route based on path and method
		key := fmt.Sprintf("%s %s", r.Method, r.URL.Path)

		if response, exists := responses[key]; exists {
			if err, isError := response.(error); isError {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "endpoint not found"})
		}
	}))
}

func TestBuildUserCreateRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected *UserCreateRequest
	}{
		{
			name: "all fields populated",
			input: map[string]interface{}{
				"user_id":         "user123",
				"user_email":      "test@example.com",
				"user_alias":      "Test User",
				"user_role":       "internal_user",
				"max_budget":      100.50,
				"budget_duration": "monthly",
				"models":          []interface{}{"gpt-4", "gpt-3.5-turbo"},
				"tpm_limit":       1000,
				"rpm_limit":       60,
				"metadata": map[string]interface{}{
					"department": "engineering",
					"team":       "ai-platform",
				},
			},
			expected: &UserCreateRequest{
				UserID:         "user123",
				UserEmail:      "test@example.com",
				UserAlias:      "Test User",
				UserRole:       "internal_user",
				MaxBudget:      100.50,
				BudgetDuration: "monthly",
				Models:         []string{"gpt-4", "gpt-3.5-turbo"},
				TPMLimit:       1000,
				RPMLimit:       60,
				Metadata: map[string]interface{}{
					"department": "engineering",
					"team":       "ai-platform",
				},
				SendInviteEmail:      false,
				AutoCreateKey:        false,
				Blocked:              false,
				Aliases:              map[string]interface{}{},
				Config:               map[string]interface{}{},
				AllowedCacheControls: []string{},
				ModelMaxBudget:       map[string]interface{}{},
				ModelRPMLimit:        map[string]interface{}{},
				ModelTPMLimit:        map[string]interface{}{},
				Prompts:              []string{},
				Organizations:        []string{},
			},
		},
		{
			name: "minimal required fields",
			input: map[string]interface{}{
				"user_id": "user456",
			},
			expected: &UserCreateRequest{
				UserID:               "user456",
				Models:               []string{"no-default-models"},
				SendInviteEmail:      false,
				AutoCreateKey:        false,
				Blocked:              false,
				Aliases:              map[string]interface{}{},
				Config:               map[string]interface{}{},
				AllowedCacheControls: []string{},
				ModelMaxBudget:       map[string]interface{}{},
				ModelRPMLimit:        map[string]interface{}{},
				ModelTPMLimit:        map[string]interface{}{},
				Prompts:              []string{},
				Organizations:        []string{},
			},
		},
		{
			name: "partial fields",
			input: map[string]interface{}{
				"user_id":    "user789",
				"user_email": "partial@example.com",
				"user_role":  "internal_user",
				"max_budget": 50.0,
			},
			expected: &UserCreateRequest{
				UserID:               "user789",
				UserEmail:            "partial@example.com",
				UserRole:             "internal_user",
				MaxBudget:            50.0,
				Models:               []string{"no-default-models"},
				SendInviteEmail:      false,
				AutoCreateKey:        false,
				Blocked:              false,
				Aliases:              map[string]interface{}{},
				Config:               map[string]interface{}{},
				AllowedCacheControls: []string{},
				ModelMaxBudget:       map[string]interface{}{},
				ModelRPMLimit:        map[string]interface{}{},
				ModelTPMLimit:        map[string]interface{}{},
				Prompts:              []string{},
				Organizations:        []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the actual ResourceUser schema
			resource := ResourceUser()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildUserCreateRequest(d)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildUserCreateRequest() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestBuildUserUpdateRequest(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		input    map[string]interface{}
		expected *UserUpdateRequest
	}{
		{
			name:   "all updatable fields",
			userID: "user123",
			input: map[string]interface{}{
				"user_email": "updated@example.com",
				"user_role":  "proxy_admin",
				"max_budget": 200.0,
				"models":     []interface{}{"gpt-4", "claude-3.5-sonnet"},
			},
			expected: &UserUpdateRequest{
				UserID:    "user123",
				UserEmail: "updated@example.com",
				UserRole:  "proxy_admin",
				MaxBudget: 200.0,
				Models:    []string{"gpt-4", "claude-3.5-sonnet"},
			},
		},
		{
			name:   "partial update",
			userID: "user456",
			input: map[string]interface{}{
				"user_role":  "team",
				"max_budget": 75.0,
			},
			expected: &UserUpdateRequest{
				UserID:    "user456",
				UserRole:  "team",
				MaxBudget: 75.0,
			},
		},
		{
			name:   "minimal update",
			userID: "user789",
			input:  map[string]interface{}{},
			expected: &UserUpdateRequest{
				UserID: "user789",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceUser()
			d := schema.TestResourceDataRaw(t, resource.Schema, tt.input)

			result := buildUserUpdateRequest(d, tt.userID)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("buildUserUpdateRequest() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestSetUserResourceData(t *testing.T) {
	tests := []struct {
		name        string
		user        *User
		expectError bool
	}{
		{
			name: "complete user data",
			user: &User{
				UserID:         "user123",
				UserEmail:      "test@example.com",
				UserAlias:      "Test User",
				UserRole:       "internal_user",
				MaxBudget:      100.0,
				BudgetDuration: "monthly",
				Models:         []string{"gpt-4", "gpt-3.5-turbo"},
				TPMLimit:       1000,
				RPMLimit:       60,
				Metadata: map[string]interface{}{
					"department": "engineering",
				},
				Spend:    25.75,
				KeyCount: 3,
			},
			expectError: false,
		},
		{
			name: "minimal user data",
			user: &User{
				UserID:   "user456",
				Spend:    0.0,
				KeyCount: 0,
			},
			expectError: false,
		},
		{
			name: "user with empty slices and maps",
			user: &User{
				UserID:   "user789",
				Models:   []string{},
				Metadata: map[string]interface{}{},
				Spend:    10.5,
				KeyCount: 1,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := ResourceUser()
			d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})

			err := setUserResourceData(d, tt.user)

			if tt.expectError && err == nil {
				t.Errorf("setUserResourceData() expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("setUserResourceData() unexpected error: %v", err)
			}

			if err == nil {
				// Verify key fields were set correctly
				if d.Get("user_id") != tt.user.UserID {
					t.Errorf("Expected user_id %s, got %s", tt.user.UserID, d.Get("user_id"))
				}
				if d.Get("spend") != tt.user.Spend {
					t.Errorf("Expected spend %f, got %v", tt.user.Spend, d.Get("spend"))
				}
				if d.Get("key_count") != tt.user.KeyCount {
					t.Errorf("Expected key_count %d, got %v", tt.user.KeyCount, d.Get("key_count"))
				}

				// Check optional fields only if they have values
				if tt.user.UserEmail != "" && d.Get("user_email") != tt.user.UserEmail {
					t.Errorf("Expected user_email %s, got %s", tt.user.UserEmail, d.Get("user_email"))
				}
				if tt.user.UserRole != "" && d.Get("user_role") != tt.user.UserRole {
					t.Errorf("Expected user_role %s, got %s", tt.user.UserRole, d.Get("user_role"))
				}
			}
		})
	}
}

func TestUserRoleValidation(t *testing.T) {
	validRoles := []string{
		"proxy_admin",
		"proxy_admin_viewer",
		"internal_user",
		"internal_user_viewer",
	}

	invalidRoles := []string{
		"invalid_role",
		"admin",
		"user",
		"",
		"PROXY_ADMIN", // case sensitive
	}

	resource := ResourceUser()
	userRoleSchema := resource.Schema["user_role"]

	// Test valid roles
	for _, role := range validRoles {
		t.Run(fmt.Sprintf("valid_role_%s", role), func(t *testing.T) {
			_, errors := userRoleSchema.ValidateFunc(role, "user_role")
			if len(errors) > 0 {
				t.Errorf("Expected role %s to be valid, but got errors: %v", role, errors)
			}
		})
	}

	// Test invalid roles
	for _, role := range invalidRoles {
		t.Run(fmt.Sprintf("invalid_role_%s", role), func(t *testing.T) {
			_, errors := userRoleSchema.ValidateFunc(role, "user_role")
			if len(errors) == 0 {
				t.Errorf("Expected role %s to be invalid, but got no errors", role)
			}
		})
	}
}

func TestResourceUserSchema(t *testing.T) {
	resource := ResourceUser()
	schema := resource.Schema

	// Test optional fields (user_id is optional with computed since it auto-generates)
	optionalFields := []string{
		"user_id", "user_email", "user_alias", "user_role", "max_budget",
		"budget_duration", "models", "tpm_limit", "rpm_limit", "metadata",
	}
	for _, field := range optionalFields {
		t.Run(fmt.Sprintf("optional_field_%s", field), func(t *testing.T) {
			fieldSchema, exists := schema[field]
			if !exists {
				t.Errorf("Optional field %s not found in schema", field)
				return
			}
			if fieldSchema.Required {
				t.Errorf("Field %s should be optional", field)
			}
		})
	}

	// Test computed fields
	computedFields := []string{"user_id", "spend", "key_count", "created_at", "updated_at", "budget_reset_at"}
	for _, field := range computedFields {
		t.Run(fmt.Sprintf("computed_field_%s", field), func(t *testing.T) {
			fieldSchema, exists := schema[field]
			if !exists {
				t.Errorf("Computed field %s not found in schema", field)
				return
			}
			if !fieldSchema.Computed {
				t.Errorf("Field %s should be computed", field)
			}
		})
	}

	// Test ForceNew fields
	forceNewFields := []string{"user_id"}
	for _, field := range forceNewFields {
		t.Run(fmt.Sprintf("force_new_field_%s", field), func(t *testing.T) {
			fieldSchema, exists := schema[field]
			if !exists {
				t.Errorf("ForceNew field %s not found in schema", field)
				return
			}
			if !fieldSchema.ForceNew {
				t.Errorf("Field %s should have ForceNew set to true", field)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		request        *UserCreateRequest
		mockResponse   *User
		mockError      error
		expectError    bool
		expectedResult *User
	}{
		{
			name: "successful user creation",
			request: &UserCreateRequest{
				UserID:    "user123",
				UserEmail: "test@example.com",
				UserRole:  "internal_user",
			},
			mockResponse: &User{
				UserID:    "user123",
				UserEmail: "test@example.com",
				UserRole:  "internal_user",
				Spend:     0.0,
				KeyCount:  0,
			},
			expectError: false,
			expectedResult: &User{
				UserID:    "user123",
				UserEmail: "test@example.com",
				UserRole:  "internal_user",
				Spend:     0.0,
				KeyCount:  0,
			},
		},
		{
			name: "API error during creation",
			request: &UserCreateRequest{
				UserID: "user456",
			},
			mockError:   fmt.Errorf("API error: user already exists"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock responses
			responses := make(map[string]interface{})
			if tt.mockError != nil {
				responses["POST /user/new"] = tt.mockError
			} else {
				responses["POST /user/new"] = tt.mockResponse
			}

			// Create mock server
			server := createMockServer(t, responses)
			defer server.Close()

			// Create client
			client := litellm.NewClient(server.URL, "test-key", true)

			// Test the function
			result, err := CreateUser(context.Background(), client, tt.request)

			if tt.expectError {
				if err == nil {
					t.Errorf("CreateUser() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("CreateUser() unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, tt.expectedResult) {
					t.Errorf("CreateUser() = %+v, want %+v", result, tt.expectedResult)
				}
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockResponse   *UserResponse
		mockError      error
		expectError    bool
		expectedResult *User
	}{
		{
			name:   "successful user retrieval",
			userID: "user123",
			mockResponse: &UserResponse{
				UserID: "user123",
				UserInfo: UserInfo{
					UserID:    "user123",
					UserEmail: "test@example.com",
					UserRole:  "internal_user",
					Spend:     25.50,
				},
				Keys: []interface{}{
					map[string]interface{}{"key_id": "key1"},
					map[string]interface{}{"key_id": "key2"},
				},
			},
			expectError: false,
			expectedResult: &User{
				UserID:    "user123",
				UserEmail: "test@example.com",
				UserRole:  "internal_user",
				Spend:     25.50,
				KeyCount:  2,
			},
		},
		{
			name:        "API error during retrieval",
			userID:      "user456",
			mockError:   fmt.Errorf("API error: user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock responses
			responses := make(map[string]interface{})
			endpoint := "GET /user/info"
			if tt.mockError != nil {
				responses[endpoint] = tt.mockError
			} else {
				responses[endpoint] = tt.mockResponse
			}

			// Create mock server
			server := createMockServer(t, responses)
			defer server.Close()

			// Create client
			client := litellm.NewClient(server.URL, "test-key", true)

			// Test the function
			result, err := GetUser(context.Background(), client, tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("GetUser() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("GetUser() unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, tt.expectedResult) {
					t.Errorf("GetUser() = %+v, want %+v", result, tt.expectedResult)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		request        *UserUpdateRequest
		mockResponse   *User
		mockError      error
		expectError    bool
		expectedResult *User
	}{
		{
			name: "successful user update",
			request: &UserUpdateRequest{
				UserID:    "user123",
				UserEmail: "updated@example.com",
				UserRole:  "proxy_admin",
				MaxBudget: 200.0,
			},
			mockResponse: &User{
				UserID:    "user123",
				UserEmail: "updated@example.com",
				UserRole:  "proxy_admin",
				MaxBudget: 200.0,
				Spend:     25.50,
				KeyCount:  2,
			},
			expectError: false,
			expectedResult: &User{
				UserID:    "user123",
				UserEmail: "updated@example.com",
				UserRole:  "proxy_admin",
				MaxBudget: 200.0,
				Spend:     25.50,
				KeyCount:  2,
			},
		},
		{
			name: "API error during update",
			request: &UserUpdateRequest{
				UserID: "user456",
			},
			mockError:   fmt.Errorf("API error: user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock responses
			responses := make(map[string]interface{})
			if tt.mockError != nil {
				responses["POST /user/update"] = tt.mockError
			} else {
				responses["POST /user/update"] = tt.mockResponse
			}

			// Create mock server
			server := createMockServer(t, responses)
			defer server.Close()

			// Create client
			client := litellm.NewClient(server.URL, "test-key", true)

			// Test the function
			result, err := UpdateUser(context.Background(), client, tt.request)

			if tt.expectError {
				if err == nil {
					t.Errorf("UpdateUser() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("UpdateUser() unexpected error: %v", err)
				}
				if !reflect.DeepEqual(result, tt.expectedResult) {
					t.Errorf("UpdateUser() = %+v, want %+v", result, tt.expectedResult)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		mockError   error
		expectError bool
	}{
		{
			name:        "successful user deletion",
			userID:      "user123",
			expectError: false,
		},
		{
			name:        "API error during deletion",
			userID:      "user456",
			mockError:   fmt.Errorf("API error: user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock responses
			responses := make(map[string]interface{})
			if tt.mockError != nil {
				responses["POST /user/delete"] = tt.mockError
			} else {
				responses["POST /user/delete"] = 1
			}

			// Create mock server
			server := createMockServer(t, responses)
			defer server.Close()

			// Create client
			client := litellm.NewClient(server.URL, "test-key", true)

			// Test the function
			err := DeleteUser(context.Background(), client, tt.userID)

			if tt.expectError {
				if err == nil {
					t.Errorf("DeleteUser() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("DeleteUser() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestUserImporter(t *testing.T) {
	importer := UserImporter()

	if importer == nil {
		t.Errorf("UserImporter() returned nil")
		return
	}

	if importer.StateContext == nil {
		t.Errorf("UserImporter() StateContext is nil")
	}
}

func TestResourceUserCRUDOperations(t *testing.T) {
	// Test that the resource has all required CRUD operations
	resource := ResourceUser()

	if resource.CreateContext == nil {
		t.Errorf("ResourceUser() CreateContext is nil")
	}
	if resource.ReadContext == nil {
		t.Errorf("ResourceUser() ReadContext is nil")
	}
	if resource.UpdateContext == nil {
		t.Errorf("ResourceUser() UpdateContext is nil")
	}
	if resource.DeleteContext == nil {
		t.Errorf("ResourceUser() DeleteContext is nil")
	}
	if resource.Importer == nil {
		t.Errorf("ResourceUser() Importer is nil")
	}
	if resource.Schema == nil {
		t.Errorf("ResourceUser() Schema is nil")
	}
}
