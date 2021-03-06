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
	Bay          uint64     `json:"bay"`
	Drawer       DiskDrawer `json:"drawer"`
	Link         RecordLink `json:"_link"`
}

// DiskDrawer - disk drawer
type DiskDrawer struct {
	ID   uint64 `json:"id"`
	Slot uint64 `json:"slot"`
}

// AggregateList - list of aggregates
type AggregateList struct {
	Records []AggregateRecord `json:"records"`
}

// AggregateRecord - aggregate data
type AggregateRecord struct {
	UUID  string                `json:"uuid"`
	Name  string                `json:"name"`
	State string                `json:"state"`
	Space *AggregateRecordSpace `json:"space"`
	Link  RecordLink            `json:"_link"`
}

// AggregateRecordSpace - used disk space of an aggregate
type AggregateRecordSpace struct {
	BlockStorage                AggregateRecordSpaceBlockStorage `json:"block_storage"`
	Efficiency                  AggregateRecordSpaceEfficiency   `json:"efficiency"`
	EfficiencyWithoutSnapshhots AggregateRecordSpaceEfficiency   `json:"efficiency_without_snapshots"`
}

// AggregateRecordSpaceBlockStorage - block storage information of an aggregate
type AggregateRecordSpaceBlockStorage struct {
	Size                 uint64 `json:"size"`
	Available            uint64 `json:"available"`
	Used                 uint64 `json:"used"`
	FullThresholdPercent uint64 `json:"full_threshold_percent"`
}

// AggregateRecordSpaceEfficiency - efficiency statistics for an aggregate
type AggregateRecordSpaceEfficiency struct {
	Savings     uint64  `json:"savings"`
	Ratio       float64 `json:"ratio"`
	LogicalUsed uint64  `json:"logical_used"`
}

// SnapmirrorRelationshipList - list of snapmirror relationships
type SnapmirrorRelationshipList struct {
	Records []SnapmirrorRelationshipRecord `json:"records"`
}

// SnapmirrorRelationshipRecord - status of a single snapmirror relationship
type SnapmirrorRelationshipRecord struct {
	UUID            string                                  `json:"uuid"`
	Source          SnapmirrorRelationshipSource            `json:"source"`
	Healthy         bool                                    `json:"healthy"`
	UnhealthyReason []SnapmirrorRelationshipUnhealthyReason `json:"unhealthy_reason"`
	Link            RecordLink                              `json:"_link"`
}

// SnapmirrorRelationshipSource - source of a snapmirror relationship
type SnapmirrorRelationshipSource struct {
	Path    string                        `json:"path"`
	SVM     SnapmirrorRelationshipSVM     `json:"svm"`
	Cluster SnapmirrorRelationshipCluster `json:"cluster"`
}

// SnapmirrorRelationshipSVM - definition of storage virtual machine (SVM)
type SnapmirrorRelationshipSVM struct {
	Name string `json:"name"`
}

// SnapmirrorRelationshipCluster - cluster information
type SnapmirrorRelationshipCluster struct {
	Name string     `json:"name"`
	UUID string     `json:"uuid"`
	Link RecordLink `json:"_link"`
}

// SnapmirrorRelationshipUnhealthyReason - reason why a snapmirror relationship is unhealthy
type SnapmirrorRelationshipUnhealthyReason struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
