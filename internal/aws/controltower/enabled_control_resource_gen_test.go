// Code generated by generators/resource/main.go; DO NOT EDIT.

package controltower_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSControlTowerEnabledControl_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ControlTower::EnabledControl", "awscc_controltower_enabled_control", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
