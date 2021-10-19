// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_subnet_network_acl_association", subnetNetworkAclAssociationResourceType)
}

// subnetNetworkAclAssociationResourceType returns the Terraform awscc_ec2_subnet_network_acl_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::SubnetNetworkAclAssociation resource type.
func subnetNetworkAclAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"network_acl_id": {
			// Property: NetworkAclId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the network ACL",
			//   "type": "string"
			// }
			Description: "The ID of the network ACL",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the subnet",
			//   "type": "string"
			// }
			Description: "The ID of the subnet",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::SubnetNetworkAclAssociation").WithTerraformTypeName("awscc_ec2_subnet_network_acl_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"network_acl_id": "NetworkAclId",
		"subnet_id":      "SubnetId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}