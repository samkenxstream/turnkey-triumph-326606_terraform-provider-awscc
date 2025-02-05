// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_database", databaseDataSource)
}

// databaseDataSource returns the Terraform awscc_lightsail_database data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::Database resource.
func databaseDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BackupRetention
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.",
		//	  "type": "boolean"
		//	}
		"backup_retention": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CaCertificateIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the certificate that needs to be associated with the database.",
		//	  "type": "string"
		//	}
		"ca_certificate_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the certificate that needs to be associated with the database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatabaseArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"database_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MasterDatabaseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"master_database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the database to create when the Lightsail database resource is created. For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MasterUserPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\". It cannot contain spaces.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"master_user_password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\". It cannot contain spaces.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MasterUsername
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the master user.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"master_username": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the master user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredBackupWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The daily time range during which automated backups are created for your new database if automated backups are enabled.",
		//	  "type": "string"
		//	}
		"preferred_backup_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The daily time range during which automated backups are created for your new database if automated backups are enabled.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredMaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The weekly time range during which system maintenance can occur on your new database.",
		//	  "type": "string"
		//	}
		"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The weekly time range during which system maintenance can occur on your new database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RelationalDatabaseBlueprintId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The blueprint ID for your new database. A blueprint describes the major engine version of a database.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"relational_database_blueprint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The blueprint ID for your new database. A blueprint describes the major engine version of a database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RelationalDatabaseBundleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The bundle ID for your new database. A bundle describes the performance specifications for your database.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"relational_database_bundle_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The bundle ID for your new database. A bundle describes the performance specifications for your database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RelationalDatabaseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to use for your new Lightsail database resource.",
		//	  "maxLength": 255,
		//	  "minLength": 2,
		//	  "pattern": "\\w[\\w\\-]*\\w",
		//	  "type": "string"
		//	}
		"relational_database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to use for your new Lightsail database resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RelationalDatabaseParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Update one or more parameters of the relational database.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes the parameters of the database.",
		//	    "properties": {
		//	      "AllowedValues": {
		//	        "description": "Specifies the valid range of values for the parameter.",
		//	        "type": "string"
		//	      },
		//	      "ApplyMethod": {
		//	        "description": "Indicates when parameter updates are applied. Can be immediate or pending-reboot.",
		//	        "type": "string"
		//	      },
		//	      "ApplyType": {
		//	        "description": "Specifies the engine-specific parameter type.",
		//	        "type": "string"
		//	      },
		//	      "DataType": {
		//	        "description": "Specifies the valid data type for the parameter.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "Provides a description of the parameter.",
		//	        "type": "string"
		//	      },
		//	      "IsModifiable": {
		//	        "description": "A Boolean value indicating whether the parameter can be modified.",
		//	        "type": "boolean"
		//	      },
		//	      "ParameterName": {
		//	        "description": "Specifies the name of the parameter.",
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "description": "Specifies the value of the parameter.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"relational_database_parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AllowedValues
					"allowed_values": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the valid range of values for the parameter.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ApplyMethod
					"apply_method": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Indicates when parameter updates are applied. Can be immediate or pending-reboot.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ApplyType
					"apply_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the engine-specific parameter type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DataType
					"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the valid data type for the parameter.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Provides a description of the parameter.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: IsModifiable
					"is_modifiable": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "A Boolean value indicating whether the parameter can be modified.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterName
					"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the name of the parameter.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies the value of the parameter.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Update one or more parameters of the relational database.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RotateMasterUserPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.",
		//	  "type": "boolean"
		//	}
		"rotate_master_user_password": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When true, the master user password is changed to a new strong password generated by Lightsail. Use the get relational database master user password operation to get the new password.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lightsail::Database",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Database").WithTerraformTypeName("awscc_lightsail_database")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_values":                   "AllowedValues",
		"apply_method":                     "ApplyMethod",
		"apply_type":                       "ApplyType",
		"availability_zone":                "AvailabilityZone",
		"backup_retention":                 "BackupRetention",
		"ca_certificate_identifier":        "CaCertificateIdentifier",
		"data_type":                        "DataType",
		"database_arn":                     "DatabaseArn",
		"description":                      "Description",
		"is_modifiable":                    "IsModifiable",
		"key":                              "Key",
		"master_database_name":             "MasterDatabaseName",
		"master_user_password":             "MasterUserPassword",
		"master_username":                  "MasterUsername",
		"parameter_name":                   "ParameterName",
		"parameter_value":                  "ParameterValue",
		"preferred_backup_window":          "PreferredBackupWindow",
		"preferred_maintenance_window":     "PreferredMaintenanceWindow",
		"publicly_accessible":              "PubliclyAccessible",
		"relational_database_blueprint_id": "RelationalDatabaseBlueprintId",
		"relational_database_bundle_id":    "RelationalDatabaseBundleId",
		"relational_database_name":         "RelationalDatabaseName",
		"relational_database_parameters":   "RelationalDatabaseParameters",
		"rotate_master_user_password":      "RotateMasterUserPassword",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
