// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codestarnotifications

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codestarnotifications_notification_rule", notificationRuleDataSource)
}

// notificationRuleDataSource returns the Terraform awscc_codestarnotifications_notification_rule data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeStarNotifications::NotificationRule resource.
func notificationRuleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "^arn:aws[^:\\s]*:codestar-notifications:[^:\\s]+:\\d{12}:notificationrule\\/(.*\\S)?$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"created_by": {
			// Property: CreatedBy
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"detail_type": {
			// Property: DetailType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "BASIC",
			//	    "FULL"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"event_type_id": {
			// Property: EventTypeId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"event_type_ids": {
			// Property: EventTypeIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "maxLength": 200,
			//	    "minLength": 1,
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "pattern": "[A-Za-z0-9\\-_ ]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "^arn:aws[^:\\s]*:[^:\\s]*:[^:\\s]*:[0-9]{12}:[^\\s]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "ENABLED",
			//	    "DISABLED"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "object"
			//	}
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"target_address": {
			// Property: TargetAddress
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "TargetAddress": {
			//	        "type": "string"
			//	      },
			//	      "TargetType": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "TargetType",
			//	      "TargetAddress"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 10,
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"target_address": {
						// Property: TargetAddress
						Type:     types.StringType,
						Computed: true,
					},
					"target_type": {
						// Property: TargetType
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::CodeStarNotifications::NotificationRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarNotifications::NotificationRule").WithTerraformTypeName("awscc_codestarnotifications_notification_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":            "Arn",
		"created_by":     "CreatedBy",
		"detail_type":    "DetailType",
		"event_type_id":  "EventTypeId",
		"event_type_ids": "EventTypeIds",
		"name":           "Name",
		"resource":       "Resource",
		"status":         "Status",
		"tags":           "Tags",
		"target_address": "TargetAddress",
		"target_type":    "TargetType",
		"targets":        "Targets",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
