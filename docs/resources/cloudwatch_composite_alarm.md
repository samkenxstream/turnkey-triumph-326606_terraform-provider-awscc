---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudwatch_composite_alarm Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
---

# awscc_cloudwatch_composite_alarm (Resource)

The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **alarm_name** (String) The name of the Composite Alarm
- **alarm_rule** (String) Expression which aggregates the state of other Alarms (Metric or Composite Alarms)

### Optional

- **actions_enabled** (Boolean) Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
- **alarm_actions** (List of String) The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
- **alarm_description** (String) The description of the alarm
- **insufficient_data_actions** (List of String) The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
- **ok_actions** (List of String) The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).

### Read-Only

- **arn** (String) Amazon Resource Name (ARN) of the alarm
- **id** (String) Uniquely identifies the resource.

