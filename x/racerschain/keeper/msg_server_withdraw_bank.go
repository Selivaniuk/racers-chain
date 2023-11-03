package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"racers-chain/x/racerschain/types"
)

func (k msgServer) WithdrawBank(goCtx context.Context, msg *types.MsgWithdrawBank) (*types.MsgWithdrawBankResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, amount)

	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawBankResponse{}, nil
}
