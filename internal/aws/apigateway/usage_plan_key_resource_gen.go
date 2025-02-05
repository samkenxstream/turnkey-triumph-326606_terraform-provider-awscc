// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_usage_plan_key", usagePlanKeyResource)
}

// usagePlanKeyResource returns the Terraform awscc_apigateway_usage_plan_key resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::UsagePlanKey resource.
func usagePlanKeyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the usage plan key.",
		//	  "type": "string"
		//	}
		"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the usage plan key.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KeyType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of usage plan key. Currently, the only valid key type is API_KEY.",
		//	  "enum": [
		//	    "API_KEY"
		//	  ],
		//	  "type": "string"
		//	}
		"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of usage plan key. Currently, the only valid key type is API_KEY.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"API_KEY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UsagePlanId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the usage plan.",
		//	  "type": "string"
		//	}
		"usage_plan_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the usage plan.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::ApiGateway::UsagePlanKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::UsagePlanKey").WithTerraformTypeName("awscc_apigateway_usage_plan_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":            "Id",
		"key_id":        "KeyId",
		"key_type":      "KeyType",
		"usage_plan_id": "UsagePlanId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
