package deprecated

import (
	"fmt"
	"testing"

	"github.com/g42cloud-terraform/terraform-provider-g42vbcloud/g42vbcloud/services/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDcsAZV1DataSource_basic(t *testing.T) {
	resourceName := "data.g42vbcloud_dcs_az.az1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDcsAZV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcsAZV1DataSourceID(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "code"),
					resource.TestCheckResourceAttrSet(resourceName, "port"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
		},
	})
}

func testAccCheckDcsAZV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find DCS az data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DCS az data source ID not set")
		}

		return nil
	}
}

var testAccDcsAZV1DataSource_basic = fmt.Sprintf(`
data "g42vbcloud_availability_zones" "test" {}

data "g42vbcloud_dcs_az" "az1" {
  code = data.g42vbcloud_availability_zones.test.names[0]
}
`)
