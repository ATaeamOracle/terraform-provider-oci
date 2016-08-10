package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVnicAttachmentsTestSuite struct {
	suite.Suite
	Client       *MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVnicAttachmentsTestSuite) SetupTest() {
	s.Client = &MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_vnic_attachments" "s" {
      compartment_id = "compartmentid"
      availability_domain = "availabilityid"
      vnic_id = "vnicid"
      instance_id = "instanceid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_vnic_attachments.s"

}

func (s *ResourceCoreVnicAttachmentsTestSuite) TestResourceReadCoreVnicAttachments() {
	opts := []baremtlsdk.CoreOptions{
		baremtlsdk.CoreOptions{
			AvailabilityDomain: "availabilityid",
			VnicID:             "vnicid",
			InstanceID:         "instanceid",
		},
	}

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts,
	).Return(
		&baremtlsdk.VnicAttachmentList{
			Attachments: []baremtlsdk.VnicAttachment{
				baremtlsdk.VnicAttachment{
					ID:                 "id1",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att1",
					InstanceID:         "instanceid",
					State:              baremtlsdk.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now(),
				},
				baremtlsdk.VnicAttachment{
					ID:                 "id2",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att2",
					InstanceID:         "instanceid",
					State:              baremtlsdk.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now().Add(fiveMinutes),
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_id", "vnicid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instance_id", "instanceid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.1.id", "id2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVnicAttachments", "compartmentid", opts)

}

func TestResourceCoreVnicAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVnicAttachmentsTestSuite))
}
