// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceTypeFactory("awscc_ecs_cluster_capacity_provider_associations", clusterCapacityProviderAssociationsResourceType)
}

// clusterCapacityProviderAssociationsResourceType returns the Terraform awscc_ecs_cluster_capacity_provider_associations resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECS::ClusterCapacityProviderAssociations resource type.
func clusterCapacityProviderAssociationsResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity_providers": {
			// Property: CapacityProviders
			// CloudFormation resource type schema:
			// {
			//   "description": "List of capacity providers to associate with the cluster",
			//   "items": {
			//     "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "List of capacity providers to associate with the cluster",
			// Ordered set.
			Type:     providertypes.OrderedSetType{ListType: types.ListType{ElemType: types.StringType}},
			Required: true,
		},
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cluster",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the cluster",
			Type:        types.StringType,
			Required:    true,
			// Cluster is a force-new attribute.
		},
		"default_capacity_provider_strategy": {
			// Property: DefaultCapacityProviderStrategy
			// CloudFormation resource type schema:
			// {
			//   "description": "List of capacity providers to associate with the cluster",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Base": {
			//         "type": "integer"
			//       },
			//       "CapacityProvider": {
			//         "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			//         "type": "string"
			//       },
			//       "Weight": {
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "CapacityProvider"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "List of capacity providers to associate with the cluster",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"base": {
						// Property: Base
						Type:     types.NumberType,
						Optional: true,
					},
					"capacity_provider": {
						// Property: CapacityProvider
						Description: "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
						Type:        types.StringType,
						Required:    true,
					},
					"weight": {
						// Property: Weight
						Type:     types.NumberType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
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
		Description: "Associate a set of ECS Capacity Providers with a specified ECS Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::ClusterCapacityProviderAssociations").WithTerraformTypeName("awscc_ecs_cluster_capacity_provider_associations").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ecs_cluster_capacity_provider_associations", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
