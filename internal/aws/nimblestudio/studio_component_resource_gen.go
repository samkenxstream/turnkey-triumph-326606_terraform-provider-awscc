// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

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
	registry.AddResourceFactory("awscc_nimblestudio_studio_component", studioComponentResource)
}

// studioComponentResource returns the Terraform awscc_nimblestudio_studio_component resource.
// This Terraform resource corresponds to the CloudFormation AWS::NimbleStudio::StudioComponent resource.
func studioComponentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe configuration of the studio component, based on component type.\u003c/p\u003e",
			//	  "properties": {
			//	    "ActiveDirectoryConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "\u003cp\u003eThe configuration for a Microsoft Active Directory (Microsoft AD) studio\n            resource.\u003c/p\u003e",
			//	      "properties": {
			//	        "ComputerAttributes": {
			//	          "description": "\u003cp\u003eA collection of custom attributes for an Active Directory computer.\u003c/p\u003e",
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "\u003cp\u003eAn LDAP attribute of an Active Directory computer account, in the form of a name:value\n            pair.\u003c/p\u003e",
			//	            "properties": {
			//	              "Name": {
			//	                "description": "\u003cp\u003eThe name for the LDAP attribute.\u003c/p\u003e",
			//	                "maxLength": 40,
			//	                "minLength": 1,
			//	                "type": "string"
			//	              },
			//	              "Value": {
			//	                "description": "\u003cp\u003eThe value for the LDAP attribute.\u003c/p\u003e",
			//	                "maxLength": 64,
			//	                "minLength": 1,
			//	                "type": "string"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "maxItems": 50,
			//	          "minItems": 0,
			//	          "type": "array"
			//	        },
			//	        "DirectoryId": {
			//	          "description": "\u003cp\u003eThe directory ID of the Directory Service for Microsoft Active Directory to access\n            using this studio component.\u003c/p\u003e",
			//	          "type": "string"
			//	        },
			//	        "OrganizationalUnitDistinguishedName": {
			//	          "description": "\u003cp\u003eThe distinguished name (DN) and organizational unit (OU) of an Active Directory\n            computer.\u003c/p\u003e",
			//	          "maxLength": 2000,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "ComputeFarmConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "\u003cp\u003eThe configuration for a render farm that is associated with a studio resource.\u003c/p\u003e",
			//	      "properties": {
			//	        "ActiveDirectoryUser": {
			//	          "description": "\u003cp\u003eThe name of an Active Directory user that is used on ComputeFarm worker\n            instances.\u003c/p\u003e",
			//	          "type": "string"
			//	        },
			//	        "Endpoint": {
			//	          "description": "\u003cp\u003eThe endpoint of the ComputeFarm that is accessed by the studio component\n            resource.\u003c/p\u003e",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "LicenseServiceConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "\u003cp\u003eThe configuration for a license service that is associated with a studio\n            resource.\u003c/p\u003e",
			//	      "properties": {
			//	        "Endpoint": {
			//	          "description": "\u003cp\u003eThe endpoint of the license service that is accessed by the studio component\n            resource.\u003c/p\u003e",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "SharedFileSystemConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "\u003cp\u003eThe configuration for a shared file storage system that is associated with a studio\n            resource.\u003c/p\u003e",
			//	      "properties": {
			//	        "Endpoint": {
			//	          "description": "\u003cp\u003eThe endpoint of the shared file system that is accessed by the studio component\n            resource.\u003c/p\u003e",
			//	          "type": "string"
			//	        },
			//	        "FileSystemId": {
			//	          "description": "\u003cp\u003eThe unique identifier for a file system.\u003c/p\u003e",
			//	          "type": "string"
			//	        },
			//	        "LinuxMountPoint": {
			//	          "description": "\u003cp\u003eThe mount location for a shared file system on a Linux virtual workstation.\u003c/p\u003e",
			//	          "maxLength": 128,
			//	          "minLength": 0,
			//	          "pattern": "^(/?|(\\$HOME)?(/[^/\\n\\s\\\\]+)*)$",
			//	          "type": "string"
			//	        },
			//	        "ShareName": {
			//	          "description": "\u003cp\u003eThe name of the file share.\u003c/p\u003e",
			//	          "type": "string"
			//	        },
			//	        "WindowsMountDrive": {
			//	          "description": "\u003cp\u003eThe mount location for a shared file system on a Windows virtual workstation.\u003c/p\u003e",
			//	          "pattern": "^[A-Z]$",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "<p>The configuration of the studio component, based on component type.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"active_directory_configuration": {
						// Property: ActiveDirectoryConfiguration
						Description: "<p>The configuration for a Microsoft Active Directory (Microsoft AD) studio\n            resource.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"computer_attributes": {
									// Property: ComputerAttributes
									Description: "<p>A collection of custom attributes for an Active Directory computer.</p>",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Description: "<p>The name for the LDAP attribute.</p>",
												Type:        types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 40),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"value": {
												// Property: Value
												Description: "<p>The value for the LDAP attribute.</p>",
												Type:        types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 64),
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
										validate.ArrayLenBetween(0, 50),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"directory_id": {
									// Property: DirectoryId
									Description: "<p>The directory ID of the Directory Service for Microsoft Active Directory to access\n            using this studio component.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"organizational_unit_distinguished_name": {
									// Property: OrganizationalUnitDistinguishedName
									Description: "<p>The distinguished name (DN) and organizational unit (OU) of an Active Directory\n            computer.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2000),
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
					"compute_farm_configuration": {
						// Property: ComputeFarmConfiguration
						Description: "<p>The configuration for a render farm that is associated with a studio resource.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"active_directory_user": {
									// Property: ActiveDirectoryUser
									Description: "<p>The name of an Active Directory user that is used on ComputeFarm worker\n            instances.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"endpoint": {
									// Property: Endpoint
									Description: "<p>The endpoint of the ComputeFarm that is accessed by the studio component\n            resource.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
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
					"license_service_configuration": {
						// Property: LicenseServiceConfiguration
						Description: "<p>The configuration for a license service that is associated with a studio\n            resource.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"endpoint": {
									// Property: Endpoint
									Description: "<p>The endpoint of the license service that is accessed by the studio component\n            resource.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
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
					"shared_file_system_configuration": {
						// Property: SharedFileSystemConfiguration
						Description: "<p>The configuration for a shared file storage system that is associated with a studio\n            resource.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"endpoint": {
									// Property: Endpoint
									Description: "<p>The endpoint of the shared file system that is accessed by the studio component\n            resource.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"file_system_id": {
									// Property: FileSystemId
									Description: "<p>The unique identifier for a file system.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"linux_mount_point": {
									// Property: LinuxMountPoint
									Description: "<p>The mount location for a shared file system on a Linux virtual workstation.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 128),
										validate.StringMatch(regexp.MustCompile("^(/?|(\\$HOME)?(/[^/\\n\\s\\\\]+)*)$"), ""),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"share_name": {
									// Property: ShareName
									Description: "<p>The name of the file share.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"windows_mount_drive": {
									// Property: WindowsMountDrive
									Description: "<p>The mount location for a shared file system on a Windows virtual workstation.</p>",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^[A-Z]$"), ""),
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
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe description.\u003c/p\u003e",
			//	  "maxLength": 256,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "<p>The description.</p>",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ec_2_security_group_ids": {
			// Property: Ec2SecurityGroupIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe EC2 security groups that control access to the studio component.\u003c/p\u003e",
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "maxItems": 30,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Description: "<p>The EC2 security groups that control access to the studio component.</p>",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 30),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"initialization_scripts": {
			// Property: InitializationScripts
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eInitialization scripts for studio components.\u003c/p\u003e",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "\u003cp\u003eInitialization scripts for studio components.\u003c/p\u003e",
			//	    "properties": {
			//	      "LaunchProfileProtocolVersion": {
			//	        "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
			//	        "maxLength": 10,
			//	        "minLength": 0,
			//	        "pattern": "^2021\\-03\\-31$",
			//	        "type": "string"
			//	      },
			//	      "Platform": {
			//	        "enum": [
			//	          "LINUX",
			//	          "WINDOWS"
			//	        ],
			//	        "type": "string"
			//	      },
			//	      "RunContext": {
			//	        "enum": [
			//	          "SYSTEM_INITIALIZATION",
			//	          "USER_INITIALIZATION"
			//	        ],
			//	        "type": "string"
			//	      },
			//	      "Script": {
			//	        "description": "\u003cp\u003eThe initialization script.\u003c/p\u003e",
			//	        "maxLength": 5120,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "<p>Initialization scripts for studio components.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"launch_profile_protocol_version": {
						// Property: LaunchProfileProtocolVersion
						Description: "<p>The version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".</p>",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 10),
							validate.StringMatch(regexp.MustCompile("^2021\\-03\\-31$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"platform": {
						// Property: Platform
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"LINUX",
								"WINDOWS",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"run_context": {
						// Property: RunContext
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SYSTEM_INITIALIZATION",
								"USER_INITIALIZATION",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"script": {
						// Property: Script
						Description: "<p>The initialization script.</p>",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 5120),
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
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe name for the studio component.\u003c/p\u003e",
			//	  "maxLength": 64,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "<p>The name for the studio component.</p>",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 64),
			},
		},
		"runtime_role_arn": {
			// Property: RuntimeRoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"script_parameters": {
			// Property: ScriptParameters
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eParameters for the studio component scripts.\u003c/p\u003e",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "\u003cp\u003eA parameter for a studio component script, in the form of a key:value pair.\u003c/p\u003e",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "\u003cp\u003eA script parameter key.\u003c/p\u003e",
			//	        "maxLength": 64,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z_][a-zA-Z0-9_]+$",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "\u003cp\u003eA script parameter value.\u003c/p\u003e",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "maxItems": 30,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Description: "<p>Parameters for the studio component scripts.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "<p>A script parameter key.</p>",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 64),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]+$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Description: "<p>A script parameter value.</p>",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
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
				validate.ArrayLenBetween(0, 30),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"secure_initialization_role_arn": {
			// Property: SecureInitializationRoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"studio_component_id": {
			// Property: StudioComponentId
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
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe studio ID. \u003c/p\u003e",
			//	  "type": "string"
			//	}
			Description: "<p>The studio ID. </p>",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"subtype": {
			// Property: Subtype
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "AWS_MANAGED_MICROSOFT_AD",
			//	    "AMAZON_FSX_FOR_WINDOWS",
			//	    "AMAZON_FSX_FOR_LUSTRE",
			//	    "CUSTOM"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AWS_MANAGED_MICROSOFT_AD",
					"AMAZON_FSX_FOR_WINDOWS",
					"AMAZON_FSX_FOR_LUSTRE",
					"CUSTOM",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "patternProperties": {
			//	    "": {
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
				resource.RequiresReplace(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "ACTIVE_DIRECTORY",
			//	    "SHARED_FILE_SYSTEM",
			//	    "COMPUTE_FARM",
			//	    "LICENSE_SERVICE",
			//	    "CUSTOM"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE_DIRECTORY",
					"SHARED_FILE_SYSTEM",
					"COMPUTE_FARM",
					"LICENSE_SERVICE",
					"CUSTOM",
				}),
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
		Description: "Represents a studio component that connects a non-Nimble Studio resource in your account to your studio",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StudioComponent").WithTerraformTypeName("awscc_nimblestudio_studio_component")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_directory_configuration":         "ActiveDirectoryConfiguration",
		"active_directory_user":                  "ActiveDirectoryUser",
		"compute_farm_configuration":             "ComputeFarmConfiguration",
		"computer_attributes":                    "ComputerAttributes",
		"configuration":                          "Configuration",
		"description":                            "Description",
		"directory_id":                           "DirectoryId",
		"ec_2_security_group_ids":                "Ec2SecurityGroupIds",
		"endpoint":                               "Endpoint",
		"file_system_id":                         "FileSystemId",
		"initialization_scripts":                 "InitializationScripts",
		"key":                                    "Key",
		"launch_profile_protocol_version":        "LaunchProfileProtocolVersion",
		"license_service_configuration":          "LicenseServiceConfiguration",
		"linux_mount_point":                      "LinuxMountPoint",
		"name":                                   "Name",
		"organizational_unit_distinguished_name": "OrganizationalUnitDistinguishedName",
		"platform":                               "Platform",
		"run_context":                            "RunContext",
		"runtime_role_arn":                       "RuntimeRoleArn",
		"script":                                 "Script",
		"script_parameters":                      "ScriptParameters",
		"secure_initialization_role_arn":         "SecureInitializationRoleArn",
		"share_name":                             "ShareName",
		"shared_file_system_configuration":       "SharedFileSystemConfiguration",
		"studio_component_id":                    "StudioComponentId",
		"studio_id":                              "StudioId",
		"subtype":                                "Subtype",
		"tags":                                   "Tags",
		"type":                                   "Type",
		"value":                                  "Value",
		"windows_mount_drive":                    "WindowsMountDrive",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
