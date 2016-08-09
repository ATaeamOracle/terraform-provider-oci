package baremtlsdk

// CreateUser is used to create a user. userName MUST be unique. description
import (
	"encoding/json"
	"net/http"
)

// contains a comment about the user. The caller can supply 0 or 1 options. Options
// MAY contain an idempotency token.
// The caller specifies this token so that subsequent calls to create user will
// be idempotent. The token expires after 30 minutes.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createUser
func (c *Client) CreateUser(userName, userDescription string, options ...Options) (user *IdentityResource, e error) {
	createRequest := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          userName,
		Description:   userDescription,
	}
	var headers http.Header
	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}
	return c.createIdentityResource(resourceUsers, createRequest, headers)
}

// DeleteUser deletes a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteUser
func (c *Client) DeleteUser(userID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceUsers, nil, userID)

	return c.api.deleteRequest(url, headers)

}

// GetUser returns a user identified by userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetUser(userID string) (user *IdentityResource, e error) {
	user, e = c.getIdentityResource(resourceUsers, userID)
	return
}

// ListUsers returns an array of users for the current tenancy.  The requestor
// MAY supply paging options.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
func (c *Client) ListUsers(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listIdentityResources(resourceUsers, options...)
}

func (c *Client) UpdateUser(userID, userDescription string, opts ...Options) (user *IdentityResource, e error) {
	headers := getUpdateHeaders(opts...)
	request := UpdateIdentityResourceRequest{
		Description: userDescription,
	}

	var resp []byte
	if resp, e = c.updateIdentityResource(resourceUsers, userID, request, headers); e != nil {
		return
	}

	user = &IdentityResource{}
	e = json.Unmarshal(resp, user)
	return

}
