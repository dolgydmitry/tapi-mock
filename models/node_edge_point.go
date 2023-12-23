package models

type TapiTopologyNodeEdgePoint struct {
	OperationalState                   TapiCommonOperationalState           `json:"operational-state"`
	LifecycleState                     TapiCommonLifecycleState             `json:"lifecycle-state"`
	AdministrativeState                TapiCommonAdministrativeState        `json:"administrative-state"`
	AvailableCapacity                  TapiCommonCapacity                   `json:"available-capacity"`
	TotalPotentialCapacity             TapiCommonCapacity                   `json:"total-potential-capacity"`
	Name                               []TapiCommonNameAndValue             `json:"name"`
	Uuid                               string                               `json:"uuid"`
	TerminationDirection               string                               `json:"termination-direction"`
	TerminationState                   string                               `json:"termination-state"`
	LinkPortRole                       string                               `json:"link-port-role"`
	MappedServiceInterfacePoint        []TapiCommonServiceInterfacePointRef `json:"mapped-service-interface-point"`
	AggregatedNodeEdgePoint            []TapiTopologyNodeEdgePointRef       `json:"aggregated-node-edge-point"`
	LayerProtocolName                  string                               `json:"layer-protocol-namet"`
	LinkPortDirection                  string                               `json:"link-port-direction"`
	SupportedCepLayerProtocolQualifier []string                             `json:"supported-cep-layer-protocol-qualifier"`
}
