package models

type TapiTopologyNetworkTopologyService struct {
	Name     []TapiCommonNameAndValue  `json:"name"`
	Uuid     string                    `json:"uuid"`
	Topology []TapiTopologyTopologyRef `json:"topology"`
}

type TapiTopologyTopologyContext struct {
	NwTopologyService TapiTopologyNetworkTopologyService `json:"nw-topology-service"`
	Topology          []TapiTopologyTopology             `json:"topology"`
}

type TapiCommonContext struct {
	Name                  []TapiCommonNameAndValue                 `json:"name"`
	Uuid                  string                                   `json:"uuid"`
	TopologyContext       TapiTopologyTopologyContext              `json:"topology-context"`
	ServiceInterfacePoint []TapiCommonContextServiceInterfacePoint `json:"service-interface-point"`
	ConnectivityContext   TapiConnectivityConnectivityContext      `json:"tapi-connectivity:connectivity-context"`
}
