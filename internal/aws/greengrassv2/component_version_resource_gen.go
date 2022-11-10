// Code generated by generators/resource/main.go; DO NOT EDIT.

package greengrassv2

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
	registry.AddResourceFactory("awscc_greengrassv2_component_version", componentVersionResource)
}

// componentVersionResource returns the Terraform awscc_greengrassv2_component_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::GreengrassV2::ComponentVersion resource.
func componentVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"component_name": {
			// Property: ComponentName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"component_version": {
			// Property: ComponentVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"inline_recipe": {
			// Property: InlineRecipe
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
				resource.RequiresReplace(),
			},
			// InlineRecipe is a write-only property.
		},
		"lambda_function": {
			// Property: LambdaFunction
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "ComponentDependencies": {
			//	      "additionalProperties": false,
			//	      "patternProperties": {
			//	        "": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "DependencyType": {
			//	              "enum": [
			//	                "SOFT",
			//	                "HARD"
			//	              ],
			//	              "type": "string"
			//	            },
			//	            "VersionRequirement": {
			//	              "type": "string"
			//	            }
			//	          },
			//	          "type": "object"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "ComponentLambdaParameters": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "EnvironmentVariables": {
			//	          "additionalProperties": false,
			//	          "patternProperties": {
			//	            "": {
			//	              "type": "string"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "EventSources": {
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "properties": {
			//	              "Topic": {
			//	                "type": "string"
			//	              },
			//	              "Type": {
			//	                "enum": [
			//	                  "PUB_SUB",
			//	                  "IOT_CORE"
			//	                ],
			//	                "type": "string"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "ExecArgs": {
			//	          "insertionOrder": true,
			//	          "items": {
			//	            "type": "string"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "InputPayloadEncodingType": {
			//	          "enum": [
			//	            "json",
			//	            "binary"
			//	          ],
			//	          "type": "string"
			//	        },
			//	        "LinuxProcessParams": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "ContainerParams": {
			//	              "additionalProperties": false,
			//	              "properties": {
			//	                "Devices": {
			//	                  "insertionOrder": false,
			//	                  "items": {
			//	                    "additionalProperties": false,
			//	                    "properties": {
			//	                      "AddGroupOwner": {
			//	                        "type": "boolean"
			//	                      },
			//	                      "Path": {
			//	                        "type": "string"
			//	                      },
			//	                      "Permission": {
			//	                        "enum": [
			//	                          "ro",
			//	                          "rw"
			//	                        ],
			//	                        "type": "string"
			//	                      }
			//	                    },
			//	                    "type": "object"
			//	                  },
			//	                  "type": "array"
			//	                },
			//	                "MemorySizeInKB": {
			//	                  "type": "integer"
			//	                },
			//	                "MountROSysfs": {
			//	                  "type": "boolean"
			//	                },
			//	                "Volumes": {
			//	                  "insertionOrder": false,
			//	                  "items": {
			//	                    "additionalProperties": false,
			//	                    "properties": {
			//	                      "AddGroupOwner": {
			//	                        "type": "boolean"
			//	                      },
			//	                      "DestinationPath": {
			//	                        "type": "string"
			//	                      },
			//	                      "Permission": {
			//	                        "enum": [
			//	                          "ro",
			//	                          "rw"
			//	                        ],
			//	                        "type": "string"
			//	                      },
			//	                      "SourcePath": {
			//	                        "type": "string"
			//	                      }
			//	                    },
			//	                    "type": "object"
			//	                  },
			//	                  "type": "array"
			//	                }
			//	              },
			//	              "type": "object"
			//	            },
			//	            "IsolationMode": {
			//	              "enum": [
			//	                "GreengrassContainer",
			//	                "NoContainer"
			//	              ],
			//	              "type": "string"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "MaxIdleTimeInSeconds": {
			//	          "type": "integer"
			//	        },
			//	        "MaxInstancesCount": {
			//	          "type": "integer"
			//	        },
			//	        "MaxQueueSize": {
			//	          "type": "integer"
			//	        },
			//	        "Pinned": {
			//	          "type": "boolean"
			//	        },
			//	        "StatusTimeoutInSeconds": {
			//	          "type": "integer"
			//	        },
			//	        "TimeoutInSeconds": {
			//	          "type": "integer"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "ComponentName": {
			//	      "type": "string"
			//	    },
			//	    "ComponentPlatforms": {
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "Attributes": {
			//	            "additionalProperties": false,
			//	            "patternProperties": {
			//	              "": {
			//	                "type": "string"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "Name": {
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "type": "array"
			//	    },
			//	    "ComponentVersion": {
			//	      "type": "string"
			//	    },
			//	    "LambdaArn": {
			//	      "pattern": "^arn:aws(-(cn|us-gov))?:lambda:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"component_dependencies": {
						// Property: ComponentDependencies
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"dependency_type": {
									// Property: DependencyType
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"SOFT",
											"HARD",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"version_requirement": {
									// Property: VersionRequirement
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"component_lambda_parameters": {
						// Property: ComponentLambdaParameters
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"environment_variables": {
									// Property: EnvironmentVariables
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"event_sources": {
									// Property: EventSources
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"topic": {
												// Property: Topic
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"type": {
												// Property: Type
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"PUB_SUB",
														"IOT_CORE",
													}),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										Multiset(),
										resource.UseStateForUnknown(),
									},
								},
								"exec_args": {
									// Property: ExecArgs
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"input_payload_encoding_type": {
									// Property: InputPayloadEncodingType
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"json",
											"binary",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"linux_process_params": {
									// Property: LinuxProcessParams
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"container_params": {
												// Property: ContainerParams
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"devices": {
															// Property: Devices
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"add_group_owner": {
																		// Property: AddGroupOwner
																		Type:     types.BoolType,
																		Optional: true,
																		Computed: true,
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																	"path": {
																		// Property: Path
																		Type:     types.StringType,
																		Optional: true,
																		Computed: true,
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																	"permission": {
																		// Property: Permission
																		Type:     types.StringType,
																		Optional: true,
																		Computed: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringInSlice([]string{
																				"ro",
																				"rw",
																			}),
																		},
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																},
															),
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																Multiset(),
																resource.UseStateForUnknown(),
															},
														},
														"memory_size_in_kb": {
															// Property: MemorySizeInKB
															Type:     types.Int64Type,
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																resource.UseStateForUnknown(),
															},
														},
														"mount_ro_sysfs": {
															// Property: MountROSysfs
															Type:     types.BoolType,
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																resource.UseStateForUnknown(),
															},
														},
														"volumes": {
															// Property: Volumes
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"add_group_owner": {
																		// Property: AddGroupOwner
																		Type:     types.BoolType,
																		Optional: true,
																		Computed: true,
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																	"destination_path": {
																		// Property: DestinationPath
																		Type:     types.StringType,
																		Optional: true,
																		Computed: true,
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																	"permission": {
																		// Property: Permission
																		Type:     types.StringType,
																		Optional: true,
																		Computed: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringInSlice([]string{
																				"ro",
																				"rw",
																			}),
																		},
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																	"source_path": {
																		// Property: SourcePath
																		Type:     types.StringType,
																		Optional: true,
																		Computed: true,
																		PlanModifiers: []tfsdk.AttributePlanModifier{
																			resource.UseStateForUnknown(),
																		},
																	},
																},
															),
															Optional: true,
															Computed: true,
															PlanModifiers: []tfsdk.AttributePlanModifier{
																Multiset(),
																resource.UseStateForUnknown(),
															},
														},
													},
												),
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"isolation_mode": {
												// Property: IsolationMode
												Type:     types.StringType,
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringInSlice([]string{
														"GreengrassContainer",
														"NoContainer",
													}),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"max_idle_time_in_seconds": {
									// Property: MaxIdleTimeInSeconds
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"max_instances_count": {
									// Property: MaxInstancesCount
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"max_queue_size": {
									// Property: MaxQueueSize
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"pinned": {
									// Property: Pinned
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"status_timeout_in_seconds": {
									// Property: StatusTimeoutInSeconds
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"timeout_in_seconds": {
									// Property: TimeoutInSeconds
									Type:     types.Int64Type,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"component_name": {
						// Property: ComponentName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"component_platforms": {
						// Property: ComponentPlatforms
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"attributes": {
									// Property: Attributes
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
						},
					},
					"component_version": {
						// Property: ComponentVersion
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"lambda_arn": {
						// Property: LambdaArn
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:lambda:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// LambdaFunction is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "patternProperties": {
			//	    "": {
			//	      "maxLength": 256,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			//
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
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
		Description: "Resource for Greengrass component version.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GreengrassV2::ComponentVersion").WithTerraformTypeName("awscc_greengrassv2_component_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_group_owner":             "AddGroupOwner",
		"arn":                         "Arn",
		"attributes":                  "Attributes",
		"component_dependencies":      "ComponentDependencies",
		"component_lambda_parameters": "ComponentLambdaParameters",
		"component_name":              "ComponentName",
		"component_platforms":         "ComponentPlatforms",
		"component_version":           "ComponentVersion",
		"container_params":            "ContainerParams",
		"dependency_type":             "DependencyType",
		"destination_path":            "DestinationPath",
		"devices":                     "Devices",
		"environment_variables":       "EnvironmentVariables",
		"event_sources":               "EventSources",
		"exec_args":                   "ExecArgs",
		"inline_recipe":               "InlineRecipe",
		"input_payload_encoding_type": "InputPayloadEncodingType",
		"isolation_mode":              "IsolationMode",
		"lambda_arn":                  "LambdaArn",
		"lambda_function":             "LambdaFunction",
		"linux_process_params":        "LinuxProcessParams",
		"max_idle_time_in_seconds":    "MaxIdleTimeInSeconds",
		"max_instances_count":         "MaxInstancesCount",
		"max_queue_size":              "MaxQueueSize",
		"memory_size_in_kb":           "MemorySizeInKB",
		"mount_ro_sysfs":              "MountROSysfs",
		"name":                        "Name",
		"path":                        "Path",
		"permission":                  "Permission",
		"pinned":                      "Pinned",
		"source_path":                 "SourcePath",
		"status_timeout_in_seconds":   "StatusTimeoutInSeconds",
		"tags":                        "Tags",
		"timeout_in_seconds":          "TimeoutInSeconds",
		"topic":                       "Topic",
		"type":                        "Type",
		"version_requirement":         "VersionRequirement",
		"volumes":                     "Volumes",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/LambdaFunction",
		"/properties/InlineRecipe",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
