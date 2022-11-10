// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_device_fleet", deviceFleetResource)
}

// deviceFleetResource returns the Terraform awscc_sagemaker_device_fleet resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::DeviceFleet resource.
func deviceFleetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Description for the edge device fleet",
			//	  "maxLength": 800,
			//	  "minLength": 0,
			//	  "pattern": "[\\S\\s]+",
			//	  "type": "string"
			//	}
			Description: "Description for the edge device fleet",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 800),
				validate.StringMatch(regexp.MustCompile("[\\S\\s]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"device_fleet_name": {
			// Property: DeviceFleetName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the edge device fleet",
			//	  "maxLength": 63,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9](-*_*[a-zA-Z0-9])*$",
			//	  "type": "string"
			//	}
			Description: "The name of the edge device fleet",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9](-*_*[a-zA-Z0-9])*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"output_config": {
			// Property: OutputConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "S3 bucket and an ecryption key id (if available) to store outputs for the fleet",
			//	  "properties": {
			//	    "KmsKeyId": {
			//	      "description": "The KMS key id used for encryption on the S3 bucket",
			//	      "maxLength": 2048,
			//	      "minLength": 1,
			//	      "pattern": "[a-zA-Z0-9:_-]+",
			//	      "type": "string"
			//	    },
			//	    "S3OutputLocation": {
			//	      "description": "The Amazon Simple Storage (S3) bucket URI",
			//	      "maxLength": 1024,
			//	      "pattern": "^s3://([^/]+)/?(.*)$",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "S3OutputLocation"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "S3 bucket and an ecryption key id (if available) to store outputs for the fleet",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Description: "The KMS key id used for encryption on the S3 bucket",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
							validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"s3_output_location": {
						// Property: S3OutputLocation
						Description: "The Amazon Simple Storage (S3) bucket URI",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(1024),
							validate.StringMatch(regexp.MustCompile("^s3://([^/]+)/?(.*)$"), ""),
						},
					},
				},
			),
			Required: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Role associated with the device fleet",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
			//	  "type": "string"
			//	}
			Description: "Role associated with the device fleet",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Associate tags with the resource",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "Key-value pair to associate as a tag for the resource",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "Associate tags with the resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::SageMaker::DeviceFleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::DeviceFleet").WithTerraformTypeName("awscc_sagemaker_device_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":        "Description",
		"device_fleet_name":  "DeviceFleetName",
		"key":                "Key",
		"kms_key_id":         "KmsKeyId",
		"output_config":      "OutputConfig",
		"role_arn":           "RoleArn",
		"s3_output_location": "S3OutputLocation",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
