package mcbroken

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCities() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCitiesRead,
		Schema: map[string]*schema.Schema{
			"cities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"city": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"broken": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"broken": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourceCitiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	providerConfig := m.(map[string]interface{})
	url := providerConfig["url"].(string)

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// Unmarshal data
	brokenCities := &Mcbroken{}
	err = json.NewDecoder(r.Body).Decode(&brokenCities)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set national average broken percentage
	if err := d.Set("broken", brokenCities.Broken); err != nil {
		return diag.FromErr(err)
	}

	cities := make([]map[string]interface{}, 0)
	// Set broken values for all currently available city data
	for _, v := range brokenCities.Cities {
		city := make(map[string]interface{})
		city["city"] = v.City
		city["broken"] = v.Broken
		cities = append(cities, city)
	}
	if err := d.Set("cities", cities); err != nil {
		return diag.FromErr(err)
	}

	// Always run to give unique id
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
