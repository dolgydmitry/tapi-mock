package controllers

import (
	"errors"
	"log"
	"tapi-controller/models"
	"tapi-controller/utils"

	"golang.org/x/net/context"
)

const (
	DesiredOrchStateActive = "active"
	OrchStateRequsted      = "requested"
	ConnSrvNotFound        = "connective-service not found"
	EmptyConnSrvList       = "input connective-services list is empty"
)

type CreateConSrvResponse struct {
	DesiredOrchState string
	OrchState        string
}

func (*TapiCtrlInMemDB) CreateConSrv(ctx context.Context, data models.TapiCommonContext, db *utils.InMemoryDb) (CreateConSrvResponse, error) {
	select {
	case <-ctx.Done():
		log.Printf("cannot create new conn-srv, err: %s", ctx.Err())
		return CreateConSrvResponse{}, ctx.Err()
	default:
		if len(data.ConnectivityContext.ConnectivityService) == 0 {
			err := errors.New(EmptyConnSrvList)
			log.Printf("cannot create new conn-srv, err: %s", err)
			return CreateConSrvResponse{}, err
		}
		log.Println("create new connective-service")
		db.TapiCommonContext.ConnectivityContext.ConnectivityService = data.ConnectivityContext.ConnectivityService
		res := CreateConSrvResponse{
			DesiredOrchState: DesiredOrchStateActive,
			OrchState:        OrchStateRequsted,
		}
		return res, nil
	}
}

func (*TapiCtrlInMemDB) GetAllConSrv(ctx context.Context, db *utils.InMemoryDb) ([]models.TapiConnectivityConnectivityService, error) {
	select {
	case <-ctx.Done():
		log.Printf("cannot get all connective-services, err: %s", ctx.Err())
		return []models.TapiConnectivityConnectivityService{}, ctx.Err()
	default:
		log.Println("get all connective-services")
		return db.TapiCommonContext.ConnectivityContext.ConnectivityService, nil
	}
}

func (*TapiCtrlInMemDB) GetConSrv(ctx context.Context, db *utils.InMemoryDb, uuid string) (models.TapiConnectivityConnectivityService, error) {
	select {
	case <-ctx.Done():
		log.Printf("cannot get connective-service by uuid: %s, err: %s", uuid, ctx.Err())
		return models.TapiConnectivityConnectivityService{}, ctx.Err()
	default:
		log.Printf("get connective-service by uuid: %s", uuid)
		for _, connSrv := range db.TapiCommonContext.ConnectivityContext.ConnectivityService {
			if connSrv.Uuid == uuid {
				return connSrv, nil
			}
		}
		return models.TapiConnectivityConnectivityService{}, errors.New(ConnSrvNotFound)
	}
}

func (*TapiCtrlInMemDB) UpdateConSrv(ctx context.Context, db *utils.InMemoryDb, uuid string, data models.TapiConnectivityConnectivityServiceData) (models.TapiConnectivityConnectivityService, error) {
	select {
	case <-ctx.Done():
		log.Printf("cannot update connective-service by uuid: %s, err: %s", uuid, ctx.Err())
		return models.TapiConnectivityConnectivityService{}, ctx.Err()
	default:
		log.Printf("update connective-service by uuid: %s", uuid)
		for index, connSrv := range db.TapiCommonContext.ConnectivityContext.ConnectivityService {
			if connSrv.Uuid == uuid {
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].AdministrativeState = data.AdministrativeState
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].AvoidTopology = data.AvoidTopology
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].Connection = data.Connection
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ConnectivityDirection = data.ConnectivityDirection
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].CorouteInclusion = data.CorouteInclusion
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].CostCharacteristic = data.CostCharacteristic
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].DiversityExclusion = data.DiversityExclusion
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].DiversityPolicy = data.DiversityPolicy
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].Endpoint = data.Endpoint
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludeLink = data.ExcludeLink
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludeNode = data.ExcludeNode
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludePath = data.ExcludePath
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].HoldOffTime = data.HoldOffTime
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IsFrozen = data.IsFrozen
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].WaitToRevertTime = data.WaitToRevertTime
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ResilienceType = data.ResilienceType
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].PreferredRestorationLayer = data.PreferredRestorationLayer
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].RestorePriority = data.RestorePriority
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ReversionMode = data.ReversionMode
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IsExclusive = data.IsExclusive
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].DiversityPolicy = data.DiversityPolicy
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].RouteObjectiveFunction = data.RouteObjectiveFunction
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].CostCharacteristic = data.CostCharacteristic
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].LatencyCharacteristic = data.LatencyCharacteristic
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].RiskDiversityCharacteristic = data.RiskDiversityCharacteristic
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].RouteDirection = data.RouteDirection
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IncludeNode = data.IncludeNode
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludeLink = data.ExcludeLink
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].AvoidTopology = data.AvoidTopology
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludePath = data.ExcludePath
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IncludeLink = data.IncludeLink
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].PreferredTransportLayer = data.PreferredTransportLayer
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].ExcludeNode = data.ExcludeNode
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IncludeTopology = data.IncludeTopology
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].IncludePath = data.IncludePath
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].Endpoint = data.Endpoint
				db.TapiCommonContext.ConnectivityContext.ConnectivityService[index].Connection = data.Connection
				return db.TapiCommonContext.ConnectivityContext.ConnectivityService[index], nil
			}
		}
		return models.TapiConnectivityConnectivityService{}, errors.New(ConnSrvNotFound)
	}
}
