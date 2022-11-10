// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_instance", instanceDataSource)
}

// instanceDataSource returns the Terraform awscc_lightsail_instance data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::Instance resource.
func instanceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"add_ons": {
			// Property: AddOns
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of objects representing the add-ons to enable for the new instance.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A addon associate with a resource.",
			//	    "properties": {
			//	      "AddOnType": {
			//	        "description": "The add-on type",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "AutoSnapshotAddOnRequest": {
			//	        "additionalProperties": false,
			//	        "description": "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
			//	        "properties": {
			//	          "SnapshotTimeOfDay": {
			//	            "description": "The daily time when an automatic snapshot will be created.",
			//	            "pattern": "^[0-9]{2}:00$",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "Status": {
			//	        "description": "Status of the Addon",
			//	        "enum": [
			//	          "Enabling",
			//	          "Disabling",
			//	          "Enabled",
			//	          "Terminating",
			//	          "Terminated",
			//	          "Disabled",
			//	          "Failed"
			//	        ],
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "AddOnType"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "An array of objects representing the add-ons to enable for the new instance.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"add_on_type": {
						// Property: AddOnType
						Description: "The add-on type",
						Type:        types.StringType,
						Computed:    true,
					},
					"auto_snapshot_add_on_request": {
						// Property: AutoSnapshotAddOnRequest
						Description: "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"snapshot_time_of_day": {
									// Property: SnapshotTimeOfDay
									Description: "The daily time when an automatic snapshot will be created.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"status": {
						// Property: Status
						Description: "Status of the Addon",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"blueprint_id": {
			// Property: BlueprintId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).",
			Type:        types.StringType,
			Computed:    true,
		},
		"bundle_id": {
			// Property: BundleId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).",
			Type:        types.StringType,
			Computed:    true,
		},
		"hardware": {
			// Property: Hardware
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Hardware of the Instance.",
			//	  "properties": {
			//	    "CpuCount": {
			//	      "description": "CPU count of the Instance.",
			//	      "type": "integer"
			//	    },
			//	    "Disks": {
			//	      "description": "Disks attached to the Instance.",
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "Disk associated with the Instance.",
			//	        "properties": {
			//	          "AttachedTo": {
			//	            "description": "Instance attached to the disk.",
			//	            "type": "string"
			//	          },
			//	          "AttachmentState": {
			//	            "description": "Attachment state of the disk.",
			//	            "type": "string"
			//	          },
			//	          "DiskName": {
			//	            "description": "The names to use for your new Lightsail disk.",
			//	            "maxLength": 254,
			//	            "minLength": 1,
			//	            "pattern": "^[a-zA-Z0-9][\\w\\-.]*[a-zA-Z0-9]$",
			//	            "type": "string"
			//	          },
			//	          "IOPS": {
			//	            "description": "IOPS of disk.",
			//	            "type": "integer"
			//	          },
			//	          "IsSystemDisk": {
			//	            "description": "Is the Attached disk is the system disk of the Instance.",
			//	            "type": "boolean"
			//	          },
			//	          "Path": {
			//	            "description": "Path of the disk attached to the instance.",
			//	            "type": "string"
			//	          },
			//	          "SizeInGb": {
			//	            "description": "Size of the disk attached to the Instance.",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "required": [
			//	          "DiskName",
			//	          "Path"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "type": "array",
			//	      "uniqueItems": true
			//	    },
			//	    "RamSizeInGb": {
			//	      "description": "RAM Size of the Instance.",
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Hardware of the Instance.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cpu_count": {
						// Property: CpuCount
						Description: "CPU count of the Instance.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"disks": {
						// Property: Disks
						Description: "Disks attached to the Instance.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"attached_to": {
									// Property: AttachedTo
									Description: "Instance attached to the disk.",
									Type:        types.StringType,
									Computed:    true,
								},
								"attachment_state": {
									// Property: AttachmentState
									Description: "Attachment state of the disk.",
									Type:        types.StringType,
									Computed:    true,
								},
								"disk_name": {
									// Property: DiskName
									Description: "The names to use for your new Lightsail disk.",
									Type:        types.StringType,
									Computed:    true,
								},
								"iops": {
									// Property: IOPS
									Description: "IOPS of disk.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"is_system_disk": {
									// Property: IsSystemDisk
									Description: "Is the Attached disk is the system disk of the Instance.",
									Type:        types.BoolType,
									Computed:    true,
								},
								"path": {
									// Property: Path
									Description: "Path of the disk attached to the instance.",
									Type:        types.StringType,
									Computed:    true,
								},
								"size_in_gb": {
									// Property: SizeInGb
									Description: "Size of the disk attached to the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"ram_size_in_gb": {
						// Property: RamSizeInGb
						Description: "RAM Size of the Instance.",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"instance_name": {
			// Property: InstanceName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The names to use for your new Lightsail instance.",
			//	  "maxLength": 254,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9][\\w\\-.]*[a-zA-Z0-9]$",
			//	  "type": "string"
			//	}
			Description: "The names to use for your new Lightsail instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"is_static_ip": {
			// Property: IsStaticIp
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Is the IP Address of the Instance is the static IP",
			//	  "type": "boolean"
			//	}
			Description: "Is the IP Address of the Instance is the static IP",
			Type:        types.BoolType,
			Computed:    true,
		},
		"key_pair_name": {
			// Property: KeyPairName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of your key pair.",
			//	  "type": "string"
			//	}
			Description: "The name of your key pair.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Location of a resource.",
			//	  "properties": {
			//	    "AvailabilityZone": {
			//	      "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			//	      "type": "string"
			//	    },
			//	    "RegionName": {
			//	      "description": "The Region Name in which to create your instance.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Location of a resource.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"availability_zone": {
						// Property: AvailabilityZone
						Description: "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
						Type:        types.StringType,
						Computed:    true,
					},
					"region_name": {
						// Property: RegionName
						Description: "The Region Name in which to create your instance.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"networking": {
			// Property: Networking
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Networking of the Instance.",
			//	  "properties": {
			//	    "MonthlyTransfer": {
			//	      "additionalProperties": false,
			//	      "description": "Monthly Transfer of the Instance.",
			//	      "properties": {
			//	        "GbPerMonthAllocated": {
			//	          "description": "GbPerMonthAllocated of the Instance.",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "Ports": {
			//	      "description": "Ports to the Instance.",
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "Port of the Instance.",
			//	        "properties": {
			//	          "AccessDirection": {
			//	            "description": "Access Direction for Protocol of the Instance(inbound/outbound).",
			//	            "type": "string"
			//	          },
			//	          "AccessFrom": {
			//	            "description": "Access From Protocol of the Instance.",
			//	            "type": "string"
			//	          },
			//	          "AccessType": {
			//	            "description": "Access Type Protocol of the Instance.",
			//	            "type": "string"
			//	          },
			//	          "CidrListAliases": {
			//	            "description": "cidr List Aliases",
			//	            "insertionOrder": false,
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "type": "array"
			//	          },
			//	          "Cidrs": {
			//	            "description": "cidrs",
			//	            "insertionOrder": false,
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "type": "array"
			//	          },
			//	          "CommonName": {
			//	            "description": "CommonName for Protocol of the Instance.",
			//	            "type": "string"
			//	          },
			//	          "FromPort": {
			//	            "description": "From Port of the Instance.",
			//	            "type": "integer"
			//	          },
			//	          "Ipv6Cidrs": {
			//	            "description": "IPv6 Cidrs",
			//	            "insertionOrder": false,
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "type": "array"
			//	          },
			//	          "Protocol": {
			//	            "description": "Port Protocol of the Instance.",
			//	            "type": "string"
			//	          },
			//	          "ToPort": {
			//	            "description": "To Port of the Instance.",
			//	            "type": "integer"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "type": "array",
			//	      "uniqueItems": true
			//	    }
			//	  },
			//	  "required": [
			//	    "Ports"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Networking of the Instance.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"monthly_transfer": {
						// Property: MonthlyTransfer
						Description: "Monthly Transfer of the Instance.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"gb_per_month_allocated": {
									// Property: GbPerMonthAllocated
									Description: "GbPerMonthAllocated of the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"ports": {
						// Property: Ports
						Description: "Ports to the Instance.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"access_direction": {
									// Property: AccessDirection
									Description: "Access Direction for Protocol of the Instance(inbound/outbound).",
									Type:        types.StringType,
									Computed:    true,
								},
								"access_from": {
									// Property: AccessFrom
									Description: "Access From Protocol of the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
								"access_type": {
									// Property: AccessType
									Description: "Access Type Protocol of the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
								"cidr_list_aliases": {
									// Property: CidrListAliases
									Description: "cidr List Aliases",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
								"cidrs": {
									// Property: Cidrs
									Description: "cidrs",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
								"common_name": {
									// Property: CommonName
									Description: "CommonName for Protocol of the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
								"from_port": {
									// Property: FromPort
									Description: "From Port of the Instance.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"ipv_6_cidrs": {
									// Property: Ipv6Cidrs
									Description: "IPv6 Cidrs",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
								"protocol": {
									// Property: Protocol
									Description: "Port Protocol of the Instance.",
									Type:        types.StringType,
									Computed:    true,
								},
								"to_port": {
									// Property: ToPort
									Description: "To Port of the Instance.",
									Type:        types.Int64Type,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"private_ip_address": {
			// Property: PrivateIpAddress
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Private IP Address of the Instance",
			//	  "type": "string"
			//	}
			Description: "Private IP Address of the Instance",
			Type:        types.StringType,
			Computed:    true,
		},
		"public_ip_address": {
			// Property: PublicIpAddress
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Public IP Address of the Instance",
			//	  "type": "string"
			//	}
			Description: "Public IP Address of the Instance",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Resource type of Lightsail instance.",
			//	  "type": "string"
			//	}
			Description: "Resource type of Lightsail instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ssh_key_name": {
			// Property: SshKeyName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "SSH Key Name of the  Lightsail instance.",
			//	  "type": "string"
			//	}
			Description: "SSH Key Name of the  Lightsail instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Current State of the Instance.",
			//	  "properties": {
			//	    "Code": {
			//	      "description": "Status code of the Instance.",
			//	      "type": "integer"
			//	    },
			//	    "Name": {
			//	      "description": "Status code of the Instance.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Current State of the Instance.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"code": {
						// Property: Code
						Description: "Status code of the Instance.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "Status code of the Instance.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"support_code": {
			// Property: SupportCode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Support code to help identify any issues",
			//	  "type": "string"
			//	}
			Description: "Support code to help identify any issues",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
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
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"user_data": {
			// Property: UserData
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update.",
			//	  "type": "string"
			//	}
			Description: "A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update.",
			Type:        types.StringType,
			Computed:    true,
		},
		"user_name": {
			// Property: UserName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Username of the  Lightsail instance.",
			//	  "type": "string"
			//	}
			Description: "Username of the  Lightsail instance.",
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
		Description: "Data Source schema for AWS::Lightsail::Instance",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Instance").WithTerraformTypeName("awscc_lightsail_instance")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_direction":             "AccessDirection",
		"access_from":                  "AccessFrom",
		"access_type":                  "AccessType",
		"add_on_type":                  "AddOnType",
		"add_ons":                      "AddOns",
		"attached_to":                  "AttachedTo",
		"attachment_state":             "AttachmentState",
		"auto_snapshot_add_on_request": "AutoSnapshotAddOnRequest",
		"availability_zone":            "AvailabilityZone",
		"blueprint_id":                 "BlueprintId",
		"bundle_id":                    "BundleId",
		"cidr_list_aliases":            "CidrListAliases",
		"cidrs":                        "Cidrs",
		"code":                         "Code",
		"common_name":                  "CommonName",
		"cpu_count":                    "CpuCount",
		"disk_name":                    "DiskName",
		"disks":                        "Disks",
		"from_port":                    "FromPort",
		"gb_per_month_allocated":       "GbPerMonthAllocated",
		"hardware":                     "Hardware",
		"instance_arn":                 "InstanceArn",
		"instance_name":                "InstanceName",
		"iops":                         "IOPS",
		"ipv_6_cidrs":                  "Ipv6Cidrs",
		"is_static_ip":                 "IsStaticIp",
		"is_system_disk":               "IsSystemDisk",
		"key":                          "Key",
		"key_pair_name":                "KeyPairName",
		"location":                     "Location",
		"monthly_transfer":             "MonthlyTransfer",
		"name":                         "Name",
		"networking":                   "Networking",
		"path":                         "Path",
		"ports":                        "Ports",
		"private_ip_address":           "PrivateIpAddress",
		"protocol":                     "Protocol",
		"public_ip_address":            "PublicIpAddress",
		"ram_size_in_gb":               "RamSizeInGb",
		"region_name":                  "RegionName",
		"resource_type":                "ResourceType",
		"size_in_gb":                   "SizeInGb",
		"snapshot_time_of_day":         "SnapshotTimeOfDay",
		"ssh_key_name":                 "SshKeyName",
		"state":                        "State",
		"status":                       "Status",
		"support_code":                 "SupportCode",
		"tags":                         "Tags",
		"to_port":                      "ToPort",
		"user_data":                    "UserData",
		"user_name":                    "UserName",
		"value":                        "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
