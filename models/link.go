package models

type TapiTopologyLink struct {
	OperationalState                     string                              `json:"operational-state"`
	LifecycleState                       string                              `json:"lifecycle-state"`
	AdministrativeState                  string                              `json:"administrative-state"`
	AvailableCapacity                    TapiCommonCapacity                  `json:"available-capacity"`
	TotalPotentialCapacity               TapiCommonCapacity                  `json:"total-potential-capacity"`
	Name                                 []TapiCommonNameAndValue            `json:"name"`
	Uuid                                 string                              `json:"uuid"`
	TransitionedLayerProtocolName        []string                            `json:"transitioned-layer-protocol-name"`
	RiskCharacteristic                   []TapiTopologyRiskCharacteristic    `json:"risk-characteristic"`
	CostCharacteristic                   []TapiTopologyCostCharacteristic    `json:"cost-characteristic"`
	ErrorCharacteristic                  string                              `json:"error-characteristic"`
	UnavailableTimeCharacteristic        string                              `json:"unavailable-time-characteristic"`
	ServerIntegrityProcessCharacteristic string                              `json:"server-integrity-process-characteristic"`
	DeliveryOrderCharacteristic          string                              `json:"delivery-order-characteristic"`
	RepeatDeliveryCharacteristic         string                              `json:"repeat-delivery-characteristic"`
	LossCharacteristic                   string                              `json:"loss-characteristic"`
	LatencyCharacteristic                []TapiTopologyLatencyCharacteristic `json:"latency-characteristic"`
	ValidationMechanism                  []TapiTopologyValidationMechanism   `json:"validation-mechanism"`
	LayerProtocolName                    []string                            `json:"layer-protocol-name"`
	ResilienceType                       TapiTopologyResilienceType          `json:"resilience-type"`
	NodeEdgePoint                        []TapiTopologyNodeEdgePointRef      `json:"node-edge-point"`
	Direction                            string                              `json:"direction"`
}
