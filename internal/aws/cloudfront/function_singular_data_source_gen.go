// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_function", functionDataSource)
}

// functionDataSource returns the Terraform awscc_cloudfront_function data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::Function resource.
func functionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_publish": {
			// Property: AutoPublish
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Computed: true,
		},
		"function_arn": {
			// Property: FunctionARN
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"function_code": {
			// Property: FunctionCode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"function_config": {
			// Property: FunctionConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Comment": {
			//	      "type": "string"
			//	    },
			//	    "Runtime": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Comment",
			//	    "Runtime"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Computed: true,
					},
					"runtime": {
						// Property: Runtime
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"function_metadata": {
			// Property: FunctionMetadata
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "FunctionARN": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"function_arn": {
						// Property: FunctionARN
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"stage": {
			// Property: Stage
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFront::Function",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::Function").WithTerraformTypeName("awscc_cloudfront_function")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_publish":      "AutoPublish",
		"comment":           "Comment",
		"function_arn":      "FunctionARN",
		"function_code":     "FunctionCode",
		"function_config":   "FunctionConfig",
		"function_metadata": "FunctionMetadata",
		"name":              "Name",
		"runtime":           "Runtime",
		"stage":             "Stage",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
