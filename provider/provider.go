package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"rio_test":  dataTestSchema(dataTest),
			"rio_error": dataTestSchema(dataTestFail),
		},
	}
}

func dataTestSchema(f schema.ReadContextFunc) *schema.Resource {
	return &schema.Resource{
		ReadContext: f,
		Schema: map[string]*schema.Schema{
			"test": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataTest(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if err := d.Set("test", true); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataTestFail(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := fmt.Errorf("This is an error")
	return diag.FromErr(err)
}
