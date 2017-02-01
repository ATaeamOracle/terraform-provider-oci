package baremetal

// UserGroupMembership returned by GetUserGroupMembership and related methods.
import (
	"net/http"
	"time"
)

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UserGroupMembership
type UserGroupMembership struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID  string    `json:"compartmentId"`
	GroupID        string    `json:"groupId"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
	UserID         string    `json:"userId"`
}

type ListUserGroupMemberships struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Memberships []UserGroupMembership
}

func (l *ListUserGroupMemberships) GetList() interface{} {
	return &l.Memberships
}

// AddUserToGroup adds a user to a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#addUserToGroup
func (c *Client) AddUserToGroup(userID, groupID string, opts *RetryTokenOptions) (res *UserGroupMembership, e error) {
	required := struct {
		GroupID string `header:"-" json:"groupId" url:"-"`
		UserID  string `header:"-" json:"userId" url:"-"`
	}{
		GroupID: groupID,
		UserID:  userID,
	}

	details := &requestDetails{
		name:     resourceUserGroupMemberships,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.identityApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &UserGroupMembership{}
	e = resp.unmarshal(res)
	return
}

// GetUserGroupMembership returns a UserGroupMembership identified by id.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getUserGroupMembership
func (c *Client) GetUserGroupMembership(id string) (res *UserGroupMembership, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceUserGroupMemberships,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &UserGroupMembership{}
	e = resp.unmarshal(res)
	return
}

// DeleteUserGroupMembership removes a user from a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#removeUserFromGroup
func (c *Client) DeleteUserGroupMembership(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceUserGroupMemberships,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

func (c *Client) ListUserGroupMemberships(opts *ListMembershipsOptions) (resources *ListUserGroupMemberships, e error) {
	details := &requestDetails{
		name:     resourceUserGroupMemberships,
		optional: opts,
		required: ocidRequirement{c.authInfo.tenancyOCID},
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListUserGroupMemberships{}
	e = resp.unmarshal(resources)
	return
}
