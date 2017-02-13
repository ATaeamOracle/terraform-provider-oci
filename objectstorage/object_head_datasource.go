// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ObjectHeadDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectHead,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content-length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"content-type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func readObjectHead(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &ObjectHeadDatasourceCrud{
		D:      d,
		Client: client,
	}

	return crud.ReadResource(reader)
}
