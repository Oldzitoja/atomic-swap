package xmrtaker

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/athanorlabs/atomic-swap/common/types"
	mcrypto "github.com/athanorlabs/atomic-swap/crypto/monero"
	"github.com/athanorlabs/atomic-swap/db"
	contracts "github.com/athanorlabs/atomic-swap/ethereum"
	"github.com/athanorlabs/atomic-swap/protocol/backend"
	pswap "github.com/athanorlabs/atomic-swap/protocol/swap"
)

func newTestInstance(t *testing.T) *Instance {
	inst, err := NewInstance(&Config{
		Backend:        newBackend(t),
		DataDir:        "",
		TransferBack:   true,
		ExternalSender: false,
	})
	require.NoError(t, err)
	return inst
}

func TestNewInstance(t *testing.T) {
	inst := newTestInstance(t)
	assert.Nil(t, inst.GetOngoingSwapState(types.EmptyHash))
	assert.Equal(t, inst.Provides(), types.ProvidesETH)
	_, err := inst.Refund(types.EmptyHash)
	assert.ErrorIs(t, err, errNoOngoingSwap)
}

func TestInstance_createOngoingSwap(t *testing.T) {
	inst := newTestInstance(t)
	rdb := inst.backend.RecoveryDB().(*backend.MockRecoveryDB)

	offer := types.NewOffer(
		types.ProvidesXMR,
		1,
		1,
		1,
		types.EthAssetETH,
	)

	s := &pswap.Info{
		ID:             offer.ID,
		Provides:       types.ProvidesXMR,
		ProvidedAmount: 1,
		ReceivedAmount: 1,
		ExchangeRate:   types.ExchangeRate(1),
		EthAsset:       types.EthAssetETH,
		Status:         types.ETHLocked,
	}

	sk, err := mcrypto.GenerateKeys()
	require.NoError(t, err)

	makerKeys, err := mcrypto.GenerateKeys()
	require.NoError(t, err)

	rdb.EXPECT().GetSharedSwapPrivateKey(s.ID).Return(nil, errors.New("some error"))
	rdb.EXPECT().GetContractSwapInfo(s.ID).Return(&db.EthereumSwapInfo{
		StartNumber:     big.NewInt(1),
		ContractAddress: inst.backend.ContractAddr(),
		Swap: contracts.SwapFactorySwap{
			Timeout0: big.NewInt(1),
			Timeout1: big.NewInt(2),
		},
	}, nil)
	rdb.EXPECT().GetSwapPrivateKey(s.ID).Return(
		sk.SpendKey(), nil,
	)
	rdb.EXPECT().GetXMRMakerSwapKeys(s.ID).Return(
		makerKeys.SpendKey().Public(), makerKeys.ViewKey(), nil,
	)

	err = inst.createOngoingSwap(*s)
	require.NoError(t, err)

	inst.swapMu.Lock()
	defer inst.swapMu.Unlock()
	close(inst.swapStates[s.ID].done)
}
