package baremtlsdk

// AvailablityDomain contains name and then tenancy ID that an
import (
	"bytes"
	"encoding/json"
	"net/url"
)

// availability domain belongs to.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#AvailabilityDomain
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

// ListAvailablityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listAvailabilityDomains
func (c *Client) ListAvailablityDomains() (ads []AvailabilityDomain, e error) {
	url := buildIdentityURL(resourceAvailabilityDomains, url.Values{
		queryCompartmentID: []string{c.authInfo.tenancyOCID},
	})
	var getResp *requestResponse

	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	e = decoder.Decode(&ads)
	return
}
