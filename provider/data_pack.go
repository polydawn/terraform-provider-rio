package provider

import (
	"context"
	"fmt"

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
			"target": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Warehouse for the packed ware. Leave empty to scan only.`,
			},
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Configure filters for file properties, such as mtime, uid, gid, etc. By default many of these attribute will be flattened.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
	target := api.WarehouseLocation(d.Get("target").(string))

	// handle filters
	filters := d.Get("filters").(map[string]interface{})

	filter := api.FilesetPackFilter_Conservative
	for k, v := range filters {
		filt, err := api.ParseFilesetPackFilter(fmt.Sprintf("%s=%s", k, v))
		if err != nil {
			return diag.FromErr(err)
		}

		filter = filter.Apply(filt)
	}

	wareId, err := packFunc(ctx, api.PackType(packType), path, filter, target, rio.Monitor{})
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
