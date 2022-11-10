// Code generated by generators/resource/main.go; DO NOT EDIT.

package codestarconnections

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
	registry.AddResourceFactory("awscc_codestarconnections_connection", connectionResource)
}

// connectionResource returns the Terraform awscc_codestarconnections_connection resource.
// This Terraform resource corresponds to the CloudFormation AWS::CodeStarConnections::Connection resource.
func connectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.",
			//	  "maxLength": 256,
			//	  "minLength": 0,
			//	  "pattern": "arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"connection_name": {
			// Property: ConnectionName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the connection. Connection names must be unique in an AWS user account.",
			//	  "maxLength": 32,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The name of the connection. Connection names must be unique in an AWS user account.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 32),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"connection_status": {
			// Property: ConnectionStatus
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The current status of the connection.",
			//	  "type": "string"
			//	}
			Description: "The current status of the connection.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"host_arn": {
			// Property: HostArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.",
			//	  "maxLength": 256,
			//	  "minLength": 0,
			//	  "pattern": "arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+",
			//	  "type": "string"
			//	}
			Description: "The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
				validate.StringMatch(regexp.MustCompile("arn:aws(-[\\w]+)*:.+:.+:[0-9]{12}:.+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"owner_account_id": {
			// Property: OwnerAccountId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.",
			//	  "maxLength": 12,
			//	  "minLength": 12,
			//	  "pattern": "[0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"provider_type": {
			// Property: ProviderType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.",
			//	  "type": "string"
			//	}
			Description: "The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the tags applied to a connection.",
			//	  "items": {
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "Specifies the tags applied to a connection.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
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
		Description: "Schema for AWS::CodeStarConnections::Connection resource which can be used to connect external source providers with AWS CodePipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarConnections::Connection").WithTerraformTypeName("awscc_codestarconnections_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_arn":    "ConnectionArn",
		"connection_name":   "ConnectionName",
		"connection_status": "ConnectionStatus",
		"host_arn":          "HostArn",
		"key":               "Key",
		"owner_account_id":  "OwnerAccountId",
		"provider_type":     "ProviderType",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
