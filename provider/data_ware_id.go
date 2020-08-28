package provider

import (
	"context"
	"fmt"
	"strings"

	api "go.polydawn.net/go-timeless-api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceParseWareIdSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataParseWareId,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The ware ID`,
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf(`The ware type. [%s]`, strings.Join(packTypes, ", ")),
			},
			"hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The hash of the ware.`,
			},
		},
	}
}

func dataParseWareId(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	wareIdString := d.Get("id").(string)
	wareId, err := api.ParseWareID(wareIdString)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", string(wareId.Type)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hash", string(wareId.Hash)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(wareIdString)

	return diags
}

func dataSourceWareIdSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataWareId,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The ware ID`,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: fmt.Sprintf(`The ware type. [%s]`, strings.Join(packTypes, ", ")),
			},
			"hash": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The hash of the ware.`,
			},
		},
	}
}

func dataWareId(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	packType := api.PackType(d.Get("type").(string))

	hash := d.Get("hash").(string)

	wareId := api.WareID{
		Type: packType,
		Hash: hash,
	}

	d.SetId(wareId.String())

	return diags
}
