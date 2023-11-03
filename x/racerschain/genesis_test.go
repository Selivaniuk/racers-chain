package racerschain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "racers-chain/testutil/keeper"
	"racers-chain/testutil/nullify"
	"racers-chain/x/racerschain"
	"racers-chain/x/racerschain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RacerschainKeeper(t)
	racerschain.InitGenesis(ctx, *k, genesisState)
	got := racerschain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
