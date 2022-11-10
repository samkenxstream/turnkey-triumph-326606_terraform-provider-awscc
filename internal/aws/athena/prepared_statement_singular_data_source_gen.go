// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_athena_prepared_statement", preparedStatementDataSource)
}

// preparedStatementDataSource returns the Terraform awscc_athena_prepared_statement data source.
// This Terraform data source corresponds to the CloudFormation AWS::Athena::PreparedStatement resource.
func preparedStatementDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the prepared statement.",
			//	  "maxLength": 1024,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The description of the prepared statement.",
			Type:        types.StringType,
			Computed:    true,
		},
		"query_statement": {
			// Property: QueryStatement
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The query string for the prepared statement.",
			//	  "maxLength": 262144,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The query string for the prepared statement.",
			Type:        types.StringType,
			Computed:    true,
		},
		"statement_name": {
			// Property: StatementName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the prepared statement.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The name of the prepared statement.",
			Type:        types.StringType,
			Computed:    true,
		},
		"work_group": {
			// Property: WorkGroup
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the workgroup to which the prepared statement belongs.",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The name of the workgroup to which the prepared statement belongs.",
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
		Description: "Data Source schema for AWS::Athena::PreparedStatement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::PreparedStatement").WithTerraformTypeName("awscc_athena_prepared_statement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":     "Description",
		"query_statement": "QueryStatement",
		"statement_name":  "StatementName",
		"work_group":      "WorkGroup",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
