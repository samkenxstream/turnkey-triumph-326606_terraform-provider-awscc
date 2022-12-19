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
	registry.AddDataSourceFactory("awscc_cloudfront_continuous_deployment_policy", continuousDeploymentPolicyDataSource)
}

// continuousDeploymentPolicyDataSource returns the Terraform awscc_cloudfront_continuous_deployment_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::ContinuousDeploymentPolicy resource.
func continuousDeploymentPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"continuous_deployment_policy_config": {
			// Property: ContinuousDeploymentPolicyConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Enabled": {
			//	      "type": "boolean"
			//	    },
			//	    "StagingDistributionDnsNames": {
			//	      "insertionOrder": true,
			//	      "items": {
			//	        "type": "string"
			//	      },
			//	      "minItems": 1,
			//	      "type": "array",
			//	      "uniqueItems": true
			//	    },
			//	    "TrafficConfig": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "SingleHeaderConfig": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "Header": {
			//	              "maxLength": 256,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            },
			//	            "Value": {
			//	              "maxLength": 1783,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "Header",
			//	            "Value"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "SingleWeightConfig": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "SessionStickinessConfig": {
			//	              "additionalProperties": false,
			//	              "properties": {
			//	                "IdleTTL": {
			//	                  "maximum": 3600,
			//	                  "minimum": 300,
			//	                  "type": "integer"
			//	                },
			//	                "MaximumTTL": {
			//	                  "maximum": 3600,
			//	                  "minimum": 300,
			//	                  "type": "integer"
			//	                }
			//	              },
			//	              "required": [
			//	                "IdleTTL",
			//	                "MaximumTTL"
			//	              ],
			//	              "type": "object"
			//	            },
			//	            "Weight": {
			//	              "maximum": 1,
			//	              "minimum": 0,
			//	              "type": "number"
			//	            }
			//	          },
			//	          "required": [
			//	            "Weight"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "Type": {
			//	          "enum": [
			//	            "SingleWeight",
			//	            "SingleHeader"
			//	          ],
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "Type"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "required": [
			//	    "Enabled",
			//	    "StagingDistributionDnsNames"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enabled": {
						// Property: Enabled
						Type:     types.BoolType,
						Computed: true,
					},
					"staging_distribution_dns_names": {
						// Property: StagingDistributionDnsNames
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"traffic_config": {
						// Property: TrafficConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"single_header_config": {
									// Property: SingleHeaderConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"header": {
												// Property: Header
												Type:     types.StringType,
												Computed: true,
											},
											"value": {
												// Property: Value
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"single_weight_config": {
									// Property: SingleWeightConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"session_stickiness_config": {
												// Property: SessionStickinessConfig
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"idle_ttl": {
															// Property: IdleTTL
															Type:     types.Int64Type,
															Computed: true,
														},
														"maximum_ttl": {
															// Property: MaximumTTL
															Type:     types.Int64Type,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"weight": {
												// Property: Weight
												Type:     types.Float64Type,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
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
		Description: "Data Source schema for AWS::CloudFront::ContinuousDeploymentPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::ContinuousDeploymentPolicy").WithTerraformTypeName("awscc_cloudfront_continuous_deployment_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"continuous_deployment_policy_config": "ContinuousDeploymentPolicyConfig",
		"enabled":                             "Enabled",
		"header":                              "Header",
		"id":                                  "Id",
		"idle_ttl":                            "IdleTTL",
		"last_modified_time":                  "LastModifiedTime",
		"maximum_ttl":                         "MaximumTTL",
		"session_stickiness_config":           "SessionStickinessConfig",
		"single_header_config":                "SingleHeaderConfig",
		"single_weight_config":                "SingleWeightConfig",
		"staging_distribution_dns_names":      "StagingDistributionDnsNames",
		"traffic_config":                      "TrafficConfig",
		"type":                                "Type",
		"value":                               "Value",
		"weight":                              "Weight",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}