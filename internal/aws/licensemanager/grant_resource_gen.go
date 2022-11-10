// Code generated by generators/resource/main.go; DO NOT EDIT.

package licensemanager

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
	registry.AddResourceFactory("awscc_licensemanager_grant", grantResource)
}

// grantResource returns the Terraform awscc_licensemanager_grant resource.
// This Terraform resource corresponds to the CloudFormation AWS::LicenseManager::Grant resource.
func grantResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"allowed_operations": {
			// Property: AllowedOperations
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// AllowedOperations is a write-only property.
		},
		"grant_arn": {
			// Property: GrantArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Arn of the grant.",
			//	  "maxLength": 2048,
			//	  "type": "string"
			//	}
			Description: "Arn of the grant.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"grant_name": {
			// Property: GrantName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Name for the created Grant.",
			//	  "type": "string"
			//	}
			Description: "Name for the created Grant.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Home region for the created grant.",
			//	  "type": "string"
			//	}
			Description: "Home region for the created grant.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"license_arn": {
			// Property: LicenseArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "License Arn for the grant.",
			//	  "maxLength": 2048,
			//	  "type": "string"
			//	}
			Description: "License Arn for the grant.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"principals": {
			// Property: Principals
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "maxLength": 2048,
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
				validate.ArrayForEach(validate.StringLenAtMost(2048)),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// Principals is a write-only property.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The version of the grant.",
			//	  "type": "string"
			//	}
			Description: "The version of the grant.",
			Type:        types.StringType,
			Computed:    true,
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LicenseManager::Grant").WithTerraformTypeName("awscc_licensemanager_grant")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_operations": "AllowedOperations",
		"grant_arn":          "GrantArn",
		"grant_name":         "GrantName",
		"home_region":        "HomeRegion",
		"license_arn":        "LicenseArn",
		"principals":         "Principals",
		"status":             "Status",
		"version":            "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Principals",
		"/properties/AllowedOperations",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
