package litellm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	APIBase    string
	APIKey     string
	httpClient *http.Client
}

// NewClient creates a new Client.
//
// WARNING: Setting insecureSkipVerify to true disables TLS certificate verification.
// This should ONLY be used in development environments or with proper security justification.
// Disabling certificate verification exposes you to man-in-the-middle attacks and other security risks.
func NewClient(apiBase, apiKey string, insecureSkipVerify bool) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	return &Client{
		APIBase:    apiBase,
		APIKey:     apiKey,
		httpClient: &http.Client{Transport: tr},
	}
}

// Team-related methods
func (c *Client) CreateTeam(team map[string]interface{}) (map[string]interface{}, error) {
	return c.sendRequest("POST", "/team/new", team)
}

func (c *Client) GetTeam(teamID string) (map[string]interface{}, error) {
	return c.sendRequest("GET", fmt.Sprintf("/team/info?team_id=%s", teamID), nil)
}

func (c *Client) UpdateTeam(team map[string]interface{}) (map[string]interface{}, error) {
	return c.sendRequest("POST", "/team/update", team)
}

func (c *Client) DeleteTeam(teamID string) error {
	payload := map[string]interface{}{
		"team_ids": []string{teamID},
	}
	_, err := c.sendRequest("POST", "/team/delete", payload)
	return err
}

// Key-related methods
func (c *Client) CreateKey(key *Key) (*Key, error) {
	resp, err := c.sendRequest("POST", "/key/generate", key)
	if err != nil {
		return nil, err
	}

	return c.parseKeyResponse(resp)
}

func (c *Client) GetKey(keyID string) (*Key, error) {
	resp, err := c.sendRequest("GET", fmt.Sprintf("/key/info?key=%s", keyID), nil)
	if err != nil {
		return nil, err
	}

	return c.parseKeyResponse(resp)
}

func (c *Client) UpdateKey(key *Key) (*Key, error) {
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

	resp, err := c.sendRequest("POST", "/key/update", updateData)
	if err != nil {
		return nil, err
	}

	return c.parseKeyResponse(resp)
}

func (c *Client) DeleteKey(keyID string) error {
	payload := map[string]interface{}{
		"keys": []string{keyID},
	}
	_, err := c.sendRequest("POST", "/key/delete", payload)
	return err
}

func (c *Client) parseKeyResponse(resp map[string]interface{}) (*Key, error) {
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

func (c *Client) sendRequest(method, path string, body interface{}) (map[string]interface{}, error) {
	url := c.APIBase + path

	var req *http.Request
	var err error

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
		log.Printf("Making %s request to %s with body:\n%s", method, url, string(jsonBody))
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	} else {
		log.Printf("Making %s request to %s", method, url)
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	log.Printf("Response status: %d", resp.StatusCode)
	log.Printf("Response body: %s", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		if method == "POST" && (len(bodyBytes) == 0 || string(bodyBytes) == "null") {
			return make(map[string]interface{}), nil
		}
		return nil, fmt.Errorf("error parsing response JSON: %v\nResponse body: %s", err, string(bodyBytes))
	}

	return result, nil
}
