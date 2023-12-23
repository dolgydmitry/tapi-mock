package models

type TapiTopologyTopologyRef struct {
	TopologyUuid string `json:"topology-uuid"`
}

type TapiTopologyNodeEdgePointRef struct {
	TopologyUuid      string `json:"topology-uuid"`
	NodeUuid          string `json:"node-uuid"`
	NodeEdgePointUuid string `json:"node-edge-point-uuid"`
}

type TapiCommonServiceInterfacePointRef struct {
	ServiceInterfacePointUuid string `json:"service-interface-point-uuid"`
}

type TapiTopologyNodeRuleGroupRef struct {
	TopologyUuid      string `json:"topology-uuid"`
	NodeUuid          string `json:"node-uuid"`
	NodeRuleGroupUuid string `json:"node-rule-group-uuid"`
}
