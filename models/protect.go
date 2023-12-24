package models

type TapiTopologyResilienceType struct {
	RestorationPolicy string `json:"restoration-policy"`
	ProtectionType    string `json:"protection-type"`
}
