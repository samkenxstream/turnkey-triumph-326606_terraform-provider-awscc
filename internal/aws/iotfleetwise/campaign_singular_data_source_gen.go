// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotfleetwise_campaign", campaignDataSource)
}

// campaignDataSource returns the Terraform awscc_iotfleetwise_campaign data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTFleetWise::Campaign resource.
func campaignDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "APPROVE",
			//	    "SUSPEND",
			//	    "RESUME",
			//	    "UPDATE"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"collection_scheme": {
			// Property: CollectionScheme
			// CloudFormation resource type schema:
			//
			//	{
			//	  "properties": {
			//	    "ConditionBasedCollectionScheme": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "ConditionLanguageVersion": {
			//	          "minimum": 1,
			//	          "type": "integer"
			//	        },
			//	        "Expression": {
			//	          "maxLength": 2048,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "MinimumTriggerIntervalMs": {
			//	          "maximum": 4294967295,
			//	          "minimum": 0,
			//	          "type": "number"
			//	        },
			//	        "TriggerMode": {
			//	          "enum": [
			//	            "ALWAYS",
			//	            "RISING_EDGE"
			//	          ],
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "Expression"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "TimeBasedCollectionScheme": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "PeriodMs": {
			//	          "maximum": 60000,
			//	          "minimum": 10000,
			//	          "type": "number"
			//	        }
			//	      },
			//	      "required": [
			//	        "PeriodMs"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"condition_based_collection_scheme": {
						// Property: ConditionBasedCollectionScheme
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"condition_language_version": {
									// Property: ConditionLanguageVersion
									Type:     types.Int64Type,
									Computed: true,
								},
								"expression": {
									// Property: Expression
									Type:     types.StringType,
									Computed: true,
								},
								"minimum_trigger_interval_ms": {
									// Property: MinimumTriggerIntervalMs
									Type:     types.Float64Type,
									Computed: true,
								},
								"trigger_mode": {
									// Property: TriggerMode
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"time_based_collection_scheme": {
						// Property: TimeBasedCollectionScheme
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"period_ms": {
									// Property: PeriodMs
									Type:     types.Float64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"compression": {
			// Property: Compression
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "OFF",
			//	  "enum": [
			//	    "OFF",
			//	    "SNAPPY"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"data_extra_dimensions": {
			// Property: DataExtraDimensions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "maxLength": 150,
			//	    "minLength": 1,
			//	    "pattern": "^[a-zA-Z0-9_.]+$",
			//	    "type": "string"
			//	  },
			//	  "maxItems": 5,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"diagnostics_mode": {
			// Property: DiagnosticsMode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "OFF",
			//	  "enum": [
			//	    "OFF",
			//	    "SEND_ACTIVE_DTCS"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"expiry_time": {
			// Property: ExpiryTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "253402243200",
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"last_modification_time": {
			// Property: LastModificationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 100,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"post_trigger_collection_duration": {
			// Property: PostTriggerCollectionDuration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": 0,
			//	  "maximum": 4294967295,
			//	  "minimum": 0,
			//	  "type": "number"
			//	}
			Type:     types.Float64Type,
			Computed: true,
		},
		"priority": {
			// Property: Priority
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": 0,
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Computed: true,
		},
		"signal_catalog_arn": {
			// Property: SignalCatalogArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"signals_to_collect": {
			// Property: SignalsToCollect
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "MaxSampleCount": {
			//	        "maximum": 4294967295,
			//	        "minimum": 1,
			//	        "type": "number"
			//	      },
			//	      "MinimumSamplingIntervalMs": {
			//	        "maximum": 4294967295,
			//	        "minimum": 0,
			//	        "type": "number"
			//	      },
			//	      "Name": {
			//	        "maxLength": 150,
			//	        "minLength": 1,
			//	        "pattern": "^[\\w|*|-]+(\\.[\\w|*|-]+)*$",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Name"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 1000,
			//	  "minItems": 0,
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"max_sample_count": {
						// Property: MaxSampleCount
						Type:     types.Float64Type,
						Computed: true,
					},
					"minimum_sampling_interval_ms": {
						// Property: MinimumSamplingIntervalMs
						Type:     types.Float64Type,
						Computed: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"spooling_mode": {
			// Property: SpoolingMode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "OFF",
			//	  "enum": [
			//	    "OFF",
			//	    "TO_DISK"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"start_time": {
			// Property: StartTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "0",
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "CREATING",
			//	    "WAITING_FOR_APPROVAL",
			//	    "RUNNING",
			//	    "SUSPENDED"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
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
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "minItems": 0,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"target_arn": {
			// Property: TargetArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTFleetWise::Campaign",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::Campaign").WithTerraformTypeName("awscc_iotfleetwise_campaign")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                            "Action",
		"arn":                               "Arn",
		"collection_scheme":                 "CollectionScheme",
		"compression":                       "Compression",
		"condition_based_collection_scheme": "ConditionBasedCollectionScheme",
		"condition_language_version":        "ConditionLanguageVersion",
		"creation_time":                     "CreationTime",
		"data_extra_dimensions":             "DataExtraDimensions",
		"description":                       "Description",
		"diagnostics_mode":                  "DiagnosticsMode",
		"expiry_time":                       "ExpiryTime",
		"expression":                        "Expression",
		"key":                               "Key",
		"last_modification_time":            "LastModificationTime",
		"max_sample_count":                  "MaxSampleCount",
		"minimum_sampling_interval_ms":      "MinimumSamplingIntervalMs",
		"minimum_trigger_interval_ms":       "MinimumTriggerIntervalMs",
		"name":                              "Name",
		"period_ms":                         "PeriodMs",
		"post_trigger_collection_duration":  "PostTriggerCollectionDuration",
		"priority":                          "Priority",
		"signal_catalog_arn":                "SignalCatalogArn",
		"signals_to_collect":                "SignalsToCollect",
		"spooling_mode":                     "SpoolingMode",
		"start_time":                        "StartTime",
		"status":                            "Status",
		"tags":                              "Tags",
		"target_arn":                        "TargetArn",
		"time_based_collection_scheme":      "TimeBasedCollectionScheme",
		"trigger_mode":                      "TriggerMode",
		"value":                             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
