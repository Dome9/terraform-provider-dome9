package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
	"testing"
)

func TestAccDataSourceAWSUnifiedOnboardingUpdateVersionStackConfogurationBasic(t *testing.T) {
	resourceTypeAndName, dataTypeAndName, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)
	log.Println("TestAccDataSourceAWSUnifiedOnboardingUpdateVersionStackConfogurationBasic ",resourceTypeAndName, dataTypeAndName, resourceName)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSUnifiedOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingUpdateVersionStackConfogurationBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName+variable.DataSourceSuffix, "ID", resourceTypeAndName, "ID"),
					resource.TestCheckResourceAttrPair(resourceName+variable.DataSourceSuffix, "stack_name", resourceTypeAndName, "stack_name"),
				),
			},
		},
	})
}

func testAccCheckAwsUnifiedOnbordingUpdateVersionStackConfogurationBasic(resourceTypeAndName string, generatedName string) string {
	res := fmt.Sprintf(`
// AwsUnifiedOnbording resource

%s

data "%s" "%s" {
  onboarding_id = "${%s.id}"
}
`,
		// continuous compliance notification resource
		getAwsUnifiedOnboardingHCL(generatedName),

		// data source variables
		resourcetype.AwsUnifiedOnboardingUpdateVersionStackConfig,
		generatedName+variable.DataSourceSuffix,
		resourceTypeAndName,
	)
	log.Printf("[INFO] testAccCheckAwsUnifiedOnbordingBasic:%+v\n", res)

	return res
}
