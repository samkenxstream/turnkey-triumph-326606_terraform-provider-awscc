// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package aps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_aps_workspace", workspaceDataSource)
}

// workspaceDataSource returns the Terraform awscc_aps_workspace data source.
// This Terraform data source corresponds to the CloudFormation AWS::APS::Workspace resource.
func workspaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"alert_manager_definition": {
			// Property: AlertManagerDefinition
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The AMP Workspace alert manager definition data",
			//	  "type": "string"
			//	}
			Description: "The AMP Workspace alert manager definition data",
			Type:        types.StringType,
			Computed:    true,
		},
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "AMP Workspace alias.",
			//	  "maxLength": 100,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "AMP Workspace alias.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Workspace arn.",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):aps:[a-z0-9-]+:[0-9]+:workspace/[a-zA-Z0-9-]+$",
			//	  "type": "string"
			//	}
			Description: "Workspace arn.",
			Type:        types.StringType,
			Computed:    true,
		},
		"logging_configuration": {
			// Property: LoggingConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Logging configuration",
			//	  "properties": {
			//	    "LogGroupArn": {
			//	      "description": "CloudWatch log group ARN",
			//	      "maxLength": 512,
			//	      "minLength": 0,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Logging configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_group_arn": {
						// Property: LogGroupArn
						Description: "CloudWatch log group ARN",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"prometheus_endpoint": {
			// Property: PrometheusEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "AMP Workspace prometheus endpoint",
			//	  "type": "string"
			//	}
			Description: "AMP Workspace prometheus endpoint",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this resource.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"workspace_id": {
			// Property: WorkspaceId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Required to identify a specific APS Workspace.",
			//	  "maxLength": 100,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9_-]{1,99}$",
			//	  "type": "string"
			//	}
			Description: "Required to identify a specific APS Workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::APS::Workspace",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::APS::Workspace").WithTerraformTypeName("awscc_aps_workspace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alert_manager_definition": "AlertManagerDefinition",
		"alias":                    "Alias",
		"arn":                      "Arn",
		"key":                      "Key",
		"log_group_arn":            "LogGroupArn",
		"logging_configuration":    "LoggingConfiguration",
		"prometheus_endpoint":      "PrometheusEndpoint",
		"tags":                     "Tags",
		"value":                    "Value",
		"workspace_id":             "WorkspaceId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
