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
	registry.AddDataSourceFactory("awscc_ec2_host", hostDataSource)
}

// hostDataSource returns the Terraform awscc_ec2_host data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Host resource.
func hostDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_placement": {
			// Property: AutoPlacement
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
			//	  "type": "string"
			//	}
			Description: "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Availability Zone in which to allocate the Dedicated Host.",
			//	  "type": "string"
			//	}
			Description: "The Availability Zone in which to allocate the Dedicated Host.",
			Type:        types.StringType,
			Computed:    true,
		},
		"host_id": {
			// Property: HostId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Id of the host created.",
			//	  "type": "string"
			//	}
			Description: "Id of the host created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"host_recovery": {
			// Property: HostRecovery
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
			//	  "type": "string"
			//	}
			Description: "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_family": {
			// Property: InstanceFamily
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
			//	  "type": "string"
			//	}
			Description: "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_type": {
			// Property: InstanceType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
			//	  "type": "string"
			//	}
			Description: "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
			Type:        types.StringType,
			Computed:    true,
		},
		"outpost_arn": {
			// Property: OutpostArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
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
		Description: "Data Source schema for AWS::EC2::Host",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Host").WithTerraformTypeName("awscc_ec2_host")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_placement":    "AutoPlacement",
		"availability_zone": "AvailabilityZone",
		"host_id":           "HostId",
		"host_recovery":     "HostRecovery",
		"instance_family":   "InstanceFamily",
		"instance_type":     "InstanceType",
		"outpost_arn":       "OutpostArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
