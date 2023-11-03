package keeper

import (
	"racers-chain/x/racerschain/types"
)

var _ types.QueryServer = Keeper{}
