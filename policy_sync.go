package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type PolicySync struct {
	d      *schema.ResourceData
	client BareMetalClient
}

func (s *PolicySync) toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicySync) Create() (res BareMetalResource, e error) {
	name := s.d.Get("name").(string)
	description := s.d.Get("description").(string)
	statements := s.toStringArray(s.d.Get("statements"))

	var raw *baremtlsdk.Policy
	raw, e = s.client.CreatePolicy(name, description, statements)
	res = &BareMetalPolicyAdapter{raw}
	return
}

func (s *PolicySync) Get() (res BareMetalResource, e error) {
	var raw *baremtlsdk.Policy
	raw, e = s.client.GetPolicy(s.d.Id())
	res = &BareMetalPolicyAdapter{raw}
	return
}

func (s *PolicySync) Update() (res BareMetalResource, e error) {
	description := s.d.Get("description").(string)
	statements := s.toStringArray(s.d.Get("statements"))
	var raw *baremtlsdk.Policy
	raw, e = s.client.UpdatePolicy(s.d.Id(), description, statements)
	res = &BareMetalPolicyAdapter{raw}
	return
}

func (s *PolicySync) SetData(res BareMetalResource) {
	a := res.(*BareMetalPolicyAdapter)
	s.d.Set("statements", a.Statements)
	s.d.Set("name", a.Name)
	s.d.Set("description", a.Description)
	s.d.Set("compartment_id", a.CompartmentID)
	s.d.Set("state", a.State)
	s.d.Set("time_modified", a.TimeModified.String())
	s.d.Set("time_created", a.TimeCreated.String())
}

func (s *PolicySync) Delete() (e error) {
	return s.client.DeletePolicy(s.d.Id())
}
