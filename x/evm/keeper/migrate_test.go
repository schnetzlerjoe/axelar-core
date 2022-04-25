package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/axelarnetwork/axelar-core/app/params"
	"github.com/axelarnetwork/axelar-core/testutils/fake"
	"github.com/axelarnetwork/axelar-core/testutils/rand"
	"github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/axelar-core/x/evm/types/mock"
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

func setup() (sdk.Context, types.BaseKeeper) {
	encCfg := params.MakeEncodingConfig()
	paramsK := paramsKeeper.NewKeeper(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("params"), sdk.NewKVStoreKey("tparams"))
	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
	keeper := NewKeeper(encCfg.Codec, sdk.NewKVStoreKey("evm"), paramsK)

	return ctx, keeper
}

func TestGetMigrationHandler_deleteUaxlToken(t *testing.T) {
	ctx, keeper := setup()
	evmChains := []exported.Chain{
		{
			Name:   "evm-1",
			Module: types.ModuleName,
		},
		{
			Name:   "evm-2",
			Module: types.ModuleName,
		},
	}

	nexus := mock.NexusMock{
		GetChainsFunc: func(_ sdk.Context) []exported.Chain {
			return evmChains
		},
	}

	uaxlToken := types.ERC20TokenMetadata{
		Asset: "uaxl",
		Details: types.TokenDetails{
			Symbol: rand.NormalizedStr(5),
		},
		Status: types.Confirmed,
	}
	otherToken := types.ERC20TokenMetadata{
		Asset: rand.NormalizedStr(5),
		Details: types.TokenDetails{
			Symbol: rand.NormalizedStr(5),
		},
		Status: types.Confirmed,
	}

	for _, chain := range evmChains {
		keeper.ForChain(chain.Name).(chainKeeper).setTokenMetadata(ctx, uaxlToken)
		keeper.ForChain(chain.Name).(chainKeeper).setTokenMetadata(ctx, otherToken)

		_, ok := keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataByAsset(ctx, uaxlToken.Asset)
		assert.True(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataBySymbol(ctx, uaxlToken.Details.Symbol)
		assert.True(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataByAsset(ctx, otherToken.Asset)
		assert.True(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataBySymbol(ctx, otherToken.Details.Symbol)
		assert.True(t, ok)
	}

	handler := GetMigrationHandler(keeper, &nexus)
	err := handler(ctx)

	assert.NoError(t, err)

	for _, chain := range evmChains {
		_, ok := keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataByAsset(ctx, uaxlToken.Asset)
		assert.False(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataBySymbol(ctx, uaxlToken.Details.Symbol)
		assert.False(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataByAsset(ctx, otherToken.Asset)
		assert.True(t, ok)
		_, ok = keeper.ForChain(chain.Name).(chainKeeper).getTokenMetadataBySymbol(ctx, otherToken.Details.Symbol)
		assert.True(t, ok)
	}
}
