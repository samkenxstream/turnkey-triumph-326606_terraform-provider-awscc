// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_kinesis_stream", streamDataSource)
}

// streamDataSource returns the Terraform awscc_kinesis_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::Kinesis::Stream resource.
func streamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon resource name (ARN) of the Kinesis stream",
			//	  "type": "string"
			//	}
			Description: "The Amazon resource name (ARN) of the Kinesis stream",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the Kinesis stream.",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9_.-]+$",
			//	  "type": "string"
			//	}
			Description: "The name of the Kinesis stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"retention_period_hours": {
			// Property: RetentionPeriodHours
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of hours for the data records that are stored in shards to remain accessible.",
			//	  "minimum": 24,
			//	  "type": "integer"
			//	}
			Description: "The number of hours for the data records that are stored in shards to remain accessible.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"shard_count": {
			// Property: ShardCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.",
			//	  "minimum": 1,
			//	  "type": "integer"
			//	}
			Description: "The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"stream_encryption": {
			// Property: StreamEncryption
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.",
			//	  "properties": {
			//	    "EncryptionType": {
			//	      "description": "The encryption type to use. The only valid value is KMS. ",
			//	      "enum": [
			//	        "KMS"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "KeyId": {
			//	      "description": "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "EncryptionType",
			//	    "KeyId"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "The encryption type to use. The only valid value is KMS. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"key_id": {
						// Property: KeyId
						Description: "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"stream_mode_details": {
			// Property: StreamModeDetails
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "default": {
			//	    "StreamMode": "PROVISIONED"
			//	  },
			//	  "description": "The mode in which the stream is running.",
			//	  "properties": {
			//	    "StreamMode": {
			//	      "description": "The mode of the stream",
			//	      "enum": [
			//	        "ON_DEMAND",
			//	        "PROVISIONED"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "StreamMode"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The mode in which the stream is running.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stream_mode": {
						// Property: StreamMode
						Description: "The mode of the stream",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 255,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Kinesis::Stream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kinesis::Stream").WithTerraformTypeName("awscc_kinesis_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"encryption_type":        "EncryptionType",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"name":                   "Name",
		"retention_period_hours": "RetentionPeriodHours",
		"shard_count":            "ShardCount",
		"stream_encryption":      "StreamEncryption",
		"stream_mode":            "StreamMode",
		"stream_mode_details":    "StreamModeDetails",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
