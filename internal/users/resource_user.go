package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

// ResourceUser returns the user resource
func ResourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer:      UserImporter(),
		Schema:        resourceUserSchema(),
	}
}

// resourceUserCreate creates a new user
func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	req := buildUserCreateRequest(d)

	// Generate UUIDv7 if user_id is not provided
	if req.UserID == "" {
		userUUID, err := uuid.NewV7()
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to generate user ID: %w", err))
		}
		req.UserID = userUUID.String()
	}

	user, err := CreateUser(ctx, client, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create user: %w", err))
	}

	d.SetId(req.UserID)

	if err := setUserResourceData(d, user); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set user resource data: %w", err))
	}

	return nil
}

// resourceUserRead reads user information
func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	userID := d.Id()
	user, err := GetUser(ctx, client, userID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read user: %w", err))
	}

	if user == nil {
		d.SetId("")
		return nil
	}

	if err := setUserResourceData(d, user); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set user resource data: %w", err))
	}

	return nil
}

// resourceUserUpdate updates an existing user
func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	userID := d.Id()
	req := buildUserUpdateRequest(d, userID)

	user, err := UpdateUser(ctx, client, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update user: %w", err))
	}

	if err := setUserResourceData(d, user); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set user resource data: %w", err))
	}

	return nil
}

// resourceUserDelete deletes a user
func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*litellm.Client)

	userID := d.Id()
	if err := DeleteUser(ctx, client, userID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete user: %w", err))
	}

	d.SetId("")
	return nil
}

// buildUserCreateRequest builds a UserCreateRequest from Terraform resource data
func buildUserCreateRequest(d *schema.ResourceData) *UserCreateRequest {
	req := &UserCreateRequest{
		UserID: d.Get("user_id").(string),
	}

	if v, ok := d.GetOk("user_email"); ok {
		req.UserEmail = v.(string)
	}
	if v, ok := d.GetOk("user_alias"); ok {
		req.UserAlias = v.(string)
	}
	if v, ok := d.GetOk("user_role"); ok {
		req.UserRole = v.(string)
	}
	if v, ok := d.GetOk("max_budget"); ok {
		req.MaxBudget = v.(float64)
	}
	if v, ok := d.GetOk("budget_duration"); ok {
		req.BudgetDuration = v.(string)
	}
	if v, ok := d.GetOk("tpm_limit"); ok {
		req.TPMLimit = v.(int)
	}
	if v, ok := d.GetOk("rpm_limit"); ok {
		req.RPMLimit = v.(int)
	}

	// Handle models slice
	if v, ok := d.GetOk("models"); ok {
		modelsList := v.([]interface{})
		models := make([]string, len(modelsList))
		for i, model := range modelsList {
			models[i] = model.(string)
		}
		req.Models = models
	}

	// Handle metadata map
	if v, ok := d.GetOk("metadata"); ok {
		metadataMap := v.(map[string]interface{})
		req.Metadata = metadataMap
	}

	return req
}

// buildUserUpdateRequest builds a UserUpdateRequest from Terraform resource data
func buildUserUpdateRequest(d *schema.ResourceData, userID string) *UserUpdateRequest {
	req := &UserUpdateRequest{
		UserID: userID,
	}

	if v, ok := d.GetOk("user_email"); ok {
		req.UserEmail = v.(string)
	}
	if v, ok := d.GetOk("user_role"); ok {
		req.UserRole = v.(string)
	}
	if v, ok := d.GetOk("max_budget"); ok {
		req.MaxBudget = v.(float64)
	}

	// Handle models slice
	if v, ok := d.GetOk("models"); ok {
		modelsList := v.([]interface{})
		models := make([]string, len(modelsList))
		for i, model := range modelsList {
			models[i] = model.(string)
		}
		req.Models = models
	}

	return req
}

// setUserResourceData sets the user data in the Terraform resource using utility functions
func setUserResourceData(d *schema.ResourceData, user *User) error {
	// Always set the user_id as it's required
	if err := d.Set("user_id", user.UserID); err != nil {
		return err
	}

	// Use SetIfNotZero for optional fields to preserve existing values when API returns zero values
	utils.SetIfNotZero(d, "user_email", user.UserEmail)
	utils.SetIfNotZero(d, "user_alias", user.UserAlias)
	utils.SetIfNotZero(d, "user_role", user.UserRole)
	utils.SetIfNotZero(d, "max_budget", user.MaxBudget)
	utils.SetIfNotZero(d, "budget_duration", user.BudgetDuration)
	utils.SetIfNotZero(d, "tpm_limit", user.TPMLimit)
	utils.SetIfNotZero(d, "rpm_limit", user.RPMLimit)

	// Handle models slice - only set if not empty
	if len(user.Models) > 0 {
		if err := d.Set("models", user.Models); err != nil {
			return err
		}
	}

	// Handle metadata map - only set if not empty
	if len(user.Metadata) > 0 {
		if err := d.Set("metadata", user.Metadata); err != nil {
			return err
		}
	}

	// Always set computed fields (spend and key_count)
	if err := d.Set("spend", user.Spend); err != nil {
		return err
	}
	if err := d.Set("key_count", user.KeyCount); err != nil {
		return err
	}

	return nil
}
