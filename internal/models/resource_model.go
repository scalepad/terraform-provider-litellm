package models

import (
	"context"
	"fmt"

	"github.com/scalepad/terraform-provider-litellm/internal/litellm"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceModel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceModelCreate,
		ReadContext:   resourceModelRead,
		UpdateContext: resourceModelUpdate,
		DeleteContext: resourceModelDelete,
		Importer:      ModelImporter(),
		Schema:        resourceModelSchema(),
	}
}

func resourceModelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"model_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"custom_llm_provider": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tpm": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"rpm": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"reasoning_effort": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				"low",
				"medium",
				"high",
			}, false),
		},
		"thinking_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"thinking_budget_tokens": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1024,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// Only include thinking_budget_tokens in the diff if thinking_enabled is true
				return !d.Get("thinking_enabled").(bool)
			},
		},
		"merge_reasoning_content_in_choices": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"model_api_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"model_api_base": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"api_version": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"base_model": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tier": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "free",
		},
		"team_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				"completion",
				"embedding",
				"image_generation",
				"chat",
				"moderation",
				"audio_transcription",
			}, false),
		},
		"input_cost_per_million_tokens": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"output_cost_per_million_tokens": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"input_cost_per_pixel": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"output_cost_per_pixel": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"input_cost_per_second": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"output_cost_per_second": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"aws_access_key_id": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"aws_secret_access_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"aws_region_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vertex_project": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"vertex_location": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"vertex_credentials": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"additional_litellm_params": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Additional parameters to pass to litellm_params beyond the standard ones",
		},
	}
}

func resourceModelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	modelData := buildModelData(d)
	model := buildModelForCreation(modelData)

	createdModel, err := createModel(ctx, c, model)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating model: %s", err))
	}

	d.SetId(createdModel.ModelInfo.ID)
	return resourceModelRead(ctx, d, m)
}

func resourceModelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	model, err := getModel(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading model: %s", err))
	}

	if model == nil {
		d.SetId("")
		return nil
	}

	if err := setModelResourceData(d, model); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceModelUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	modelData := buildModelData(d)
	model := buildModelForCreation(modelData)
	model.ModelInfo.ID = d.Id() // Set the model ID for update

	_, err := updateModel(ctx, c, model)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating model: %s", err))
	}

	return resourceModelRead(ctx, d, m)
}

func resourceModelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*litellm.Client)

	err := deleteModel(ctx, c, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting model: %s", err))
	}

	d.SetId("")
	return nil
}
