// Code generated by generators/resource/main.go; DO NOT EDIT.

package rekognition

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
	registry.AddResourceFactory("awscc_rekognition_project", projectResource)
}

// projectResource returns the Terraform awscc_rekognition_project resource.
// This Terraform resource corresponds to the CloudFormation AWS::Rekognition::Project resource.
func projectResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "pattern": "(^arn:[a-z\\d-]+:rekognition:[a-z\\d-]+:\\d{12}:project/[a-zA-Z0-9_.\\-]{1,255}/[0-9]+$)",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the project",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9][a-zA-Z0-9_\\-]*",
			//	  "type": "string"
			//	}
			Description: "The name of the project",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9][a-zA-Z0-9_\\-]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "The AWS::Rekognition::Project type creates an Amazon Rekognition CustomLabels Project. A project is a grouping of the resources needed to create and manage Dataset and ProjectVersions.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Rekognition::Project").WithTerraformTypeName("awscc_rekognition_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":          "Arn",
		"project_name": "ProjectName",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(15).WithDeleteTimeoutInMinutes(15)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
