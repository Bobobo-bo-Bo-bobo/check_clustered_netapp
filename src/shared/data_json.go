package shared

// VolumeList - list of defined volumes
type VolumeList struct {
	Records []VolumeRecord `json:"records"`
}

// VolumeRecord - single record from volume list
type VolumeRecord struct {
	UUID string           `json:"uuid"`
	Name string           `json:"name"`
	Link VolumeRecordLink `json:"_link"`
	// Note: some fields are considered "expensive" and must be explicitely queried.
	// See https://library.netapp.com/ecmdocs/ECMLP2856304/html/index.html#/docs-docs-Requesting-specific-fields
	Space *VolumeRecordSpace `json:"space"`
}

// VolumeRecordLink - _link dict from VolumeRecord
type VolumeRecordLink struct {
	Self VolumeRecordLinkSelf `json:"self"`
}

// VolumeRecordLinkSelf - self record from VolumeRecordLink
type VolumeRecordLinkSelf struct {
	Href string `json:"href"`
}

// VolumeRecordSpace - volume utilisation record
type VolumeRecordSpace struct {
	Size            uint64                  `json:"size"`
	Available       uint64                  `json:"available"`
	Used            uint64                  `json:"used"`
	Footprint       uint64                  `json:"footprint"`
	OverProvisioned uint64                  `json:"over_provisioned"`
	Metadata        uint64                  `json:"metadata"`
	LogicalSpace    *VolumeRecordSpaceUsage `json:"logical_space"`
	Snapshot        *VolumeRecordSnapshot   `json:"snapshot"`
}

// VolumeRecordSpaceUsage - used disk space of a volume
type VolumeRecordSpaceUsage struct {
	Reporting   bool   `json:"reporting"`
	Enforcement bool   `json:"enforcement"`
	UsedByAFS   uint64 `json:"used_by_afs"`
}

// VolumeRecordSnapshot - snapshot usage of a volume
type VolumeRecordSnapshot struct {
	Used              uint64 `json:"used"`
	ReservePercent    uint64 `json:"reserve_percent"`
	AutodeleteEnabled bool   `json:"autodelete_enabled"`
}
