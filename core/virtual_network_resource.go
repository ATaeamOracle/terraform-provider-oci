package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCoreVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createVirtualNetwork,
		Read:   readVirtualNetwork,
		Delete: deleteVirtualNetwork,
		Schema: map[string]*schema.Schema{
			"cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_routing_table_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_security_list_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkSync{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkSync{D: d, Client: client}
	return crud.ReadResource(sync)
}

func deleteVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkSync{D: d, Client: client}
	return crud.DeleteResource(sync)
}
