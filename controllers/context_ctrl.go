package controllers

import (
	"log"
	"tapi-controller/models"
	"tapi-controller/utils"

	"golang.org/x/net/context"
)

func (*TapiCtrlInMemDB) GetContext(ctx context.Context, db utils.InMemoryDb, ctxId string) (models.TapiCommonContext, error) {
	for {
		select {
		case <-ctx.Done():
			return models.TapiCommonContext{}, ctx.Err()
		default:
			log.Printf("get context by Id: %s", ctxId)
			return db.TapiCommonContext, nil
		}
	}
}
