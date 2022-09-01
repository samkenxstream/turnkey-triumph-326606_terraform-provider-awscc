// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package supportapp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_supportapp_slack_channel_configuration", slackChannelConfigurationDataSourceType)
}

// slackChannelConfigurationDataSourceType returns the Terraform awscc_supportapp_slack_channel_configuration data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SupportApp::SlackChannelConfiguration resource type.
func slackChannelConfigurationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"channel_id": {
			// Property: ChannelId
			// CloudFormation resource type schema:
			// {
			//   "description": "The channel ID in Slack, which identifies a channel within a workspace.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^\\S+$",
			//   "type": "string"
			// }
			Description: "The channel ID in Slack, which identifies a channel within a workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"channel_name": {
			// Property: ChannelName
			// CloudFormation resource type schema:
			// {
			//   "description": "The channel name in Slack.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^.+$",
			//   "type": "string"
			// }
			Description: "The channel name in Slack.",
			Type:        types.StringType,
			Computed:    true,
		},
		"channel_role_arn": {
			// Property: ChannelRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.",
			//   "maxLength": 2048,
			//   "minLength": 31,
			//   "pattern": "^arn:aws[-a-z0-9]*:iam::[0-9]{12}:role\\/(.+)$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.",
			Type:        types.StringType,
			Computed:    true,
		},
		"notify_on_add_correspondence_to_case": {
			// Property: NotifyOnAddCorrespondenceToCase
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to notify when a correspondence is added to a case.",
			//   "type": "boolean"
			// }
			Description: "Whether to notify when a correspondence is added to a case.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"notify_on_case_severity": {
			// Property: NotifyOnCaseSeverity
			// CloudFormation resource type schema:
			// {
			//   "description": "The severity level of a support case that a customer wants to get notified for.",
			//   "enum": [
			//     "none",
			//     "all",
			//     "high"
			//   ],
			//   "type": "string"
			// }
			Description: "The severity level of a support case that a customer wants to get notified for.",
			Type:        types.StringType,
			Computed:    true,
		},
		"notify_on_create_or_reopen_case": {
			// Property: NotifyOnCreateOrReopenCase
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to notify when a case is created or reopened.",
			//   "type": "boolean"
			// }
			Description: "Whether to notify when a case is created or reopened.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"notify_on_resolve_case": {
			// Property: NotifyOnResolveCase
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to notify when a case is resolved.",
			//   "type": "boolean"
			// }
			Description: "Whether to notify when a case is resolved.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"team_id": {
			// Property: TeamId
			// CloudFormation resource type schema:
			// {
			//   "description": "The team ID in Slack, which uniquely identifies a workspace.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^\\S+$",
			//   "type": "string"
			// }
			Description: "The team ID in Slack, which uniquely identifies a workspace.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SupportApp::SlackChannelConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SupportApp::SlackChannelConfiguration").WithTerraformTypeName("awscc_supportapp_slack_channel_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"channel_id":                           "ChannelId",
		"channel_name":                         "ChannelName",
		"channel_role_arn":                     "ChannelRoleArn",
		"notify_on_add_correspondence_to_case": "NotifyOnAddCorrespondenceToCase",
		"notify_on_case_severity":              "NotifyOnCaseSeverity",
		"notify_on_create_or_reopen_case":      "NotifyOnCreateOrReopenCase",
		"notify_on_resolve_case":               "NotifyOnResolveCase",
		"team_id":                              "TeamId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
