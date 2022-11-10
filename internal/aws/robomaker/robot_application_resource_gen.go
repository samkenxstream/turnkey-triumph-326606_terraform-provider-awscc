// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_robomaker_robot_application", robotApplicationResource)
}

// robotApplicationResource returns the Terraform awscc_robomaker_robot_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::RoboMaker::RobotApplication resource.
func robotApplicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"current_revision_id": {
			// Property: CurrentRevisionId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The revision ID of robot application.",
			//	  "maxLength": 40,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The revision ID of robot application.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 40),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"environment": {
			// Property: Environment
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The URI of the Docker image for the robot application.",
			//	  "type": "string"
			//	}
			Description: "The URI of the Docker image for the robot application.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the robot application.",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The name of the robot application.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"robot_software_suite": {
			// Property: RobotSoftwareSuite
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The robot software suite used by the robot application.",
			//	  "properties": {
			//	    "Name": {
			//	      "description": "The name of robot software suite.",
			//	      "enum": [
			//	        "ROS",
			//	        "ROS2",
			//	        "General"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "Version": {
			//	      "description": "The version of robot software suite.",
			//	      "enum": [
			//	        "Kinetic",
			//	        "Melodic",
			//	        "Dashing"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Name"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The robot software suite used by the robot application.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of robot software suite.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ROS",
								"ROS2",
								"General",
							}),
						},
					},
					"version": {
						// Property: Version
						Description: "The version of robot software suite.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"Kinetic",
								"Melodic",
								"Dashing",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"sources": {
			// Property: Sources
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The sources of the robot application.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Architecture": {
			//	        "description": "The architecture of robot application.",
			//	        "enum": [
			//	          "X86_64",
			//	          "ARM64",
			//	          "ARMHF"
			//	        ],
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "S3Bucket": {
			//	        "description": "The Arn of the S3Bucket that stores the robot application source.",
			//	        "type": "string"
			//	      },
			//	      "S3Key": {
			//	        "description": "The s3 key of robot application source.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "S3Bucket",
			//	      "S3Key",
			//	      "Architecture"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The sources of the robot application.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"architecture": {
						// Property: Architecture
						Description: "The architecture of robot application.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
							validate.StringInSlice([]string{
								"X86_64",
								"ARM64",
								"ARMHF",
							}),
						},
					},
					"s3_bucket": {
						// Property: S3Bucket
						Description: "The Arn of the S3Bucket that stores the robot application source.",
						Type:        types.StringType,
						Required:    true,
					},
					"s3_key": {
						// Property: S3Key
						Description: "The s3 key of robot application source.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "A key-value pair to associate with a resource.",
			//	  "patternProperties": {
			//	    "": {
			//	      "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	      "maxLength": 256,
			//	      "minLength": 1,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
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
		Description: "AWS::RoboMaker::RobotApplication resource creates an AWS RoboMaker RobotApplication. Robot application can be used in AWS RoboMaker Simulation Jobs.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::RobotApplication").WithTerraformTypeName("awscc_robomaker_robot_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":         "Architecture",
		"arn":                  "Arn",
		"current_revision_id":  "CurrentRevisionId",
		"environment":          "Environment",
		"name":                 "Name",
		"robot_software_suite": "RobotSoftwareSuite",
		"s3_bucket":            "S3Bucket",
		"s3_key":               "S3Key",
		"sources":              "Sources",
		"tags":                 "Tags",
		"version":              "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
