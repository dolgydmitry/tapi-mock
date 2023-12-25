package models

type TapiCommonNameAndValue struct {
	ValueName string `json:"value-name"`
	Value     string `json:"value"`
}

type TapiTopologyCostCharacteristic struct {
	CostValue     string `json:"cost-value"`
	CostAlgorithm string `json:"cost-algoritm"`
	CostName      string `json:"cost-name"`
}

type TapiTopologyLatencyCharacteristic struct {
	TrafficPropertyName         string `json:"traffic-property-name"`
	JitterCharacteristic        string `json:"jitter-characteristic"`
	FixedLatencyCharacteristic  string `json:"fixed-latency-characteristic"`
	WanderCharacteristic        string `json:"wander-characteristic"`
	QueingLatencyCharacteristic string `json:"queing-latency-characteristic"`
}

type TapiCommonTimeRange struct {
	EndTime   string `json:"end-time"`
	StartTime string `json:"start-time"`
}
