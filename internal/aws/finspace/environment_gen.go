// Code generated by generators/resource/main.go; DO NOT EDIT.

package finspace

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_finspace_environment", environmentResourceType)
}

// environmentResourceType returns the Terraform awscc_finspace_environment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FinSpace::Environment resource type.
func environmentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS account ID associated with the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "AWS account ID associated with the Environment",
			Type:        types.StringType,
			Computed:    true,
		},
		"dedicated_service_account_id": {
			// Property: DedicatedServiceAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "ID for FinSpace created account used to store Environment artifacts",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ID for FinSpace created account used to store Environment artifacts",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Description of the Environment",
			Type:        types.StringType,
			Optional:    true,
		},
		"environment_arn": {
			// Property: EnvironmentArn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ARN of the Environment",
			Type:        types.StringType,
			Computed:    true,
		},
		"environment_id": {
			// Property: EnvironmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier for representing FinSpace Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Unique identifier for representing FinSpace Environment",
			Type:        types.StringType,
			Computed:    true,
		},
		"environment_url": {
			// Property: EnvironmentUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "URL used to login to the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "URL used to login to the Environment",
			Type:        types.StringType,
			Computed:    true,
		},
		"federation_mode": {
			// Property: FederationMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Federation mode used with the Environment",
			//   "enum": [
			//     "LOCAL",
			//     "FEDERATED"
			//   ],
			//   "type": "string"
			// }
			Description: "Federation mode used with the Environment",
			Type:        types.StringType,
			Optional:    true,
		},
		"federation_parameters": {
			// Property: FederationParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "Additional parameters to identify Federation mode",
			//   "properties": {
			//     "ApplicationCallBackURL": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "AttributeMap": {
			//       "description": "Attribute map for SAML configuration",
			//       "type": "object"
			//     },
			//     "FederationProviderName": {
			//       "description": "Federation provider name to link with the Environment",
			//       "maxLength": 32,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "FederationURN": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SamlMetadataDocument": {
			//       "description": "SAML metadata document to link the federation provider to the Environment",
			//       "maxLength": 10000000,
			//       "minLength": 1000,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SamlMetadataURL": {
			//       "description": "SAML metadata URL to link with the Environment",
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Additional parameters to identify Federation mode",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"application_call_back_url": {
						// Property: ApplicationCallBackURL
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
					"attribute_map": {
						// Property: AttributeMap
						Description: "Attribute map for SAML configuration",
						Type:        types.MapType{ElemType: types.StringType},
						Optional:    true,
					},
					"federation_provider_name": {
						// Property: FederationProviderName
						Description: "Federation provider name to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
					"federation_urn": {
						// Property: FederationURN
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
					"saml_metadata_document": {
						// Property: SamlMetadataDocument
						Description: "SAML metadata document to link the federation provider to the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
					"saml_metadata_url": {
						// Property: SamlMetadataURL
						Description: "SAML metadata URL to link with the Environment",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "KMS key used to encrypt customer data within FinSpace Environment infrastructure",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// KmsKeyId is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the Environment",
			Type:        types.StringType,
			Required:    true,
		},
		"sage_maker_studio_domain_url": {
			// Property: SageMakerStudioDomainUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "SageMaker Studio Domain URL associated with the Environment",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "SageMaker Studio Domain URL associated with the Environment",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the Environment",
			//   "enum": [
			//     "CREATE_REQUESTED",
			//     "CREATING",
			//     "CREATED",
			//     "DELETE_REQUESTED",
			//     "DELETING",
			//     "DELETED",
			//     "FAILED_CREATION",
			//     "FAILED_DELETION",
			//     "RETRY_DELETION",
			//     "SUSPENDED"
			//   ],
			//   "type": "string"
			// }
			Description: "State of the Environment",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FinSpace::Environment").WithTerraformTypeName("awscc_finspace_environment").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_finspace_environment", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
