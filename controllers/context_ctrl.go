package controllers

import (
	"context"
	"log"
	"tapi-controller/models"
	"tapi-controller/utils"
)

func GetContext(ctx context.Context, db utils.InMemoryDb, ctxId string) (models.TapiCommonContext, error) {
	select {
	case <-ctx.Done():
		return models.TapiCommonContext{}, ctx.Err()
	default:
		log.Printf("get context by Id: %s", ctxId)
		return db.TapiCommonContext, nil
	}
}
