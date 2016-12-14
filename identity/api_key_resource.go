package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func APIKeyResource() *schema.Resource {
	return &schema.Resource{
		Create: createAPIKey,
		Read:   readAPIKey,
		Delete: deleteAPIKey,
		Schema: map[string]*schema.Schema{
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func deleteAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{D: d, Client: client}
	return crud.DeleteResource(sync)
}
