// Code generated by generators/resource/main.go; DO NOT EDIT.

package kms

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
	registry.AddResourceTypeFactory("awscc_kms_key", keyResourceType)
}

// keyResourceType returns the Terraform awscc_kms_key resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::KMS::Key resource type.
func keyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			//   "maxLength": 8192,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.",
			Type:        types.StringType,
			Optional:    true,
		},
		"enable_key_rotation": {
			// Property: EnableKeyRotation
			// CloudFormation resource type schema:
			// {
			//   "description": "Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.",
			//   "type": "boolean"
			// }
			Description: "Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"key_id": {
			// Property: KeyId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"key_policy": {
			// Property: KeyPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			//   "type": "string"
			// }
			Description: "The key policy that authorizes use of the CMK. The key policy must observe the following rules.",
			Type:        types.StringType,
			Required:    true,
		},
		"key_spec": {
			// Property: KeySpec
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.",
			//   "enum": [
			//     "SYMMETRIC_DEFAULT",
			//     "RSA_2048",
			//     "RSA_3072",
			//     "RSA_4096",
			//     "ECC_NIST_P256",
			//     "ECC_NIST_P384",
			//     "ECC_NIST_P521",
			//     "ECC_SECG_P256K1"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.",
			Type:        types.StringType,
			Optional:    true,
		},
		"key_usage": {
			// Property: KeyUsage
			// CloudFormation resource type schema:
			// {
			//   "description": "Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.",
			//   "enum": [
			//     "ENCRYPT_DECRYPT",
			//     "SIGN_VERIFY"
			//   ],
			//   "type": "string"
			// }
			Description: "Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.",
			Type:        types.StringType,
			Optional:    true,
		},
		"multi_region": {
			// Property: MultiRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"pending_window_in_days": {
			// Property: PendingWindowInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			//   "type": "integer"
			// }
			Description: "Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
			Type:        types.NumberType,
			Optional:    true,
			// PendingWindowInDays is a write-only attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::KMS::Key resource specifies a customer master key (CMK) in AWS Key Management Service (AWS KMS). Authorized users can use the CMK to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use CMKs to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::Key").WithTerraformTypeName("awscc_kms_key").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/PendingWindowInDays",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_kms_key", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
