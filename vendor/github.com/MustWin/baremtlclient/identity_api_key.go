package baremtlsdk

// APIKey is returned for operations that create or modify user API keys.
import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type CreateAPIKeyRequest struct {
	Key string `json:"key"`
}

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#ApiKey
type APIKey struct {
	KeyID        string    `json:"keyId"`
	KeyValue     string    `json:"keyValue"`
	Fingerprint  string    `json:"fingerprint"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
}

// ListAPIKeyResponse contains a list of API keys
type ListAPIKeyResponse struct {
	OPCRequestID string
	Keys         []APIKey
}

// Deletes an API key belonging to a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteApiKey
func (c *Client) DeleteAPIKey(userID, fingerprint string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, fingerprint)
	return c.api.deleteRequest(url, headers)

}

// ListAPIKeys returns information about a user's API keys.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listApiKeys
func (c *Client) ListAPIKeys(userID string) (response *ListAPIKeyResponse, e error) {
	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, "/")
	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}
	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var keys []APIKey

	if e = decoder.Decode(&keys); e != nil {
		return
	}

	response = &ListAPIKeyResponse{
		Keys:         keys,
		OPCRequestID: getResp.header.Get(headerOPCRequestID),
	}

	return

}

// UploadAPIKey - add an API signing key for user. The key must be an RSA public
// key in pem format.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#uploadApiKey
func (c *Client) UploadAPIKey(userID, key string, opts ...Options) (apiKey *APIKey, e error) {
	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, "/")
	request := CreateAPIKeyRequest{
		Key: key,
	}

	var headers http.Header

	if len(opts) > 0 {
		if opts[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, opts[0].OPCIdempotencyToken)
		}
	}

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, url, request, headers); e != nil {
		return
	}

	apiKey = &APIKey{}
	e = json.Unmarshal(resp.body, apiKey)
	return

}
