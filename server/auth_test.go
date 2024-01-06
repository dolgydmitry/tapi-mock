package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"tapi-controller/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	config, _ := utils.LoadConfigForTest()

	testCases := []struct {
		name    string
		payload map[string]string
		checker func(T *testing.T, statusCode int, resp *bytes.Buffer)
	}{
		{
			name:    "good day",
			payload: map[string]string{"username": "test", "password": "test"},
			checker: func(t *testing.T, statusCode int, resp *bytes.Buffer) {
				require.Equal(t, http.StatusOK, statusCode)
				var respToken string
				byteValue, err := io.ReadAll(resp)
				require.NoError(t, err)
				json.Unmarshal(byteValue, &respToken)
				require.NotEmpty(t, respToken)
			},
		},
		{
			name:    "user unautorised",
			payload: map[string]string{"username": "test", "password": "wrong password"},
			checker: func(t *testing.T, statusCode int, resp *bytes.Buffer) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
			},
		},
		{
			name:    "bad request",
			payload: map[string]string{"temp": "test", "password": "wrong password"},
			checker: func(t *testing.T, statusCode int, resp *bytes.Buffer) {
				require.Equal(t, http.StatusBadRequest, statusCode)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			// Create test server
			server, err := InitServer(config, inMemoryDB)
			server.AddRoute()
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			url := "/login/"

			inputBody, err := json.Marshal(tc.payload)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(inputBody))
			require.NoError(t, err)
			server.engine.ServeHTTP(recorder, request)

			// check response
			tc.checker(t, recorder.Code, recorder.Body)
		})
	}

}
