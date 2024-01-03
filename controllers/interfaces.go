package controllers

import (
	"tapi-controller/models"
	"tapi-controller/utils"

	"golang.org/x/net/context"
)

type TapiCtrl interface {
	// return tapi context, correspond to TapiCommonContext model
	GetContext(ctx context.Context, db utils.InMemoryDb, ctxId string) (models.TapiCommonContext, error)

	/*
		Create list of TapiConnectivityConnectivityService in TapiCommonContext
		Input data this is object type of TapiCommonContext contain of list of TapiConnectivityConnectivityService
		If list is empty it return error "list empty"
		Return data is a object CreateConSrvResponse
	*/
	CreateConSrv(ctx context.Context, data models.TapiCommonContext, db *utils.InMemoryDb) (CreateConSrvResponse, error)

	// return list of connective-services
	GetAllConSrv(ctx context.Context, db *utils.InMemoryDb) ([]models.TapiConnectivityConnectivityService, error)

	// return specific connective-service by uuid
	GetConSrv(ctx context.Context, db *utils.InMemoryDb, uuid string) (models.TapiConnectivityConnectivityService, error)

	// update specific connective-service by uuid
	// return updated TapiConnectivityConnectivityService
	UpdateConSrv(ctx context.Context, db *utils.InMemoryDb, uuid string, data models.TapiConnectivityConnectivityServiceData) (models.TapiConnectivityConnectivityService, error)
}
