// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type DBVersion struct {
	Version string `json:"version"`
}

type ListDBVersions struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DBVersions []DBVersion
}

func (l *ListDBVersions) GetList() interface{} {
	return &l.DBVersions
}

// ListVersions returns a list of supported Oracle database versions. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/DbVersion/ListDbVersions
func (c *Client) ListDBVersions(compartmentID string, limit uint64, opts *PageListOptions) (resources *ListDBVersions, e error) {
	required := struct {
		listOCIDRequirement
		Limit uint64 `header:"-" json:"-" url:"limit"`
	}{
		Limit: limit,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDBVersions,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDBVersions{}
	e = resp.unmarshal(resources)
	return
}
