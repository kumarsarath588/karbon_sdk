package v3

//import "time"

// KarbonClusterIntentInput ...
type KarbonClusterIntentInput struct {
	Name *string `json:"name"`

	Description *string `json:"description,omitempty"`

	VmNetwork *string `json:"vm_network"`

	K8sConfig *K8sConfig `json:"k8s_config"`

	ClusterReference *string `json:"cluster_ref"`

	LoggingConfig *LoggingConfig `json:"logging_config"`

	StorageClassConfig *StorageClassConfig `json:"storage_class_config"`

	EtcdConfig *EtcdConfig `json:"etcd_config"`
}

// LoggingConfig
type LoggingConfig struct {
	EnableAppLogging *bool `json:"enable_app_logging"`
}

//EtcdConfig
type EtcdConfig struct {
	NumInstances *int `json:"num_instances"`

	Name *string `json:"name"`

	ResourceConfig *ResourceConfig `json:"resource_config"`
}

// ResourceConfig
type ResourceConfig struct {
	Cpu *int `json:"cpu"`

	MemoryMib *int64 `json:"memory_mib"`

	Image *string `json:"image"`

	DiskMib *int64 `json:"disk_mib"`
}

// K8sConfig
type K8sConfig struct {
	ServiceClusterIpRange *string `json:"service_cluster_ip_range"`

	NetworkCidr *string `json:"network_cidr"`

	Fqdn *string `json:"fqdn"`

	OsFlavor *string `json:"os_flavor"`

	NetworkSubnet *int `json:"network_subnet_len"`

	Version *string `json:"version"`

	Workers []*ResourceConfig `json:"workers"`

	Masters []*ResourceConfig `json:"masters"`
}

// StorageClassConfig
type StorageClassConfig struct {
	Metadata *StorageClassMetadata `json:"metadata"`

	Spec *StorageClassSpec `json:"spec"`
}

//StorageClassMetadata
type StorageClassMetadata struct {
	Name *string `json:"name"`
}

//StorageClassSpec
type StorageClassSpec struct {
	ClusterRef *string `json:"cluster_ref"`

	User *string `json:"user"`

	Password *string `json:"password"`

	StorageContainer *string `json:"storage_container"`

	FileSystem *string `json:"file_system"`

	FlashMode *bool `json:"flash_mode"`
}

//KarbonClusterCreateDeleteResponse
type KarbonClusterCreateDeleteResponse struct {
	UUID *string `json:"cluster_uuid"`

	TaskUUID *string `json:"task_uuid"`
}

// KarbonClusterIntentResponse Response object for intentful operation of karbon cluster
type KarbonClusterIntentResponse struct {
	 Name *string `json:"name"`

	 UUID *string `json:"uuid"`

	 K8sConfig *K8sConfig `json:"k8s_config"`

	 EtcdConfig *EtcdConfig `json:"etcd_config"`

	 AddonsConfig *AddonsConfig `json:"addons_config"`
}

// AddonsConfig
type AddonsConfig struct {
	LoggingConfig *ResponseLoggingConfig `json:"logging_config"`
}

// ResponseLoggingConfig
type ResponseLoggingConfig struct {
	State *string `json:"state"`

	StorageSizeMib *int64 `json:"storage_size_mib"`

	Version *string `json:"version"`
}

//KarbonClusterIntentResponseList
type KarbonClusterIntentResponseList struct {
	ClusterMetadata *KarbonClusterIntentResponse `json:"cluster_metadata"`

	TaskProgressMessage *string `json:"task_progress_message"`

	TaskProgressPercentage *int `json:"task_progress_percent"`

	TaskStatus *int `json:"task_status"`

	TaskType *string `json:"task_type"`

	TaskUuid *string `json:"task_uuid"`
}

//KarbonClusterListIntentResponse
type KarbonClusterListIntentResponse []*KarbonClusterIntentResponseList

//KarbonClusterKubeconfigIntentResponse
type KarbonClusterKubeconfigIntentResponse struct {
	UUID *string `json:"cluster_uuid"`

	YmlConfig *string `json:"yml_config"`
}
