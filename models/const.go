package models

type TapiCommonOperationalState struct {
	Disable string
	Enable  string
}

func InitTapiCommonOperationalState() TapiCommonOperationalState {
	return TapiCommonOperationalState{
		Disable: "DISABLED",
		Enable:  "ENABLED",
	}
}

type TapiCommonLifecycleState struct {
	Planned           string
	PotentialAvaiable string
	PotentialBusy     string
	Installed         string
	PendingRemoval    string
}

func InitTapiCommonLifecycleState() TapiCommonLifecycleState {
	return TapiCommonLifecycleState{
		Planned:           "PLANNED",
		PotentialAvaiable: "POTENTIAL_AVAILABLE",
		PotentialBusy:     "POTENTIAL_BUSY",
		Installed:         "INSTALLED",
		PendingRemoval:    "PENDING_REMOVAL",
	}
}

type TapiCommonAdministrativeState struct {
	Locked   string
	Unlocked string
}

func InitTapiCommonAdministrativeState() TapiCommonAdministrativeState {
	return TapiCommonAdministrativeState{
		Locked:   "LOCKED",
		Unlocked: "UNLOCKED",
	}
}

type TapiCommonLayerProtocolName struct {
	Odu           string
	Eth           string
	Dsr           string
	PhoyonicMedia string
}

func InitTapiCommonLayerProtocolName() TapiCommonLayerProtocolName {
	return TapiCommonLayerProtocolName{
		Odu:           "ODU",
		Eth:           "ETH",
		Dsr:           "DSR",
		PhoyonicMedia: "PHOTONIC_MEDIA",
	}
}

type TapiCommonTerminationDirection struct {
	Bidirectional   string
	Sink            string
	Source          string
	UndefinedUnknow string
}

func InitTapiCommonTerminationDirection() TapiCommonTerminationDirection {
	return TapiCommonTerminationDirection{
		Bidirectional:   "BIDIRECTIONAL",
		Sink:            "SINK",
		Source:          "SOURCE",
		UndefinedUnknow: "UNDEFINED_OR_UNKNOWN",
	}
}

type TapiCommonTerminationState struct {
	LpCanNeverTerminate          string
	LtNotTerminated              string
	TerminatedServerToClientFlow string
	TerminatedClientToServerFlow string
	TerminatedBidirectional      string
	LtPermenantlyTreminated      string
	TerminationStateUnknow       string
}

func InitTapiCommonTerminationState() TapiCommonTerminationState {
	return TapiCommonTerminationState{
		LpCanNeverTerminate:          "LP_CAN_NEVER_TERMINATE",
		LtNotTerminated:              "LT_NOT_TERMINATED",
		TerminatedServerToClientFlow: "TERMINATED_SERVER_TO_CLIENT_FLOW",
		TerminatedClientToServerFlow: "TERMINATED_CLIENT_TO_SERVER_FLOW",
		TerminatedBidirectional:      "TERMINATED_BIDIRECTIONAL",
		LtPermenantlyTreminated:      "LT_PERMENANTLY_TERMINATED",
		TerminationStateUnknow:       "TERMINATION_STATE_UNKNOWN",
	}
}

type TapiCommonPortRole struct {
	Symmetric string
	Root      string
	Leaf      string
	Trunk     string
	Unknow    string
}

func InitTapiCommonPortRole() TapiCommonPortRole {
	return TapiCommonPortRole{
		Symmetric: "SYMMETRIC",
		Root:      "ROOT",
		Leaf:      "LEAF",
		Trunk:     "TRUNK",
		Unknow:    "UNKNOWN",
	}
}

type TapiCommonPortDirection struct {
	Bidirectional   string
	Input           string
	Output          string
	UndefinedUnknow string
}

func InitTapiCommonPortDirection() TapiCommonPortDirection {
	return TapiCommonPortDirection{
		Bidirectional:   "BIDIRECTIONAL",
		Input:           "INPUT",
		Output:          "OUTPUT",
		UndefinedUnknow: "UNIDENTIFIED_OR_UNKNOWN",
	}
}

type TapiCommonForwardingDirection struct {
	Bidirectional   string
	Uniderectional  string
	UndefinedUnknow string
}

func InitTapiCommonForwardingDirection() TapiCommonForwardingDirection {
	return TapiCommonForwardingDirection{
		Bidirectional:   "BIDIRECTIONAL",
		Uniderectional:  "UNIDIRECTIONAL",
		UndefinedUnknow: "UNIDENTIFIED_OR_UNKNOWN",
	}
}
