package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"tapi-controller/token"
	"tapi-controller/utils"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type addAuthParams struct {
	t                 *testing.T
	request           *http.Request
	tokenMaker        token.TokenMaker
	authorizationType string
	username          string
	duration          time.Duration
}

func addAuthoization(params *addAuthParams) {
	fmt.Println("add auth")
	fmt.Printf("username: %s", params.username)
	token, _ := params.tokenMaker.CreateToken(params.username, params.duration)
	authorizationHeader := fmt.Sprintf("%s %s", params.authorizationType, token)
	params.request.Header.Set(authorizationHeaderKey, authorizationHeader)

}

func TestAuthMiddleare(t *testing.T) {
	inMemoryDB := utils.InitMemoryDb()
	config, _ := utils.LoadConfigForTest()

	testCases := []struct {
		name      string
		setupAuth func(t *testing.T, request http.Request, tokenMaker token.TokenMaker)
		checker   func(t *testing.T, statusCode int, res *bytes.Buffer)
	}{
		{
			name: "good day",
			setupAuth: func(t *testing.T, request http.Request, tokenMaker token.TokenMaker) {
				addAuthoization(&addAuthParams{
					t:                 t,
					request:           &request,
					tokenMaker:        tokenMaker,
					authorizationType: authorizationTypeBearer,
					username:          "test",
					duration:          time.Second * 20,
				})
			},
			checker: func(t *testing.T, statusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusOK, statusCode)
			},
		},
		{
			name:      "unautorised",
			setupAuth: func(t *testing.T, request http.Request, tokenMaker token.TokenMaker) {},
			checker: func(t *testing.T, statusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
			},
		},
		{
			name: "authorization type incorrect",
			setupAuth: func(t *testing.T, request http.Request, tokenMaker token.TokenMaker) {
				addAuthoization(&addAuthParams{
					t:                 t,
					request:           &request,
					tokenMaker:        tokenMaker,
					authorizationType: "",
					username:          "test",
					duration:          time.Second * 20,
				})
			},
			checker: func(t *testing.T, statusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
			},
		},
		{
			name: "unsupported authorization",
			setupAuth: func(t *testing.T, request http.Request, tokenMaker token.TokenMaker) {
				addAuthoization(&addAuthParams{
					t:                 t,
					request:           &request,
					tokenMaker:        tokenMaker,
					authorizationType: "unsupported",
					username:          "test",
					duration:          time.Second * 20,
				})
			},
			checker: func(t *testing.T, statusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
			},
		},
		{
			name: "token expired",
			setupAuth: func(t *testing.T, request http.Request, tokenMaker token.TokenMaker) {
				addAuthoization(&addAuthParams{
					t:                 t,
					request:           &request,
					tokenMaker:        tokenMaker,
					authorizationType: authorizationTypeBearer,
					username:          "test",
					duration:          -time.Minute,
				})
			},
			checker: func(t *testing.T, statusCode int, res *bytes.Buffer) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			// Create test server
			server, err := InitServer(config, inMemoryDB)
			require.NoError(t, err)
			authPath := "/auth"
			server.engine.GET(
				authPath,
				authMiddleware(server.token),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, *request, server.token)
			server.engine.ServeHTTP(recorder, request)
			tc.checker(t, recorder.Code, recorder.Body)
		})
	}
}
