package models

type TapiConnectivityConnection struct {
	Name                []TapiCommonNameAndValue                `json:"name"`
	Uuid                string                                  `json:"uuid"`
	OperationalState    TapiCommonOperationalState              `json:"operational-state"`
	LifecycleState      TapiCommonLifecycleState                `json:"lifecycle-state"`
	SupportedClientLink []TapiTopologyLinkRef                   `json:"supported-client-link"`
	LowerConnection     []TapiConnectivityConnectionRef         `json:"lower-connection"`
	SwitchControl       []TapiConnectivitySwitchControl         `json:"switch-control"`
	Route               []TapiConnectivityRoute                 `json:"route"`
	LayerProtocolName   string                                  `json:"layer-protocol-name"`
	ConnectionEndpoint  []TapiConnectivityConnectionEndPointRef `json:"connection-end-point"`
	Direction           string                                  `json:"direction"`
}

type TapiConnectivityRoute struct {
	Name               []TapiCommonNameAndValue                `json:"name"`
	LocalId            string                                  `json:"local-id"`
	ConnectionEndpoint []TapiConnectivityConnectionEndPointRef `json:"connection-end-point"`
}

type TapiConnectivitySwitchControl struct {
	Name                           []TapiCommonNameAndValue           `json:"name"`
	Uuid                           string                             `json:"uuid"`
	IsLockOut                      bool                               `json:"is-lock-out"`
	MaxSwitchTimes                 int                                `json:"max-switch-times"`
	RestorationCoordinateType      string                             `json:"restoration-coordinate-type"`
	IsCoordinatedSwitchingBothEnds bool                               `json:"is-coordinated-switching-both-ends"`
	HoldOffTime                    int                                `json:"hold-off-time"`
	IsFrozen                       bool                               `json:"is-frozen"`
	WaitToRevertTime               int                                `json:"wait-to-revert-time"`
	ResilienceType                 TapiTopologyResilienceType         `json:"resilience-type"`
	PreferredRestorationLayer      []string                           `json:"preferred-restoration-layer"`
	RestorePriority                int                                `json:"restore-priority"`
	ReversionMode                  string                             `json:"reversion-mode"`
	SubSwitchControl               []TapiConnectivitySwitchControlRef `json:"sub-switch-control"`
	Switch                         []TapiConnectivitySwitch           `json:"switch"`
}

type TapiConnectivitySwitch struct {
	Name                       []TapiCommonNameAndValue                `json:"name"`
	LocalId                    string                                  `json:"local-id"`
	SelectedConnectionEndpoint []TapiConnectivityConnectionEndPointRef `json:"selected-connection-end-point"`
	SelectedRoute              []TapiConnectivityRouteRef              `json:"selected-route"`
	SelectionControl           string                                  `json:"selection-control"`
	SelectionReason            string                                  `json:"selection-reason"`
	SwitchDirection            string                                  `json:"switch-direction"`
}
