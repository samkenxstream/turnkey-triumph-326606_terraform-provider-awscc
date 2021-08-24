// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_rds_db_proxy_target_group", dBProxyTargetGroupResourceType)
}

// dBProxyTargetGroupResourceType returns the Terraform awscc_rds_db_proxy_target_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RDS::DBProxyTargetGroup resource type.
func dBProxyTargetGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"connection_pool_configuration_info": {
			// Property: ConnectionPoolConfigurationInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ConnectionBorrowTimeout": {
			//       "description": "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
			//       "type": "integer"
			//     },
			//     "InitQuery": {
			//       "description": "One or more SQL statements for the proxy to run when opening each new database connection.",
			//       "type": "string"
			//     },
			//     "MaxConnectionsPercent": {
			//       "description": "The maximum size of the connection pool for each target in a target group.",
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MaxIdleConnectionsPercent": {
			//       "description": "Controls how actively the proxy closes idle database connections in the connection pool.",
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "SessionPinningFilters": {
			//       "description": "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"connection_borrow_timeout": {
						// Property: ConnectionBorrowTimeout
						Description: "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"init_query": {
						// Property: InitQuery
						Description: "One or more SQL statements for the proxy to run when opening each new database connection.",
						Type:        types.StringType,
						Optional:    true,
					},
					"max_connections_percent": {
						// Property: MaxConnectionsPercent
						Description: "The maximum size of the connection pool for each target in a target group.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 100),
						},
					},
					"max_idle_connections_percent": {
						// Property: MaxIdleConnectionsPercent
						Description: "Controls how actively the proxy closes idle database connections in the connection pool.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 100),
						},
					},
					"session_pinning_filters": {
						// Property: SessionPinningFilters
						Description: "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"db_cluster_identifiers": {
			// Property: DBClusterIdentifiers
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"db_instance_identifiers": {
			// Property: DBInstanceIdentifiers
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"db_proxy_name": {
			// Property: DBProxyName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the proxy.",
			//   "maxLength": 64,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The identifier for the proxy.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 64),
			},
			// DBProxyName is a force-new attribute.
		},
		"target_group_arn": {
			// Property: TargetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) representing the target group.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) representing the target group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_group_name": {
			// Property: TargetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the DBProxyTargetGroup",
			//   "enum": [
			//     "default"
			//   ],
			//   "type": "string"
			// }
			Description: "The identifier for the DBProxyTargetGroup",
			Type:        types.StringType,
			Required:    true,
			// TargetGroupName is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::RDS::DBProxyTargetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyTargetGroup").WithTerraformTypeName("awscc_rds_db_proxy_target_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_borrow_timeout":          "ConnectionBorrowTimeout",
		"connection_pool_configuration_info": "ConnectionPoolConfigurationInfo",
		"db_cluster_identifiers":             "DBClusterIdentifiers",
		"db_instance_identifiers":            "DBInstanceIdentifiers",
		"db_proxy_name":                      "DBProxyName",
		"init_query":                         "InitQuery",
		"max_connections_percent":            "MaxConnectionsPercent",
		"max_idle_connections_percent":       "MaxIdleConnectionsPercent",
		"session_pinning_filters":            "SessionPinningFilters",
		"target_group_arn":                   "TargetGroupArn",
		"target_group_name":                  "TargetGroupName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_rds_db_proxy_target_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
