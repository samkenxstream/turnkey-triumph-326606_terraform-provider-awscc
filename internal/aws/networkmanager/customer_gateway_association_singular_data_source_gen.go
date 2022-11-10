// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_customer_gateway_association", customerGatewayAssociationDataSource)
}

// customerGatewayAssociationDataSource returns the Terraform awscc_networkmanager_customer_gateway_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::CustomerGatewayAssociation resource.
func customerGatewayAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"customer_gateway_arn": {
			// Property: CustomerGatewayArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the customer gateway.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the customer gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the device",
			//	  "type": "string"
			//	}
			Description: "The ID of the device",
			Type:        types.StringType,
			Computed:    true,
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the global network.",
			//	  "type": "string"
			//	}
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Computed:    true,
		},
		"link_id": {
			// Property: LinkId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the link",
			//	  "type": "string"
			//	}
			Description: "The ID of the link",
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
		Description: "Data Source schema for AWS::NetworkManager::CustomerGatewayAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::CustomerGatewayAssociation").WithTerraformTypeName("awscc_networkmanager_customer_gateway_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_gateway_arn": "CustomerGatewayArn",
		"device_id":            "DeviceId",
		"global_network_id":    "GlobalNetworkId",
		"link_id":              "LinkId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
