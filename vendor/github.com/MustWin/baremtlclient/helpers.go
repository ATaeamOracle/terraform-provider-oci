package baremtlsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CreateAPIKeyRequest struct {
	Key string `json:"key"`
}

type CreatePolicyRequest struct {
	CreateResourceRequest
	Statements []string `json:"statements"`
}

type CreateResourceRequest struct {
	CompartmentID string `json:"compartmentId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

type UpdateUIPasswordRequest struct {
	Password string `json:"password"`
}

type UpdateResourceRequest struct {
	Description string `json:"description"`
}

type UpdatePolicyRequest struct {
	UpdateResourceRequest
	Statements []string `json:"statements"`
}

func buildCoreURL(resource resourceName, query url.Values, ids ...interface{}) string {
	urlStr := fmt.Sprintf("%s/%s/%s", coreServiceAPI, coreServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildURL(urlStr string, query url.Values, ids ...interface{}) string {
	const seperator = "/"
	for _, id := range ids {
		var strVal string

		switch id := id.(type) {
		default:
			panic("Unsupported type")
		case bool:
			strVal = strconv.FormatBool(id)
		case uint64:
			strVal = strconv.FormatUint(id, 10)
		case string:
			strVal = id
		}

		if strVal != seperator {
			urlStr += seperator
		}

		urlStr += strVal
	}

	u, _ := url.Parse(urlStr)
	if query != nil {
		q := u.Query()
		for key, vals := range query {
			for _, val := range vals {
				q.Add(key, val)
			}
		}

		u.RawQuery += q.Encode()
	}

	return u.String()

}

func buildIdentityURL(resource resourceName, query url.Values, ids ...interface{}) string {
	urlStr := fmt.Sprintf("%s/%s/%s", identityServiceAPI, identityServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func (c *Client) createResource(resourceType resourceName, request CreateResourceRequest, headers http.Header) (resource *Resource, e error) {
	urlStr := buildIdentityURL(resourceType, nil)

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, urlStr, request, headers); e != nil {
		return
	}

	resource = &Resource{}
	e = json.Unmarshal(resp.body, resource)
	return

}

func (c *Client) getIdentity(resource resourceName, ids ...interface{}) (item *Resource, e error) {
	url := buildIdentityURL(resource, nil, ids...)
	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	item = &Resource{}
	e = decoder.Decode(item)
	return

}

func getUpdateHeaders(options ...Options) http.Header {
	var headers http.Header
	if len(options) > 0 {
		if options[0].IfMatch != "" {
			headers := &http.Header{}
			headers.Set(headerIfMatch, options[0].IfMatch)
		}
	}
	return headers
}

func (c *Client) listItems(resource resourceName, options ...ListOptions) (resp *ListResourceResponse, e error) {

	q := url.Values{}
	q.Set(queryCompartmentID, c.authInfo.tenancyOCID)
	if len(options) > 0 {
		q.Set(queryLimit, strconv.FormatUint(options[0].Limit, 10))
		q.Set(queryPage, options[0].Page)
	}

	resourceURL := buildIdentityURL(resource, q)

	var getResp *requestResponse
	if getResp, e = c.api.getRequest(resourceURL, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var items []Resource
	if e = decoder.Decode(&items); e != nil {
		return
	}

	resp = &ListResourceResponse{
		Page:  getResp.header.Get(headerOPCNextPage),
		Items: items,
	}

	return
}

func (c *Client) updateResource(resource resourceName, resourceID string, request interface{}, headers http.Header) (resp []byte, e error) {
	url := buildIdentityURL(resource, nil, resourceID)
	var r *requestResponse
	if r, e = c.api.request(http.MethodPut, url, request, headers); e != nil {
		return
	}
	resp = r.body
	return
}
