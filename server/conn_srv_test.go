package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"tapi-controller/controllers"
	"tapi-controller/models"
	"tapi-controller/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateConnSrvRoute(t *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	config, _ := utils.LoadConfig()

	testCases := []struct {
		name    string
		req     models.TapiCommonContext
		pattern controllers.CreateConSrvResponse
		checker func(t *testing.T, statusCode int, req models.TapiCommonContext, res *bytes.Buffer, pattern controllers.CreateConSrvResponse)
	}{
		{
			name: "good day",
			req: models.TapiCommonContext{
				ConnectivityContext: models.TapiConnectivityConnectivityContext{
					ConnectivityService: []models.TapiConnectivityConnectivityService{
						{
							Uuid: "test-1",
							Name: []models.TapiCommonNameAndValue{
								{
									ValueName: "SERVICE_NAME",
									Value:     "linevo-Iskitim-100G",
								},
							},
							AdministrativeState:   "UNLOCKED",
							ServiceLayer:          "DSR",
							ConnectivityDirection: "BIDIRECTIONAL",
							ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
							Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-a",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-z",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
							},
						},
					},
				},
			},
			pattern: controllers.CreateConSrvResponse{
				DesiredOrchState: controllers.DesiredOrchStateActive,
				OrchState:        controllers.OrchStateRequsted,
			},
			checker: func(t *testing.T, statusCode int, req models.TapiCommonContext, res *bytes.Buffer, pattern controllers.CreateConSrvResponse) {
				require.Equal(t, http.StatusCreated, statusCode)
				var resObject controllers.CreateConSrvResponse
				byteValue, err := io.ReadAll(res)
				require.NoError(t, err)
				json.Unmarshal(byteValue, &resObject)
				require.NotEmpty(t, resObject)
				require.Equal(t, pattern.DesiredOrchState, resObject.DesiredOrchState)
				require.Equal(t, pattern.OrchState, resObject.OrchState)
			},
		},
		{
			name: "internal service error",
			pattern: controllers.CreateConSrvResponse{
				DesiredOrchState: controllers.DesiredOrchStateActive,
				OrchState:        controllers.OrchStateRequsted,
			},
			checker: func(t *testing.T, statusCode int, req models.TapiCommonContext, res *bytes.Buffer, pattern controllers.CreateConSrvResponse) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				var resObject interface{}
				byteValue, err := io.ReadAll(res)
				require.NoError(t, err)
				json.Unmarshal(byteValue, &resObject)
				resErr, ok := resObject.(map[string]any)
				require.True(t, ok)
				require.Equal(t, controllers.EmptyConnSrvList, resErr["error"])
			},
		},
		{
			name: "bad request",
			pattern: controllers.CreateConSrvResponse{
				DesiredOrchState: controllers.DesiredOrchStateActive,
				OrchState:        controllers.OrchStateRequsted,
			},
			checker: func(t *testing.T, statusCode int, req models.TapiCommonContext, res *bytes.Buffer, pattern controllers.CreateConSrvResponse) {
				require.Equal(t, http.StatusBadRequest, statusCode)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			// Create test server
			server, err := InitServer(config, &inMemoryDB)
			server.AddRoute()
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			url := "/tapi-common:context/tapi-connectivity:connectivity-context/"

			inputBody, err := json.Marshal(tc.req)
			require.NoError(t, err)

			if tc.name == "bad request" {
				request, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)
				server.engine.ServeHTTP(recorder, request)
			} else {
				request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(inputBody))
				require.NoError(t, err)
				server.engine.ServeHTTP(recorder, request)
			}

			// check response
			tc.checker(t, recorder.Code, tc.req, recorder.Body, tc.pattern)
		})
	}
}

func TestGetAllConnSrvRoute(t *testing.T) {
	config, _ := utils.LoadConfig()

	testCases := []struct {
		name     string
		prepare  bool
		testLoad models.TapiCommonContext
		checker  func(t *testing.T, respStatusCode int, res *bytes.Buffer)
	}{
		{
			name:    "good day",
			prepare: true,
			testLoad: models.TapiCommonContext{
				ConnectivityContext: models.TapiConnectivityConnectivityContext{
					ConnectivityService: []models.TapiConnectivityConnectivityService{
						{
							Uuid: "test-1",
							Name: []models.TapiCommonNameAndValue{
								{
									ValueName: "SERVICE_NAME",
									Value:     "linevo-Iskitim-100G",
								},
							},
							AdministrativeState:   "UNLOCKED",
							ServiceLayer:          "DSR",
							ConnectivityDirection: "BIDIRECTIONAL",
							ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
							Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-a",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-z",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
							},
						},
						{
							Uuid: "test-2",
							Name: []models.TapiCommonNameAndValue{
								{
									ValueName: "SERVICE_NAME",
									Value:     "linevo-Iskitim-100G",
								},
							},
							AdministrativeState:   "UNLOCKED",
							ServiceLayer:          "DSR",
							ConnectivityDirection: "BIDIRECTIONAL",
							ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
							Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-a",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-z",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
							},
						},
						{
							Uuid: "test-3",
							Name: []models.TapiCommonNameAndValue{
								{
									ValueName: "SERVICE_NAME",
									Value:     "linevo-Iskitim-100G",
								},
							},
							AdministrativeState:   "UNLOCKED",
							ServiceLayer:          "DSR",
							ConnectivityDirection: "BIDIRECTIONAL",
							ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
							Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-a",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
								{
									Role: "SYMMETRIC",
									ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
										ServiceInterfacePointUuid: "test-ref-endpoint-z",
									},
									LayerProtocolName:      "DSR",
									LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
									Direction:              "BIDIRECTIONAL",
								},
							},
						},
					},
				},
			},
			checker: func(t *testing.T, respStatusCode int, res *bytes.Buffer) {
				require.Equal(t, respStatusCode, http.StatusOK)
				var resList []models.TapiConnectivityConnectivityService
				byteValue, err := io.ReadAll(res)
				json.Unmarshal(byteValue, &resList)
				require.NoError(t, err)
				require.NotEmpty(t, resList)
				require.Len(t, resList, 3)
				for _, i := range resList {
					require.NotEmpty(t, i)
				}
			},
		},
		{
			name:    "not found",
			prepare: false,
			checker: func(t *testing.T, respStatusCode int, res *bytes.Buffer) {
				require.Equal(t, respStatusCode, http.StatusOK)
				var resList []models.TapiConnectivityConnectivityService
				byteValue, err := io.ReadAll(res)
				json.Unmarshal(byteValue, &resList)
				require.NoError(t, err)
				require.Len(t, resList, 0)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			inMemoryDB := utils.InitMemoryDb()

			// Create test server
			server, err := InitServer(config, &inMemoryDB)
			if tc.prepare {
				server.tapiCtrl.CreateConSrv(context.Background(), tc.testLoad, &inMemoryDB)
			}
			server.AddRoute()
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			url := "/tapi-common:context/tapi-connectivity:connectivity-context/"

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.engine.ServeHTTP(recorder, request)

			// check response
			tc.checker(t, recorder.Code, recorder.Body)
		})
	}
}

func TestGetConnSrvRoute(t *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	config, _ := utils.LoadConfig()
	testLoad := models.TapiCommonContext{
		ConnectivityContext: models.TapiConnectivityConnectivityContext{
			ConnectivityService: []models.TapiConnectivityConnectivityService{
				{
					Uuid: "test-1",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
				{
					Uuid: "test-2",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
				{
					Uuid: "test-3",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
			},
		},
	}
	ctrl := new(controllers.TapiCtrlInMemDB)
	ctrl.CreateConSrv(context.Background(), testLoad, &inMemoryDB)

	testCases := []struct {
		name    string
		uuid    string
		checker func(t *testing.T, uuid string, respStatusCode int, res *bytes.Buffer)
	}{
		{
			name: "good day",
			uuid: "test-1",
			checker: func(t *testing.T, uuid string, respStatusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusOK, respStatusCode)
				var resObj models.TapiConnectivityConnectivityService
				byteValue, err := io.ReadAll(res)
				json.Unmarshal(byteValue, &resObj)
				require.NoError(t, err)
				require.Equal(t, uuid, resObj.Uuid)
			},
		},
		{
			name: "not found",
			uuid: "wrong",
			checker: func(t *testing.T, uuid string, respStatusCode int, res *bytes.Buffer) {
				require.Equal(t, respStatusCode, http.StatusNotFound)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		log.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {

			// Create test server
			server, err := InitServer(config, &inMemoryDB)
			server.AddRoute()
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/tapi-common:context/tapi-connectivity:connectivity-context/%s", tc.uuid)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.engine.ServeHTTP(recorder, request)

			// check response
			tc.checker(t, tc.uuid, recorder.Code, recorder.Body)
		})
	}
}

func TestUpdateConSrvRoute(t *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	config, _ := utils.LoadConfig()
	testLoad := models.TapiCommonContext{
		ConnectivityContext: models.TapiConnectivityConnectivityContext{
			ConnectivityService: []models.TapiConnectivityConnectivityService{
				{
					Uuid: "test-1",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
				{
					Uuid: "test-2",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
				{
					Uuid: "test-3",
					Name: []models.TapiCommonNameAndValue{
						{
							ValueName: "SERVICE_NAME",
							Value:     "linevo-Iskitim-100G",
						},
					},
					AdministrativeState:   "UNLOCKED",
					ServiceLayer:          "DSR",
					ConnectivityDirection: "BIDIRECTIONAL",
					ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
					Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-a",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
						{
							Role: "SYMMETRIC",
							ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
								ServiceInterfacePointUuid: "test-ref-endpoint-z",
							},
							LayerProtocolName:      "DSR",
							LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
							Direction:              "BIDIRECTIONAL",
						},
					},
				},
			},
		},
	}
	ctrl := new(controllers.TapiCtrlInMemDB)
	ctrl.CreateConSrv(context.Background(), testLoad, &inMemoryDB)

	testCases := []struct {
		name    string
		uuid    string
		req     models.TapiConnectivityConnectivityServiceData
		checker func(t *testing.T, uuid string, respStatusCode int, inputData models.TapiConnectivityConnectivityServiceData, res *bytes.Buffer)
	}{
		{
			name: "good day",
			uuid: "test-1",
			req: models.TapiConnectivityConnectivityServiceData{
				Name: []models.TapiCommonNameAndValue{
					{
						ValueName: "SERVICE_NAME",
						Value:     "linevo-Iskitim-100G",
					},
				},
				AdministrativeState:   "LOCKED",
				ServiceLayer:          "DSR",
				ConnectivityDirection: "BIDIRECTIONAL",
				ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
				Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
					{
						Role: "SYMMETRIC",
						ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
							ServiceInterfacePointUuid: "test-ref-endpoint-a",
						},
						LayerProtocolName:      "DSR",
						LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
						Direction:              "BIDIRECTIONAL",
					},
					{
						Role: "SYMMETRIC",
						ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
							ServiceInterfacePointUuid: "test-ref-endpoint-z",
						},
						LayerProtocolName:      "DSR",
						LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
						Direction:              "BIDIRECTIONAL",
					},
				},
			},
			checker: func(t *testing.T, uuid string, respStatusCode int, inputData models.TapiConnectivityConnectivityServiceData, res *bytes.Buffer) {
				require.Equal(t, http.StatusOK, respStatusCode)
				var resObject models.TapiConnectivityConnectivityService
				byteValue, err := io.ReadAll(res)
				require.NoError(t, err)
				json.Unmarshal(byteValue, &resObject)
				require.NotEmpty(t, resObject)
				require.Equal(t, inputData.AdministrativeState, resObject.AdministrativeState)
			},
		},
		{
			name: "not found",
			uuid: "test-wrong",
			req: models.TapiConnectivityConnectivityServiceData{
				Name: []models.TapiCommonNameAndValue{
					{
						ValueName: "SERVICE_NAME",
						Value:     "linevo-Iskitim-100G",
					},
				},
				AdministrativeState:   "LOCKED",
				ServiceLayer:          "DSR",
				ConnectivityDirection: "BIDIRECTIONAL",
				ServiceType:           "POINT_TO_POINT_CONNECTIVITY",
				Endpoint: []models.TapiConnectivityConnectivityserviceEndPoint{
					{
						Role: "SYMMETRIC",
						ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
							ServiceInterfacePointUuid: "test-ref-endpoint-a",
						},
						LayerProtocolName:      "DSR",
						LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
						Direction:              "BIDIRECTIONAL",
					},
					{
						Role: "SYMMETRIC",
						ServiceInterfacePoint: models.TapiCommonServiceInterfacePointRef{
							ServiceInterfacePointUuid: "test-ref-endpoint-z",
						},
						LayerProtocolName:      "DSR",
						LayerProtocolQualifier: "tapi-dsr:DIGITAL_SIGNAL_TYPE_10_GigE_LAN",
						Direction:              "BIDIRECTIONAL",
					},
				},
			},
			checker: func(t *testing.T, uuid string, respStatusCode int, inputData models.TapiConnectivityConnectivityServiceData, res *bytes.Buffer) {
				require.Equal(t, http.StatusNotFound, respStatusCode)
			},
		},
		{
			name: "bad request",
			uuid: "test-wrong",
			checker: func(t *testing.T, uuid string, respStatusCode int, inputData models.TapiConnectivityConnectivityServiceData, res *bytes.Buffer) {
				require.Equal(t, http.StatusBadRequest, respStatusCode)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		log.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {

			// Create test server
			server, err := InitServer(config, &inMemoryDB)
			server.AddRoute()
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			inputBody, err := json.Marshal(tc.req)
			require.NoError(t, err)

			url := fmt.Sprintf("/tapi-common:context/tapi-connectivity:connectivity-context/%s", tc.uuid)
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(inputBody))
			require.NoError(t, err)

			server.engine.ServeHTTP(recorder, request)

			// check response
			tc.checker(t, tc.uuid, recorder.Code, tc.req, recorder.Body)
		})
	}
}
