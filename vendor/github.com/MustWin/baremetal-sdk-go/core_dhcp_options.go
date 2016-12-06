package baremetal

import "net/http"

// DHCPDNSOption specifies how DNS (host name resolution) is handled in the VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPDNSOption/
type DHCPDNSOption struct {
	Type             string   `json:"type"`
	CustomDNSServers []string `json:"customDnsServers"`
	ServerType       string   `json:"serverType"`
}

// DHCPOptions contains a set of dhcp options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/
type DHCPOptions struct {
	RequestableResource
	ETaggedResource
	CompartmentID string          `json:"compartmentId"`
	DisplayName   string          `json:"displayName"`
	ID            string          `json:"id"`
	Options       []DHCPDNSOption `json:"options"`
	State         string          `json:"lifecycleState"`
	TimeCreated   Time            `json:"timeCreated"`
}

// ListDHCPOptions contains a list of dhcp options
//
type ListDHCPOptions struct {
	ResourceContainer
	DHCPOptions []DHCPOptions
}

func (l *ListDHCPOptions) GetList() interface{} {
	return &l.DHCPOptions
}

// CreateDHCPOptions creates a new set of DHCP options for the specified VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/CreateDHCPOptions
func (c *Client) CreateDHCPOptions(compartmentID, vcnID string, dhcpOptions []DHCPDNSOption, opts *CreateOptions) (res *DHCPOptions, e error) {
	required := struct {
		ocidRequirement
		Options []DHCPDNSOption `json:"options" url:"-"`
		VcnID   string          `json:"vcnId" url:"-"`
	}{
		Options: dhcpOptions,
		VcnID:   vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDHCPOptions,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = response.unmarshal(res)
	return
}

// GetDHCPOptions gets the specified set of DHCP options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/GetDHCPOptions
func (c *Client) GetDHCPOptions(id string) (res *DHCPOptions, e error) {
	details := &requestDetails{
		name: resourceDHCPOptions,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = resp.unmarshal(res)
	return
}

// UpdateDHCPOptions updates the specified set of DHCP options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/UpdateDHCPOptions
func (c *Client) UpdateDHCPOptions(id string, opts *UpdateDHCPDNSOptions) (res *DHCPOptions, e error) {
	details := &requestDetails{
		name:     resourceDHCPOptions,
		ids:      urlParts{id},
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = response.unmarshal(res)
	return
}

// DeleteDHCPOptions deletes the specified set of DHCP options, but only if it's
// not in use by a subnet
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/DeleteDHCPOptions
func (c *Client) DeleteDHCPOptions(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceDHCPOptions,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListDHCPOptions gets a list of the sets of DHCP options in the specified VCN
// and specified compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/ListDHCPOptions
func (c *Client) ListDHCPOptions(compartmentID, vcnID string, opts *ListOptions) (res *ListDHCPOptions, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDHCPOptions,
		required: required,
		optional: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDHCPOptions{}
	e = resp.unmarshal(res)
	return
}
