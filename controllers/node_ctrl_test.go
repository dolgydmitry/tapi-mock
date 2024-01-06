package controllers

import (
	"context"
	"tapi-controller/models"
	"tapi-controller/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllNode(T *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	ctx := context.Background()

	testCases := []struct {
		name    string
		ctx     context.Context
		uuid    string
		db      *utils.InMemoryDb
		checker func(t *testing.T, res []models.TapiTopologyNode, err error)
	}{
		{
			name: "Good day",
			ctx:  ctx,
			uuid: "top-tpd",
			db:   inMemoryDB,
			checker: func(t *testing.T, res []models.TapiTopologyNode, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
				require.Len(t, res, 3)
			},
		},
		{
			name: "Not found",
			ctx:  ctx,
			uuid: "temp",
			db:   inMemoryDB,
			checker: func(t *testing.T, res []models.TapiTopologyNode, err error) {
				require.EqualError(t, err, TopologyNotFoundErr)
				require.Len(t, res, 0)
			},
		},
	}
	for _, testCase := range testCases {
		res, err := GetAllNode(testCase.ctx, testCase.uuid, testCase.db)
		testCase.checker(T, res, err)
	}
}

func TestGetNodeUuid(T *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	ctx := context.Background()
	testCases := []struct {
		name     string
		ctx      context.Context
		db       *utils.InMemoryDb
		topoUuid string
		nodeUuid string
		checker  func(t *testing.T, res models.TapiTopologyNode, err error)
	}{
		{
			name:     "good day",
			ctx:      ctx,
			db:       inMemoryDB,
			topoUuid: "top-tpd",
			nodeUuid: "node-tpd-1",
			checker: func(t *testing.T, res models.TapiTopologyNode, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
				require.Equal(t, "node-tpd-1", res.Uuid)
				require.Len(t, res.OwnedNodeEdgePoint, 3)
			},
		},
		{
			name:     "topo not found",
			ctx:      ctx,
			db:       inMemoryDB,
			topoUuid: "temp",
			nodeUuid: "node-tpd-1",
			checker: func(t *testing.T, res models.TapiTopologyNode, err error) {
				require.EqualError(t, err, TopologyNotFoundErr)
				require.Empty(t, res)
			},
		},
		{
			name:     "node not found",
			ctx:      ctx,
			db:       inMemoryDB,
			topoUuid: "top-tpd",
			nodeUuid: "temp",
			checker: func(t *testing.T, res models.TapiTopologyNode, err error) {
				require.EqualError(t, err, NodeNotFoundErr)
				require.Empty(t, res)
			},
		},
	}
	for _, testCase := range testCases {
		res, err := GetNodeUuid(testCase.ctx, testCase.topoUuid, testCase.nodeUuid, testCase.db)
		testCase.checker(T, res, err)
	}
}
