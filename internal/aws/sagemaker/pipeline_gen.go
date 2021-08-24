// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("awscc_sagemaker_pipeline", pipelineResourceType)
}

// pipelineResourceType returns the Terraform awscc_sagemaker_pipeline resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Pipeline resource type.
func pipelineResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"pipeline_definition": {
			// Property: PipelineDefinition
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Required: true,
		},
		"pipeline_description": {
			// Property: PipelineDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the Pipeline.",
			//   "maxLength": 3072,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the Pipeline.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 3072),
			},
		},
		"pipeline_display_name": {
			// Property: PipelineDisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The display name of the Pipeline.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"pipeline_name": {
			// Property: PipelineName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Pipeline.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			// PipelineName is a force-new attribute.
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SageMaker::Pipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Pipeline").WithTerraformTypeName("awscc_sagemaker_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                   "Key",
		"pipeline_definition":   "PipelineDefinition",
		"pipeline_description":  "PipelineDescription",
		"pipeline_display_name": "PipelineDisplayName",
		"pipeline_name":         "PipelineName",
		"role_arn":              "RoleArn",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_sagemaker_pipeline", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
