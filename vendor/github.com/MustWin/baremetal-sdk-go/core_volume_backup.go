package baremetal

import "net/http"

// VolumeBackup describe a point-in-time copy of a volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/
type VolumeBackup struct {
	ETaggedResource
	CompartmentID       string `json:"compartmentId"`
	DisplayName         string `json:"displayName"`
	ID                  string `json:"id"`
	SizeInMBs           uint64 `json:"sizeInMBs"`
	State               string `json:"lifecycleState"`
	TimeCreated         Time   `json:"timeCreated"`
	TimeRequestReceived Time   `json:"timeRequestReceived"`
	UniqueSizeInMBs     uint64 `json:"uniqueSizeInMBs"`
	VolumeID            string `json:"volumeId"`
}

// ListVolumeBackups contains a list of volume backups
//
type ListVolumeBackups struct {
	ResourceContainer
	VolumeBackups []VolumeBackup
}

func (l *ListVolumeBackups) GetList() interface{} {
	return &l.VolumeBackups
}

// CreateVolumeBackup Creates a new backup of the specified volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/CreateVolumeBackup
func (c *Client) CreateVolumeBackup(volumeID string, opts *CreateOptions) (vol *VolumeBackup, e error) {
	required := struct {
		VolumeID string `json:"volumeId" url:"-"`
	}{
		VolumeID: volumeID,
	}

	details := &requestDetails{
		name:     resourceVolumeBackups,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = response.unmarshal(vol)
	return
}

// GetVolumeBackup gets information for the specified volumeBackup
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/GetVolumeBackup
func (c *Client) GetVolumeBackup(id string) (vol *VolumeBackup, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceVolumeBackups,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = resp.unmarshal(vol)
	return
}

// UpdateVolumeBackup updates a volume's display name
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/UpdateVolumeBackup
func (c *Client) UpdateVolumeBackup(id string, opts *UpdateBackupOptions) (vol *VolumeBackup, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumeBackups,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = response.unmarshal(vol)
	return
}

// DeleteVolumeBackup deletes a volumeBackup
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/DeleteVolumeBackup
func (c *Client) DeleteVolumeBackup(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumeBackups,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListVolumeBackups returns a list of volumes for a particular compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/ListVolumeBackups
func (c *Client) ListVolumeBackups(compartmentID string, opts *ListBackupsOptions) (vols *ListVolumeBackups, e error) {
	details := &requestDetails{
		name:     resourceVolumeBackups,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vols = &ListVolumeBackups{}
	e = resp.unmarshal(vols)
	return
}
