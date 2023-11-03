package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"racers-chain/x/racerschain/types"
)

func (k msgServer) BankReplenishment(goCtx context.Context, msg *types.MsgBankReplenishment) (*types.MsgBankReplenishmentResponse, error) {
	//TODO: использовать методы Burn и Mint у банка
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, amount)
	if sdkError != nil {
		return nil, sdkError
	}

	return &types.MsgBankReplenishmentResponse{}, nil
}
