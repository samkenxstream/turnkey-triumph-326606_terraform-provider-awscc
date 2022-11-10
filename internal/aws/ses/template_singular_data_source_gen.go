// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ses_template", templateDataSource)
}

// templateDataSource returns the Terraform awscc_ses_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::SES::Template resource.
func templateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"template": {
			// Property: Template
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The content of the email, composed of a subject line, an HTML part, and a text-only part",
			//	  "properties": {
			//	    "HtmlPart": {
			//	      "description": "The HTML body of the email.",
			//	      "type": "string"
			//	    },
			//	    "SubjectPart": {
			//	      "description": "The subject line of the email.",
			//	      "type": "string"
			//	    },
			//	    "TemplateName": {
			//	      "description": "The name of the template.",
			//	      "maxLength": 64,
			//	      "minLength": 1,
			//	      "pattern": "^[a-zA-Z0-9_-]{1,64}$",
			//	      "type": "string"
			//	    },
			//	    "TextPart": {
			//	      "description": "The email body that is visible to recipients whose email clients do not display HTML content.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "SubjectPart"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The content of the email, composed of a subject line, an HTML part, and a text-only part",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"html_part": {
						// Property: HtmlPart
						Description: "The HTML body of the email.",
						Type:        types.StringType,
						Computed:    true,
					},
					"subject_part": {
						// Property: SubjectPart
						Description: "The subject line of the email.",
						Type:        types.StringType,
						Computed:    true,
					},
					"template_name": {
						// Property: TemplateName
						Description: "The name of the template.",
						Type:        types.StringType,
						Computed:    true,
					},
					"text_part": {
						// Property: TextPart
						Description: "The email body that is visible to recipients whose email clients do not display HTML content.",
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
		Description: "Data Source schema for AWS::SES::Template",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::Template").WithTerraformTypeName("awscc_ses_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"html_part":     "HtmlPart",
		"id":            "Id",
		"subject_part":  "SubjectPart",
		"template":      "Template",
		"template_name": "TemplateName",
		"text_part":     "TextPart",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
