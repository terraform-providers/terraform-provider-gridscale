package gridscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccdataSourceGridscaleIPv4_basic(t *testing.T) {
	name := fmt.Sprintf("object-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGridscaleIpv4DestroyCheck,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckDataSourceIPv4Config_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.gridscale_ipv4.foo", "id"),
					resource.TestCheckResourceAttr("data.gridscale_ipv4.foo", "name", name),
				),
			},
		},
	})

}

func testAccCheckDataSourceIPv4Config_basic(name string) string {
	return fmt.Sprintf(`

resource "gridscale_ipv4" "foo" {
	name   = "%s"
}


data "gridscale_ipv4" "foo" {
	resource_id   = gridscale_ipv4.foo.id
}

`, name)
}
