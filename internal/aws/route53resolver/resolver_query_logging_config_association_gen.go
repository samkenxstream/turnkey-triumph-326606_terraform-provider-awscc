// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver

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
	registry.AddResourceTypeFactory("awscc_route53resolver_resolver_query_logging_config_association", resolverQueryLoggingConfigAssociationResourceType)
}

// resolverQueryLoggingConfigAssociationResourceType returns the Terraform awscc_route53resolver_resolver_query_logging_config_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation resource type.
func resolverQueryLoggingConfigAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "Rfc3339TimeString",
			//   "maxLength": 40,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 40),
			},
		},
		"error": {
			// Property: Error
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigAssociationError",
			//   "enum": [
			//     "NONE",
			//     "DESTINATION_NOT_FOUND",
			//     "ACCESS_DENIED"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigAssociationError",
			Type:        types.StringType,
			Computed:    true,
		},
		"error_message": {
			// Property: ErrorMessage
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigAssociationErrorMessage",
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigAssociationErrorMessage",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Id",
			Type:        types.StringType,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
		},
		"resolver_query_log_config_id": {
			// Property: ResolverQueryLogConfigId
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigId",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			// ResolverQueryLogConfigId is a force-new attribute.
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "ResourceId",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "ResourceId",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			// ResourceId is a force-new attribute.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "ResolverQueryLogConfigAssociationStatus",
			//   "enum": [
			//     "CREATING",
			//     "ACTIVE",
			//     "ACTION_NEEDED",
			//     "DELETING",
			//     "FAILED",
			//     "OVERRIDDEN"
			//   ],
			//   "type": "string"
			// }
			Description: "ResolverQueryLogConfigAssociationStatus",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation").WithTerraformTypeName("awscc_route53resolver_resolver_query_logging_config_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                "CreationTime",
		"error":                        "Error",
		"error_message":                "ErrorMessage",
		"id":                           "Id",
		"resolver_query_log_config_id": "ResolverQueryLogConfigId",
		"resource_id":                  "ResourceId",
		"status":                       "Status",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53resolver_resolver_query_logging_config_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
