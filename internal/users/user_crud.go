package users

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
)

// CreateUser creates a new user via the LiteLLM API
func CreateUser(ctx context.Context, client *litellm.Client, req *UserCreateRequest) (*User, error) {
	user, err := litellm.SendRequestTyped[UserCreateRequest, User](ctx, client, http.MethodPost, "/user/new", req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

// GetUser retrieves user information by user_id
func GetUser(ctx context.Context, client *litellm.Client, userID string) (*User, error) {
	path := fmt.Sprintf("/user/info?user_id=%s", url.QueryEscape(userID))
	userResponse, err := litellm.SendRequestTyped[any, UserResponse](ctx, client, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Convert UserInfo to User and add computed fields
	user := &User{
		UserID:               userResponse.UserInfo.UserID,
		UserEmail:            userResponse.UserInfo.UserEmail,
		UserAlias:            userResponse.UserInfo.UserAlias,
		UserRole:             userResponse.UserInfo.UserRole,
		MaxBudget:            userResponse.UserInfo.MaxBudget,
		BudgetDuration:       userResponse.UserInfo.BudgetDuration,
		Models:               userResponse.UserInfo.Models,
		TPMLimit:             userResponse.UserInfo.TPMLimit,
		RPMLimit:             userResponse.UserInfo.RPMLimit,
		Metadata:             userResponse.UserInfo.Metadata,
		Spend:                userResponse.UserInfo.Spend,
		SSOUserID:            userResponse.UserInfo.SSOUserID,
		MaxParallelRequests:  userResponse.UserInfo.MaxParallelRequests,
		BudgetResetAt:        userResponse.UserInfo.BudgetResetAt,
		AllowedCacheControls: userResponse.UserInfo.AllowedCacheControls,
		ModelMaxBudget:       userResponse.UserInfo.ModelMaxBudget,
		CreatedAt:            userResponse.UserInfo.CreatedAt,
		UpdatedAt:            userResponse.UserInfo.UpdatedAt,
		KeyCount:             len(userResponse.Keys),
	}

	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(ctx context.Context, client *litellm.Client, req *UserUpdateRequest) (*User, error) {
	user, err := litellm.SendRequestTyped[UserUpdateRequest, User](ctx, client, http.MethodPost, "/user/update", req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return user, nil
}

// DeleteUser deletes a user by user_id
func DeleteUser(ctx context.Context, client *litellm.Client, userID string) error {
	deleteReq := &UserDeleteRequest{
		UserIDs: []string{userID},
	}

	// For delete operations, we don't need the response, just check for errors
	_, err := litellm.SendRequestTyped[UserDeleteRequest, int](ctx, client, http.MethodPost, "/user/delete", deleteReq)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
