package baremetal

import "net/http"

// Cpe describes customer premise equipment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Cpe/
type Cpe struct {
	ETaggedResource
	ID            string `json:"id"`
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	IPAddress     string `json:"ipAddress"`
	TimeCreated   Time   `json:"timeCreated"`
}

// CpeList contains a list of customer premise equipment
//
type ListCpes struct {
	ResourceContainer
	Cpes []Cpe
}

func (l *ListCpes) GetList() interface{} {
	return &l.Cpes
}

// ListCpes returns a list of customer premise equipment for a particular compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Cpe/ListCpes
func (c *Client) ListCpes(compartmentID string, opts *ListOptions) (cpes *ListCpes, e error) {
	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		required: listOCIDRequirement{CompartmentID: compartmentID},
		optional: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	cpes = &ListCpes{}
	e = resp.unmarshal(cpes)
	return
}

// CreateCpe is used to define customer premise equipment such as routers
// in the Bare Metal cloud
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Cpe/CreateCpe
func (c *Client) CreateCpe(compartmentID, ipAddress string, opts *CreateOptions) (cpe *Cpe, e error) {
	required := struct {
		ocidRequirement
		IPAddress string `json:"ipAddress" url:"-"`
	}{
		IPAddress: ipAddress,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		optional: opts,
		required: required,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)
	return
}

// GetCpe retrieves information on a customer premise equipment resource.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Cpe/GetCpe
func (c *Client) GetCpe(id string) (cpe *Cpe, e error) {
	details := &requestDetails{
		name: resourceCustomerPremiseEquipment,
		ids:  urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)
	return
}

// DeleteCpe removes customer premise equipment resource
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Cpe/DeleteCpe
func (c *Client) DeleteCpe(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}
