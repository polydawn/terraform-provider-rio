package provider

import (
	"context"

	api "go.polydawn.net/go-timeless-api"
	"go.polydawn.net/go-timeless-api/rio"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePackSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataPack,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of pack format to use. ["tar", "zip"]`,
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The fileset to scan and pack (absolute path).`,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The id of the packed fileset`,
			},
		},
	}
}

// Currently support scanning only
func dataPack(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	packType := d.Get("type").(string)
	packFunc, err := demuxPackTool(packType)
	if err != nil {
		return diag.FromErr(err)
	}
	path := d.Get("path").(string)

	wareId, err := packFunc(ctx, api.PackType(packType), path, api.FilesetPackFilter_Conservative, "", rio.Monitor{})
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("id", wareId.String()); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(wareId.String())
	// d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
