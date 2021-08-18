// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_mediaconnect_flow_vpc_interface", flowVpcInterfaceResourceType)
}

// flowVpcInterfaceResourceType returns the Terraform aws_mediaconnect_flow_vpc_interface resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaConnect::FlowVpcInterface resource type.
func flowVpcInterfaceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Required:    true,
			// FlowArn is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
			//   "type": "string"
			// }
			Description: "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"network_interface_ids": {
			// Property: NetworkInterfaceIds
			// CloudFormation resource type schema:
			// {
			//   "description": "IDs of the network interfaces created in customer's account by MediaConnect.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "IDs of the network interfaces created in customer's account by MediaConnect.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
			//   "type": "string"
			// }
			Description: "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
			Type:        types.StringType,
			Required:    true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "Security Group IDs to be used on ENI.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Security Group IDs to be used on ENI.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "Subnet must be in the AZ of the Flow",
			//   "type": "string"
			// }
			Description: "Subnet must be in the AZ of the Flow",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::FlowVpcInterface",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowVpcInterface").WithTerraformTypeName("aws_mediaconnect_flow_vpc_interface").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_mediaconnect_flow_vpc_interface", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}