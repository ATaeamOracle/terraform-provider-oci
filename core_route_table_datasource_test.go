package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreRouteTablesTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreRouteTablesTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_route_tables" "t" {
      compartment_id = "compartment_id"
      vcn_id = "vcn_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_route_tables.t"

}

func (s *ResourceCoreRouteTablesTestSuite) TestResourceListRouteTables() {
	opts := []baremetal.Options{}

	s.Client.On(
		"ListRouteTables",
		"compartment_id",
		"vcn_id",
		opts,
	).Return(
		&baremetal.ListRouteTables{
			RouteTables: []baremetal.RouteTable{
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id1",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id2",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListRouteTables", "compartment_id", "vcn_id", opts)

}

func (s *ResourceCoreRouteTablesTestSuite) TestResourceListRouteTablesPaged() {
	opts := []baremetal.Options{}

	s.Client.On(
		"ListRouteTables",
		"compartment_id",
		"vcn_id",
		opts,
	).Return(
		&baremetal.ListRouteTables{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			RouteTables: []baremetal.RouteTable{
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id1",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id2",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	opts2 := []baremetal.Options{baremetal.Options{Page: "nextpage"}}

	s.Client.On(
		"ListRouteTables",
		"compartment_id",
		"vcn_id",
		opts2,
	).Return(
		&baremetal.ListRouteTables{
			RouteTables: []baremetal.RouteTable{
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id3",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
				baremetal.RouteTable{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id4",
					RouteRules: []baremetal.RouteRule{
						baremetal.RouteRule{
							CidrBlock:         "cidr_block",
							DisplayName:       "display_name",
							NetworkEntityID:   "network_entity_id",
							NetworkEntityType: "network_entity_type",
							TimeCreated:       baremetal.Time{Time: time.Now()},
						},
					},
					TimeModified: baremetal.Time{Time: time.Now()},
					State:        baremetal.ResourceAvailable,
					TimeCreated:  baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListRouteTables", "compartment_id", "vcn_id", opts2)

}

func TestResourceCoreRouteTablesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTablesTestSuite))
}
