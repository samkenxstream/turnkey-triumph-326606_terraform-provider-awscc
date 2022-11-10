// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_acmpca_certificate", certificateDataSource)
}

// certificateDataSource returns the Terraform awscc_acmpca_certificate data source.
// This Terraform data source corresponds to the CloudFormation AWS::ACMPCA::Certificate resource.
func certificateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_passthrough": {
			// Property: ApiPassthrough
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.",
			//	  "properties": {
			//	    "Extensions": {
			//	      "additionalProperties": false,
			//	      "description": "Structure that contains X.500 extensions for a Certificate.",
			//	      "properties": {
			//	        "CertificatePolicies": {
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "Structure that contains X.509 Policy information.",
			//	            "properties": {
			//	              "CertPolicyId": {
			//	                "description": "String that contains X.509 ObjectIdentifier information.",
			//	                "type": "string"
			//	              },
			//	              "PolicyQualifiers": {
			//	                "items": {
			//	                  "additionalProperties": false,
			//	                  "description": "Structure that contains X.509 Policy qualifier information.",
			//	                  "properties": {
			//	                    "PolicyQualifierId": {
			//	                      "type": "string"
			//	                    },
			//	                    "Qualifier": {
			//	                      "additionalProperties": false,
			//	                      "description": "Structure that contains a X.509 policy qualifier.",
			//	                      "properties": {
			//	                        "CpsUri": {
			//	                          "type": "string"
			//	                        }
			//	                      },
			//	                      "required": [
			//	                        "CpsUri"
			//	                      ],
			//	                      "type": "object"
			//	                    }
			//	                  },
			//	                  "required": [
			//	                    "PolicyQualifierId",
			//	                    "Qualifier"
			//	                  ],
			//	                  "type": "object"
			//	                },
			//	                "type": "array"
			//	              }
			//	            },
			//	            "required": [
			//	              "CertPolicyId"
			//	            ],
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "CustomExtensions": {
			//	          "description": "Array of X.509 extensions for a certificate.",
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "Structure that contains X.509 extension information for a certificate.",
			//	            "properties": {
			//	              "Critical": {
			//	                "type": "boolean"
			//	              },
			//	              "ObjectIdentifier": {
			//	                "description": "String that contains X.509 ObjectIdentifier information.",
			//	                "type": "string"
			//	              },
			//	              "Value": {
			//	                "type": "string"
			//	              }
			//	            },
			//	            "required": [
			//	              "ObjectIdentifier",
			//	              "Value"
			//	            ],
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "ExtendedKeyUsage": {
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "Structure that contains X.509 ExtendedKeyUsage information.",
			//	            "properties": {
			//	              "ExtendedKeyUsageObjectIdentifier": {
			//	                "description": "String that contains X.509 ObjectIdentifier information.",
			//	                "type": "string"
			//	              },
			//	              "ExtendedKeyUsageType": {
			//	                "type": "string"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "KeyUsage": {
			//	          "additionalProperties": false,
			//	          "description": "Structure that contains X.509 KeyUsage information.",
			//	          "properties": {
			//	            "CRLSign": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "DataEncipherment": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "DecipherOnly": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "DigitalSignature": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "EncipherOnly": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "KeyAgreement": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "KeyCertSign": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "KeyEncipherment": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            },
			//	            "NonRepudiation": {
			//	              "default": false,
			//	              "type": "boolean"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "SubjectAlternativeNames": {
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
			//	            "properties": {
			//	              "DirectoryName": {
			//	                "additionalProperties": false,
			//	                "description": "Structure that contains X.500 distinguished name information.",
			//	                "properties": {
			//	                  "CommonName": {
			//	                    "type": "string"
			//	                  },
			//	                  "Country": {
			//	                    "type": "string"
			//	                  },
			//	                  "CustomAttributes": {
			//	                    "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
			//	                    "items": {
			//	                      "additionalProperties": false,
			//	                      "description": "Structure that contains X.500 attribute type and value.",
			//	                      "properties": {
			//	                        "ObjectIdentifier": {
			//	                          "description": "String that contains X.509 ObjectIdentifier information.",
			//	                          "type": "string"
			//	                        },
			//	                        "Value": {
			//	                          "type": "string"
			//	                        }
			//	                      },
			//	                      "required": [
			//	                        "ObjectIdentifier",
			//	                        "Value"
			//	                      ],
			//	                      "type": "object"
			//	                    },
			//	                    "type": "array"
			//	                  },
			//	                  "DistinguishedNameQualifier": {
			//	                    "type": "string"
			//	                  },
			//	                  "GenerationQualifier": {
			//	                    "type": "string"
			//	                  },
			//	                  "GivenName": {
			//	                    "type": "string"
			//	                  },
			//	                  "Initials": {
			//	                    "type": "string"
			//	                  },
			//	                  "Locality": {
			//	                    "type": "string"
			//	                  },
			//	                  "Organization": {
			//	                    "type": "string"
			//	                  },
			//	                  "OrganizationalUnit": {
			//	                    "type": "string"
			//	                  },
			//	                  "Pseudonym": {
			//	                    "type": "string"
			//	                  },
			//	                  "SerialNumber": {
			//	                    "type": "string"
			//	                  },
			//	                  "State": {
			//	                    "type": "string"
			//	                  },
			//	                  "Surname": {
			//	                    "type": "string"
			//	                  },
			//	                  "Title": {
			//	                    "type": "string"
			//	                  }
			//	                },
			//	                "type": "object"
			//	              },
			//	              "DnsName": {
			//	                "description": "String that contains X.509 DnsName information.",
			//	                "type": "string"
			//	              },
			//	              "EdiPartyName": {
			//	                "additionalProperties": false,
			//	                "description": "Structure that contains X.509 EdiPartyName information.",
			//	                "properties": {
			//	                  "NameAssigner": {
			//	                    "type": "string"
			//	                  },
			//	                  "PartyName": {
			//	                    "type": "string"
			//	                  }
			//	                },
			//	                "required": [
			//	                  "PartyName",
			//	                  "NameAssigner"
			//	                ],
			//	                "type": "object"
			//	              },
			//	              "IpAddress": {
			//	                "description": "String that contains X.509 IpAddress information.",
			//	                "type": "string"
			//	              },
			//	              "OtherName": {
			//	                "additionalProperties": false,
			//	                "description": "Structure that contains X.509 OtherName information.",
			//	                "properties": {
			//	                  "TypeId": {
			//	                    "description": "String that contains X.509 ObjectIdentifier information.",
			//	                    "type": "string"
			//	                  },
			//	                  "Value": {
			//	                    "type": "string"
			//	                  }
			//	                },
			//	                "required": [
			//	                  "TypeId",
			//	                  "Value"
			//	                ],
			//	                "type": "object"
			//	              },
			//	              "RegisteredId": {
			//	                "description": "String that contains X.509 ObjectIdentifier information.",
			//	                "type": "string"
			//	              },
			//	              "Rfc822Name": {
			//	                "description": "String that contains X.509 Rfc822Name information.",
			//	                "type": "string"
			//	              },
			//	              "UniformResourceIdentifier": {
			//	                "description": "String that contains X.509 UniformResourceIdentifier information.",
			//	                "type": "string"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        }
			//	      },
			//	      "type": "object"
			//	    },
			//	    "Subject": {
			//	      "additionalProperties": false,
			//	      "description": "Structure that contains X.500 distinguished name information.",
			//	      "properties": {
			//	        "CommonName": {
			//	          "type": "string"
			//	        },
			//	        "Country": {
			//	          "type": "string"
			//	        },
			//	        "CustomAttributes": {
			//	          "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "description": "Structure that contains X.500 attribute type and value.",
			//	            "properties": {
			//	              "ObjectIdentifier": {
			//	                "description": "String that contains X.509 ObjectIdentifier information.",
			//	                "type": "string"
			//	              },
			//	              "Value": {
			//	                "type": "string"
			//	              }
			//	            },
			//	            "required": [
			//	              "ObjectIdentifier",
			//	              "Value"
			//	            ],
			//	            "type": "object"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "DistinguishedNameQualifier": {
			//	          "type": "string"
			//	        },
			//	        "GenerationQualifier": {
			//	          "type": "string"
			//	        },
			//	        "GivenName": {
			//	          "type": "string"
			//	        },
			//	        "Initials": {
			//	          "type": "string"
			//	        },
			//	        "Locality": {
			//	          "type": "string"
			//	        },
			//	        "Organization": {
			//	          "type": "string"
			//	        },
			//	        "OrganizationalUnit": {
			//	          "type": "string"
			//	        },
			//	        "Pseudonym": {
			//	          "type": "string"
			//	        },
			//	        "SerialNumber": {
			//	          "type": "string"
			//	        },
			//	        "State": {
			//	          "type": "string"
			//	        },
			//	        "Surname": {
			//	          "type": "string"
			//	        },
			//	        "Title": {
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "These are fields to be overridden in a certificate at the time of issuance. These requires an API_Passthrough template be used or they will be ignored.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"extensions": {
						// Property: Extensions
						Description: "Structure that contains X.500 extensions for a Certificate.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"certificate_policies": {
									// Property: CertificatePolicies
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"cert_policy_id": {
												// Property: CertPolicyId
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"policy_qualifiers": {
												// Property: PolicyQualifiers
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"policy_qualifier_id": {
															// Property: PolicyQualifierId
															Type:     types.StringType,
															Computed: true,
														},
														"qualifier": {
															// Property: Qualifier
															Description: "Structure that contains a X.509 policy qualifier.",
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"cps_uri": {
																		// Property: CpsUri
																		Type:     types.StringType,
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
										},
									),
									Computed: true,
								},
								"custom_extensions": {
									// Property: CustomExtensions
									Description: "Array of X.509 extensions for a certificate.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"critical": {
												// Property: Critical
												Type:     types.BoolType,
												Computed: true,
											},
											"object_identifier": {
												// Property: ObjectIdentifier
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Computed:    true,
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
								"extended_key_usage": {
									// Property: ExtendedKeyUsage
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"extended_key_usage_object_identifier": {
												// Property: ExtendedKeyUsageObjectIdentifier
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"extended_key_usage_type": {
												// Property: ExtendedKeyUsageType
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"key_usage": {
									// Property: KeyUsage
									Description: "Structure that contains X.509 KeyUsage information.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"crl_sign": {
												// Property: CRLSign
												Type:     types.BoolType,
												Computed: true,
											},
											"data_encipherment": {
												// Property: DataEncipherment
												Type:     types.BoolType,
												Computed: true,
											},
											"decipher_only": {
												// Property: DecipherOnly
												Type:     types.BoolType,
												Computed: true,
											},
											"digital_signature": {
												// Property: DigitalSignature
												Type:     types.BoolType,
												Computed: true,
											},
											"encipher_only": {
												// Property: EncipherOnly
												Type:     types.BoolType,
												Computed: true,
											},
											"key_agreement": {
												// Property: KeyAgreement
												Type:     types.BoolType,
												Computed: true,
											},
											"key_cert_sign": {
												// Property: KeyCertSign
												Type:     types.BoolType,
												Computed: true,
											},
											"key_encipherment": {
												// Property: KeyEncipherment
												Type:     types.BoolType,
												Computed: true,
											},
											"non_repudiation": {
												// Property: NonRepudiation
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"subject_alternative_names": {
									// Property: SubjectAlternativeNames
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"directory_name": {
												// Property: DirectoryName
												Description: "Structure that contains X.500 distinguished name information.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"common_name": {
															// Property: CommonName
															Type:     types.StringType,
															Computed: true,
														},
														"country": {
															// Property: Country
															Type:     types.StringType,
															Computed: true,
														},
														"custom_attributes": {
															// Property: CustomAttributes
															Description: "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"object_identifier": {
																		// Property: ObjectIdentifier
																		Description: "String that contains X.509 ObjectIdentifier information.",
																		Type:        types.StringType,
																		Computed:    true,
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
														"distinguished_name_qualifier": {
															// Property: DistinguishedNameQualifier
															Type:     types.StringType,
															Computed: true,
														},
														"generation_qualifier": {
															// Property: GenerationQualifier
															Type:     types.StringType,
															Computed: true,
														},
														"given_name": {
															// Property: GivenName
															Type:     types.StringType,
															Computed: true,
														},
														"initials": {
															// Property: Initials
															Type:     types.StringType,
															Computed: true,
														},
														"locality": {
															// Property: Locality
															Type:     types.StringType,
															Computed: true,
														},
														"organization": {
															// Property: Organization
															Type:     types.StringType,
															Computed: true,
														},
														"organizational_unit": {
															// Property: OrganizationalUnit
															Type:     types.StringType,
															Computed: true,
														},
														"pseudonym": {
															// Property: Pseudonym
															Type:     types.StringType,
															Computed: true,
														},
														"serial_number": {
															// Property: SerialNumber
															Type:     types.StringType,
															Computed: true,
														},
														"state": {
															// Property: State
															Type:     types.StringType,
															Computed: true,
														},
														"surname": {
															// Property: Surname
															Type:     types.StringType,
															Computed: true,
														},
														"title": {
															// Property: Title
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"dns_name": {
												// Property: DnsName
												Description: "String that contains X.509 DnsName information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"edi_party_name": {
												// Property: EdiPartyName
												Description: "Structure that contains X.509 EdiPartyName information.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"name_assigner": {
															// Property: NameAssigner
															Type:     types.StringType,
															Computed: true,
														},
														"party_name": {
															// Property: PartyName
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"ip_address": {
												// Property: IpAddress
												Description: "String that contains X.509 IpAddress information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"other_name": {
												// Property: OtherName
												Description: "Structure that contains X.509 OtherName information.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"type_id": {
															// Property: TypeId
															Description: "String that contains X.509 ObjectIdentifier information.",
															Type:        types.StringType,
															Computed:    true,
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
											"registered_id": {
												// Property: RegisteredId
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"rfc_822_name": {
												// Property: Rfc822Name
												Description: "String that contains X.509 Rfc822Name information.",
												Type:        types.StringType,
												Computed:    true,
											},
											"uniform_resource_identifier": {
												// Property: UniformResourceIdentifier
												Description: "String that contains X.509 UniformResourceIdentifier information.",
												Type:        types.StringType,
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
					"subject": {
						// Property: Subject
						Description: "Structure that contains X.500 distinguished name information.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"common_name": {
									// Property: CommonName
									Type:     types.StringType,
									Computed: true,
								},
								"country": {
									// Property: Country
									Type:     types.StringType,
									Computed: true,
								},
								"custom_attributes": {
									// Property: CustomAttributes
									Description: "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"object_identifier": {
												// Property: ObjectIdentifier
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Computed:    true,
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
								"distinguished_name_qualifier": {
									// Property: DistinguishedNameQualifier
									Type:     types.StringType,
									Computed: true,
								},
								"generation_qualifier": {
									// Property: GenerationQualifier
									Type:     types.StringType,
									Computed: true,
								},
								"given_name": {
									// Property: GivenName
									Type:     types.StringType,
									Computed: true,
								},
								"initials": {
									// Property: Initials
									Type:     types.StringType,
									Computed: true,
								},
								"locality": {
									// Property: Locality
									Type:     types.StringType,
									Computed: true,
								},
								"organization": {
									// Property: Organization
									Type:     types.StringType,
									Computed: true,
								},
								"organizational_unit": {
									// Property: OrganizationalUnit
									Type:     types.StringType,
									Computed: true,
								},
								"pseudonym": {
									// Property: Pseudonym
									Type:     types.StringType,
									Computed: true,
								},
								"serial_number": {
									// Property: SerialNumber
									Type:     types.StringType,
									Computed: true,
								},
								"state": {
									// Property: State
									Type:     types.StringType,
									Computed: true,
								},
								"surname": {
									// Property: Surname
									Type:     types.StringType,
									Computed: true,
								},
								"title": {
									// Property: Title
									Type:     types.StringType,
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
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of the issued certificate.",
			//	  "type": "string"
			//	}
			Description: "The ARN of the issued certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"certificate": {
			// Property: Certificate
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The issued certificate in base 64 PEM-encoded format.",
			//	  "type": "string"
			//	}
			Description: "The issued certificate in base 64 PEM-encoded format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"certificate_authority_arn": {
			// Property: CertificateAuthorityArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) for the private CA to issue the certificate.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) for the private CA to issue the certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"certificate_signing_request": {
			// Property: CertificateSigningRequest
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The certificate signing request (CSR) for the Certificate.",
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The certificate signing request (CSR) for the Certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"signing_algorithm": {
			// Property: SigningAlgorithm
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the algorithm that will be used to sign the Certificate.",
			//	  "type": "string"
			//	}
			Description: "The name of the algorithm that will be used to sign the Certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"template_arn": {
			// Property: TemplateArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.",
			//	  "type": "string"
			//	}
			Description: "Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.",
			Type:        types.StringType,
			Computed:    true,
		},
		"validity": {
			// Property: Validity
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The time before which the Certificate will be valid.",
			//	  "properties": {
			//	    "Type": {
			//	      "type": "string"
			//	    },
			//	    "Value": {
			//	      "type": "number"
			//	    }
			//	  },
			//	  "required": [
			//	    "Value",
			//	    "Type"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The time before which the Certificate will be valid.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.Float64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"validity_not_before": {
			// Property: ValidityNotBefore
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The time after which the Certificate will be valid.",
			//	  "properties": {
			//	    "Type": {
			//	      "type": "string"
			//	    },
			//	    "Value": {
			//	      "type": "number"
			//	    }
			//	  },
			//	  "required": [
			//	    "Value",
			//	    "Type"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The time after which the Certificate will be valid.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.Float64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ACMPCA::Certificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::Certificate").WithTerraformTypeName("awscc_acmpca_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_passthrough":                      "ApiPassthrough",
		"arn":                                  "Arn",
		"cert_policy_id":                       "CertPolicyId",
		"certificate":                          "Certificate",
		"certificate_authority_arn":            "CertificateAuthorityArn",
		"certificate_policies":                 "CertificatePolicies",
		"certificate_signing_request":          "CertificateSigningRequest",
		"common_name":                          "CommonName",
		"country":                              "Country",
		"cps_uri":                              "CpsUri",
		"critical":                             "Critical",
		"crl_sign":                             "CRLSign",
		"custom_attributes":                    "CustomAttributes",
		"custom_extensions":                    "CustomExtensions",
		"data_encipherment":                    "DataEncipherment",
		"decipher_only":                        "DecipherOnly",
		"digital_signature":                    "DigitalSignature",
		"directory_name":                       "DirectoryName",
		"distinguished_name_qualifier":         "DistinguishedNameQualifier",
		"dns_name":                             "DnsName",
		"edi_party_name":                       "EdiPartyName",
		"encipher_only":                        "EncipherOnly",
		"extended_key_usage":                   "ExtendedKeyUsage",
		"extended_key_usage_object_identifier": "ExtendedKeyUsageObjectIdentifier",
		"extended_key_usage_type":              "ExtendedKeyUsageType",
		"extensions":                           "Extensions",
		"generation_qualifier":                 "GenerationQualifier",
		"given_name":                           "GivenName",
		"initials":                             "Initials",
		"ip_address":                           "IpAddress",
		"key_agreement":                        "KeyAgreement",
		"key_cert_sign":                        "KeyCertSign",
		"key_encipherment":                     "KeyEncipherment",
		"key_usage":                            "KeyUsage",
		"locality":                             "Locality",
		"name_assigner":                        "NameAssigner",
		"non_repudiation":                      "NonRepudiation",
		"object_identifier":                    "ObjectIdentifier",
		"organization":                         "Organization",
		"organizational_unit":                  "OrganizationalUnit",
		"other_name":                           "OtherName",
		"party_name":                           "PartyName",
		"policy_qualifier_id":                  "PolicyQualifierId",
		"policy_qualifiers":                    "PolicyQualifiers",
		"pseudonym":                            "Pseudonym",
		"qualifier":                            "Qualifier",
		"registered_id":                        "RegisteredId",
		"rfc_822_name":                         "Rfc822Name",
		"serial_number":                        "SerialNumber",
		"signing_algorithm":                    "SigningAlgorithm",
		"state":                                "State",
		"subject":                              "Subject",
		"subject_alternative_names":            "SubjectAlternativeNames",
		"surname":                              "Surname",
		"template_arn":                         "TemplateArn",
		"title":                                "Title",
		"type":                                 "Type",
		"type_id":                              "TypeId",
		"uniform_resource_identifier":          "UniformResourceIdentifier",
		"validity":                             "Validity",
		"validity_not_before":                  "ValidityNotBefore",
		"value":                                "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
