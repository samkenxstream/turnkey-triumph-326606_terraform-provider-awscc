---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sso_instance_access_control_attribute_configuration Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SSO::InstanceAccessControlAttributeConfiguration
---

# awscc_sso_instance_access_control_attribute_configuration (Data Source)

Data Source schema for AWS::SSO::InstanceAccessControlAttributeConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **access_control_attributes** (Attributes List) (see [below for nested schema](#nestedatt--access_control_attributes))
- **instance_access_control_attribute_configuration** (Attributes) The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead. (see [below for nested schema](#nestedatt--instance_access_control_attribute_configuration))
- **instance_arn** (String) The ARN of the AWS SSO instance under which the operation will be executed.

<a id="nestedatt--access_control_attributes"></a>
### Nested Schema for `access_control_attributes`

Read-Only:

- **key** (String)
- **value** (Attributes) (see [below for nested schema](#nestedatt--access_control_attributes--value))

<a id="nestedatt--access_control_attributes--value"></a>
### Nested Schema for `access_control_attributes.value`

Read-Only:

- **source** (List of String)



<a id="nestedatt--instance_access_control_attribute_configuration"></a>
### Nested Schema for `instance_access_control_attribute_configuration`

Read-Only:

- **access_control_attributes** (Attributes List) (see [below for nested schema](#nestedatt--instance_access_control_attribute_configuration--access_control_attributes))

<a id="nestedatt--instance_access_control_attribute_configuration--access_control_attributes"></a>
### Nested Schema for `instance_access_control_attribute_configuration.access_control_attributes`

Read-Only:

- **key** (String)
- **value** (Attributes) (see [below for nested schema](#nestedatt--instance_access_control_attribute_configuration--access_control_attributes--value))

<a id="nestedatt--instance_access_control_attribute_configuration--access_control_attributes--value"></a>
### Nested Schema for `instance_access_control_attribute_configuration.access_control_attributes.value`

Read-Only:

- **source** (List of String)

