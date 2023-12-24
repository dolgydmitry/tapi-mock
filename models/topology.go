package models

type TapiTopologyTopology struct {
	Name              []TapiCommonNameAndValue      `json:"name"`
	Uuid              string                        `json:"uuid"`
	LayerProtocolName []TapiCommonLayerProtocolName `json:"layer-protocol-name"`
	Link              []TapiTopologyLink            `json:"link"`
	Node              []TapiTopologyNode            `json:"node"`
}
