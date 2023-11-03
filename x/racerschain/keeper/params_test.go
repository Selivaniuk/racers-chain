package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "racers-chain/testutil/keeper"
	"racers-chain/x/racerschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RacerschainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
