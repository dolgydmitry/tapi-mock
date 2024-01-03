package models

type TapiConnectivityConnectivityContext struct {
	ConnectivityService []TapiConnectivityConnectivityService `json:"connectivity-service"`
	Connection          []TapiConnectivityConnection          `json:"connection"`
}

type TapiConnectivityConnectivityService struct {
	OperationalState               string                                        `json:"operational-state"`
	LifecycleState                 string                                        `json:"lifecycle-state"`
	AdministrativeState            string                                        `json:"administrative-state" binding:"required"`
	Name                           []TapiCommonNameAndValue                      `json:"name" binding:"required"`
	Uuid                           string                                        `json:"uuid" binding:"required"`
	ServiceLayer                   string                                        `json:"service-layer" binding:"required"`
	Schedule                       TapiCommonTimeRange                           `json:"schedule"`
	ConnectivityDirection          string                                        `json:"connectivity-direction"`
	RequestedCapacity              TapiCommonCapacity                            `json:"requested-capacity"`
	DiversityExclusion             []TapiConnectivityConnectivityServiceRef      `json:"diversity-exclusion"`
	ServiceLevel                   string                                        `json:"service-level"`
	ServiceType                    string                                        `json:"service-type"`
	CorouteInclusion               TapiConnectivityConnectivityServiceRef        `json:"coroute-inclusion"`
	IsLockOut                      bool                                          `json:"is-lock-out"`
	MaxSwitchTimes                 int                                           `json:"'max-switch-times'"`
	RestorationCoordinateType      string                                        `json:"restoration-coordinate-type"`
	IsCoordinatedSwitchingBothEnds bool                                          `json:"is-coordinated-switching-both-ends"`
	HoldOffTime                    int                                           `json:"hold-off-time"`
	IsFrozen                       bool                                          `json:"is-frozen"`
	WaitToRevertTime               int                                           `json:"wait-to-revert-time"`
	ResilienceType                 TapiTopologyResilienceType                    `json:"resilience-type"`
	PreferredRestorationLayer      []string                                      `json:"preferred-restoration-layer"`
	RestorePriority                int                                           `json:"restore-priority"`
	ReversionMode                  string                                        `json:"reversion-mode"`
	IsExclusive                    bool                                          `json:"is-exclusive"`
	DiversityPolicy                string                                        `json:"diversity-policy"`
	RouteObjectiveFunction         string                                        `json:"route-objective-function"`
	CostCharacteristic             []TapiTopologyCostCharacteristic              `json:"cost-characteristic"`
	LatencyCharacteristic          []TapiTopologyLatencyCharacteristic           `json:"latency-characteristic"`
	RiskDiversityCharacteristic    []TapiTopologyRiskCharacteristic              `json:"risk-diversity-characteristic"`
	RouteDirection                 string                                        `json:"route-direction"`
	IncludeNode                    []TapiTopologyNodeRef                         `json:"include-node"`
	ExcludeLink                    []TapiTopologyLinkRef                         `json:"exclude-link"`
	AvoidTopology                  []TapiTopologyTopologyRef                     `json:"avoid-topology"`
	ExcludePath                    []TapiPathComputationPathRef                  `json:"exclude-path"`
	IncludeLink                    []TapiTopologyLinkRef                         `json:"include-link"`
	PreferredTransportLayer        string                                        `json:"preferred-transport-layer"`
	ExcludeNode                    []TapiTopologyNodeRef                         `json:"exclude-node"`
	IncludeTopology                []TapiTopologyTopologyRef                     `json:"include-topology"`
	IncludePath                    []TapiPathComputationPathRef                  `json:"include-path"`
	Endpoint                       []TapiConnectivityConnectivityserviceEndPoint `json:"end-point" binding:"required"`
	Connection                     []TapiConnectivityConnectionRef               `json:"connection"`
}

type TapiConnectivityConnectivityServiceData struct {
	OperationalState               string                                        `json:"operational-state"`
	LifecycleState                 string                                        `json:"lifecycle-state"`
	AdministrativeState            string                                        `json:"administrative-state" binding:"required"`
	Name                           []TapiCommonNameAndValue                      `json:"name" binding:"required"`
	Uuid                           string                                        `json:"uuid"`
	ServiceLayer                   string                                        `json:"service-layer" binding:"required"`
	Schedule                       TapiCommonTimeRange                           `json:"schedule"`
	ConnectivityDirection          string                                        `json:"connectivity-direction"`
	RequestedCapacity              TapiCommonCapacity                            `json:"requested-capacity"`
	DiversityExclusion             []TapiConnectivityConnectivityServiceRef      `json:"diversity-exclusion"`
	ServiceLevel                   string                                        `json:"service-level"`
	ServiceType                    string                                        `json:"service-type"`
	CorouteInclusion               TapiConnectivityConnectivityServiceRef        `json:"coroute-inclusion"`
	IsLockOut                      bool                                          `json:"is-lock-out"`
	MaxSwitchTimes                 int                                           `json:"'max-switch-times'"`
	RestorationCoordinateType      string                                        `json:"restoration-coordinate-type"`
	IsCoordinatedSwitchingBothEnds bool                                          `json:"is-coordinated-switching-both-ends"`
	HoldOffTime                    int                                           `json:"hold-off-time"`
	IsFrozen                       bool                                          `json:"is-frozen"`
	WaitToRevertTime               int                                           `json:"wait-to-revert-time"`
	ResilienceType                 TapiTopologyResilienceType                    `json:"resilience-type"`
	PreferredRestorationLayer      []string                                      `json:"preferred-restoration-layer"`
	RestorePriority                int                                           `json:"restore-priority"`
	ReversionMode                  string                                        `json:"reversion-mode"`
	IsExclusive                    bool                                          `json:"is-exclusive"`
	DiversityPolicy                string                                        `json:"diversity-policy"`
	RouteObjectiveFunction         string                                        `json:"route-objective-function"`
	CostCharacteristic             []TapiTopologyCostCharacteristic              `json:"cost-characteristic"`
	LatencyCharacteristic          []TapiTopologyLatencyCharacteristic           `json:"latency-characteristic"`
	RiskDiversityCharacteristic    []TapiTopologyRiskCharacteristic              `json:"risk-diversity-characteristic"`
	RouteDirection                 string                                        `json:"route-direction"`
	IncludeNode                    []TapiTopologyNodeRef                         `json:"include-node"`
	ExcludeLink                    []TapiTopologyLinkRef                         `json:"exclude-link"`
	AvoidTopology                  []TapiTopologyTopologyRef                     `json:"avoid-topology"`
	ExcludePath                    []TapiPathComputationPathRef                  `json:"exclude-path"`
	IncludeLink                    []TapiTopologyLinkRef                         `json:"include-link"`
	PreferredTransportLayer        string                                        `json:"preferred-transport-layer"`
	ExcludeNode                    []TapiTopologyNodeRef                         `json:"exclude-node"`
	IncludeTopology                []TapiTopologyTopologyRef                     `json:"include-topology"`
	IncludePath                    []TapiPathComputationPathRef                  `json:"include-path"`
	Endpoint                       []TapiConnectivityConnectivityserviceEndPoint `json:"end-point" binding:"required"`
	Connection                     []TapiConnectivityConnectionRef               `json:"connection"`
}

type TapiConnectivityConnectivityserviceEndPoint struct {
	OperationalState                      TapiCommonOperationalState                             `json:"operational-state"`
	LifecycleState                        TapiCommonLifecycleState                               `json:"lifecycle-state"`
	AdministrativeState                   TapiCommonAdministrativeState                          `json:"administrative-state"`
	Name                                  []TapiCommonNameAndValue                               `json:"name"`
	LocalId                               string                                                 `json:"local-id"`
	ProtectionRole                        string                                                 `json:"protection-role"`
	Role                                  string                                                 `json:"role"`
	ServiceInterfacePoint                 TapiCommonServiceInterfacePointRef                     `json:"service-interface-point"`
	LayerProtocolName                     string                                                 `json:"layer-protocol-name"`
	LayerProtocolQualifier                string                                                 `json:"layer-protocol-qualifier"`
	ConnectionEndpoint                    []TapiConnectivityConnectionEndPointRef                `json:"connection-end-point"`
	Direction                             string                                                 `json:"direction"`
	Capacity                              TapiCommonCapacity                                     `json:"capacity"`
	MediaChannelServiceInterfacePointSpec TapiPhotonicMediaMediaChannelServiceInterfacePointSpec `json:"media-channel-service-interface-point-spec"`
	OtsiConnectivityServiceEndpointSpec   TapiPhotonicMediaOtsiConnectivityServiceEndPointSpec   `json:"otsi-connectivity-service-end-point-spec"`
}

type TapiPhotonicMediaOtsiConnectivityServiceEndPointSpec struct {
	OtsiConfig TapiPhotonicMediaOtsiTerminationConfigPac `json:"otsi-config"`
}

type TapiPhotonicMediaOtsiTerminationConfigPac struct {
	ApplicationIdentifier        TapiPhotonicMediaApplicationIdentifier `json:"application-identifier"`
	CentralFrequency             TapiPhotonicMediaCentralFrequency      `json:"central-frequency"`
	Modulation                   string                                 `json:"modulation"`
	Spectrum                     TapiPhotonicMediaSpectrumBand          `json:"spectrum"`
	LaserControl                 string                                 `json:"laser-control"`
	TotalPowerWarnThresholdLower string                                 `json:"total-power-warn-threshold-lower"`
	TotalPowerWarnThresholdUpper string                                 `json:"total-power-warn-threshold-upper"`
	TransmitPower                TapiPhotonicMediaPowerPropertiesPac    `json:"transmit-power"`
}

type TapiPhotonicMediaPowerPropertiesPac struct {
	PowerSpectralDensity string `json:"power-spectral-density"`
	TotalPower           string `json:"total-power"`
}
