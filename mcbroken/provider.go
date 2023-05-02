package mcbroken

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "https://mcbroken.com/stats.json",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
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
	providerConfig["url"] = d.Get("url").(string)
	return providerConfig, diags
}

func hashCityString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
