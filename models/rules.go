package models

type TapiTopologyRule struct {
	Name             []TapiCommonNameAndValue `json:"name"`
	LocalId          string                   `json:"local-id"`
	OverridePriority int                      `json:"override-priority"`
	ForwardingRule   string                   `json:"forwarding-rule"`
	RuleType         string                   `json:"rule-type"`
}

type TapiTopologyInterRuleGroup struct {
	AvailableCapacity       TapiCommonCapacity                  `json:"available-capacity"`
	TotalPotentialCapacity  TapiCommonCapacity                  `json:"total-potential-capacity"`
	Name                    []TapiCommonNameAndValue            `json:"name"`
	Uuid                    string                              `json:"uuid"`
	RiskCharacteristic      []TapiTopologyRiskCharacteristic    `json:"risk-characteristic"`
	CostCharacteristic      []TapiTopologyCostCharacteristic    `json:"cost-characteristic"`
	LatencyCharacteristic   []TapiTopologyLatencyCharacteristic `json:"latency-characteristic"`
	AssociatedNodeRuleGroup []TapiTopologyNodeRuleGroupRef      `json:"associated-node-rule-group"`
	Rule                    []TapiTopologyRule                  `json:"rule"`
}

type TapiTopologyNodeRuleGroup struct {
	AvailableCapacity      TapiCommonCapacity                  `json:"available-capacity"`
	TotalPotentialCapacity TapiCommonCapacity                  `json:"total-potential-capacity"`
	Name                   []TapiCommonNameAndValue            `json:"name"`
	Uuid                   string                              `json:"uuid"`
	RiskCharacteristic     []TapiTopologyRiskCharacteristic    `json:"risk-characteristic"`
	CostCharacteristic     []TapiTopologyCostCharacteristic    `json:"cost-characteristic"`
	LatencyCharacteristic  []TapiTopologyLatencyCharacteristic `json:"latency-characteristic"`
	InterRuleGroup         []TapiTopologyInterRuleGroup        `json:"inter-rule-group"`
	Rule                   []TapiTopologyRule                  `json:"rule"`
	ComposedRuleGroup      []TapiTopologyNodeRuleGroupRef      `json:"composed-rule-group"`
	NodeEdgePoint          []TapiTopologyNodeEdgePointRef      `json:"node-edge-point"`
}
