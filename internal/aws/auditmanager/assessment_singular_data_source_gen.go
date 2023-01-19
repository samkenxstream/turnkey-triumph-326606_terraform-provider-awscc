// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package auditmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_auditmanager_assessment", assessmentDataSource)
}

// assessmentDataSource returns the Terraform awscc_auditmanager_assessment data source.
// This Terraform data source corresponds to the CloudFormation AWS::AuditManager::Assessment resource.
func assessmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the assessment.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:.*:auditmanager:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssessmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"assessment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AssessmentReportsDestination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The destination in which evidence reports are stored for the specified assessment.",
		//	  "properties": {
		//	    "Destination": {
		//	      "description": "The URL of the specified Amazon S3 bucket.",
		//	      "type": "string"
		//	    },
		//	    "DestinationType": {
		//	      "description": "The destination type, such as Amazon S3.",
		//	      "enum": [
		//	        "S3"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"assessment_reports_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Destination
				"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URL of the specified Amazon S3 bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DestinationType
				"destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The destination type, such as Amazon S3.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The destination in which evidence reports are stored for the specified assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AwsAccount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The AWS account associated with the assessment.",
		//	  "properties": {
		//	    "EmailAddress": {
		//	      "description": "The unique identifier for the email account.",
		//	      "maxLength": 320,
		//	      "minLength": 1,
		//	      "pattern": "^.*@.*$",
		//	      "type": "string"
		//	    },
		//	    "Id": {
		//	      "description": "The identifier for the specified AWS account.",
		//	      "maxLength": 12,
		//	      "minLength": 12,
		//	      "pattern": "^[0-9]{12}$",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "The name of the specified AWS account.",
		//	      "maxLength": 50,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"aws_account": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EmailAddress
				"email_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The unique identifier for the email account.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Id
				"id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier for the specified AWS account.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the specified AWS account.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The AWS account associated with the assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sequence of characters that identifies when the event occurred.",
		//	  "type": "number"
		//	}
		"creation_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The sequence of characters that identifies when the event occurred.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Delegations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of delegations.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The assignment of a control set to a delegate for review.",
		//	    "properties": {
		//	      "AssessmentId": {
		//	        "maxLength": 36,
		//	        "minLength": 36,
		//	        "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	        "type": "string"
		//	      },
		//	      "AssessmentName": {
		//	        "description": "The name of the related assessment.",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9-_\\.]+$",
		//	        "type": "string"
		//	      },
		//	      "Comment": {
		//	        "description": "The comment related to the delegation.",
		//	        "maxLength": 350,
		//	        "pattern": "^[\\w\\W\\s\\S]*$",
		//	        "type": "string"
		//	      },
		//	      "ControlSetId": {
		//	        "description": "The identifier for the specified control set.",
		//	        "maxLength": 300,
		//	        "minLength": 1,
		//	        "pattern": "^[\\w\\W\\s\\S]*$",
		//	        "type": "string"
		//	      },
		//	      "CreatedBy": {
		//	        "description": "The IAM user or role that performed the action.",
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "pattern": "^arn:.*:*:.*",
		//	        "type": "string"
		//	      },
		//	      "CreationTime": {
		//	        "description": "The sequence of characters that identifies when the event occurred.",
		//	        "type": "number"
		//	      },
		//	      "Id": {
		//	        "maxLength": 36,
		//	        "minLength": 36,
		//	        "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	        "type": "string"
		//	      },
		//	      "LastUpdated": {
		//	        "description": "The sequence of characters that identifies when the event occurred.",
		//	        "type": "number"
		//	      },
		//	      "RoleArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "pattern": "^arn:.*:iam:.*",
		//	        "type": "string"
		//	      },
		//	      "RoleType": {
		//	        "description": " The IAM role type.",
		//	        "enum": [
		//	          "PROCESS_OWNER",
		//	          "RESOURCE_OWNER"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Status": {
		//	        "description": "The status of the delegation.",
		//	        "enum": [
		//	          "IN_PROGRESS",
		//	          "UNDER_REVIEW",
		//	          "COMPLETE"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"delegations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AssessmentId
					"assessment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: AssessmentName
					"assessment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the related assessment.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Comment
					"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The comment related to the delegation.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ControlSetId
					"control_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The identifier for the specified control set.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CreatedBy
					"created_by": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The IAM user or role that performed the action.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: CreationTime
					"creation_time": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "The sequence of characters that identifies when the event occurred.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LastUpdated
					"last_updated": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "The sequence of characters that identifies when the event occurred.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: RoleArn
					"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: RoleType
					"role_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: " The IAM role type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The status of the delegation.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of delegations.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the specified assessment.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the specified assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FrameworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the specified framework.",
		//	  "maxLength": 36,
		//	  "minLength": 32,
		//	  "pattern": "^([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|.*\\S.*)$",
		//	  "type": "string"
		//	}
		"framework_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the specified framework.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the related assessment.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-_\\.]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the related assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Roles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of roles for the specified assessment.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The wrapper that contains AWS Audit Manager role information, such as the role type and IAM ARN.",
		//	    "properties": {
		//	      "RoleArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "pattern": "^arn:.*:iam:.*",
		//	        "type": "string"
		//	      },
		//	      "RoleType": {
		//	        "description": " The IAM role type.",
		//	        "enum": [
		//	          "PROCESS_OWNER",
		//	          "RESOURCE_OWNER"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"roles": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RoleArn
					"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: RoleType
					"role_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: " The IAM role type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of roles for the specified assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Scope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
		//	  "properties": {
		//	    "AwsAccounts": {
		//	      "description": "The AWS accounts included in scope.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The AWS account associated with the assessment.",
		//	        "properties": {
		//	          "EmailAddress": {
		//	            "description": "The unique identifier for the email account.",
		//	            "maxLength": 320,
		//	            "minLength": 1,
		//	            "pattern": "^.*@.*$",
		//	            "type": "string"
		//	          },
		//	          "Id": {
		//	            "description": "The identifier for the specified AWS account.",
		//	            "maxLength": 12,
		//	            "minLength": 12,
		//	            "pattern": "^[0-9]{12}$",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the specified AWS account.",
		//	            "maxLength": 50,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "AwsServices": {
		//	      "description": "The AWS services included in scope.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An AWS service such as Amazon S3, AWS CloudTrail, and so on.",
		//	        "properties": {
		//	          "ServiceName": {
		//	            "description": "The name of the AWS service.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scope": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AwsAccounts
				"aws_accounts": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: EmailAddress
							"email_address": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The unique identifier for the email account.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Id
							"id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The identifier for the specified AWS account.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the specified AWS account.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The AWS accounts included in scope.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AwsServices
				"aws_services": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ServiceName
							"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the AWS service.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The AWS services included in scope.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the specified assessment. ",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "INACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the specified assessment. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags associated with the assessment.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags associated with the assessment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AuditManager::Assessment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AuditManager::Assessment").WithTerraformTypeName("awscc_auditmanager_assessment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"assessment_id":                  "AssessmentId",
		"assessment_name":                "AssessmentName",
		"assessment_reports_destination": "AssessmentReportsDestination",
		"aws_account":                    "AwsAccount",
		"aws_accounts":                   "AwsAccounts",
		"aws_services":                   "AwsServices",
		"comment":                        "Comment",
		"control_set_id":                 "ControlSetId",
		"created_by":                     "CreatedBy",
		"creation_time":                  "CreationTime",
		"delegations":                    "Delegations",
		"description":                    "Description",
		"destination":                    "Destination",
		"destination_type":               "DestinationType",
		"email_address":                  "EmailAddress",
		"framework_id":                   "FrameworkId",
		"id":                             "Id",
		"key":                            "Key",
		"last_updated":                   "LastUpdated",
		"name":                           "Name",
		"role_arn":                       "RoleArn",
		"role_type":                      "RoleType",
		"roles":                          "Roles",
		"scope":                          "Scope",
		"service_name":                   "ServiceName",
		"status":                         "Status",
		"tags":                           "Tags",
		"value":                          "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
