// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_resiliencehub_resiliency_policy", resiliencyPolicyDataSourceType)
}

// resiliencyPolicyDataSourceType returns the Terraform awscc_resiliencehub_resiliency_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ResilienceHub::ResiliencyPolicy resource type.
func resiliencyPolicyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_location_constraint": {
			// Property: DataLocationConstraint
			// CloudFormation resource type schema:
			// {
			//   "description": "Data Location Constraint of the Policy.",
			//   "enum": [
			//     "AnyLocation",
			//     "SameContinent",
			//     "SameCountry"
			//   ],
			//   "type": "string"
			// }
			Description: "Data Location Constraint of the Policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "additionalProperties": false,
			//       "description": "Failure Policy.",
			//       "properties": {
			//         "RpoInSecs": {
			//           "description": "RPO in seconds.",
			//           "type": "integer"
			//         },
			//         "RtoInSecs": {
			//           "description": "RTO in seconds.",
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "RtoInSecs",
			//         "RpoInSecs"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Attributes: tfsdk.MapNestedAttributes(
				map[string]tfsdk.Attribute{
					"rpo_in_secs": {
						// Property: RpoInSecs
						Description: "RPO in seconds.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"rto_in_secs": {
						// Property: RtoInSecs
						Description: "RTO in seconds.",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
				tfsdk.MapNestedAttributesOptions{},
			),
			Computed: true,
		},
		"policy_arn": {
			// Property: PolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_description": {
			// Property: PolicyDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of Resiliency Policy.",
			//   "maxLength": 500,
			//   "type": "string"
			// }
			Description: "Description of Resiliency Policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of Resiliency Policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"tier": {
			// Property: Tier
			// CloudFormation resource type schema:
			// {
			//   "description": "Resiliency Policy Tier.",
			//   "enum": [
			//     "MissionCritical",
			//     "Critical",
			//     "Important",
			//     "CoreServices",
			//     "NonCritical"
			//   ],
			//   "type": "string"
			// }
			Description: "Resiliency Policy Tier.",
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
		Description: "Data Source schema for AWS::ResilienceHub::ResiliencyPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::ResiliencyPolicy").WithTerraformTypeName("awscc_resiliencehub_resiliency_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location_constraint": "DataLocationConstraint",
		"policy":                   "Policy",
		"policy_arn":               "PolicyArn",
		"policy_description":       "PolicyDescription",
		"policy_name":              "PolicyName",
		"rpo_in_secs":              "RpoInSecs",
		"rto_in_secs":              "RtoInSecs",
		"tags":                     "Tags",
		"tier":                     "Tier",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}