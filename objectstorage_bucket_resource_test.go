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

type ResourceObjectstorageBucketTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Bucket
	Namespace    baremetal.Namespace
}

func (s *ResourceObjectstorageBucketTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_objectstorage_bucket" "t" {
			compartment_id = "compartment_id"
			name = "name"
			namespace = "namespace"
			metadata = {
				"foo" = "bar"
			}
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_objectstorage_bucket.t"
	metadata := map[string]string{
		"foo": "bar",
	}
	s.Namespace = baremetal.Namespace("namespace")
	s.Res = &baremetal.Bucket{
		CompartmentID: "compartment_id",
		Name:          "name",
		Namespace:     s.Namespace,
		Metadata:      metadata,
		CreatedBy:     "created_by",
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	opts := &baremetal.CreateBucketOptions{
		Metadata: metadata,
	}
	s.Client.On(
		"CreateBucket",
		"compartment_id",
		"name",
		string(s.Namespace),
		opts).Return(s.Res, nil)
	s.Client.On("DeleteBucket", "name", string(s.Namespace), (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceObjectstorageBucketTestSuite) TestCreateResourceObjectstorageBucket() {
	s.Client.On("GetBucket", "name", string(s.Namespace)).Return(s.Res, nil).Times(2)
	s.Client.On("GetBucket", "name", string(s.Namespace)).Return(nil, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", string(s.Res.Namespace)),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageBucketTestSuite) TestUpdateResourceObjectstorageBucket() {
	s.Client.On("GetBucket", "name", string(s.Namespace)).Return(s.Res, nil).Times(2)

	config := `
		resource "baremetal_objectstorage_bucket" "t" {
			compartment_id = "compartment_id"
			name = "new_name"
			namespace = "namespace"
			metadata = {
				"foo" = "bar"
			}
		}
	`
	config += testProviderConfig
	metadata := map[string]string{
		"foo": "bar",
	}

	res := &baremetal.Bucket{
		CompartmentID: "compartment_id",
		Name:          "new_name",
		Namespace:     "namespace",
		Metadata:      metadata,
		CreatedBy:     "created_by",
		TimeCreated:   s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateBucketOptions{
		Metadata: metadata,
	}
	s.Client.On("UpdateBucket",
		res.CompartmentID, "new_name", string(res.Namespace), opts).Return(res, nil)
	s.Client.On("GetBucket", "new_name", string(s.Namespace)).Return(res, nil)
	s.Client.On("DeleteBucket", "new_name", string(s.Namespace), (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", res.Name),
				),
			},
		},
	})
}

func (s *ResourceObjectstorageBucketTestSuite) TestDeleteResourceObjectstorageBucket() {
	s.Client.On("GetBucket", "name", string(s.Namespace)).Return(s.Res, nil).Times(2)
	s.Client.On("GetBucket", "name", string(s.Namespace)).Return(nil, nil)
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})
	s.Client.AssertCalled(s.T(), "DeleteBucket", "name", string(s.Namespace), (*baremetal.IfMatchOptions)(nil))
}

func TestResourceobjectstorageBucketTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectstorageBucketTestSuite))
}
