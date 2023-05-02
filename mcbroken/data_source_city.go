package mcbroken

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCity() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCityRead,
		Schema: map[string]*schema.Schema{
			"city": {
				Type:     schema.TypeString,
				Required: true,
			},
			"broken": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourceCityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var httpTimeout time.Duration = 10
	client := &http.Client{Timeout: httpTimeout * time.Second}
	providerConfig := m.(map[string]interface{})
	url := providerConfig["url"].(string)
	userChosenCity := d.Get("city").(string)

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

	// In the event a city isn't found, broken value will be -1
	if err := d.Set("broken", -1); err != nil {
		return diag.FromErr(err)
	}

	// Search all available cities for user provided one
	for _, v := range brokenCities.Cities {
		if strings.EqualFold(v.City, userChosenCity) {
			if err := d.Set("broken", v.Broken); err != nil {
				return diag.FromErr(err)
			}
			break
		}
	}

	d.SetId(hashCityString(userChosenCity))

	return diags
}
