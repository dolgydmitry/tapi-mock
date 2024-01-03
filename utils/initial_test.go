package utils

import (
	"tapi-controller/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitTokenMaker(t *testing.T) {
	config, err := LoadConfigForTest()
	require.NoError(t, err)
	tokenMaker, err := InitTokenMaker(config)
	require.NoError(t, err)
	require.IsType(t, &token.JwtMaker{}, tokenMaker)
}
