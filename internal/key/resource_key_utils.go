package key

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func buildKeyData(d *schema.ResourceData) map[string]interface{} {
	keyData := make(map[string]interface{})

	// String list fields
	utils.GetStringListValue(d, "models", keyData)
	utils.GetStringListValue(d, "allowed_cache_controls", keyData)
	utils.GetStringListValue(d, "guardrails", keyData)
	utils.GetStringListValue(d, "tags", keyData)

	// Float64 fields
	utils.GetValueDefault[float64](d, "max_budget", keyData)
	utils.GetValueDefault[float64](d, "soft_budget", keyData)

	// String fields
	utils.GetValueDefault[string](d, "user_id", keyData)
	utils.GetValueDefault[string](d, "team_id", keyData)
	utils.GetValueDefault[string](d, "budget_duration", keyData)
	utils.GetValueDefault[string](d, "key_alias", keyData)
	utils.GetValueDefault[string](d, "duration", keyData)

	// Int fields
	utils.GetValueDefault[int](d, "max_parallel_requests", keyData)
	utils.GetValueDefault[int](d, "tpm_limit", keyData)
	utils.GetValueDefault[int](d, "rpm_limit", keyData)

	// Bool fields
	utils.GetValueDefault[bool](d, "blocked", keyData)
	utils.GetValueDefault[bool](d, "send_invite_email", keyData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "metadata", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "aliases", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "config", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "permissions", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "model_max_budget", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "model_rpm_limit", keyData)
	utils.GetValueDefault[map[string]interface{}](d, "model_tpm_limit", keyData)

	return keyData
}

func setKeyResourceData(d *schema.ResourceData, key *Key) error {
	fields := map[string]interface{}{
		"key":                    key.Key,
		"models":                 key.Models,
		"spend":                  key.Spend,
		"max_budget":             key.MaxBudget,
		"user_id":                key.UserID,
		"team_id":                key.TeamID,
		"max_parallel_requests":  key.MaxParallelRequests,
		"metadata":               key.Metadata,
		"tpm_limit":              key.TPMLimit,
		"rpm_limit":              key.RPMLimit,
		"budget_duration":        key.BudgetDuration,
		"allowed_cache_controls": key.AllowedCacheControls,
		"soft_budget":            key.SoftBudget,
		"key_alias":              key.KeyAlias,
		"duration":               key.Duration,
		"aliases":                key.Aliases,
		"config":                 key.Config,
		"permissions":            key.Permissions,
		"model_max_budget":       key.ModelMaxBudget,
		"model_rpm_limit":        key.ModelRPMLimit,
		"model_tpm_limit":        key.ModelTPMLimit,
		"guardrails":             key.Guardrails,
		"blocked":                key.Blocked,
		"tags":                   key.Tags,
		"send_invite_email":      key.SendInviteEmail,
	}

	for field, value := range fields {
		if err := d.Set(field, value); err != nil {
			log.Printf("[WARN] Error setting %s: %s", field, err)
			return fmt.Errorf("error setting %s: %s", field, err)
		}
	}

	return nil
}

func mapToKey(data map[string]interface{}) *Key {
	key := &Key{}
	for k, v := range data {
		switch k {
		case "key":
			key.Key = v.(string)
		case "models":
			key.Models = v.([]string)
		case "max_budget":
			key.MaxBudget = v.(float64)
		case "user_id":
			key.UserID = v.(string)
		case "team_id":
			key.TeamID = v.(string)
		case "max_parallel_requests":
			key.MaxParallelRequests = v.(int)
		case "metadata":
			key.Metadata = v.(map[string]interface{})
		case "tpm_limit":
			key.TPMLimit = v.(int)
		case "rpm_limit":
			key.RPMLimit = v.(int)
		case "budget_duration":
			key.BudgetDuration = v.(string)
		case "allowed_cache_controls":
			key.AllowedCacheControls = v.([]string)
		case "soft_budget":
			key.SoftBudget = v.(float64)
		case "key_alias":
			key.KeyAlias = v.(string)
		case "duration":
			key.Duration = v.(string)
		case "aliases":
			key.Aliases = v.(map[string]interface{})
		case "config":
			key.Config = v.(map[string]interface{})
		case "permissions":
			key.Permissions = v.(map[string]interface{})
		case "model_max_budget":
			key.ModelMaxBudget = v.(map[string]interface{})
		case "model_rpm_limit":
			key.ModelRPMLimit = v.(map[string]interface{})
		case "model_tpm_limit":
			key.ModelTPMLimit = v.(map[string]interface{})
		case "guardrails":
			key.Guardrails = v.([]string)
		case "blocked":
			key.Blocked = v.(bool)
		case "tags":
			key.Tags = v.([]string)
		case "send_invite_email":
			key.SendInviteEmail = v.(bool)
		}
	}
	return key
}

func buildKeyForCreation(data map[string]interface{}) *Key {
	return mapToKey(data)
}
