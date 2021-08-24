// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_xray_sampling_rule", samplingRuleResourceType)
}

// samplingRuleResourceType returns the Terraform awscc_xray_sampling_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::XRay::SamplingRule resource type.
func samplingRuleResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"rule_arn": {
			// Property: RuleARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//   "type": "string"
			// }
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_name": {
			// Property: RuleName
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//   "maxLength": 32,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 32),
			},
		},
		"sampling_rule": {
			// Property: SamplingRule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Attributes": {
			//       "additionalProperties": false,
			//       "$comment": "String to string map",
			//       "description": "Matches attributes derived from the request.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "FixedRate": {
			//       "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			//       "maximum": 1,
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "HTTPMethod": {
			//       "description": "Matches the HTTP method from a request URL.",
			//       "maxLength": 10,
			//       "type": "string"
			//     },
			//     "Host": {
			//       "description": "Matches the hostname from a request URL.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "Priority": {
			//       "description": "The priority of the sampling rule.",
			//       "maximum": 9999,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "ReservoirSize": {
			//       "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "ResourceARN": {
			//       "description": "Matches the ARN of the AWS resource on which the service runs.",
			//       "maxLength": 500,
			//       "type": "string"
			//     },
			//     "RuleARN": {
			//       "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//       "type": "string"
			//     },
			//     "RuleName": {
			//       "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//       "maxLength": 32,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "ServiceName": {
			//       "description": "Matches the name that the service uses to identify itself in segments.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "ServiceType": {
			//       "description": "Matches the origin that the service uses to identify its type in segments.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "URLPath": {
			//       "description": "Matches the path from a request URL.",
			//       "maxLength": 128,
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "The version of the sampling rule format (1)",
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Description: "Matches attributes derived from the request.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"fixed_rate": {
						// Property: FixedRate
						Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"http_method": {
						// Property: HTTPMethod
						Description: "Matches the HTTP method from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 10),
						},
					},
					"host": {
						// Property: Host
						Description: "Matches the hostname from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"priority": {
						// Property: Priority
						Description: "The priority of the sampling rule.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 9999),
						},
					},
					"reservoir_size": {
						// Property: ReservoirSize
						Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
					"resource_arn": {
						// Property: ResourceARN
						Description: "Matches the ARN of the AWS resource on which the service runs.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 500),
						},
					},
					"rule_arn": {
						// Property: RuleARN
						Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_name": {
						// Property: RuleName
						Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 32),
						},
					},
					"service_name": {
						// Property: ServiceName
						Description: "Matches the name that the service uses to identify itself in segments.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"service_type": {
						// Property: ServiceType
						Description: "Matches the origin that the service uses to identify its type in segments.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"url_path": {
						// Property: URLPath
						Description: "Matches the path from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 128),
						},
					},
					"version": {
						// Property: Version
						Description: "The version of the sampling rule format (1)",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(1),
						},
					},
				},
			),
			Optional: true,
		},
		"sampling_rule_record": {
			// Property: SamplingRuleRecord
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CreatedAt": {
			//       "description": "When the rule was created, in Unix time seconds.",
			//       "type": "string"
			//     },
			//     "ModifiedAt": {
			//       "description": "When the rule was modified, in Unix time seconds.",
			//       "type": "string"
			//     },
			//     "SamplingRule": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Attributes": {
			//           "additionalProperties": false,
			//           "$comment": "String to string map",
			//           "description": "Matches attributes derived from the request.",
			//           "patternProperties": {
			//             "": {
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "FixedRate": {
			//           "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			//           "maximum": 1,
			//           "minimum": 0,
			//           "type": "number"
			//         },
			//         "HTTPMethod": {
			//           "description": "Matches the HTTP method from a request URL.",
			//           "maxLength": 10,
			//           "type": "string"
			//         },
			//         "Host": {
			//           "description": "Matches the hostname from a request URL.",
			//           "maxLength": 64,
			//           "type": "string"
			//         },
			//         "Priority": {
			//           "description": "The priority of the sampling rule.",
			//           "maximum": 9999,
			//           "minimum": 1,
			//           "type": "integer"
			//         },
			//         "ReservoirSize": {
			//           "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			//           "minimum": 0,
			//           "type": "integer"
			//         },
			//         "ResourceARN": {
			//           "description": "Matches the ARN of the AWS resource on which the service runs.",
			//           "maxLength": 500,
			//           "type": "string"
			//         },
			//         "RuleARN": {
			//           "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//           "type": "string"
			//         },
			//         "RuleName": {
			//           "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "ServiceName": {
			//           "description": "Matches the name that the service uses to identify itself in segments.",
			//           "maxLength": 64,
			//           "type": "string"
			//         },
			//         "ServiceType": {
			//           "description": "Matches the origin that the service uses to identify its type in segments.",
			//           "maxLength": 64,
			//           "type": "string"
			//         },
			//         "URLPath": {
			//           "description": "Matches the path from a request URL.",
			//           "maxLength": 128,
			//           "type": "string"
			//         },
			//         "Version": {
			//           "description": "The version of the sampling rule format (1)",
			//           "minimum": 1,
			//           "type": "integer"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"created_at": {
						// Property: CreatedAt
						Description: "When the rule was created, in Unix time seconds.",
						Type:        types.StringType,
						Optional:    true,
					},
					"modified_at": {
						// Property: ModifiedAt
						Description: "When the rule was modified, in Unix time seconds.",
						Type:        types.StringType,
						Optional:    true,
					},
					"sampling_rule": {
						// Property: SamplingRule
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"attributes": {
									// Property: Attributes
									Description: "Matches attributes derived from the request.",
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"fixed_rate": {
									// Property: FixedRate
									Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
									Type:        types.NumberType,
									Optional:    true,
								},
								"http_method": {
									// Property: HTTPMethod
									Description: "Matches the HTTP method from a request URL.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 10),
									},
								},
								"host": {
									// Property: Host
									Description: "Matches the hostname from a request URL.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 64),
									},
								},
								"priority": {
									// Property: Priority
									Description: "The priority of the sampling rule.",
									Type:        types.NumberType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 9999),
									},
								},
								"reservoir_size": {
									// Property: ReservoirSize
									Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
									Type:        types.NumberType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntAtLeast(0),
									},
								},
								"resource_arn": {
									// Property: ResourceARN
									Description: "Matches the ARN of the AWS resource on which the service runs.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 500),
									},
								},
								"rule_arn": {
									// Property: RuleARN
									Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
									Type:        types.StringType,
									Optional:    true,
								},
								"rule_name": {
									// Property: RuleName
									Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 32),
									},
								},
								"service_name": {
									// Property: ServiceName
									Description: "Matches the name that the service uses to identify itself in segments.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 64),
									},
								},
								"service_type": {
									// Property: ServiceType
									Description: "Matches the origin that the service uses to identify its type in segments.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 64),
									},
								},
								"url_path": {
									// Property: URLPath
									Description: "Matches the path from a request URL.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 128),
									},
								},
								"version": {
									// Property: Version
									Description: "The version of the sampling rule format (1)",
									Type:        types.NumberType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntAtLeast(1),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"sampling_rule_update": {
			// Property: SamplingRuleUpdate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Attributes": {
			//       "additionalProperties": false,
			//       "$comment": "String to string map",
			//       "description": "Matches attributes derived from the request.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "FixedRate": {
			//       "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			//       "maximum": 1,
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "HTTPMethod": {
			//       "description": "Matches the HTTP method from a request URL.",
			//       "maxLength": 10,
			//       "type": "string"
			//     },
			//     "Host": {
			//       "description": "Matches the hostname from a request URL.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "Priority": {
			//       "description": "The priority of the sampling rule.",
			//       "maximum": 9999,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "ReservoirSize": {
			//       "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "ResourceARN": {
			//       "description": "Matches the ARN of the AWS resource on which the service runs.",
			//       "maxLength": 500,
			//       "type": "string"
			//     },
			//     "RuleARN": {
			//       "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//       "type": "string"
			//     },
			//     "RuleName": {
			//       "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			//       "maxLength": 32,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "ServiceName": {
			//       "description": "Matches the name that the service uses to identify itself in segments.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "ServiceType": {
			//       "description": "Matches the origin that the service uses to identify its type in segments.",
			//       "maxLength": 64,
			//       "type": "string"
			//     },
			//     "URLPath": {
			//       "description": "Matches the path from a request URL.",
			//       "maxLength": 128,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Description: "Matches attributes derived from the request.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"fixed_rate": {
						// Property: FixedRate
						Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"http_method": {
						// Property: HTTPMethod
						Description: "Matches the HTTP method from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 10),
						},
					},
					"host": {
						// Property: Host
						Description: "Matches the hostname from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"priority": {
						// Property: Priority
						Description: "The priority of the sampling rule.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 9999),
						},
					},
					"reservoir_size": {
						// Property: ReservoirSize
						Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
					"resource_arn": {
						// Property: ResourceARN
						Description: "Matches the ARN of the AWS resource on which the service runs.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 500),
						},
					},
					"rule_arn": {
						// Property: RuleARN
						Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_name": {
						// Property: RuleName
						Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 32),
						},
					},
					"service_name": {
						// Property: ServiceName
						Description: "Matches the name that the service uses to identify itself in segments.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"service_type": {
						// Property: ServiceType
						Description: "Matches the origin that the service uses to identify its type in segments.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 64),
						},
					},
					"url_path": {
						// Property: URLPath
						Description: "Matches the path from a request URL.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 128),
						},
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::SamplingRule").WithTerraformTypeName("awscc_xray_sampling_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":           "Attributes",
		"created_at":           "CreatedAt",
		"fixed_rate":           "FixedRate",
		"host":                 "Host",
		"http_method":          "HTTPMethod",
		"key":                  "Key",
		"modified_at":          "ModifiedAt",
		"priority":             "Priority",
		"reservoir_size":       "ReservoirSize",
		"resource_arn":         "ResourceARN",
		"rule_arn":             "RuleARN",
		"rule_name":            "RuleName",
		"sampling_rule":        "SamplingRule",
		"sampling_rule_record": "SamplingRuleRecord",
		"sampling_rule_update": "SamplingRuleUpdate",
		"service_name":         "ServiceName",
		"service_type":         "ServiceType",
		"tags":                 "Tags",
		"url_path":             "URLPath",
		"value":                "Value",
		"version":              "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_xray_sampling_rule", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
