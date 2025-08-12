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
		UserID:               d.Get("user_id").(string),
		Aliases:              make(map[string]interface{}),
		Config:               make(map[string]interface{}),
		AllowedCacheControls: make([]string, 0),
		Guardrails:           make([]string, 0),
		Permissions:          make(map[string]interface{}),
		ModelMaxBudget:       make(map[string]interface{}),
		ModelRPMLimit:        make(map[string]interface{}),
		ModelTPMLimit:        make(map[string]interface{}),
		ObjectPermission:     make(map[string]interface{}),
		Prompts:              make([]string, 0),
		Organizations:        make([]string, 0),
	}

	// Basic string fields
	if v, ok := d.GetOk("user_email"); ok {
		req.UserEmail = v.(string)
	}
	if v, ok := d.GetOk("user_alias"); ok {
		req.UserAlias = v.(string)
	}
	if v, ok := d.GetOk("user_role"); ok {
		req.UserRole = v.(string)
	}
	if v, ok := d.GetOk("budget_duration"); ok {
		req.BudgetDuration = v.(string)
	}
	if v, ok := d.GetOk("duration"); ok {
		req.Duration = v.(string)
	}
	if v, ok := d.GetOk("key_alias"); ok {
		req.KeyAlias = v.(string)
	}
	if v, ok := d.GetOk("sso_user_id"); ok {
		req.SSOUserID = v.(string)
	}

	// Numeric fields
	if v, ok := d.GetOk("max_budget"); ok {
		req.MaxBudget = v.(float64)
	}
	if v, ok := d.GetOk("soft_budget"); ok {
		req.SoftBudget = v.(float64)
	}
	if v, ok := d.GetOk("tpm_limit"); ok {
		req.TPMLimit = v.(int)
	}
	if v, ok := d.GetOk("rpm_limit"); ok {
		req.RPMLimit = v.(int)
	}
	if v, ok := d.GetOk("max_parallel_requests"); ok {
		req.MaxParallelRequests = v.(int)
	}

	// Boolean fields - use d.Get() to include false values
	req.SendInviteEmail = d.Get("send_invite_email").(bool)
	req.AutoCreateKey = d.Get("auto_create_key").(bool)
	req.Blocked = d.Get("blocked").(bool)

	// Handle string slices
	if v, ok := d.GetOk("models"); ok {
		modelsList := v.([]interface{})
		models := make([]string, len(modelsList))
		for i, model := range modelsList {
			models[i] = model.(string)
		}
		req.Models = models
	}

	if v, ok := d.GetOk("allowed_cache_controls"); ok {
		controlsList := v.([]interface{})
		controls := make([]string, len(controlsList))
		for i, control := range controlsList {
			controls[i] = control.(string)
		}
		req.AllowedCacheControls = controls
	}

	if v, ok := d.GetOk("guardrails"); ok {
		guardrailsList := v.([]interface{})
		guardrails := make([]string, len(guardrailsList))
		for i, guardrail := range guardrailsList {
			guardrails[i] = guardrail.(string)
		}
		req.Guardrails = guardrails
	}

	if v, ok := d.GetOk("prompts"); ok {
		promptsList := v.([]interface{})
		prompts := make([]string, len(promptsList))
		for i, prompt := range promptsList {
			prompts[i] = prompt.(string)
		}
		req.Prompts = prompts
	}

	if v, ok := d.GetOk("organizations"); ok {
		orgsList := v.([]interface{})
		orgs := make([]string, len(orgsList))
		for i, org := range orgsList {
			orgs[i] = org.(string)
		}
		req.Organizations = orgs
	}

	// Handle maps
	if v, ok := d.GetOk("metadata"); ok {
		req.Metadata = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("aliases"); ok {
		req.Aliases = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("config"); ok {
		req.Config = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("permissions"); ok {
		req.Permissions = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_max_budget"); ok {
		req.ModelMaxBudget = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_rpm_limit"); ok {
		req.ModelRPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_tpm_limit"); ok {
		req.ModelTPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("object_permission"); ok {
		req.ObjectPermission = v.(map[string]interface{})
	}

	return req
}

// buildUserUpdateRequest builds a UserUpdateRequest from Terraform resource data
func buildUserUpdateRequest(d *schema.ResourceData, userID string) *UserUpdateRequest {
	req := &UserUpdateRequest{
		UserID: userID,
	}

	// Basic string fields
	if v, ok := d.GetOk("user_email"); ok {
		req.UserEmail = v.(string)
	}
	if v, ok := d.GetOk("user_alias"); ok {
		req.UserAlias = v.(string)
	}
	if v, ok := d.GetOk("user_role"); ok {
		req.UserRole = v.(string)
	}
	if v, ok := d.GetOk("budget_duration"); ok {
		req.BudgetDuration = v.(string)
	}
	if v, ok := d.GetOk("sso_user_id"); ok {
		req.SSOUserID = v.(string)
	}

	// Numeric fields
	if v, ok := d.GetOk("max_budget"); ok {
		req.MaxBudget = v.(float64)
	}
	if v, ok := d.GetOk("soft_budget"); ok {
		req.SoftBudget = v.(float64)
	}
	if v, ok := d.GetOk("tpm_limit"); ok {
		req.TPMLimit = v.(int)
	}
	if v, ok := d.GetOk("rpm_limit"); ok {
		req.RPMLimit = v.(int)
	}
	if v, ok := d.GetOk("max_parallel_requests"); ok {
		req.MaxParallelRequests = v.(int)
	}

	// Boolean fields - use d.Get() to include false values
	req.Blocked = d.Get("blocked").(bool)

	// Handle string slices
	if v, ok := d.GetOk("models"); ok {
		modelsList := v.([]interface{})
		models := make([]string, len(modelsList))
		for i, model := range modelsList {
			models[i] = model.(string)
		}
		req.Models = models
	}

	if v, ok := d.GetOk("allowed_cache_controls"); ok {
		controlsList := v.([]interface{})
		controls := make([]string, len(controlsList))
		for i, control := range controlsList {
			controls[i] = control.(string)
		}
		req.AllowedCacheControls = controls
	}

	if v, ok := d.GetOk("guardrails"); ok {
		guardrailsList := v.([]interface{})
		guardrails := make([]string, len(guardrailsList))
		for i, guardrail := range guardrailsList {
			guardrails[i] = guardrail.(string)
		}
		req.Guardrails = guardrails
	}

	if v, ok := d.GetOk("prompts"); ok {
		promptsList := v.([]interface{})
		prompts := make([]string, len(promptsList))
		for i, prompt := range promptsList {
			prompts[i] = prompt.(string)
		}
		req.Prompts = prompts
	}

	if v, ok := d.GetOk("organizations"); ok {
		orgsList := v.([]interface{})
		orgs := make([]string, len(orgsList))
		for i, org := range orgsList {
			orgs[i] = org.(string)
		}
		req.Organizations = orgs
	}

	// Handle maps
	if v, ok := d.GetOk("metadata"); ok {
		req.Metadata = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("aliases"); ok {
		req.Aliases = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("config"); ok {
		req.Config = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("permissions"); ok {
		req.Permissions = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_max_budget"); ok {
		req.ModelMaxBudget = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_rpm_limit"); ok {
		req.ModelRPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("model_tpm_limit"); ok {
		req.ModelTPMLimit = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("object_permission"); ok {
		req.ObjectPermission = v.(map[string]interface{})
	}

	return req
}

// setUserResourceData sets the user data in the Terraform resource using utility functions
func setUserResourceData(d *schema.ResourceData, user *User) error {
	// Always set the user_id as it's required
	if err := d.Set("user_id", user.UserID); err != nil {
		return err
	}

	// Use SetIfNotZero for optional string fields to preserve existing values when API returns zero values
	utils.SetIfNotZero(d, "user_email", user.UserEmail)
	utils.SetIfNotZero(d, "user_alias", user.UserAlias)
	utils.SetIfNotZero(d, "user_role", user.UserRole)
	utils.SetIfNotZero(d, "budget_duration", user.BudgetDuration)
	utils.SetIfNotZero(d, "duration", user.Duration)
	utils.SetIfNotZero(d, "key_alias", user.KeyAlias)
	utils.SetIfNotZero(d, "sso_user_id", user.SSOUserID)

	// Use SetIfNotZero for numeric fields
	utils.SetIfNotZero(d, "max_budget", user.MaxBudget)
	utils.SetIfNotZero(d, "soft_budget", user.SoftBudget)
	utils.SetIfNotZero(d, "tpm_limit", user.TPMLimit)
	utils.SetIfNotZero(d, "rpm_limit", user.RPMLimit)
	utils.SetIfNotZero(d, "max_parallel_requests", user.MaxParallelRequests)

	// Handle boolean fields - always set them (including false values)
	utils.SetIfNotZero(d, "send_invite_email", user.SendInviteEmail)
	utils.SetIfNotZero(d, "auto_create_key", user.AutoCreateKey)
	utils.SetIfNotZero(d, "blocked", user.Blocked)

	// Handle string slices - only set if not empty
	if len(user.Models) > 0 {
		if err := d.Set("models", user.Models); err != nil {
			return err
		}
	}

	if len(user.AllowedCacheControls) > 0 {
		if err := d.Set("allowed_cache_controls", user.AllowedCacheControls); err != nil {
			return err
		}
	}

	if len(user.Guardrails) > 0 {
		if err := d.Set("guardrails", user.Guardrails); err != nil {
			return err
		}
	}

	if len(user.Prompts) > 0 {
		if err := d.Set("prompts", user.Prompts); err != nil {
			return err
		}
	}

	if len(user.Organizations) > 0 {
		if err := d.Set("organizations", user.Organizations); err != nil {
			return err
		}
	}

	// Handle maps - only set if not empty
	if len(user.Metadata) > 0 {
		if err := d.Set("metadata", user.Metadata); err != nil {
			return err
		}
	}

	if len(user.Aliases) > 0 {
		if err := d.Set("aliases", user.Aliases); err != nil {
			return err
		}
	}

	if len(user.Config) > 0 {
		if err := d.Set("config", user.Config); err != nil {
			return err
		}
	}

	if len(user.Permissions) > 0 {
		if err := d.Set("permissions", user.Permissions); err != nil {
			return err
		}
	}

	if len(user.ModelMaxBudget) > 0 {
		if err := d.Set("model_max_budget", user.ModelMaxBudget); err != nil {
			return err
		}
	}

	if len(user.ModelRPMLimit) > 0 {
		if err := d.Set("model_rpm_limit", user.ModelRPMLimit); err != nil {
			return err
		}
	}

	if len(user.ModelTPMLimit) > 0 {
		if err := d.Set("model_tpm_limit", user.ModelTPMLimit); err != nil {
			return err
		}
	}

	if len(user.ObjectPermission) > 0 {
		if err := d.Set("object_permission", user.ObjectPermission); err != nil {
			return err
		}
	}

	// Always set computed fields
	if err := d.Set("spend", user.Spend); err != nil {
		return err
	}
	if err := d.Set("key_count", user.KeyCount); err != nil {
		return err
	}

	// Set computed timestamp fields
	utils.SetIfNotZero(d, "created_at", user.CreatedAt)
	utils.SetIfNotZero(d, "updated_at", user.UpdatedAt)
	utils.SetIfNotZero(d, "budget_reset_at", user.BudgetResetAt)

	return nil
}
