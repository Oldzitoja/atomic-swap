package xmrtaker

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/atomic-swap/common/types"
)

func newTestXMRTaker(t *testing.T) *Instance {
	b := newBackend(t)
	cfg := &Config{
		Backend: b,
		DataDir: path.Join(t.TempDir(), "xmrtaker"),
	}

	xmrtaker, err := NewInstance(cfg)
	require.NoError(t, err)
	return xmrtaker
}

func TestXMRTaker_InitiateProtocol(t *testing.T) {
	a := newTestXMRTaker(t)
	offer := types.NewOffer(types.ProvidesETH, 0, 0, 1, types.EthAssetETH)
	s, err := a.InitiateProtocol(3.33, offer)
	require.NoError(t, err)
	require.Equal(t, a.swapStates[offer.ID], s)
}
