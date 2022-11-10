// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_elasticache_user", userResource)
}

// userResource returns the Terraform awscc_elasticache_user resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElastiCache::User resource.
func userResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_string": {
			// Property: AccessString
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Access permissions string used for this user account.",
			//	  "type": "string"
			//	}
			Description: "Access permissions string used for this user account.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// AccessString is a write-only property.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the user account.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the user account.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"engine": {
			// Property: Engine
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Must be redis.",
			//	  "enum": [
			//	    "redis"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Must be redis.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"redis",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"no_password_required": {
			// Property: NoPasswordRequired
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates a password is not required for this user account.",
			//	  "type": "boolean"
			//	}
			Description: "Indicates a password is not required for this user account.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// NoPasswordRequired is a write-only property.
		},
		"passwords": {
			// Property: Passwords
			// CloudFormation resource type schema:
			//
			//	{
			//	  "$comment": "List of passwords.",
			//	  "description": "Passwords used for this user account. You can create up to two passwords for each user.",
			//	  "insertionOrder": true,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "Passwords used for this user account. You can create up to two passwords for each user.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Passwords is a write-only property.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			//	  "type": "string"
			//	}
			Description: "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"user_id": {
			// Property: UserId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the user.",
			//	  "pattern": "[a-z][a-z0-9\\\\-]*",
			//	  "type": "string"
			//	}
			Description: "The ID of the user.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[a-z][a-z0-9\\\\-]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"user_name": {
			// Property: UserName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The username of the user.",
			//	  "type": "string"
			//	}
			Description: "The username of the user.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::ElastiCache::User",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::User").WithTerraformTypeName("awscc_elasticache_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_string":        "AccessString",
		"arn":                  "Arn",
		"engine":               "Engine",
		"no_password_required": "NoPasswordRequired",
		"passwords":            "Passwords",
		"status":               "Status",
		"user_id":              "UserId",
		"user_name":            "UserName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Passwords",
		"/properties/NoPasswordRequired",
		"/properties/AccessString",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
