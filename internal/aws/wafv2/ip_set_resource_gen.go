// Code generated by generators/resource/main.go; DO NOT EDIT.

package wafv2

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
	registry.AddResourceFactory("awscc_wafv2_ip_set", iPSetResource)
}

// iPSetResource returns the Terraform awscc_wafv2_ip_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::WAFv2::IPSet resource.
func iPSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"addresses": {
			// Property: Addresses
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of IPAddresses.",
			//	  "items": {
			//	    "description": "IP address",
			//	    "maxLength": 50,
			//	    "minLength": 1,
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "List of IPAddresses.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringLenBetween(1, 50)),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "ARN of the WAF entity.",
			//	  "type": "string"
			//	}
			Description: "ARN of the WAF entity.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Description of the entity.",
			//	  "pattern": "^[a-zA-Z0-9=:#@/\\-,.][a-zA-Z0-9+=:#@/\\-,.\\s]+[a-zA-Z0-9+=:#@/\\-,.]{1,256}$",
			//	  "type": "string"
			//	}
			Description: "Description of the entity.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9=:#@/\\-,.][a-zA-Z0-9+=:#@/\\-,.\\s]+[a-zA-Z0-9+=:#@/\\-,.]{1,256}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ip_address_version": {
			// Property: IPAddressVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			//	  "enum": [
			//	    "IPV4",
			//	    "IPV6"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"IPV4",
					"IPV6",
				}),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Id of the IPSet",
			//	  "pattern": "^[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}$",
			//	  "type": "string"
			//	}
			Description: "Id of the IPSet",
			Type:        types.StringType,
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
			//	  "description": "Name of the IPSet.",
			//	  "pattern": "^[0-9A-Za-z_-]{1,128}$",
			//	  "type": "string"
			//	}
			Description: "Name of the IPSet.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[0-9A-Za-z_-]{1,128}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			//	  "enum": [
			//	    "CLOUDFRONT",
			//	    "REGIONAL"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CLOUDFRONT",
					"REGIONAL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::IPSet").WithTerraformTypeName("awscc_wafv2_ip_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addresses":          "Addresses",
		"arn":                "Arn",
		"description":        "Description",
		"id":                 "Id",
		"ip_address_version": "IPAddressVersion",
		"key":                "Key",
		"name":               "Name",
		"scope":              "Scope",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
