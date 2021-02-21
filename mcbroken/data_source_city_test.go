package mcbroken

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMcbrokenCity(t *testing.T) {
	brokenNumberRegex := regexp.MustCompile(`\d{1,3}(.\d{1,2})?`)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { /* no precheck needed testAccPreCheck(t) */ },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// valid city
			{
				Config: testAccCheckCity("Dallas"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.mcbroken_city.dallas", "city", "dallas"),
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_city.dallas", "broken"),
					resource.TestMatchResourceAttr(
						"data.mcbroken_city.dallas", "broken", brokenNumberRegex),
				),
			},
			// invalid city
			{
				Config: testAccCheckCity("not_found"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.mcbroken_city.not_found", "city", "not_found"),
					resource.TestCheckResourceAttr(
						"data.mcbroken_city.not_found", "broken", "-1"),
				),
			},
		},
	})
}

func testAccCheckCity(city string) string {
	return fmt.Sprintf(`
data "mcbroken_city" "%[1]v" {
  city = "%[1]v"
}
`, strings.ToLower(city))
}
