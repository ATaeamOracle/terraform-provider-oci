package baremetal

import "time"

type Database struct {
	ETaggedResource
	ocidRequirement
	DBHomeID     string    `json:"dbHomeId"`
	DBName       string    `json:"dbHomeId"`
	DBUniqueName string    `json:"dbUniqueName"`
	ID           string    `json:"id"`
	State        string    `json:"lifecycleState"`
	TimeCreated  time.Time `json:"timeCreated"`
}

type ListDatabases struct {
	ResourceContainer
	Databases []Database
}

func (l *ListDatabases) GetList() interface{} {
	return &l.Databases
}

// GetDatabase retrieves information about a Database
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/Database/GetDatabase
func (c *Client) GetDatabase(id string) (res *Database, e error) {
	details := &requestDetails{
		name: resourceDatabases,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	res = &Database{}
	e = resp.unmarshal(res)
	return
}

// ListDatabases returns a list of supported Oracle database versions. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/database/20160918/Database/ListDatabases
func (c *Client) ListDatabases(compartmentID, dbHomeID string, limit uint64, opts *PageListOptions) (resources *ListDatabases, e error) {
	required := struct {
		listOCIDRequirement
		DBHomeID string `json:"-" url:"dbHomeId"`
		Limit    uint64 `json:"-" url:"limit"`
	}{
		DBHomeID: dbHomeID,
		Limit:    limit,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDatabases,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.databaseApi.getRequest(details); e != nil {
		return
	}

	resources = &ListDatabases{}
	e = response.unmarshal(resources)
	return
}
