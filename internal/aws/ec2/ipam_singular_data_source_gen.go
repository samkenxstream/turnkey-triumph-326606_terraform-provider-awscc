// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_ipam", iPAMDataSource)
}

// iPAMDataSource returns the Terraform awscc_ec2_ipam data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAM resource.
func iPAMDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the IPAM.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the IPAM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"ipam_id": {
			// Property: IpamId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Id of the IPAM.",
			//	  "type": "string"
			//	}
			Description: "Id of the IPAM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"operating_regions": {
			// Property: OperatingRegions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
			//	    "properties": {
			//	      "RegionName": {
			//	        "description": "The name of the region.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "RegionName"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"region_name": {
						// Property: RegionName
						Description: "The name of the region.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"private_default_scope_id": {
			// Property: PrivateDefaultScopeId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			//	  "type": "string"
			//	}
			Description: "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_default_scope_id": {
			// Property: PublicDefaultScopeId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			//	  "maxLength": 255,
			//	  "type": "string"
			//	}
			Description: "The Id of the default scope for publicly routable IP space, created with this IPAM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"scope_count": {
			// Property: ScopeCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of scopes that currently exist in this IPAM.",
			//	  "type": "integer"
			//	}
			Description: "The number of scopes that currently exist in this IPAM.",
			Type:        types.Int64Type,
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::IPAM",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAM").WithTerraformTypeName("awscc_ec2_ipam")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"description":              "Description",
		"ipam_id":                  "IpamId",
		"key":                      "Key",
		"operating_regions":        "OperatingRegions",
		"private_default_scope_id": "PrivateDefaultScopeId",
		"public_default_scope_id":  "PublicDefaultScopeId",
		"region_name":              "RegionName",
		"scope_count":              "ScopeCount",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
