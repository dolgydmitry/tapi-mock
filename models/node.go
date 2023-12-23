package models

type TapiTopologyNode struct {
	OperationalState                     TapiCommonOperationalState          `json:"operational-state"`
	LifecycleState                       TapiCommonLifecycleState            `json:"lifecycle-state"`
	AdministrativeState                  TapiCommonAdministrativeState       `json:"administrative-state"`
	AvailableCapacity                    TapiCommonCapacity                  `json:"available-capacity"`
	TotalPotentialCapacity               TapiCommonCapacity                  `json:"total-potential-capacity"`
	Name                                 []TapiCommonNameAndValue            `json:"name"`
	Uuid                                 string                              `json:"uuid"`
	CostCharacteristic                   []TapiTopologyCostCharacteristic    `json:"cost-characteristic"`
	ErrorCharacteristic                  string                              `json:"error-characteristic"`
	UnavailableTimeCharacteristic        string                              `json:"unavailable-time-characteristic"`
	ServerIntegrityProcessCharacteristic string                              `json:"server-integrity-process-characteristic"`
	DeliveryOrderCharacteristic          string                              `json:"delivery-order-characteristic"`
	RepeatDeliveryCharacteristic         string                              `json:"repeat-delivery-characteristic"`
	LossCharacteristic                   string                              `json:"loss-characteristic"`
	LatencyCharacteristic                []TapiTopologyLatencyCharacteristic `json:"latency-characteristic"`
	LayerProtocolName                    []string                            `json:"layer-protocol-name"`
	EncapTopology                        TapiTopologyTopologyRef             `json:"encap-topology"`
	OwnedNodeEdgePoint                   []TapiTopologyNodeEdgePoint         `json:"owned-node-edge-point"`
	NodeRuleGroup                        []TapiTopologyNodeRuleGroup         `json:"node-rule-group"`
	AggregatedNodeEdgePoint              []TapiTopologyNodeEdgePointRef      `json:"aggregated-node-edge-point"`
}
