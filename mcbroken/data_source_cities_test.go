package mcbroken

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMcbrokenCities(t *testing.T) {
	brokenNumberRegex := regexp.MustCompile(`\d{1,3}(.\d{1,2})?`)
	cityRegex := regexp.MustCompile(`[a-zA-Z]+$`)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { /* no precheck needed testAccPreCheck(t) */ },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCities(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_cities.all", "broken"),
					resource.TestMatchResourceAttr(
						"data.mcbroken_cities.all", "broken", brokenNumberRegex),
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_cities.all", "cities.0.city"),
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_cities.all", "cities.0.broken"),
					resource.TestMatchResourceAttr(
						"data.mcbroken_cities.all", "cities.0.city", cityRegex),
					resource.TestMatchResourceAttr(
						"data.mcbroken_cities.all", "cities.0.broken", brokenNumberRegex),
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_cities.all", "cities.1.city"),
					resource.TestCheckResourceAttrSet(
						"data.mcbroken_cities.all", "cities.1.broken"),
					resource.TestMatchResourceAttr(
						"data.mcbroken_cities.all", "cities.1.city", cityRegex),
					resource.TestMatchResourceAttr(
						"data.mcbroken_cities.all", "cities.1.broken", brokenNumberRegex),
				),
			},
		},
	})
}

func testAccCheckCities() string {
	return `data "mcbroken_cities" "all" {}`
}
