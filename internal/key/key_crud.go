package key

import (
	"context"
	"fmt"
	"github.com/scalepad/terraform-provider-litellm/internal/litellm"
	"net/http"
)

func createKey(ctx context.Context, c *litellm.Client, key *Key) (*Key, error) {
	resp, err := c.SendRequest(ctx, http.MethodPost, "/key/generate", key)
	if err != nil {
		return nil, err
	}

	return parseKeyResponse(resp)
}

func getKey(ctx context.Context, c *litellm.Client, keyID string) (*Key, error) {
	resp, err := c.SendRequest(ctx, http.MethodGet, fmt.Sprintf("/key/info?key=%s", keyID), nil)
	if err != nil {
		return nil, err
	}

	return parseKeyResponse(resp)
}

func parseKeyResponse(resp map[string]interface{}) (*Key, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	createdKey := &Key{}

	for k, v := range resp {
		if v == nil {
			continue
		}

		switch k {
		case "key":
			if s, ok := v.(string); ok {
				createdKey.Key = s
			}
		case "models":
			if models, ok := v.([]interface{}); ok {
				createdKey.Models = make([]string, len(models))
				for i, model := range models {
					if s, ok := model.(string); ok {
						createdKey.Models[i] = s
					}
				}
			}
		case "spend":
			if f, ok := v.(float64); ok {
				createdKey.Spend = f
			}
		case "max_budget":
			if f, ok := v.(float64); ok {
				createdKey.MaxBudget = f
			}
		case "user_id":
			if s, ok := v.(string); ok {
				createdKey.UserID = s
			}
		case "team_id":
			if s, ok := v.(string); ok {
				createdKey.TeamID = s
			}
		case "max_parallel_requests":
			if i, ok := v.(float64); ok {
				createdKey.MaxParallelRequests = int(i)
			}
		case "metadata":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.Metadata = m
			}
		case "tpm_limit":
			if i, ok := v.(float64); ok {
				createdKey.TPMLimit = int(i)
			}
		case "rpm_limit":
			if i, ok := v.(float64); ok {
				createdKey.RPMLimit = int(i)
			}
		case "budget_duration":
			if s, ok := v.(string); ok {
				createdKey.BudgetDuration = s
			}
		case "soft_budget":
			if f, ok := v.(float64); ok {
				createdKey.SoftBudget = f
			}
		case "key_alias":
			if s, ok := v.(string); ok {
				createdKey.KeyAlias = s
			}
		case "duration":
			if s, ok := v.(string); ok {
				createdKey.Duration = s
			}
		case "aliases":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.Aliases = m
			}
		case "config":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.Config = m
			}
		case "permissions":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.Permissions = m
			}
		case "model_max_budget":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.ModelMaxBudget = m
			}
		case "model_rpm_limit":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.ModelRPMLimit = m
			}
		case "model_tpm_limit":
			if m, ok := v.(map[string]interface{}); ok {
				createdKey.ModelTPMLimit = m
			}
		case "guardrails":
			if guardrails, ok := v.([]interface{}); ok {
				createdKey.Guardrails = make([]string, len(guardrails))
				for i, guardrail := range guardrails {
					if s, ok := guardrail.(string); ok {
						createdKey.Guardrails[i] = s
					}
				}
			}
		case "blocked":
			if b, ok := v.(bool); ok {
				createdKey.Blocked = b
			}
		case "tags":
			if tags, ok := v.([]interface{}); ok {
				createdKey.Tags = make([]string, len(tags))
				for i, tag := range tags {
					if s, ok := tag.(string); ok {
						createdKey.Tags[i] = s
					}
				}
			}
		}
	}

	return createdKey, nil
}

func updateKey(ctx context.Context, c *litellm.Client, key *Key) (*Key, error) {
	// Create a new map with only the fields that can be updated
	updateData := map[string]interface{}{
		"key":                   key.Key,
		"models":                key.Models,
		"max_budget":            key.MaxBudget,
		"team_id":               key.TeamID,
		"max_parallel_requests": key.MaxParallelRequests,
		"metadata":              key.Metadata,
		"tpm_limit":             key.TPMLimit,
		"rpm_limit":             key.RPMLimit,
		"budget_duration":       key.BudgetDuration,
		"key_alias":             key.KeyAlias,
		"aliases":               key.Aliases,
		"permissions":           key.Permissions,
		"model_max_budget":      key.ModelMaxBudget,
		"model_rpm_limit":       key.ModelRPMLimit,
		"model_tpm_limit":       key.ModelTPMLimit,
		"guardrails":            key.Guardrails,
		"blocked":               key.Blocked,
	}

	resp, err := c.SendRequest(ctx, http.MethodPost, "/key/update", updateData)
	if err != nil {
		return nil, err
	}

	return parseKeyResponse(resp)
}

func deleteKey(ctx context.Context, c *litellm.Client, keyID string) error {
	payload := map[string]interface{}{
		"keys": []string{keyID},
	}
	_, err := c.SendRequest(ctx, http.MethodPost, "/key/delete", payload)
	return err
}
