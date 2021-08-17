// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ssmincidents_replication_set", replicationSetResourceType)
}

// replicationSetResourceType returns the Terraform awscc_ssmincidents_replication_set resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMIncidents::ReplicationSet resource type.
func replicationSetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the ReplicationSet.",
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the ReplicationSet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"deletion_protected": {
			// Property: DeletionProtected
			// CloudFormation resource type schema:
			// {
			//   "description": "Configures the ReplicationSet deletion protection.",
			//   "type": "boolean"
			// }
			Description: "Configures the ReplicationSet deletion protection.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"regions": {
			// Property: Regions
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The ReplicationSet regional configuration.",
			//     "properties": {
			//       "RegionConfiguration": {
			//         "additionalProperties": false,
			//         "description": "The ReplicationSet regional configuration.",
			//         "properties": {
			//           "SseKmsKeyId": {
			//             "description": "The ARN of the ReplicationSet.",
			//             "maxLength": 1000,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "SseKmsKeyId"
			//         ],
			//         "type": "object"
			//       },
			//       "RegionName": {
			//         "description": "The AWS region name.",
			//         "maxLength": 20,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 3,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"region_configuration": {
						// Property: RegionConfiguration
						Description: "The ReplicationSet regional configuration.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"sse_kms_key_id": {
									// Property: SseKmsKeyId
									Description: "The ARN of the ReplicationSet.",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
					"region_name": {
						// Property: RegionName
						Description: "The AWS region name.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 3,
				},
			),
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource type definition for AWS::SSMIncidents::ReplicationSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMIncidents::ReplicationSet").WithTerraformTypeName("awscc_ssmincidents_replication_set").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ssmincidents_replication_set", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
