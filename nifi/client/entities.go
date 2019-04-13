package client

type RevisionDTO struct {
	// A client identifier used to make a request. By including a client identifier, the API can allow multiple
	// requests without needing the current revision. Due to the asynchronous nature of requests/responses this
	// was implemented to allow the client to make numerous requests without having to wait for the previous
	// response to come back
	ClientID string `json:"clientId"`
	// NiFi employs an optimistic locking strategy where the client must include a revision in their request
	// when performing an update. In a response to a mutable flow request, this field represents the updated
	// base version.
	Version int64 `json:"version"`
	// The user that last modified the flow. This property is read only.
	LastModifier string `json:"lastModifier"`
}

type PositionDTO struct {
	// The x coordinate.
	X float64 `json:"x"`
	// The y coordinate.
	Y float64 `json:"y"`
}

type PermissionsDTO struct {
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
	// Indicates whether the user can write a given resource. This property is read only.
	CanWrite bool `json:"canWrite"`
}

type BulletinDTO struct {
	// The id of the bulletin.
	ID int64 `json:"id"`
	// If clustered, the address of the node from which the bulletin originated.
	NodeAddress string `json:"nodeAddress"`
	// The category of this bulletin.
	Category string `json:"category"`
	// The group id of the source component.
	GroupID string `json:"groupId"`
	// The id of the source component.
	SourceID string `json:"sourceId"`
	// The name of the source component.
	SourceName string `json:"sourceName"`
	// The level of the bulletin.
	Level string `json:"level"`
	// The bulletin message.
	Message string `json:"message"`
	// When this bulletin was generated.
	Timestamp string `json:"timestamp"`
}

type BulletinEntity struct {
	// The id of the bulletin.
	ID int64 `json:"id"`
	// The group id of the source component.
	GroupID string `json:"groupId"`
	// The id of the source component.
	SourceID string `json:"sourceId"`
	// When this bulletin was generated.
	Timestamp string `json:"timestamp"`
	// If clustered, the address of the node from which the bulletin originated.
	NodeAddress string `json:"nodeAddress"`
	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead"`
	// The bulletin.
	Bulletin BulletinDTO `json:"bulletin"`
}

type ProcessGroupDTO struct {
	// The id of the component.
	ID string `json:"id"`
	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId"`
	// The position of this component in the UI if applicable.
	Position PositionDTO `json:"position"`
	// The name of the process group.
	Name string `json:"name"`
	// The comments for the process group.
	Comments string `json:"comments"`
	// The number of running components in this process group.
	RunningCount int `json:"runningCount"`
	// The number of stopped components in the process group.
	StoppedCount int `json:"stoppedCount"`
	// The number of invalid components in the process group.
	InvalidCount int `json:"invalidCount"`
	// The number of disabled components in the process group.
	DisabledCount int `json:"disabledCount"`
	// The number of active remote ports in the process group.
	ActiveRemotePortCount int `json:"activeRemotePortCount"`
	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount int `json:"inactiveRemotePortCount"`
	// The number of input ports in the process group.
	InputPortCount int `json:"inputPortCount"`
	// The number of output ports in the process group.
	OutputPortCount int `json:"outputPortCount"`
}

type ConnectionStatusSnapshotDTO struct {
	// The id of the connection.
	ID string `json:"id"`
	// The id of the process group the connection belongs to.
	GroupID string `json:"groupId"`
	// The name of the connection.
	Name string `json:"name"`
	// The id of the source of the connection.
	SourceID string `json:"sourceId"`
	// The name of the source of the connection.
	SourceName string `json:"sourceName"`
	// The id of the destination of the connection.
	DestinationID string `json:"destinationId"`
	// The name of the destination of the connection.
	DestinationName string `json:"destinationName"`
	// The number of FlowFiles that have come into the connection in the last 5 minutes.
	FlowFilesIn int `json:"flowFilesIn"`
	// The size of the FlowFiles that have come into the connection in the last 5 minutes.
	BytesIn int64 `json:"bytesIn"`
	// The input count/size for the connection in the last 5 minutes, pretty printed.
	Input string `json:"input"`
	// The number of FlowFiles that have left the connection in the last 5 minutes.
	FlowFilesOut int `json:"flowFilesOut"`
	// The number of bytes that have left the connection in the last 5 minutes.
	BytesOut int64 `json:"bytesOut"`
	// The output count/sie for the connection in the last 5 minutes, pretty printed.
	Output string `json:"output"`
	// The number of FlowFiles that are currently queued in the connection.
	FlowFilesQueued int `json:"flowFilesQueued"`
	// The size of the FlowFiles that are currently queued in the connection.
	BytesQueued int64 `json:"bytesQueued"`
	// The total count and size of queued flowfiles formatted.
	Queued string `json:"queued"`
	// The total size of flowfiles that are queued formatted.
	QueuedSize string `json:"queuedSize"`
	// The number of flowfiles that are queued, pretty printed.
	QueuedCount string `json:"queuedCount"`
	// Connection percent use regarding queued flow files count and backpressure threshold if configured.
	PercentUseCount int `json:"percentUseCount"`
	// Connection percent use regarding queued flow files size and backpressure threshold if configured.
	PercentUseBytes int `json:"percentUseBytes"`
}

type ConnectionStatusSnapshotEntity struct {
	// The id of the connection.
	ID string `json:"id"`
	// The connection status snapshot.
	ConnectionStatusSnapshot ConnectionStatusSnapshotDTO `json:"connectionStatusSnapshot"`
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
}

type ProcessorStatusSnapshotDTO struct {
	// 	The id of the processor.
	ID string `json:"id"`
	// The id of the parent process group to which the processor belongs.
	GroupID string `json:"groupId"`
	// The name of the prcessor.
	Name string `json:"name"`
	// The type of the processor.
	Type string `json:"type"`
	// The state of the processor. Allowable values: RUNNING, STOPPED, DISABLED, INVALID
	RunStatus string `json:"runStatus"`
	// The number of bytes read by this Processor in the last 5 mintues
	BytesRead int64 `json:"bytesRead"`
	// The number of bytes written by this Processor in the last 5 minutes
	BytesWritten int64 `json:"bytesWritten"`
	// The number of bytes read in the last 5 minutes.
	Read string `json:"read"`
	// The number of bytes written in the last 5 minutes.
	Written string `json:"written"`
	// The number of FlowFiles that have been accepted in the last 5 minutes
	FlowFilesIn int `json:"flowFilesIn"`
	// The size of the FlowFiles that have been accepted in the last 5 minutes
	BytesIn int64 `json:"bytesIn"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input string `json:"input"`
	// The number of FlowFiles transferred to a Connection in the last 5 minutes
	FlowFilesOut int `json:"flowFilesOut"`
	// The size of the FlowFiles transferred to a Connection in the last 5 minutes
	BytesOut int64 `json:"bytesOut"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output string `json:"output"`
	// The number of times this Processor has run in the last 5 minutes
	TaskCount int `json:"taskCount"`
	// The number of nanoseconds that this Processor has spent running in the last 5 minutes
	TasksDurationNanos int64 `json:"tasksDurationNanos"`
	// The total number of task this connectable has completed over the last 5 minutes.
	Tasks string `json:"tasks"`
	// The total duration of all tasks for this connectable over the last 5 minutes.
	TasksDuration string `json:"tasksDuration"`
	// The number of threads currently executing in the processor.
	ActiveThreadCount int `json:"activeThreadCount"`
}

type ProcessorStatusSnapshotEntity struct {
	// The id of the processor.
	ID string `json:"id"`
	// The processor status snapshot.
	ProcessorStatusSnapshot ProcessorStatusSnapshotDTO `json:"processorStatusSnapshot"`
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
}

type ProcessGroupStatusSnapshotEntity struct {
	// The id of the process group.
	ID string `json:"id"`
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
}

type RemoteProcessGroupStatusSnapshotDTO struct {
	// The id of the remote process group.
	ID string `json:"id"`
	// The id of the parent process group the remote process group resides in.
	GroupID string `json:"groupId"`
	// The name of the remote process group.
	Name string `json:"name"`
	// The URI of the target system.
	TargetURI string `json:"targetUri"`
	// The transmission status of the remote process group.
	TransmissionStatus string `json:"transmissionStatus"`
	// The number of active threads for the remote process group.
	ActiveThreadCount int `json:"activeThreadCount"`
	// The number of FlowFiles sent to the remote process group in the last 5 minutes.
	FlowFilesSent int `json:"flowFilesSent"`
	// The size of the FlowFiles sent to the remote process group in the last 5 minutes.
	BytesSent int64 `json:"bytesSent"`
	// The count/size of the flowfiles sent to the remote process group in the last 5 minutes.
	Sent string `json:"sent"`
	// The number of FlowFiles received from the remote process group in the last 5 minutes.
	FlowFilesReceived int `json:"flowFilesReceived"`
	// The size of the FlowFiles received from the remote process group in the last 5 minutes.
	BytesReceived int64 `json:"bytesReceived"`
	// The count/size of the flowfiles received from the remote process group in the last 5 minutes.
	Received string `json:"received"`
}

type RemoteProcessGroupStatusSnapshotEntity struct {
	// The id of the remote process group.
	ID string `json:"id"`
	// The remote process group status snapshot.
	RemoteProcessGroupStatusSnapshot RemoteProcessGroupStatusSnapshotDTO `json:"remoteProcessGroupStatusSnapshot"`
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
}

type PortStatusSnapshotDTO struct {
	// The id of the port.
	ID string `json:"id"`
	// The id of the parent process group of the port.
	GroupID string `json:"groupId"`
	// The name of the port.
	Name string `json:"name"`
	// The active thread count for the port.
	ActiveThreadCount int `json:"activeThreadCount"`
	// The number of FlowFiles that have been accepted in the last 5 minutes.
	FlowFilesIn int `json:"flowFilesIn"`
	// The size of hte FlowFiles that have been accepted in the last 5 minutes.
	BytesIn int64 `json:"bytesIn"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input string `json:"input"`
	// The number of FlowFiles that have been processed in the last 5 minutes.
	FlowFilesOut int `json:"flowFilesOut"`
	// The number of bytes that have been processed in the last 5 minutes.
	BytesOut int64 `json:"bytesOut"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output string `json:"output"`
	// Whether the port has incoming or outgoing connections to a remote NiFi.
	Transmitting bool `json:"transmitting"`
	// The run status of the port.
	RunStatus string `json:"runStatus"`
}

type PortStatusSnapshotEntity struct {
	// The id of the port .
	ID string `json:"id"`
	// The port status snapshot.
	PortStatusSnapshot PortStatusSnapshotDTO `json:"portStatusSnapshot"`
	// Indicates whether the user can read a given resource. This property is read only.
	CanRead bool `json:"canRead"`
}

type ProcessGroupStatusSnapshotDTO struct {
	// The id of the process group.
	ID string `json:"id"`
	// The name of this process group.
	Name string `json:"name"`
	// The status of all conenctions in the process group.
	ConnectionStatusSnapshots []ConnectionStatusSnapshotEntity `json:"connectionStatusSnapshots"`
	// The status of all processors in the process group.
	ProcessorStatusSnapshots []ProcessorStatusSnapshotEntity `json:"processorStatusSnapshots"`
	// The status of all process groups in the process group.
	ProcessGroupStatusSnapshots []ProcessGroupStatusSnapshotEntity `json:"processGroupStatusSnapshots"`
	// The status of all remote process groups in the process group.
	RemoteProcessGroupStatusSnapshots []RemoteProcessGroupStatusSnapshotEntity `json:"remoteProcessGroupStatusSnapshots"`
	// The status of all input ports in the process group.
	InputPortStatusSnapshots []PortStatusSnapshotEntity `json:"inputPortStatusSnapshots"`
	// The status of all output ports in the process group.
	OutputPortStatusSnapshots []PortStatusSnapshotEntity `json:"outputPortStatusSnapshots"`
	// The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes
	FlowFilesIn int `json:"flowFilesIn"`
	// The number of bytes that have come into this ProcessGroup in the last 5 minutes
	BytesIn int64 `json:"bytesIn"`
	// The input count/size for the process group in the last 5 minutes (pretty printed).
	Input string `json:"input"`
	// The number of FlowFiles that are queued up in this ProcessGroup right now
	FlowFilesQueued int `json:"flowFilesQueued"`
	// The number of bytes that are queued up in this ProcessGroup right now
	BytesQueued int64 `json:"bytesQueued"`
	// The count/size that is queued in the the process group.
	Queued string `json:"queued"`
	// The count that is queued for the process group.
	QueuedCount string `json:"queuedCount"`
	// The size that is queued for the process group.
	QueuedSize string `json:"queuedSize"`
	// The number of bytes read by components in this ProcessGroup in the last 5 minutes
	BytesRead int64 `json:"bytesRead"`
	// The number of bytes read in the last 5 minutes.
	Read string `json:"read"`
	// The number of bytes written by components in this ProcessGroup in the last 5 minutes
	BytesWritten int64 `json:"bytesWritten"`
	// The number of bytes written in the last 5 minutes.
	Written string `json:"written"`
	// The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes
	FlowFilesOut int `json:"flowFilesOut"`
	// The number of bytes transferred out of this ProcessGroup in the last 5 minutes
	BytesOut int64 `json:"bytesOut"`
	// The output count/size for the process group in the last 5 minutes.
	Output string `json:"output"`
	// The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes
	FlowFilesTransferred int `json:"flowFilesTransferred"`
	// The number of bytes transferred in this ProcessGroup in the last 5 minutes
	BytesTransferred int64 `json:"bytesTransferred"`
	// The count/size transferred to/from queues in the process group in the last 5 minutes.
	Transferred string `json:"transferred"`
	// The number of bytes received from external sources by components within this ProcessGroup in
	// the last 5 minutes
	BytesReceived int64 `json:"bytesReceived"`
	// The number of FlowFiles received from external sources by components within this ProcessGroup in
	// the last 5 minutes
	FlowFilesReceived int `json:"flowFilesReceived"`
	// The count/size sent to the process group in the last 5 minutes.
	Received string `json:"received"`
	// The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes
	BytesSent int64 `json:"bytesSent"`
	// The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes
	FlowFilesSent int `json:"flowFilesSent"`
	// The count/size sent from this process group in the last 5 minutes.
	Sent string `json:"sent"`
	// The active thread count for this process group.
	ActiveThreadCount int `json:"activeThreadCount"`
}

type NodeProcessGroupStatusSnapshotDTO struct {
	// The unique ID that identifies the node
	NodeID string `json:"nodeId"`
	// The API address of the node
	Address string `json:"address"`
	// The API port used to communicate with the node
	ApiPort int `json:"apiPort"`
	// The process group status snapshot from the node.
	StatusSnapshot ProcessGroupStatusSnapshotDTO `json:"statusSnapshot"`
}

type ProcessGroupStatusDTO struct {
	// The ID of the Process Group
	ID string `json:"id"`
	// The name of the Process Group
	Name string `json:"name"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed"`
	// The aggregate status of all nodes in the cluster
	AggregateSnapshot *ProcessGroupStatusSnapshotDTO `json:"aggregateSnapshot"`
	// The status reported by each node in the cluster. If the NiFi instance is a standalone instance,
	// rather than a clustered instance, this value may be null.
	NodeSnapshots []NodeProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`
}

type ProcessGroupEntity struct {
	// The revision for this request/response. The revision is required for any mutable flow requests and
	// is included in all responses.
	Revision RevisionDTO `json:"revision"`
	// The id of the component.
	ID string `json:"id"`
	// The URI for futures requests to the component.
	URI string `json:"uri"`
	// The position of this component in the UI if applicable.
	Position PositionDTO `json:"position"`
	// The permissions for this component.
	Permissions PermissionsDTO `json:"permissions"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins"`
	// The process group.
	Component ProcessGroupDTO `json:"component"`
	// The status of the process group.
	Status ProcessGroupStatusDTO `json:"status"`
	// The number of running components in this process group.
	RunningCount int `json:"runningCount"`
	// The number of stopped components in the process group.
	StoppedCount int `json:"stoppedCount"`
	// The number of invalid components in the process group.
	InvalidCount int `json:"invalidCount"`
	// The number of disabled components in the process group.
	DisabledCount int `json:"disabledCount"`
	// The number of active remote ports in the process group.
	ActiveRemotePortCount int `json:"activeRemotePortCount"`
	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount int `json:"inactiveRemotePortCount"`
	// The number of input ports in the process group.
	InputPortCount int `json:"inputPortCount"`
	// The number of output ports in the process group.
	OutputPortCount int `json:"outputPortCount"`
}

type ProcessGroupsEntity struct {
	ProcessGroups []ProcessGroupEntity `json:"processGroups"`
}

// ConnectionEndPointDTO represents a connection's source and/or destination
type ConnectionEndPointDTO struct {
	//The ID
	ID string `json:"id"`
	// The type
	Type string `json:"type"`
	// The group ID
	GroupID string `json:"groupId"`
	// The name
	Name string `json:"name"`
	//Is this connection running?
	Running bool `json:"running"`
	// User comments on this connection
	Comments string `json:"comments"`
}

// ConnectionDTO is the connection component
type ConnectionDTO struct {
	// The id of the component.
	ID string `json:"id"`
	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId"`

	// The source of the connection
	Source ConnectionEndPointDTO `json:"source"`
	// The destination of the connection
	Destination ConnectionEndPointDTO `json:"destination"`
	// The Name
	Name string `json:"name"`
	// The label Index
	LabelIndex int `json:"labelIndex"`
	// The zindex of this connection
	ZIndex int `json:"zIndex"`
	// Valid relationships for this connection
	SelectedRelationships []string `json:"selectedRelationships"`
	// Potential relationships for this connection
	AvailableRelationships []string `json:"availableRelationships"`
	// Threshold for back pressure objects
	BackPressureObjectThreshold int `json:"backPressureObjectThreshold"`
	// Threshold for backpressure size
	BackPressureDataSizeThreshold string `json:"backPressureDataSizeThreshold"`
	// When to the flows
	FlowFileExpiration string `json:"flowFileExpiration"`
	// The bends in the connection layout
	Bends []PositionDTO `json:"bends"`
	// The comparators used to prioritize the queue
	Prioritizers []string `json:"prioritizers"`
}

// ConnectionStatusDTO represents the status of a connection
type ConnectionStatusDTO struct {
	// The ID of the Connection
	ID string `json:"id"`
	// The GroupID of the Connection
	GroupID string `json:"groupId"`
	// The name of the Connection
	Name string `json:"name"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed"`
	SourceID           string `json:"sourceId"`
	SourceName         string `json:"sourceName"`
	DestinationID      string `json:"destinationId"`
	DestinationName    string `json:"destinationName"`

	// The aggregate status of all nodes in the cluster
	AggregateSnapshot *ConnectionStatusSnapshotDTO `json:"aggregateSnapshot"`
	// The status reported by each node in the cluster. If the NiFi instance is a standalone instance,
	// rather than a clustered instance, this value may be null.
	NodeSnapshots []NodeProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`
}

// ConnectionEntity represents a NiFi connection
type ConnectionEntity struct {
	// The revision for this request/response. The revision is required for any mutable flow requests and
	// is included in all responses.
	Revision RevisionDTO `json:"revision"`
	// The id of the component.
	ID string `json:"id"`
	// The URI for futures requests to the component.
	URI string `json:"uri"`
	// The permissions for this component.
	Permissions PermissionsDTO `json:"permissions"`
	// The Connection.
	Component ConnectionDTO `json:"component"`
	// The status of the process group.
	Status ConnectionStatusDTO `json:"status"`
	// The bends in the connection layout
	Bends []PositionDTO `json:"bends"`

	// The label Index
	LabelIndex int `json:"labelIndex"`
	// The zindex of this connection
	ZIndex int `json:"zIndex"`
	// Connection source id
	SourceID string `json:"sourceId"`
	// Connection source group id
	SourceGroupID string `json:"sourceGroupId"`
	// Connection source type
	SourceType string `json:"sourceType"`
	// Connection destination id
	DestinationID string `json:"destinationId"`
	// Connection destination group id
	DestinationGroupID string `json:"destinationGroupId"`
	// Connection destination type
	DestinationType string `json:"destinationType"`
}

type ConnectionsEntity struct {
	Connections []ConnectionEntity `json:"connections"`
}

type CounterDTO struct {
	// The id of the counter.
	ID string `json:"id"`
	// The context of the counter.
	Context string `json:"context"`
	// The name of the counter.
	Name string `json:"name"`
	// The value count.
	ValueCount int64 `json:"valueCount"`
	// The value of the counter.
	Value string `json:"value"`
}

type CountersSnapshotDTO struct {
	// The timestamp when the report was generated.
	Generated string `json:"generated"`
	// All counters in the NiFi.
	Counters []CounterDTO `json:"counters"`
}

type NodeCountersSnapshotDTO struct {
	// The unique ID that identifies the node
	NodeID string `json:"nodeId"`
	// The API address of the node
	Address string `json:"address"`
	// The API port used to communicate with the node
	ApiPort int `json:"apiPort"`
	// The counters from the node.
	Snapshot CountersSnapshotDTO `json:"snapshot"`
}

type CountersDTO struct {
	// A Counters snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi instance
	// is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	AggregateSnapshot *CountersSnapshotDTO `json:"aggregateSnapshot"`
	// A Counters snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather
	// than a cluster, this may be null.
	NodeSnapshots []NodeCountersSnapshotDTO `json:"nodeSnapshots"`
}

type CountersEntity struct {
	Counters CountersDTO `json:"counters"`
}

type StorageUsageDTO struct {
	// The identifier of this storage location. The identifier will correspond to the identifier keyed in the
	// storage configuration.
	Identifier string `json:"identifier"`
	// Amount of free space.
	FreeSpace string `json:"freeSpace"`
	// Amount of total space.
	TotalSpace string `json:"totalSpace"`
	// Amount of used space.
	UsedSpace string `json:"usedSpace"`
	// The number of bytes of free space.
	FreeSpaceBytes int64 `json:"freeSpaceBytes"`
	// The number of bytes of total space.
	TotalSpaceBytes int64 `json:"totalSpaceBytes"`
	// The number of bytes of used space.
	UsedSpaceBytes int64 `json:"usedSpaceBytes"`
	// Utilization of this storage location.
	Utilization string `json:"utilization"`
}

type GarbageCollectionDTO struct {
	// The name of the garbage collectors.
	Name string `json:"name"`
	// The number of times garbage collection has run.
	CollectionCount int64 `json:"collectionCount"`
	// The total amount of time spent garbage collecting.
	CollectionTime string `json:"collectionTime"`
	// The total number of milliseconds spent garbage collecting.
	CollectionMillis int64 `json:"collectionMillis"`
}

type VersionInfoDTO struct {
	// The version of this NiFi.
	NiFiVersion string `json:"niFiVersion"`
	// Java JVM vendor
	JavaVendor string `json:"javaVendor"`
	// Java version
	JavaVersion string `json:"javaVersion"`
	// Host operating system name
	OsName string `json:"osName"`
	// Host operating system version
	OsVersion string `json:"osVersion"`
	// Host operating system architecture
	OsArchitecture string `json:"osArchitecture"`
	// Build tag
	BuildTag string `json:"buildTag"`
	// Build revision or commit hash
	BuildRevision string `json:"buildRevision"`
	// Build branch
	BuildBranch string `json:"buildBranch"`
	// Build timestamp
	BuildTimestamp string `json:"buildTimestamp"`
}

type SystemDiagnosticsSnapshotDTO struct {
	// Total size of non heap.
	TotalNonHeap string `json:"totalNonHeap"`
	// Total number of bytes allocated to the JVM not used for heap
	TotalNonHeapBytes int64 `json:"totalNonHeapBytes"`
	// Amount of use non heap.
	UsedNonHeap string `json:"usedNonHeap"`
	// Total number of bytes used by the JVM not in the heap space
	UsedNonHeapBytes int64 `json:"usedNonHeapBytes"`
	// Amount of free non heap.
	FreeNonHeap string `json:"freeNonHeap"`
	// Total number of free non-heap bytes available to the JVM
	FreeNonHeapBytes int64 `json:"freeNonHeapBytes"`
	// Maximum size of non heap.
	MaxNonHeap string `json:"maxNonHeap"`
	// The maximum number of bytes that the JVM can use for non-heap purposes
	MaxNonHeapBytes int64 `json:"maxNonHeapBytes"`
	// Utilization of non heap.
	NonHeapUtilization string `json:"nonHeapUtilization"`
	// Total size of heap.
	TotalHeap string `json:"totalHeap"`
	// The total number of bytes that are available for the JVM heap to use
	TotalHeapBytes int64 `json:"totalHeapBytes"`
	// Amount of used heap.
	UsedHeap string `json:"usedHeap"`
	// The number of bytes of JVM heap that are currently being used
	UsedHeapBytes int64 `json:"usedHeapBytes"`
	// Amount of free heap
	FreeHeap string `json:"freeHeap"`
	// The number of bytes that are allocated to the JVM heap but not currently being used
	FreeHeapBytes int64 `json:"freeHeapBytes"`
	// Maximum size of heap
	MaxHeap string `json:"maxHeap"`
	// The maximum number of bytes that can be used by the JVM
	MaxHeapBytes int64 `json:"maxHeapBytes"`
	// Utilization of heap
	HeapUtilization string `json:"heapUtilization"`
	// Number of available processors if supported by the underlying system
	AvailableProcessors int `json:"availableProcessors"`
	// The processor load average if supported by the underlying system
	ProcessorLoadAverage float64 `json:"processorLoadAverage"`
	// Total number of threads
	TotalThreads int `json:"totalThreads"`
	// Number of daemon threads
	DaemonThreads int `json:"daemonThreads"`
	// The uptime of the Java virtual machine
	Uptime string `json:"uptime"`
	// The flowfile repository storage usage
	FlowFileRepositoryStorageUsage StorageUsageDTO `json:"flowFileRepositoryStorageUsage"`
	// The content repository storage usage
	ContentRepositoryStorageUsage []StorageUsageDTO `json:"contentRepositoryStorageUsage"`
	// The provenance repository storage usage.
	ProvenanceRepositoryStorageUsage []StorageUsageDTO `json:"provenanceRepositoryStorageUsage"`
	// The garbage collection details
	GarbageCollection []GarbageCollectionDTO `json:"garbageCollection"`
	// When the diagnostics were generated
	StatsLastRefreshed string `json:"statsLastRefreshed"`
	// The nifi, os, java, and build version information
	VersionInfo VersionInfoDTO `json:"versionInfo"`
}

type NodeSystemDiagnosticsSnapshotDTO struct {
	// The unique ID that identifies the node
	NodeID string `json:"nodeId"`
	// The API address of the node
	Address string `json:"address"`
	// The API port used to communicate with the node
	ApiPort int `json:"apiPort"`
	// The System Diagnostics snapshot from the node.
	Snapshot SystemDiagnosticsSnapshotDTO `json:"snapshot"`
}

type SystemDiagnosticsDTO struct {
	// A systems diagnostic snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi
	// instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	AggregateSnapshot *SystemDiagnosticsSnapshotDTO `json:"aggregateSnapshot"`
	// A systems diagnostics snapshot for each node in the cluster. If the NiFi instance is a standalone instance,
	// rather than a cluster, this may be null.
	NodeSnapshots []NodeSystemDiagnosticsSnapshotDTO `json:"nodeSnapshots"`
}

type SystemDiagnosticsEntity struct {
	SystemDiagnostics SystemDiagnosticsDTO `json:"systemDiagnostics"`
}
