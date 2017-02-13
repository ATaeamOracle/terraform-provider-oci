// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

var transportSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
}

var icmpSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
}

func SecurityListResource() *schema.Resource {
	return &schema.Resource{
		Create: createSecurityList,
		Read:   readSecurityList,
		Update: updateSecurityList,
		Delete: deleteSecurityList,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"egress_security_rules": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"icmp_options": icmpSchema,
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
					},
				},
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_security_rules": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_options": icmpSchema,
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
					},
				},
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, crd)
}

func readSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{D: d, Client: client}
	return crud.ReadResource(crd)
}

func updateSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, crd)
}

func deleteSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{D: d, Client: client}
	return crud.DeleteResource(crd)
}
