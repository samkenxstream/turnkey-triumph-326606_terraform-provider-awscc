// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_networkfirewall_rule_group", ruleGroupResourceType)
}

// ruleGroupResourceType returns the Terraform awscc_networkfirewall_rule_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkFirewall::RuleGroup resource type.
func ruleGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity": {
			// Property: Capacity
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Required: true,
			// Capacity is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"rule_group": {
			// Property: RuleGroup
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "RuleVariables": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "IPSets": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Definition": {
			//                   "insertionOrder": false,
			//                   "items": {
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   },
			//                   "type": "array",
			//                   "uniqueItems": true
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "PortSets": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Definition": {
			//                   "insertionOrder": false,
			//                   "items": {
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   },
			//                   "type": "array",
			//                   "uniqueItems": true
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "RulesSource": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "RulesSourceList": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "GeneratedRulesType": {
			//               "enum": [
			//                 "ALLOWLIST",
			//                 "DENYLIST"
			//               ],
			//               "type": "string"
			//             },
			//             "TargetTypes": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "enum": [
			//                   "TLS_SNI",
			//                   "HTTP_HOST"
			//                 ],
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "Targets": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "Targets",
			//             "TargetTypes",
			//             "GeneratedRulesType"
			//           ],
			//           "type": "object"
			//         },
			//         "RulesString": {
			//           "maxLength": 1000000,
			//           "minLength": 0,
			//           "type": "string"
			//         },
			//         "StatefulRules": {
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Action": {
			//                 "enum": [
			//                   "PASS",
			//                   "DROP",
			//                   "ALERT"
			//                 ],
			//                 "type": "string"
			//               },
			//               "Header": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Destination": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   },
			//                   "DestinationPort": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   },
			//                   "Direction": {
			//                     "enum": [
			//                       "FORWARD",
			//                       "ANY"
			//                     ],
			//                     "type": "string"
			//                   },
			//                   "Protocol": {
			//                     "enum": [
			//                       "IP",
			//                       "TCP",
			//                       "UDP",
			//                       "ICMP",
			//                       "HTTP",
			//                       "FTP",
			//                       "TLS",
			//                       "SMB",
			//                       "DNS",
			//                       "DCERPC",
			//                       "SSH",
			//                       "SMTP",
			//                       "IMAP",
			//                       "MSN",
			//                       "KRB5",
			//                       "IKEV2",
			//                       "TFTP",
			//                       "NTP",
			//                       "DHCP"
			//                     ],
			//                     "type": "string"
			//                   },
			//                   "Source": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   },
			//                   "SourcePort": {
			//                     "maxLength": 1024,
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "Protocol",
			//                   "Source",
			//                   "SourcePort",
			//                   "Direction",
			//                   "Destination",
			//                   "DestinationPort"
			//                 ],
			//                 "type": "object"
			//               },
			//               "RuleOptions": {
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "Keyword": {
			//                       "maxLength": 128,
			//                       "minLength": 1,
			//                       "pattern": "",
			//                       "type": "string"
			//                     },
			//                     "Settings": {
			//                       "insertionOrder": false,
			//                       "items": {
			//                         "maxLength": 8192,
			//                         "minLength": 1,
			//                         "pattern": "",
			//                         "type": "string"
			//                       },
			//                       "type": "array",
			//                       "uniqueItems": true
			//                     }
			//                   },
			//                   "required": [
			//                     "Keyword"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": true
			//               }
			//             },
			//             "required": [
			//               "Action",
			//               "Header",
			//               "RuleOptions"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "StatelessRulesAndCustomActions": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "CustomActions": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "ActionDefinition": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "PublishMetricAction": {
			//                         "additionalProperties": false,
			//                         "properties": {
			//                           "Dimensions": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "Value": {
			//                                   "maxLength": 128,
			//                                   "minLength": 1,
			//                                   "pattern": "",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "Value"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           }
			//                         },
			//                         "required": [
			//                           "Dimensions"
			//                         ],
			//                         "type": "object"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "ActionName": {
			//                     "maxLength": 128,
			//                     "minLength": 1,
			//                     "pattern": "",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "ActionName",
			//                   "ActionDefinition"
			//                 ],
			//                 "type": "object"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "StatelessRules": {
			//               "insertionOrder": false,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "Priority": {
			//                     "type": "integer"
			//                   },
			//                   "RuleDefinition": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Actions": {
			//                         "insertionOrder": false,
			//                         "items": {
			//                           "type": "string"
			//                         },
			//                         "type": "array",
			//                         "uniqueItems": true
			//                       },
			//                       "MatchAttributes": {
			//                         "additionalProperties": false,
			//                         "properties": {
			//                           "DestinationPorts": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "FromPort": {
			//                                   "type": "integer"
			//                                 },
			//                                 "ToPort": {
			//                                   "type": "integer"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "FromPort",
			//                                 "ToPort"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           },
			//                           "Destinations": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "AddressDefinition": {
			//                                   "maxLength": 255,
			//                                   "minLength": 1,
			//                                   "pattern": "",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "AddressDefinition"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           },
			//                           "Protocols": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "type": "integer"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           },
			//                           "SourcePorts": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "FromPort": {
			//                                   "type": "integer"
			//                                 },
			//                                 "ToPort": {
			//                                   "type": "integer"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "FromPort",
			//                                 "ToPort"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           },
			//                           "Sources": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "AddressDefinition": {
			//                                   "maxLength": 255,
			//                                   "minLength": 1,
			//                                   "pattern": "",
			//                                   "type": "string"
			//                                 }
			//                               },
			//                               "required": [
			//                                 "AddressDefinition"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           },
			//                           "TCPFlags": {
			//                             "insertionOrder": false,
			//                             "items": {
			//                               "additionalProperties": false,
			//                               "properties": {
			//                                 "Flags": {
			//                                   "insertionOrder": false,
			//                                   "items": {
			//                                     "enum": [
			//                                       "FIN",
			//                                       "SYN",
			//                                       "RST",
			//                                       "PSH",
			//                                       "ACK",
			//                                       "URG",
			//                                       "ECE",
			//                                       "CWR"
			//                                     ],
			//                                     "type": "string"
			//                                   },
			//                                   "type": "array",
			//                                   "uniqueItems": true
			//                                 },
			//                                 "Masks": {
			//                                   "insertionOrder": false,
			//                                   "items": {
			//                                     "enum": [
			//                                       "FIN",
			//                                       "SYN",
			//                                       "RST",
			//                                       "PSH",
			//                                       "ACK",
			//                                       "URG",
			//                                       "ECE",
			//                                       "CWR"
			//                                     ],
			//                                     "type": "string"
			//                                   },
			//                                   "type": "array",
			//                                   "uniqueItems": true
			//                                 }
			//                               },
			//                               "required": [
			//                                 "Flags"
			//                               ],
			//                               "type": "object"
			//                             },
			//                             "type": "array",
			//                             "uniqueItems": true
			//                           }
			//                         },
			//                         "type": "object"
			//                       }
			//                     },
			//                     "required": [
			//                       "MatchAttributes",
			//                       "Actions"
			//                     ],
			//                     "type": "object"
			//                   }
			//                 },
			//                 "required": [
			//                   "RuleDefinition",
			//                   "Priority"
			//                 ],
			//                 "type": "object"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "StatelessRules"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "RulesSource"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"rule_variables": {
						// Property: RuleVariables
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"ip_sets": {
									// Property: IPSets
									// Pattern: ""
									Attributes: tfsdk.MapNestedAttributes(
										map[string]tfsdk.Attribute{
											"definition": {
												// Property: Definition
												Type:     providertypes.SetType{ElemType: types.StringType},
												Optional: true,
											},
										},
										tfsdk.MapNestedAttributesOptions{},
									),
									Optional: true,
								},
								"port_sets": {
									// Property: PortSets
									// Pattern: ""
									Attributes: tfsdk.MapNestedAttributes(
										map[string]tfsdk.Attribute{
											"definition": {
												// Property: Definition
												Type:     providertypes.SetType{ElemType: types.StringType},
												Optional: true,
											},
										},
										tfsdk.MapNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"rules_source": {
						// Property: RulesSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"rules_source_list": {
									// Property: RulesSourceList
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"generated_rules_type": {
												// Property: GeneratedRulesType
												Type:     types.StringType,
												Required: true,
											},
											"target_types": {
												// Property: TargetTypes
												Type:     providertypes.SetType{ElemType: types.StringType},
												Required: true,
											},
											"targets": {
												// Property: Targets
												Type:     providertypes.SetType{ElemType: types.StringType},
												Required: true,
											},
										},
									),
									Optional: true,
								},
								"rules_string": {
									// Property: RulesString
									Type:     types.StringType,
									Optional: true,
								},
								"stateful_rules": {
									// Property: StatefulRules
									Attributes: providertypes.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"action": {
												// Property: Action
												Type:     types.StringType,
												Required: true,
											},
											"header": {
												// Property: Header
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"destination": {
															// Property: Destination
															Type:     types.StringType,
															Required: true,
														},
														"destination_port": {
															// Property: DestinationPort
															Type:     types.StringType,
															Required: true,
														},
														"direction": {
															// Property: Direction
															Type:     types.StringType,
															Required: true,
														},
														"protocol": {
															// Property: Protocol
															Type:     types.StringType,
															Required: true,
														},
														"source": {
															// Property: Source
															Type:     types.StringType,
															Required: true,
														},
														"source_port": {
															// Property: SourcePort
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Required: true,
											},
											"rule_options": {
												// Property: RuleOptions
												Attributes: providertypes.SetNestedAttributes(
													map[string]tfsdk.Attribute{
														"keyword": {
															// Property: Keyword
															Type:     types.StringType,
															Required: true,
														},
														"settings": {
															// Property: Settings
															Type:     providertypes.SetType{ElemType: types.StringType},
															Optional: true,
														},
													},
													providertypes.SetNestedAttributesOptions{},
												),
												Required: true,
											},
										},
										providertypes.SetNestedAttributesOptions{},
									),
									Optional: true,
								},
								"stateless_rules_and_custom_actions": {
									// Property: StatelessRulesAndCustomActions
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"custom_actions": {
												// Property: CustomActions
												Attributes: providertypes.SetNestedAttributes(
													map[string]tfsdk.Attribute{
														"action_definition": {
															// Property: ActionDefinition
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"publish_metric_action": {
																		// Property: PublishMetricAction
																		Attributes: tfsdk.SingleNestedAttributes(
																			map[string]tfsdk.Attribute{
																				"dimensions": {
																					// Property: Dimensions
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"value": {
																								// Property: Value
																								Type:     types.StringType,
																								Required: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Required: true,
																				},
																			},
																		),
																		Optional: true,
																	},
																},
															),
															Required: true,
														},
														"action_name": {
															// Property: ActionName
															Type:     types.StringType,
															Required: true,
														},
													},
													providertypes.SetNestedAttributesOptions{},
												),
												Optional: true,
											},
											"stateless_rules": {
												// Property: StatelessRules
												Attributes: providertypes.SetNestedAttributes(
													map[string]tfsdk.Attribute{
														"priority": {
															// Property: Priority
															Type:     types.NumberType,
															Required: true,
														},
														"rule_definition": {
															// Property: RuleDefinition
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"actions": {
																		// Property: Actions
																		Type:     providertypes.SetType{ElemType: types.StringType},
																		Required: true,
																	},
																	"match_attributes": {
																		// Property: MatchAttributes
																		Attributes: tfsdk.SingleNestedAttributes(
																			map[string]tfsdk.Attribute{
																				"destination_ports": {
																					// Property: DestinationPorts
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"from_port": {
																								// Property: FromPort
																								Type:     types.NumberType,
																								Required: true,
																							},
																							"to_port": {
																								// Property: ToPort
																								Type:     types.NumberType,
																								Required: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Optional: true,
																				},
																				"destinations": {
																					// Property: Destinations
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"address_definition": {
																								// Property: AddressDefinition
																								Type:     types.StringType,
																								Required: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Optional: true,
																				},
																				"protocols": {
																					// Property: Protocols
																					Type:     providertypes.SetType{ElemType: types.NumberType},
																					Optional: true,
																				},
																				"source_ports": {
																					// Property: SourcePorts
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"from_port": {
																								// Property: FromPort
																								Type:     types.NumberType,
																								Required: true,
																							},
																							"to_port": {
																								// Property: ToPort
																								Type:     types.NumberType,
																								Required: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Optional: true,
																				},
																				"sources": {
																					// Property: Sources
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"address_definition": {
																								// Property: AddressDefinition
																								Type:     types.StringType,
																								Required: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Optional: true,
																				},
																				"tcp_flags": {
																					// Property: TCPFlags
																					Attributes: providertypes.SetNestedAttributes(
																						map[string]tfsdk.Attribute{
																							"flags": {
																								// Property: Flags
																								Type:     providertypes.SetType{ElemType: types.StringType},
																								Required: true,
																							},
																							"masks": {
																								// Property: Masks
																								Type:     providertypes.SetType{ElemType: types.StringType},
																								Optional: true,
																							},
																						},
																						providertypes.SetNestedAttributesOptions{},
																					),
																					Optional: true,
																				},
																			},
																		),
																		Required: true,
																	},
																},
															),
															Required: true,
														},
													},
													providertypes.SetNestedAttributesOptions{},
												),
												Required: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Computed: true,
			// RuleGroup is a force-new attribute.
		},
		"rule_group_arn": {
			// Property: RuleGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_group_id": {
			// Property: RuleGroupId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"rule_group_name": {
			// Property: RuleGroupName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// RuleGroupName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
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
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "STATELESS",
			//     "STATEFUL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// Type is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::RuleGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::RuleGroup").WithTerraformTypeName("awscc_networkfirewall_rule_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkfirewall_rule_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
