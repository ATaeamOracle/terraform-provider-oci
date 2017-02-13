// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func BucketResource() *schema.Resource {
	return &schema.Resource{
		Create: createBucket,
		Read:   readBucket,
		Update: updateBucket,
		Delete: deleteBucket,
		Schema: bucketSchema,
	}
}

func createBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{D: d, Client: m.(client.BareMetalClient)}
	return crud.CreateResource(d, sync)
}

func readBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{D: d, Client: m.(client.BareMetalClient)}
	return crud.ReadResource(sync)
}

func updateBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{D: d, Client: m.(client.BareMetalClient)}
	return crud.UpdateResource(d, sync)
}

func deleteBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{D: d, Client: m.(client.BareMetalClient)}
	return crud.DeleteResource(sync)
}
