package litellm

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Client struct {
	APIBase        string
	APIKey         string
	httpClient     *http.Client
	rateLimitedMux sync.Mutex
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

// SendRequest sends an HTTP request to the LiteLLM API and returns the response as a map.
func (c *Client) SendRequest(ctx context.Context, method, path string, body interface{}) (map[string]interface{}, error) {
	url := c.APIBase + path

	var req *http.Request
	var err error

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
		tflog.Debug(ctx, "Making request with body", map[string]interface{}{
			"method": method,
			"url":    url,
			"body":   string(jsonBody),
		})
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
	} else {
		tflog.Debug(ctx, "Making request", map[string]interface{}{
			"method": method,
			"url":    url,
		})
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
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

	tflog.Debug(ctx, "Received response", map[string]interface{}{
		"status_code": resp.StatusCode,
		"body":        string(bodyBytes),
	})

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

// SendRequestTyped sends an HTTP request to the LiteLLM API with typed request and response.
func SendRequestTyped[TRequest any, TResponse any](ctx context.Context, c *Client, method, path string, body *TRequest) (*TResponse, error) {
	url := c.APIBase + path

	var req *http.Request
	var err error

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
		tflog.Debug(ctx, "Making typed request with body", map[string]interface{}{
			"method": method,
			"url":    url,
			"body":   string(jsonBody),
		})
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
	} else {
		tflog.Debug(ctx, "Making typed request", map[string]interface{}{
			"method": method,
			"url":    url,
		})
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
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

	tflog.Debug(ctx, "Received typed response", map[string]interface{}{
		"status_code": resp.StatusCode,
		"body":        string(bodyBytes),
	})

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result TResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		if method == "POST" && (len(bodyBytes) == 0 || string(bodyBytes) == "null") {
			return &result, nil
		}
		return nil, fmt.Errorf("error parsing response JSON: %v\nResponse body: %s", err, string(bodyBytes))
	}

	return &result, nil
}

// SendRequestTypedRateLimited sends an HTTP request to the LiteLLM API with typed request and response,
// ensuring only one request is processed at a time using a mutex. This is useful for operations
// that need to be serialized to prevent race conditions or API rate limiting issues.
func SendRequestTypedRateLimited[TRequest any, TResponse any](ctx context.Context, c *Client, method, path string, body *TRequest) (*TResponse, error) {
	// Acquire mutex to ensure only one request at a time
	c.rateLimitedMux.Lock()
	defer c.rateLimitedMux.Unlock()

	tflog.Debug(ctx, "Acquired rate-limited mutex for request", map[string]interface{}{
		"method": method,
		"path":   path,
	})

	// Use the existing SendRequestTyped method for the actual request
	return SendRequestTyped[TRequest, TResponse](ctx, c, method, path, body)
}
