// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumesTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_volumes" "t" {
      availability_domain = "availability_domain"
      compartment_id = "${var.compartment_id}"
      limit = 1
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_volumes.t"
}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumes() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),

					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "2"),
				),
			},
		},
	},
	)

}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumesWithPagination() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),

					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "4"),
				),
			},
		},
	},
	)

}

func TestResourceCoreVolumesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumesTestSuite))
}
