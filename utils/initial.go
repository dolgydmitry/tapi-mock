package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"tapi-controller/models"
)

const (
	dbDataPath = "/db_data/context.json"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type InMemoryDb struct {
	TapiCommonOperationalState     models.TapiCommonOperationalState
	TapiCommonLifecycleState       models.TapiCommonLifecycleState
	TapiCommonAdministrativeState  models.TapiCommonAdministrativeState
	TapiCommonLayerProtocolName    models.TapiCommonLayerProtocolName
	TapiCommonTerminationDirection models.TapiCommonTerminationDirection
	TapiCommonTerminationState     models.TapiCommonTerminationState
	TapiCommonPortRole             models.TapiCommonPortRole
	TapiCommonPortDirection        models.TapiCommonPortDirection
	TapiCommonForwardingDirection  models.TapiCommonForwardingDirection
	TapiCommonContext              models.TapiCommonContext
}

func InitMemoryDb() InMemoryDb {
	dbDataPathRelative := fmt.Sprintf("%s%s", basepath, dbDataPath)
	log.Println("Start to init in memory DB")
	jsonFile, err := os.Open(dbDataPathRelative)
	if err != nil {
		log.Panicf("Cannot open file: %s, with error: %x", dbDataPathRelative, err)
	}
	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Panicf("Cannot read file: %s, with error: %x", dbDataPathRelative, err)
	}
	var context models.TapiCommonContext
	json.Unmarshal(byteJson, &context)
	res := InMemoryDb{
		TapiCommonOperationalState:     models.InitTapiCommonOperationalState(),
		TapiCommonLifecycleState:       models.InitTapiCommonLifecycleState(),
		TapiCommonAdministrativeState:  models.InitTapiCommonAdministrativeState(),
		TapiCommonLayerProtocolName:    models.InitTapiCommonLayerProtocolName(),
		TapiCommonTerminationDirection: models.InitTapiCommonTerminationDirection(),
		TapiCommonTerminationState:     models.InitTapiCommonTerminationState(),
		TapiCommonPortRole:             models.InitTapiCommonPortRole(),
		TapiCommonPortDirection:        models.InitTapiCommonPortDirection(),
		TapiCommonForwardingDirection:  models.InitTapiCommonForwardingDirection(),
		TapiCommonContext:              context,
	}
	log.Println("In memroy DB created")
	return res

}
