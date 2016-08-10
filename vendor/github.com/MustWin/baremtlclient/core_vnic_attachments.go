package baremtlsdk

import (
	"encoding/json"
	"net/url"
	"time"
)

// VnicAttachment Vnic information for a particular instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#VnicAttachment
type VnicAttachment struct {
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	State              string    `json:"state"`
	SubnetID           string    `json:"subnetId"`
	TimeCreated        time.Time `json:"TimeCreated"`
	VnicID             string    `json:"vnicId"`
}

// VnicAttachmentList list of VnicAttachments as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating results.
// OPCRequestID is used to identify the request for support issues.
type VnicAttachmentList struct {
	OPCNextPage  string
	OPCRequestID string
	Attachments  []VnicAttachment
}

// ListVnicAttachments returns a list of VnicAttachments with matching compartmentID
// and optionally instanceId, vnicId, and/or availabilityDomain. Optional parameters
// are assigned to the optional CoreOptions argument.  Page and Limit can also
// be supplied to support pagination.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listVnicAttachments
func (c *Client) ListVnicAttachments(compartmentID string, opts ...CoreOptions) (res *VnicAttachmentList, e error) {
	query := url.Values{}
	query.Set(queryCompartmentID, compartmentID)
	c.setCoreOptions(query, opts...)
	urlStr := buildCoreURL(resourceVnicAttachments, query)

	var resp *requestResponse
	if resp, e = c.api.getRequest(urlStr, nil); e != nil {
		return
	}

	res = &VnicAttachmentList{
		OPCNextPage:  resp.header.Get(headerOPCNextPage),
		OPCRequestID: resp.header.Get(headerOPCRequestID),
	}

	if e = json.Unmarshal(resp.body, &res.Attachments); e != nil {
		return
	}

	return
}
