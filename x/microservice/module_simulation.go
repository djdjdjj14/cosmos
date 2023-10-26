package microservice

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"microservice/testutil/sample"
	microservicesimulation "microservice/x/microservice/simulation"
	"microservice/x/microservice/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = microservicesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateLog = "op_weight_msg_log"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLog int = 100

	opWeightMsgUpdateLog = "op_weight_msg_log"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLog int = 100

	opWeightMsgDeleteLog = "op_weight_msg_log"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLog int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	microserviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		LogList: []types.Log{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		LogCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&microserviceGenesis)
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

	var weightMsgCreateLog int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLog, &weightMsgCreateLog, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLog = defaultWeightMsgCreateLog
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLog,
		microservicesimulation.SimulateMsgCreateLog(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLog int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLog, &weightMsgUpdateLog, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLog = defaultWeightMsgUpdateLog
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLog,
		microservicesimulation.SimulateMsgUpdateLog(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLog int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteLog, &weightMsgDeleteLog, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLog = defaultWeightMsgDeleteLog
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLog,
		microservicesimulation.SimulateMsgDeleteLog(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateLog,
			defaultWeightMsgCreateLog,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				microservicesimulation.SimulateMsgCreateLog(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLog,
			defaultWeightMsgUpdateLog,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				microservicesimulation.SimulateMsgUpdateLog(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteLog,
			defaultWeightMsgDeleteLog,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				microservicesimulation.SimulateMsgDeleteLog(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
