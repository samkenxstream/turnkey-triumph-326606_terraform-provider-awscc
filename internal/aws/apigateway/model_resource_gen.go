// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_model", modelResource)
}

// modelResource returns the Terraform awscc_apigateway_model resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::Model resource.
func modelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"content_type": {
			// Property: ContentType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The content type for the model.",
			//	  "type": "string"
			//	}
			Description: "The content type for the model.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description that identifies this model.",
			//	  "type": "string"
			//	}
			Description: "A description that identifies this model.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.",
			//	  "type": "string"
			//	}
			Description: "A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of a REST API with which to associate this model.",
			//	  "type": "string"
			//	}
			Description: "The ID of a REST API with which to associate this model.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"schema": {
			// Property: Schema
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.",
			//	  "type": "string"
			//	}
			Description: "The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ApiGateway::Model",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Model").WithTerraformTypeName("awscc_apigateway_model")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"content_type": "ContentType",
		"description":  "Description",
		"name":         "Name",
		"rest_api_id":  "RestApiId",
		"schema":       "Schema",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
