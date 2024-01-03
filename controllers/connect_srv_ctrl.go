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
		log.Printf("cannot create new conn-srv, err: %x", ctx.Err())
		return CreateConSrvResponse{}, ctx.Err()
	default:
		if len(data.ConnectivityContext.ConnectivityService) == 0 {
			err := errors.New(EmptyConnSrvList)
			log.Printf("cannot create new conn-srv, err: %x", err)
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
		log.Printf("cannot get all connective-services, err: %x", ctx.Err())
		return []models.TapiConnectivityConnectivityService{}, ctx.Err()
	default:
		log.Println("get all connective-services")
		return db.TapiCommonContext.ConnectivityContext.ConnectivityService, nil
	}
}

func (*TapiCtrlInMemDB) GetConSrv(ctx context.Context, db *utils.InMemoryDb, uuid string) (models.TapiConnectivityConnectivityService, error) {
	select {
	case <-ctx.Done():
		log.Printf("cannot get connective-service by uuid: %s, err: %x", uuid, ctx.Err())
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
		log.Printf("cannot update connective-service by uuid: %s, err: %x", uuid, ctx.Err())
		return models.TapiConnectivityConnectivityService{}, ctx.Err()
	default:
		log.Printf("update connective-service by uuid: %s", uuid)
		for _, connSrv := range db.TapiCommonContext.ConnectivityContext.ConnectivityService {
			if connSrv.Uuid == uuid {
				connSrv.AdministrativeState = data.AdministrativeState
				connSrv.AvoidTopology = data.AvoidTopology
				connSrv.Connection = data.Connection
				connSrv.ConnectivityDirection = data.ConnectivityDirection
				connSrv.CorouteInclusion = data.CorouteInclusion
				connSrv.CostCharacteristic = data.CostCharacteristic
				connSrv.DiversityExclusion = data.DiversityExclusion
				connSrv.DiversityPolicy = data.DiversityPolicy
				connSrv.Endpoint = data.Endpoint
				connSrv.ExcludeLink = data.ExcludeLink
				connSrv.ExcludeNode = data.ExcludeNode
				connSrv.ExcludePath = data.ExcludePath
				connSrv.HoldOffTime = data.HoldOffTime
				connSrv.IsFrozen = data.IsFrozen
				connSrv.WaitToRevertTime = data.WaitToRevertTime
				connSrv.ResilienceType = data.ResilienceType
				connSrv.PreferredRestorationLayer = data.PreferredRestorationLayer
				connSrv.RestorePriority = data.RestorePriority
				connSrv.ReversionMode = data.ReversionMode
				connSrv.IsExclusive = data.IsExclusive
				connSrv.DiversityPolicy = data.DiversityPolicy
				connSrv.RouteObjectiveFunction = data.RouteObjectiveFunction
				connSrv.CostCharacteristic = data.CostCharacteristic
				connSrv.LatencyCharacteristic = data.LatencyCharacteristic
				connSrv.RiskDiversityCharacteristic = data.RiskDiversityCharacteristic
				connSrv.RouteDirection = data.RouteDirection
				connSrv.IncludeNode = data.IncludeNode
				connSrv.ExcludeLink = data.ExcludeLink
				connSrv.AvoidTopology = data.AvoidTopology
				connSrv.ExcludePath = data.ExcludePath
				connSrv.IncludeLink = data.IncludeLink
				connSrv.PreferredTransportLayer = data.PreferredTransportLayer
				connSrv.ExcludeNode = data.ExcludeNode
				connSrv.IncludeTopology = data.IncludeTopology
				connSrv.IncludePath = data.IncludePath
				connSrv.Endpoint = data.Endpoint
				connSrv.Connection = data.Connection
				return connSrv, nil
			}
		}
		return models.TapiConnectivityConnectivityService{}, errors.New(ConnSrvNotFound)
	}
}
