package models

type TapiTopologyValidationMechanism struct {
	LayerProtocolAdjacencyValidated string `json:"layer-protocol-adjacency-validated"`
	ValidationMechanism             string `json:"validation-mechanism"`
	ValidationRobustness            string `json:"validation-robustness"`
}
