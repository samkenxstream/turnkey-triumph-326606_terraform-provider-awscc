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
	registry.AddResourceFactory("awscc_licensemanager_license", licenseResource)
}

// licenseResource returns the Terraform awscc_licensemanager_license resource.
// This Terraform resource corresponds to the CloudFormation AWS::LicenseManager::License resource.
func licenseResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"beneficiary": {
			// Property: Beneficiary
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Beneficiary of the license.",
			//	  "type": "string"
			//	}
			Description: "Beneficiary of the license.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"consumption_configuration": {
			// Property: ConsumptionConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "BorrowConfiguration": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "AllowEarlyCheckIn": {
			//	          "type": "boolean"
			//	        },
			//	        "MaxTimeToLiveInMinutes": {
			//	          "type": "integer"
			//	        }
			//	      },
			//	      "required": [
			//	        "MaxTimeToLiveInMinutes",
			//	        "AllowEarlyCheckIn"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "ProvisionalConfiguration": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "MaxTimeToLiveInMinutes": {
			//	          "type": "integer"
			//	        }
			//	      },
			//	      "required": [
			//	        "MaxTimeToLiveInMinutes"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "RenewType": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"borrow_configuration": {
						// Property: BorrowConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"allow_early_check_in": {
									// Property: AllowEarlyCheckIn
									Type:     types.BoolType,
									Required: true,
								},
								"max_time_to_live_in_minutes": {
									// Property: MaxTimeToLiveInMinutes
									Type:     types.Int64Type,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"provisional_configuration": {
						// Property: ProvisionalConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"max_time_to_live_in_minutes": {
									// Property: MaxTimeToLiveInMinutes
									Type:     types.Int64Type,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"renew_type": {
						// Property: RenewType
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"entitlements": {
			// Property: Entitlements
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "AllowCheckIn": {
			//	        "type": "boolean"
			//	      },
			//	      "MaxCount": {
			//	        "type": "integer"
			//	      },
			//	      "Name": {
			//	        "type": "string"
			//	      },
			//	      "Overage": {
			//	        "type": "boolean"
			//	      },
			//	      "Unit": {
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Name",
			//	      "Unit"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"allow_check_in": {
						// Property: AllowCheckIn
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_count": {
						// Property: MaxCount
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
					"overage": {
						// Property: Overage
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"unit": {
						// Property: Unit
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Home region for the created license.",
			//	  "type": "string"
			//	}
			Description: "Home region for the created license.",
			Type:        types.StringType,
			Required:    true,
		},
		"issuer": {
			// Property: Issuer
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Name": {
			//	      "type": "string"
			//	    },
			//	    "SignKey": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Name"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
					"sign_key": {
						// Property: SignKey
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"license_arn": {
			// Property: LicenseArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Amazon Resource Name is a unique name for each resource.",
			//	  "maxLength": 2048,
			//	  "type": "string"
			//	}
			Description: "Amazon Resource Name is a unique name for each resource.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"license_metadata": {
			// Property: LicenseMetadata
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Name": {
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Name",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"license_name": {
			// Property: LicenseName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Name for the created license.",
			//	  "type": "string"
			//	}
			Description: "Name for the created license.",
			Type:        types.StringType,
			Required:    true,
		},
		"product_name": {
			// Property: ProductName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Product name for the created license.",
			//	  "type": "string"
			//	}
			Description: "Product name for the created license.",
			Type:        types.StringType,
			Required:    true,
		},
		"product_sku": {
			// Property: ProductSKU
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "ProductSKU of the license.",
			//	  "maxLength": 1024,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "ProductSKU of the license.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1024),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			// Status is a write-only property.
		},
		"validity": {
			// Property: Validity
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Begin": {
			//	      "description": "Validity begin date for the license.",
			//	      "format": "date-time",
			//	      "type": "string"
			//	    },
			//	    "End": {
			//	      "description": "Validity begin date for the license.",
			//	      "format": "date-time",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Begin",
			//	    "End"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"begin": {
						// Property: Begin
						Description: "Validity begin date for the license.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IsRFC3339Time(),
						},
					},
					"end": {
						// Property: End
						Description: "Validity begin date for the license.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IsRFC3339Time(),
						},
					},
				},
			),
			Required: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The version of the license.",
			//	  "type": "string"
			//	}
			Description: "The version of the license.",
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
		Description: "Resource Type definition for AWS::LicenseManager::License",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LicenseManager::License").WithTerraformTypeName("awscc_licensemanager_license")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_check_in":              "AllowCheckIn",
		"allow_early_check_in":        "AllowEarlyCheckIn",
		"begin":                       "Begin",
		"beneficiary":                 "Beneficiary",
		"borrow_configuration":        "BorrowConfiguration",
		"consumption_configuration":   "ConsumptionConfiguration",
		"end":                         "End",
		"entitlements":                "Entitlements",
		"home_region":                 "HomeRegion",
		"issuer":                      "Issuer",
		"license_arn":                 "LicenseArn",
		"license_metadata":            "LicenseMetadata",
		"license_name":                "LicenseName",
		"max_count":                   "MaxCount",
		"max_time_to_live_in_minutes": "MaxTimeToLiveInMinutes",
		"name":                        "Name",
		"overage":                     "Overage",
		"product_name":                "ProductName",
		"product_sku":                 "ProductSKU",
		"provisional_configuration":   "ProvisionalConfiguration",
		"renew_type":                  "RenewType",
		"sign_key":                    "SignKey",
		"status":                      "Status",
		"unit":                        "Unit",
		"validity":                    "Validity",
		"value":                       "Value",
		"version":                     "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Status",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
