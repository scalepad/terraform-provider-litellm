package models

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scalepad/terraform-provider-litellm/internal/utils"
)

func buildModelData(d *schema.ResourceData) map[string]interface{} {
	modelData := make(map[string]interface{})

	// String fields
	utils.GetValueDefault[string](d, "model_name", modelData)
	utils.GetValueDefault[string](d, "custom_llm_provider", modelData)
	utils.GetValueDefault[string](d, "base_model", modelData)
	utils.GetValueDefault[string](d, "tier", modelData)
	utils.GetValueDefault[string](d, "mode", modelData)
	utils.GetValueDefault[string](d, "team_id", modelData)
	utils.GetValueDefault[string](d, "model_api_key", modelData)
	utils.GetValueDefault[string](d, "model_api_base", modelData)
	utils.GetValueDefault[string](d, "api_version", modelData)
	utils.GetValueDefault[string](d, "aws_access_key_id", modelData)
	utils.GetValueDefault[string](d, "aws_secret_access_key", modelData)
	utils.GetValueDefault[string](d, "aws_region_name", modelData)
	utils.GetValueDefault[string](d, "vertex_project", modelData)
	utils.GetValueDefault[string](d, "vertex_location", modelData)
	utils.GetValueDefault[string](d, "vertex_credentials", modelData)
	utils.GetValueDefault[string](d, "reasoning_effort", modelData)

	// Int fields
	utils.GetValueDefault[int](d, "tpm", modelData)
	utils.GetValueDefault[int](d, "rpm", modelData)
	utils.GetValueDefault[int](d, "thinking_budget_tokens", modelData)

	// Float64 fields
	utils.GetValueDefault[float64](d, "input_cost_per_million_tokens", modelData)
	utils.GetValueDefault[float64](d, "output_cost_per_million_tokens", modelData)
	utils.GetValueDefault[float64](d, "input_cost_per_pixel", modelData)
	utils.GetValueDefault[float64](d, "output_cost_per_pixel", modelData)
	utils.GetValueDefault[float64](d, "input_cost_per_second", modelData)
	utils.GetValueDefault[float64](d, "output_cost_per_second", modelData)

	// Bool fields
	utils.GetValueDefault[bool](d, "thinking_enabled", modelData)
	utils.GetValueDefault[bool](d, "merge_reasoning_content_in_choices", modelData)

	// Map fields
	utils.GetValueDefault[map[string]interface{}](d, "additional_litellm_params", modelData)

	return modelData
}

func setModelResourceData(d *schema.ResourceData, model *ModelResponse) error {
	fields := map[string]interface{}{
		"model_name":                         model.ModelName,
		"custom_llm_provider":                model.LiteLLMParams.CustomLLMProvider,
		"tpm":                                model.LiteLLMParams.TPM,
		"rpm":                                model.LiteLLMParams.RPM,
		"model_api_base":                     model.LiteLLMParams.APIBase,
		"api_version":                        model.LiteLLMParams.APIVersion,
		"base_model":                         model.ModelInfo.BaseModel,
		"tier":                               model.ModelInfo.Tier,
		"mode":                               model.ModelInfo.Mode,
		"team_id":                            model.ModelInfo.TeamID,
		"aws_region_name":                    model.LiteLLMParams.AWSRegionName,
		"reasoning_effort":                   model.LiteLLMParams.ReasoningEffort,
		"merge_reasoning_content_in_choices": model.LiteLLMParams.MergeReasoningContentInChoices,
	}

	for field, value := range fields {
		// Use SetIfNotZero to preserve existing values when API doesn't return them
		utils.SetIfNotZero(d, field, value)
	}

	// Handle thinking configuration
	if model.LiteLLMParams.Thinking != nil {
		if thinkingType, ok := model.LiteLLMParams.Thinking["type"].(string); ok && thinkingType == "enabled" {
			d.Set("thinking_enabled", true)
			if budgetTokens, ok := model.LiteLLMParams.Thinking["budget_tokens"].(float64); ok {
				d.Set("thinking_budget_tokens", int(budgetTokens))
			}
		} else {
			d.Set("thinking_enabled", false)
		}
	}

	// Preserve sensitive fields from state
	d.Set("model_api_key", d.Get("model_api_key"))
	d.Set("aws_access_key_id", d.Get("aws_access_key_id"))
	d.Set("aws_secret_access_key", d.Get("aws_secret_access_key"))

	// Preserve cost information from state
	d.Set("input_cost_per_million_tokens", d.Get("input_cost_per_million_tokens"))
	d.Set("output_cost_per_million_tokens", d.Get("output_cost_per_million_tokens"))

	// Preserve additional_litellm_params from state
	if _, ok := d.GetOk("additional_litellm_params"); ok {
		d.Set("additional_litellm_params", d.Get("additional_litellm_params"))
	}

	return nil
}

func parseModelAPIResponse(resp map[string]interface{}) (*ModelResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("received nil response")
	}

	modelResp := &ModelResponse{}

	// Parse basic fields
	if v, ok := resp["model_name"].(string); ok {
		modelResp.ModelName = v
	}

	// Parse model_info
	if modelInfoData, ok := resp["model_info"].(map[string]interface{}); ok {
		if v, ok := modelInfoData["id"].(string); ok {
			modelResp.ModelInfo.ID = v
		}
		if v, ok := modelInfoData["base_model"].(string); ok {
			modelResp.ModelInfo.BaseModel = v
		}
		if v, ok := modelInfoData["tier"].(string); ok {
			modelResp.ModelInfo.Tier = v
		}
		if v, ok := modelInfoData["mode"].(string); ok {
			modelResp.ModelInfo.Mode = v
		}
		if v, ok := modelInfoData["team_id"].(string); ok {
			modelResp.ModelInfo.TeamID = v
		}
		if v, ok := modelInfoData["db_model"].(bool); ok {
			modelResp.ModelInfo.DBModel = v
		}
	}

	// Parse litellm_params
	if litellmParamsData, ok := resp["litellm_params"].(map[string]interface{}); ok {
		if v, ok := litellmParamsData["custom_llm_provider"].(string); ok {
			modelResp.LiteLLMParams.CustomLLMProvider = v
		}
		if v, ok := litellmParamsData["tpm"].(float64); ok {
			modelResp.LiteLLMParams.TPM = int(v)
		}
		if v, ok := litellmParamsData["rpm"].(float64); ok {
			modelResp.LiteLLMParams.RPM = int(v)
		}
		if v, ok := litellmParamsData["api_base"].(string); ok {
			modelResp.LiteLLMParams.APIBase = v
		}
		if v, ok := litellmParamsData["api_version"].(string); ok {
			modelResp.LiteLLMParams.APIVersion = v
		}
		if v, ok := litellmParamsData["aws_region_name"].(string); ok {
			modelResp.LiteLLMParams.AWSRegionName = v
		}
		if v, ok := litellmParamsData["reasoning_effort"].(string); ok {
			modelResp.LiteLLMParams.ReasoningEffort = v
		}
		if v, ok := litellmParamsData["merge_reasoning_content_in_choices"].(bool); ok {
			modelResp.LiteLLMParams.MergeReasoningContentInChoices = v
		}
		if v, ok := litellmParamsData["thinking"].(map[string]interface{}); ok {
			modelResp.LiteLLMParams.Thinking = v
		}
	}

	// Parse additional fields
	if v, ok := resp["additional"].(map[string]interface{}); ok {
		modelResp.Additional = v
	}

	return modelResp, nil
}

func buildModelForCreation(data map[string]interface{}) *Model {
	model := &Model{}

	if v, ok := data["model_name"].(string); ok {
		model.ModelName = v
	}

	// Build LiteLLMParams
	litellmParams := make(map[string]interface{})

	// Convert cost per million tokens to cost per token
	if v, ok := data["input_cost_per_million_tokens"].(float64); ok && v > 0 {
		litellmParams["input_cost_per_token"] = v / 1000000.0
	}
	if v, ok := data["output_cost_per_million_tokens"].(float64); ok && v > 0 {
		litellmParams["output_cost_per_token"] = v / 1000000.0
	}

	// Add other LiteLLM params
	if v, ok := data["custom_llm_provider"].(string); ok {
		litellmParams["custom_llm_provider"] = v
		// Construct the model name in the format "custom_llm_provider/base_model"
		if baseModel, ok := data["base_model"].(string); ok {
			litellmParams["model"] = fmt.Sprintf("%s/%s", v, baseModel)
		}
	}

	// Add optional parameters
	for _, field := range []string{"tpm", "rpm", "model_api_key", "model_api_base", "api_version",
		"input_cost_per_pixel", "output_cost_per_pixel", "input_cost_per_second", "output_cost_per_second",
		"aws_access_key_id", "aws_secret_access_key", "aws_region_name",
		"vertex_project", "vertex_location", "vertex_credentials", "reasoning_effort"} {
		if v, ok := data[field]; ok {
			switch field {
			case "model_api_key":
				litellmParams["api_key"] = v
			case "model_api_base":
				litellmParams["api_base"] = v
			default:
				litellmParams[field] = v
			}
		}
	}

	// Handle thinking configuration
	if thinkingEnabled, ok := data["thinking_enabled"].(bool); ok && thinkingEnabled {
		thinking := map[string]interface{}{
			"type": "enabled",
		}
		if budgetTokens, ok := data["thinking_budget_tokens"].(int); ok {
			thinking["budget_tokens"] = budgetTokens
		}
		litellmParams["thinking"] = thinking
	}

	// Handle merge_reasoning_content_in_choices
	if v, ok := data["merge_reasoning_content_in_choices"].(bool); ok {
		litellmParams["merge_reasoning_content_in_choices"] = v
	}

	// Add additional parameters if provided
	if additionalParams, ok := data["additional_litellm_params"].(map[string]interface{}); ok {
		for key, value := range additionalParams {
			// Convert string values to appropriate types where possible
			if strValue, ok := value.(string); ok {
				// Try to convert boolean strings
				if strValue == "true" {
					litellmParams[key] = true
				} else if strValue == "false" {
					litellmParams[key] = false
				} else {
					// Try to convert numeric strings
					if intValue, err := strconv.Atoi(strValue); err == nil {
						litellmParams[key] = intValue
					} else if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
						litellmParams[key] = floatValue
					} else {
						// Keep as string
						litellmParams[key] = strValue
					}
				}
			} else {
				litellmParams[key] = value
			}
		}
	}

	model.LiteLLMParams = litellmParams

	// Build ModelInfo
	modelInfo := ModelInfo{
		DBModel: true,
	}
	if v, ok := data["base_model"].(string); ok {
		modelInfo.BaseModel = v
	}
	if v, ok := data["tier"].(string); ok {
		modelInfo.Tier = v
	}
	if v, ok := data["mode"].(string); ok {
		modelInfo.Mode = v
	}
	if v, ok := data["team_id"].(string); ok {
		modelInfo.TeamID = v
	}

	model.ModelInfo = modelInfo
	model.Additional = make(map[string]interface{})

	return model
}
