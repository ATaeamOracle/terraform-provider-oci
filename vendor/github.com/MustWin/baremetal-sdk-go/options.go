package baremetal

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// To get the body, optional and required are marshalled and merged.
// To get the query string, optional and required are merged.
// To get the header,
//   optional is type asserted as a HeaderGenerator
//   required is type asserted as a HeaderGenerator
//   both results are combined
// Required options get built inline within funcs based on args.
// A single options struct gets passed in as optional.
// Both need to explicitly handle json and url tags, excluding appropriately.

type HeaderGenerator interface {
	Header() http.Header
}

type IfMatchOptions struct {
	IfMatch string `json:"-" url:"-"`
}

func (opt IfMatchOptions) Header() http.Header {
	header := http.Header{}
	if opt.IfMatch != "" {
		header.Set(headerIfMatch, opt.IfMatch)
	}
	return header
}

type RetryTokenOptions struct {
	RetryToken string `json:"-" url:"-"`
}

func (opt RetryTokenOptions) Header() http.Header {
	header := http.Header{}
	if opt.RetryToken != "" {
		header.Set(headerRetryToken, opt.RetryToken)
	}
	return header
}

// Both header options cannot be embedded into the same struct separately.
// Duplicating here keeps the API clean, allowing this option to be embedded in
// non-header structs with easy-to-set struct values.
type HeaderOptions struct {
	IfMatch    string `json:"-" url:"-"`
	RetryToken string `json:"-" url:"-"`
}

func (opt HeaderOptions) Header() http.Header {
	header := http.Header{}
	if opt.IfMatch != "" {
		header.Set(headerIfMatch, opt.IfMatch)
	}
	if opt.RetryToken != "" {
		header.Set(headerRetryToken, opt.RetryToken)
	}
	return header
}

type ocidRequirement struct {
	CompartmentID string `json:"compartmentId" url:"-"`
}

type identityCreationRequirement struct {
	CompartmentID string `json:"compartmentId" url:"-"`
	Description   string `json:"description" url:"-"`
	Name          string `json:"name" url:"-"`
}

type DisplayNameOptions struct {
	DisplayName string `json:"displayName,omitempty" url:"-"`
}

type VersionDateOptions struct {
	VersionDate string `json:"versionDate,omitempty" url:"-"`
}

type CreateOptions struct {
	RetryTokenOptions
	DisplayNameOptions
}

type UpdateOptions struct {
	HeaderOptions
	DisplayNameOptions
}

type UpdateBackupOptions struct {
	IfMatchOptions
	DisplayNameOptions
}

type UpdateIdentityOptions struct {
	IfMatchOptions
	Description string `json:"description,omitempty" url:"-"`
}

type UpdatePolicyOptions struct {
	UpdateIdentityOptions
	VersionDateOptions
	Statements []string `json:"statements,omitempty" url:"-"`
}

type UpdateDHCPDNSOptions struct {
	CreateOptions
	Options []DHCPDNSOption `json:"options,omitempty" url:"-"`
}

type LaunchInstanceOptions struct {
	CreateOptions
	Metadata map[string]string `json:"metadata,omitempty" url:"-"`
}

type UpdateGatewayOptions struct {
	IfMatchOptions
	DisplayNameOptions
	IsEnabled bool `json:"isEnabled,omitempty" url:"-"`
}

type UpdateRouteTableOptions struct {
	CreateOptions
	RouteRules []RouteRule `json:"routeRules,omitempty" url:"-"`
}

type UpdateSecurityListOptions struct {
	CreateOptions
	EgressRules  []EgressSecurityRule  `json:"egressSecurityRules,omitempty" url:"-"`
	IngressRules []IngressSecurityRule `json:"ingressSecurityRules,omitempty" url:"-"`
}

type CreateSubnetOptions struct {
	CreateOptions
	DHCPOptionsID   string   `json:"dhcpOptionsId,omitempty" url:"-"`
	RouteTableID    string   `json:"routeTableId,omitempty" url:"-"`
	SecurityListIDs []string `json:"securityListIds,omitempty" url:"-"`
}

type CreateVolumeOptions struct {
	CreateOptions
	VolumeBackupID string `json:"volumeBackupId,omitempty" url:"-"`
}

type CreatePolicyOptions struct {
	RetryTokenOptions
	VersionDateOptions
}

// ----- Options for listing resources ---

type listOCIDRequirement struct {
	CompartmentID string `json:"-" url:"compartmentId"`
}

type ListOptions struct {
	Limit uint64 `json:"-" url:"limit,omitempty"`
	Page  string `json:"-" url:"page,omitempty"`
}

type DisplayNameListOptions struct {
	DisplayName string `json:"-" url:"displayName,omitempty"`
}

type AvailabilityDomainListOptions struct {
	AvailabilityDomain string `json:"-" url:"availabilityDomain,omitempty"`
}

type DrgIDListOptions struct {
	DrgID string `json:"-" url:"drgId,omitempty"`
}

type InstanceIDListOptions struct {
	InstanceID string `json:"-" url:"instanceId,omitempty"`
}

type ListInstancesOptions struct {
	AvailabilityDomainListOptions
	DisplayNameListOptions
	ListOptions
}

type ListConsoleHistoriesOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
}

type ListDrgAttachmentsOptions struct {
	DrgIDListOptions
	ListOptions
	VcnID string `json:"-" url:"vcnId,omitempty"`
}

type ListImagesOptions struct {
	DisplayNameListOptions
	ListOptions
	OperatingSystem        string `json:"-" url:"operatingSystem,omitempty"`
	OperatingSystemVersion string `json:"-" url:"operatingSystemVersion,omitempty"`
}

type ListIPSecConnsOptions struct {
	DrgIDListOptions
	ListOptions
	CpeID string `json:"-" url:"cpeId,omitempty"`
}

type ListShapesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
	ImageID string `json:"-" url:"imageId,omitempty"`
}

type ListVnicAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VnicID string `json:"-" url:"vnicId,omitempty"`
}

type ListVolumesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
}

type ListVolumeAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VolumeID string `json:"-" url:"volumeId,omitempty"`
}

type ListBackupsOptions struct {
	ListOptions
	VolumeID string `json:"-" url:"volumeId,omitempty"`
}

type ListMembershipsOptions struct {
	ListOptions
	GroupID string `json:"-" url:"groupId,omitempty"`
	UserID  string `json:"-" url:"userId,omitempty"`
}

// -------- Misc options -----

type ConsoleHistoryDataOptions struct {
	Length uint64 `json:"-" url:"length,omitempty"`
	Offset uint64 `json:"-" url:"offset,omitempty"`
}

// -------------

type urlParts []interface{}

type requestOptions interface {
	url(urlBuilderFn) (val string, e error)
	header() http.Header
	getBody() ([]byte, error)
}

type requestDetails struct {
	name     resourceName
	ids      urlParts
	required interface{}
	optional interface{}
}

func (r *requestDetails) query() (vals url.Values, e error) {
	var rVals url.Values
	if rVals, e = query.Values(r.required); e != nil {
		return
	}
	if vals, e = query.Values(r.optional); e != nil {
		return
	}
	for k, v := range rVals {
		vals[k] = v
	}
	return
}

func (r *requestDetails) url(urlFn urlBuilderFn) (val string, e error) {
	var q url.Values
	if q, e = r.query(); e != nil {
		return
	}
	val = urlFn(r.name, q, r.ids...)
	return
}

func (r *requestDetails) header() http.Header {
	var rHeader, oHeader http.Header
	if rhd, ok := r.required.(HeaderGenerator); ok == true {
		rHeader = rhd.Header()
	} else {
		rHeader = http.Header{}
	}
	if ohd, ok := r.optional.(HeaderGenerator); ok == true {
		oHeader = ohd.Header()
	} else {
		oHeader = http.Header{}
	}
	for k, v := range rHeader {
		oHeader[k] = v
	}
	return oHeader
}

func (r *requestDetails) getBody() (marshaled []byte, e error) {
	if marshaled, e = json.Marshal(r.required); e != nil {
		return
	}
	var oBody []byte
	if oBody, e = json.Marshal(r.optional); e != nil {
		return
	}
	marshaled = append(marshaled, oBody...)
	return
}
