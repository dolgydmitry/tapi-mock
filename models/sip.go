package models

type TapiPhotonicMediaFrequencyConstraint struct {
	AdjustmentGranularity string `json:"adjustment-granularity"`
	GridType              string `json:"grid-type"`
}

type TapiPhotonicMediaSpectrumBand struct {
	LowerFrequency      int                                  `json:"lower-frequency"`
	UpperFrequency      int                                  `json:"upper-frequency"`
	FrequencyConstraint TapiPhotonicMediaFrequencyConstraint `json:"frequency-constraint"`
}

type TapiPhotonicMediaMediaChannelPoolCapabilityPac struct {
	AvailableSpectrum   []TapiPhotonicMediaSpectrumBand `json:"available-spectrum"`
	SupportableSpectrum []TapiPhotonicMediaSpectrumBand `json:"supportable-spectrum"`
	OccupiedSpectrum    []TapiPhotonicMediaSpectrumBand `json:"occupied-spectrum"`
}

type TapiPhotonicMediaMediaChannelServiceInterfacePointSpec struct {
	McPool TapiPhotonicMediaMediaChannelPoolCapabilityPac `json:"mc-pool"`
}

type TapiPhotonicMediaCentralFrequency struct {
	CentralFrequency    int                                  `json:"central-frequency"`
	FrequencyConstraint TapiPhotonicMediaFrequencyConstraint `json:"frequency-constraint"`
}

type TapiPhotonicMediaApplicationIdentifier struct {
	ApplicationIdentifierType string `json:"application-identifier-type"`
	ApplicationCode           string `json:"application-code"`
}

type TapiPhotonicMediaTotalPowerThresholdPac struct {
	TotalPowerUpperWarntThresholdDefault string `json:"total-power-upper-warn-threshold-default"`
	TotalPowerLowerWarnThresholdMin      string `json:"total-power-lower-warn-threshold-min"`
	TotalPowerUpperWarnThresholdMin      string `json:"total-power-upper-warn-threshold-min"`
	TotalPowerUpperWarnThresholdMax      string `json:"total-power-upper-warn-threshold-max"`
	TotalPowerLowerWarnThresholdMax      string `json:"total-power-lower-warn-threshold-max"`
	TotalPowerLowerWarntThresholdDefault string `json:"total-power-lower-warn-threshold-default"`
}

type TapiPhotonicMediaOtsiCapabilityPac struct {
	SupportableLowerCentralFrequency []TapiPhotonicMediaCentralFrequency      `json:"supportable-lower-central-frequency"`
	SupportableApplicationIdentifier []TapiPhotonicMediaApplicationIdentifier `json:"supportable-application-identifier"`
	SupportableModulation            []string                                 `json:"supportable-modulation"`
	TotalPowerWarnThreshold          TapiPhotonicMediaTotalPowerThresholdPac  `json:"total-power-warn-threshold"`
	SupportableUpperCentralFrequency []TapiPhotonicMediaCentralFrequency      `json:"supportable-upper-central-frequency"`
}

type TapiPhotonicMediaOtsiServiceInterfacePointSpec struct {
	OtsiCapability TapiPhotonicMediaOtsiCapabilityPac `json:"otsi-capability"`
}

type TapiCommonContextServiceInterfacePoint struct {
	OperationalState                      string                                                 `json:"operational-state"`
	LifecycleState                        string                                                 `json:"lifecycle-state"`
	AdministrativeState                   string                                                 `json:"administrative-state"`
	AvailableCapacity                     TapiCommonCapacity                                     `json:"available-capacity"`
	TotalPotentialCapacity                TapiCommonCapacity                                     `json:"total-potential-capacity"`
	Name                                  []TapiCommonNameAndValue                               `json:"name"`
	Uuid                                  string                                                 `json:"uuid"`
	SupportedLayerProtocolQualifier       []string                                               `json:"supported-layer-protocol-qualifier"`
	LayerProtocolName                     []string                                               `json:"layer-protocol-name"`
	MediaChannelServiceInterfacePointSpec TapiPhotonicMediaMediaChannelServiceInterfacePointSpec `json:"media-channel-service-interface-point-spec"`
	OtsiServiceInterfacePointSpec         TapiPhotonicMediaOtsiServiceInterfacePointSpec         `json:"otsi-service-interface-point-spec"`
}
