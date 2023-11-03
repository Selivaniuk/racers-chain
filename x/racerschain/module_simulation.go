package racerschain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"racers-chain/testutil/sample"
	racerschainsimulation "racers-chain/x/racerschain/simulation"
	"racers-chain/x/racerschain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = racerschainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgBankReplenishment = "op_weight_msg_bank_replenishment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBankReplenishment int = 100

	opWeightMsgWithdrawBank = "op_weight_msg_withdraw_bank"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawBank int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	racerschainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&racerschainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgBankReplenishment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBankReplenishment, &weightMsgBankReplenishment, nil,
		func(_ *rand.Rand) {
			weightMsgBankReplenishment = defaultWeightMsgBankReplenishment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBankReplenishment,
		racerschainsimulation.SimulateMsgBankReplenishment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawBank int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawBank, &weightMsgWithdrawBank, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawBank = defaultWeightMsgWithdrawBank
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawBank,
		racerschainsimulation.SimulateMsgWithdrawBank(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgBankReplenishment,
			defaultWeightMsgBankReplenishment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				racerschainsimulation.SimulateMsgBankReplenishment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdrawBank,
			defaultWeightMsgWithdrawBank,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				racerschainsimulation.SimulateMsgWithdrawBank(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
