package baremetal

import "time"

// VnicAttachment Vnic information for a particular instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VnicAttachment/
type VnicAttachment struct {
	AvailabilityDomain string    `json:"availabilityDomain"`
	CompartmentID      string    `json:"compartmentId"`
	DisplayName        string    `json:"displayName"`
	ID                 string    `json:"id"`
	InstanceID         string    `json:"instanceId"`
	State              string    `json:"lifecycleState"`
	SubnetID           string    `json:"subnetId"`
	TimeCreated        time.Time `json:"TimeCreated"`
	VnicID             string    `json:"vnicId"`
}

// ListVnicAttachments list of VnicAttachments as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating results.
// OPCRequestID is used to identify the request for support issues.
type ListVnicAttachments struct {
	ResourceContainer
	Attachments []VnicAttachment
}

func (l *ListVnicAttachments) GetList() interface{} {
	return &l.Attachments
}

// ListVnicAttachments returns a list of VnicAttachments with matching compartmentID
// and optionally instanceId, vnicId, and/or availabilityDomain. Optional parameters
// are assigned to the optional CoreOptions argument.  Page and Limit can also
// be supplied to support pagination.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VnicAttachment/ListVnicAttachments
func (c *Client) ListVnicAttachments(compartmentID string, opts ...Options) (res *ListVnicAttachments, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVnicAttachments,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListVnicAttachments{}
	e = resp.unmarshal(res)
	return
}
