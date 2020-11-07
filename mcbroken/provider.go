package mcbroken

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://mcbroken.com/stats.json",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"mcbroken_cities": dataSourceCities(),
			"mcbroken_city":   dataSourceCity(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	providerConfig := make(map[string]interface{})
	mcbrokenEndpoint := d.Get("url").(string)

	_, err := url.Parse(mcbrokenEndpoint)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to parse url",
			Detail:   "Mcbroken endpoint provided is not a valid url",
		})

	}

	providerConfig["url"] = mcbrokenEndpoint
	return providerConfig, diags
}
