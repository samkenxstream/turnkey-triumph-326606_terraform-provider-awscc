// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lightsail_container", containerResourceType)
}

// containerResourceType returns the Terraform awscc_lightsail_container resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lightsail::Container resource type.
func containerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"container_arn": {
			// Property: ContainerArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"container_service_deployment": {
			// Property: ContainerServiceDeployment
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Describes a container deployment configuration of an Amazon Lightsail container service.",
			//   "properties": {
			//     "Containers": {
			//       "description": "An object that describes the configuration for the containers of the deployment.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.",
			//         "properties": {
			//           "Command": {
			//             "description": "The launch command for the container.",
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "ContainerName": {
			//             "description": "The name of the container.",
			//             "type": "string"
			//           },
			//           "Environment": {
			//             "description": "The environment variables of the container.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Value": {
			//                   "type": "string"
			//                 },
			//                 "Variable": {
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "Image": {
			//             "description": "The name of the image used for the container.",
			//             "type": "string"
			//           },
			//           "Ports": {
			//             "description": "The open firewall ports of the container.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Port": {
			//                   "type": "string"
			//                 },
			//                 "Protocol": {
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "PublicEndpoint": {
			//       "additionalProperties": false,
			//       "description": "An object that describes the endpoint of the deployment.",
			//       "properties": {
			//         "ContainerName": {
			//           "description": "The name of the container for the endpoint.",
			//           "type": "string"
			//         },
			//         "ContainerPort": {
			//           "description": "The port of the container to which traffic is forwarded to.",
			//           "type": "integer"
			//         },
			//         "HealthCheckConfig": {
			//           "additionalProperties": false,
			//           "description": "An object that describes the health check configuration of the container.",
			//           "properties": {
			//             "HealthyThreshold": {
			//               "description": "The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.",
			//               "type": "integer"
			//             },
			//             "IntervalSeconds": {
			//               "description": "The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.",
			//               "type": "integer"
			//             },
			//             "Path": {
			//               "description": "The path on the container on which to perform the health check. The default value is /.",
			//               "type": "string"
			//             },
			//             "SuccessCodes": {
			//               "description": "The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).",
			//               "type": "string"
			//             },
			//             "TimeoutSeconds": {
			//               "description": "The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.",
			//               "type": "integer"
			//             },
			//             "UnhealthyThreshold": {
			//               "description": "The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.",
			//               "type": "integer"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Describes a container deployment configuration of an Amazon Lightsail container service.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"containers": {
						// Property: Containers
						Description: "An object that describes the configuration for the containers of the deployment.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"command": {
									// Property: Command
									Description: "The launch command for the container.",
									Type:        types.SetType{ElemType: types.StringType},
									Optional:    true,
								},
								"container_name": {
									// Property: ContainerName
									Description: "The name of the container.",
									Type:        types.StringType,
									Optional:    true,
								},
								"environment": {
									// Property: Environment
									Description: "The environment variables of the container.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"value": {
												// Property: Value
												Type:     types.StringType,
												Optional: true,
											},
											"variable": {
												// Property: Variable
												Type:     types.StringType,
												Optional: true,
											},
										},
										tfsdk.SetNestedAttributesOptions{},
									),
									Optional: true,
								},
								"image": {
									// Property: Image
									Description: "The name of the image used for the container.",
									Type:        types.StringType,
									Optional:    true,
								},
								"ports": {
									// Property: Ports
									Description: "The open firewall ports of the container.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"port": {
												// Property: Port
												Type:     types.StringType,
												Optional: true,
											},
											"protocol": {
												// Property: Protocol
												Type:     types.StringType,
												Optional: true,
											},
										},
										tfsdk.SetNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
					"public_endpoint": {
						// Property: PublicEndpoint
						Description: "An object that describes the endpoint of the deployment.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"container_name": {
									// Property: ContainerName
									Description: "The name of the container for the endpoint.",
									Type:        types.StringType,
									Optional:    true,
								},
								"container_port": {
									// Property: ContainerPort
									Description: "The port of the container to which traffic is forwarded to.",
									Type:        types.Int64Type,
									Optional:    true,
								},
								"health_check_config": {
									// Property: HealthCheckConfig
									Description: "An object that describes the health check configuration of the container.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"healthy_threshold": {
												// Property: HealthyThreshold
												Description: "The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.",
												Type:        types.Int64Type,
												Optional:    true,
											},
											"interval_seconds": {
												// Property: IntervalSeconds
												Description: "The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.",
												Type:        types.Int64Type,
												Optional:    true,
											},
											"path": {
												// Property: Path
												Description: "The path on the container on which to perform the health check. The default value is /.",
												Type:        types.StringType,
												Optional:    true,
											},
											"success_codes": {
												// Property: SuccessCodes
												Description: "The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).",
												Type:        types.StringType,
												Optional:    true,
											},
											"timeout_seconds": {
												// Property: TimeoutSeconds
												Description: "The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.",
												Type:        types.Int64Type,
												Optional:    true,
											},
											"unhealthy_threshold": {
												// Property: UnhealthyThreshold
												Description: "The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.",
												Type:        types.Int64Type,
												Optional:    true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"is_disabled": {
			// Property: IsDisabled
			// CloudFormation resource type schema:
			// {
			//   "description": "A Boolean value to indicate whether the container service is disabled.",
			//   "type": "boolean"
			// }
			Description: "A Boolean value to indicate whether the container service is disabled.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"power": {
			// Property: Power
			// CloudFormation resource type schema:
			// {
			//   "description": "The power specification for the container service.",
			//   "type": "string"
			// }
			Description: "The power specification for the container service.",
			Type:        types.StringType,
			Required:    true,
		},
		"public_domain_names": {
			// Property: PublicDomainNames
			// CloudFormation resource type schema:
			// {
			//   "description": "The public domain names to use with the container service, such as example.com and www.example.com.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The public domain name to use with the container service, such as example.com and www.example.com.",
			//     "properties": {
			//       "CertificateName": {
			//         "type": "string"
			//       },
			//       "DomainNames": {
			//         "description": "An object that describes the configuration for the containers of the deployment.",
			//         "insertionOrder": false,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The public domain names to use with the container service, such as example.com and www.example.com.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"certificate_name": {
						// Property: CertificateName
						Type:     types.StringType,
						Optional: true,
					},
					"domain_names": {
						// Property: DomainNames
						Description: "An object that describes the configuration for the containers of the deployment.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"scale": {
			// Property: Scale
			// CloudFormation resource type schema:
			// {
			//   "description": "The scale specification for the container service.",
			//   "maximum": 20,
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Description: "The scale specification for the container service.",
			Type:        types.Int64Type,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(1, 20),
			},
		},
		"service_name": {
			// Property: ServiceName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the container service.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the container service.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			// {
			//   "description": "The publicly accessible URL of the container service.",
			//   "type": "string"
			// }
			Description: "The publicly accessible URL of the container service.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lightsail::Container",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Container").WithTerraformTypeName("awscc_lightsail_container")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_name":             "CertificateName",
		"command":                      "Command",
		"container_arn":                "ContainerArn",
		"container_name":               "ContainerName",
		"container_port":               "ContainerPort",
		"container_service_deployment": "ContainerServiceDeployment",
		"containers":                   "Containers",
		"domain_names":                 "DomainNames",
		"environment":                  "Environment",
		"health_check_config":          "HealthCheckConfig",
		"healthy_threshold":            "HealthyThreshold",
		"image":                        "Image",
		"interval_seconds":             "IntervalSeconds",
		"is_disabled":                  "IsDisabled",
		"key":                          "Key",
		"path":                         "Path",
		"port":                         "Port",
		"ports":                        "Ports",
		"power":                        "Power",
		"protocol":                     "Protocol",
		"public_domain_names":          "PublicDomainNames",
		"public_endpoint":              "PublicEndpoint",
		"scale":                        "Scale",
		"service_name":                 "ServiceName",
		"success_codes":                "SuccessCodes",
		"tags":                         "Tags",
		"timeout_seconds":              "TimeoutSeconds",
		"unhealthy_threshold":          "UnhealthyThreshold",
		"url":                          "Url",
		"value":                        "Value",
		"variable":                     "Variable",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}