package controllers

import (
	"context"
	"errors"
	"log"
	"tapi-controller/models"
	"tapi-controller/utils"
)

const (
	TopologyNotFoundErr = "topology not found"
	NodeNotFoundErr     = "node not found"
)

func GetAllNode(ctx context.Context, topologyUuid string, db *utils.InMemoryDb) ([]models.TapiTopologyNode, error) {
	select {
	case <-ctx.Done():
		return []models.TapiTopologyNode{}, ctx.Err()
	default:
		log.Printf("Get all nodes in topoly: %s", topologyUuid)
		topologyList := db.TapiCommonContext.TopologyContext.Topology
		for _, topology := range topologyList {
			if topology.Uuid == topologyUuid {
				res := topology.Node
				return res, nil
			}
		}
		log.Printf("Topology by uuid: %s not found", topologyUuid)
		return []models.TapiTopologyNode{}, errors.New(TopologyNotFoundErr)
	}
}

func GetNodeUuid(ctx context.Context, topologyUuid string, nodeUuid string, db *utils.InMemoryDb) (models.TapiTopologyNode, error) {
	select {
	case <-ctx.Done():
		return models.TapiTopologyNode{}, ctx.Err()
	default:
		log.Printf("Get nodes in topoly: %s by node uuid: %s", topologyUuid, nodeUuid)
		topologyList := db.TapiCommonContext.TopologyContext.Topology
		var topologyListRes []models.TapiTopologyTopology
		for _, topology := range topologyList {
			if topology.Uuid == topologyUuid {
				topologyListRes = append(topologyListRes, topology)
			}
			if len(topologyListRes) != 1 {
				log.Printf("Topology by uuid: %s not found", topologyUuid)
				return models.TapiTopologyNode{}, errors.New(TopologyNotFoundErr)
			}
			for _, node := range topologyListRes[0].Node {
				if node.Uuid == nodeUuid {
					return node, nil
				}
			}
		}
		log.Printf("Node by uuid: %s in topolgy: %s not found", nodeUuid, topologyUuid)
		return models.TapiTopologyNode{}, errors.New(NodeNotFoundErr)
	}
}
