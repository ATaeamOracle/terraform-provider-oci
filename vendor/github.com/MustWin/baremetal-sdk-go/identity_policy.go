package baremetal

// CreatePolicy creates a new policy.
import "net/http"

type CreatePolicyRequest struct {
	CreateIdentityResourceRequest
	Statements []string `json:"statements"`
}

type UpdatePolicyRequest struct {
	UpdateIdentityResourceRequest
	Statements []string `json:"statements"`
}

// Policy returned by GetPolicy and other policy related methods.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#Policy
type Policy struct {
	IdentityResource
	Statements []string `json:"statements"`
}

// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Policy/CreatePolicy
func (c *Client) CreatePolicy(policyName, policyDescription string, statements []string, options ...Options) (resource *Policy, e error) {
	var body CreatePolicyRequest
	body.CompartmentID = c.authInfo.tenancyOCID
	body.Name = policyName
	body.Description = policyDescription
	body.Statements = statements

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourcePolicies,
		options: options,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	resource = &Policy{}
	e = response.unmarshal(resource)
	return
}

// DeletePolicy removes a policy identified by policyID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Policy/DeletePolicy
func (c *Client) DeletePolicy(policyID string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourcePolicies,
		options: opts,
		ids:     urlParts{policyID},
	}
	return c.identityApi.deleteRequest(reqOpts)
}

// GetPolicy returns a policy identified by a policyID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Policy/GetPolicy
func (c *Client) GetPolicy(policyID string) (resource *Policy, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourcePolicies,
		ids:  urlParts{policyID},
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	resource = &Policy{}
	e = response.unmarshal(resource)
	return
}

// UpdatePolicy can be called to modify the description and statements of an existing policy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Policy/UpdatePolicy
func (c *Client) UpdatePolicy(policyID, policyDescription string, policyStatements []string, opts ...Options) (resource *Policy, e error) {
	var body UpdatePolicyRequest
	body.Description = policyDescription
	body.Statements = policyStatements

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourcePolicies,
		options: opts,
		ids:     urlParts{policyID},
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	resource = &Policy{}
	e = response.unmarshal(resource)
	return
}
