// Code generated by generators/resource/main.go; DO NOT EDIT.

package rolesanywhere

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_rolesanywhere_profile", profileResource)
}

// profileResource returns the Terraform awscc_rolesanywhere_profile resource.
// This Terraform resource corresponds to the CloudFormation AWS::RolesAnywhere::Profile resource.
func profileResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"duration_seconds": {
			// Property: DurationSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maximum": 43200,
			//	  "minimum": 900,
			//	  "type": "number"
			//	}
			Type:     types.Float64Type,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.FloatBetween(900.000000, 43200.000000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"managed_policy_arns": {
			// Property: ManagedPolicyArns
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"profile_arn": {
			// Property: ProfileArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"profile_id": {
			// Property: ProfileId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "[a-f0-9]{8}-([a-z0-9]{4}-){3}[a-z0-9]{12}",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"require_instance_properties": {
			// Property: RequireInstanceProperties
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"role_arns": {
			// Property: RoleArns
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "maxLength": 1011,
			//	    "minLength": 1,
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenBetween(1, 1011)),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"session_policy": {
			// Property: SessionPolicy
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 200,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 200),
			},
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
		Description: "Definition of AWS::RolesAnywhere::Profile Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RolesAnywhere::Profile").WithTerraformTypeName("awscc_rolesanywhere_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"duration_seconds":            "DurationSeconds",
		"enabled":                     "Enabled",
		"key":                         "Key",
		"managed_policy_arns":         "ManagedPolicyArns",
		"name":                        "Name",
		"profile_arn":                 "ProfileArn",
		"profile_id":                  "ProfileId",
		"require_instance_properties": "RequireInstanceProperties",
		"role_arns":                   "RoleArns",
		"session_policy":              "SessionPolicy",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
