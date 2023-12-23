package models

type TapiCommonCapacityValue struct {
	Value int    `json:"value"`
	Unit  string `json:"unit" validator:"min=3,max=40,regexp=TB|TBPS|GB|GBPS|MB|MBS|KB|KBS|GHZ|MHZ"`
}

type TapiCommonBandwidthProfile struct {
	CommittedInformationRate TapiCommonCapacityValue `json:"committed-information-rate"`
	CouplingFlag             bool                    `json:"coupling-flag"`
	BwProfileType            string                  `json:"bw-profile-type" validator:"MEF_10.x|RFC_2697|RFC_2698|RFC_4115"`
	PeakInformationRate      TapiCommonCapacityValue `json:"peak-information-rate"`
	CommittedBurstSize       TapiCommonCapacityValue `json:"committed-burst-size"`
	PeakBurstSize            TapiCommonCapacityValue `json:"peak-burst-size"`
	ColorAware               bool                    `json:"color-aware"`
}

type TapiCommonCapacity struct {
	BandwidthProfile TapiCommonBandwidthProfile `json:"bandwidth-profile"`
	TotalSize        TapiCommonCapacityValue    `json:"total-size"`
}
