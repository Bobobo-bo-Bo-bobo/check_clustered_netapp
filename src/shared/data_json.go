package shared

// VolumeList - list of defined volumes
type VolumeList struct {
	Records []VolumeRecord `json:"records"`
}

// VolumeRecord - single record from volume list
type VolumeRecord struct {
	UUID       string                  `json:"uuid"`
	Name       string                  `json:"name"`
	State      string                  `json:"state"`
	ErrorState *VolumeRecordErrorState `json:"error_state"`
	Link       RecordLink              `json:"_link"`
	// Note: some fields are considered "expensive" and must be explicitely queried.
	// See https://library.netapp.com/ecmdocs/ECMLP2856304/html/index.html#/docs-docs-Requesting-specific-fields
	Space *VolumeRecordSpace `json:"space"`
}

// RecordLink - _link dict from VolumeRecord
type RecordLink struct {
	Self RecordLinkSelf `json:"self"`
}

// RecordLinkSelf - self record from VolumeRecordLink
type RecordLinkSelf struct {
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

// VolumeRecordErrorState - Error states for volumes
type VolumeRecordErrorState struct {
	HasBadBlocks  bool `json:"has_bad_blocks"`
	IsInconsitent bool `json:"is_inconsistent"`
}

// ShelfList - list of shelves
type ShelfList struct {
	Records []ShelfRecord `json:"records"`
}

// ShelfRecord - single shelf record
type ShelfRecord struct {
	UUID           string     `json:"uid"`
	Name           string     `json:"name"`
	ID             string     `json:"id"`
	SerialNumber   string     `json:"serial_number"`
	Model          string     `json:"model"`
	ModuleType     string     `json:"module_type"`
	Internal       bool       `json:"internal"`
	State          string     `json:"state"`
	ConnectionType string     `json:"connection_type"`
	DiskCount      uint64     `json:"disk_count"`
	Link           RecordLink `json:"_link"`
}

// DiskList - list of disks
type DiskList struct {
	Records []DiskRecord `json:"records"`
}

// DiskRecord - record of a single disk
type DiskRecord struct {
	UUID         string     `json:"uid"`
	Name         string     `json:"name"`
	Vendor       string     `json:"vendor"`
	SerialNumber string     `json:"serial_number"`
	Model        string     `json:"model"`
	State        string     `json:"state"`
	Bay          string     `json:"bay"`
	Drawer       DiskDrawer `json:"drawer"`
}

// DiskDrawer - disk drawer
type DiskDrawer struct {
	ID   uint64 `json:"id"`
	Slot uint64 `json:"slot"`
}
