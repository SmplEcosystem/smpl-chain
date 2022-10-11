package simplusd

import (
	"math/rand"

	"github.com/Smpl-Finance/smpl-chain/testutil/sample"
	simplusdsimulation "github.com/Smpl-Finance/smpl-chain/x/simplusd/simulation"
	"github.com/Smpl-Finance/smpl-chain/x/simplusd/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = simplusdsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgBurnsusd = "op_weight_msg_burnsusd"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnsusd int = 100

	opWeightMsgMintsusd = "op_weight_msg_mintsusd"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintsusd int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	simplusdGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&simplusdGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgBurnsusd int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnsusd, &weightMsgBurnsusd, nil,
		func(_ *rand.Rand) {
			weightMsgBurnsusd = defaultWeightMsgBurnsusd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnsusd,
		simplusdsimulation.SimulateMsgBurnsusd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintsusd int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintsusd, &weightMsgMintsusd, nil,
		func(_ *rand.Rand) {
			weightMsgMintsusd = defaultWeightMsgMintsusd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintsusd,
		simplusdsimulation.SimulateMsgMintsusd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
