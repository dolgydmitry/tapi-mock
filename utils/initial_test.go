package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadKeys(t *testing.T) {
	config, err := LoadConfigForTest()
	require.NoError(t, err)
	tokenMaker, err := LoadKeys(config)
	require.NoError(t, err)
	require.IsType(t, &Keys{}, tokenMaker)
}
