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

type TapiConnectivityConnectivityServiceRef struct {
	ConnectivityServiceUuid string `json:"connectivity-service-uuid"`
}

type TapiTopologyNodeRef struct {
	TopologyUuid string `json:"topology-uuid"`
	NodeUuid     string `json:"node-uuid"`
}

type TapiTopologyLinkRef struct {
	TopologyUuid string `json:"topology-uuid"`
	LinkUuid     string `json:"link-uuid"`
}

type TapiPathComputationPathRef struct {
	PathUuid string `json:"path-uuid"`
}

type TapiConnectivityConnectionEndPointRef struct {
	TopologyUuid           string `json:"topology-uuid"`
	NodeUuid               string `json:"node-uuid"`
	NodeEdgePointUuid      string `json:"node-edge-point-uuid"`
	ConnectionEndpointUuid string `json:"connection-end-point-uuid"`
}

type TapiConnectivityConnectionRef struct {
	ConnectionUuid string `json:"connection-uuid"`
}

type TapiConnectivitySwitchControlRef struct {
	ConnectionUuid    string `json:"connection-uuid"`
	SwitchControlUuid string `json:"switch-control-uuid"`
}

type TapiConnectivityRouteRef struct {
	ConnectionUuid string `json:"connection-uuid"`
	RouteLocalId   string `json:"route-local-id"`
}
