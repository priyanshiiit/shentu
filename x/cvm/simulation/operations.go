package simulation

import (
	"encoding/hex"
	"math/rand"
	"strconv"

	"github.com/hyperledger/burrow/txs/payload"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/certikfoundation/shentu/v2/x/cvm/keeper"
	"github.com/certikfoundation/shentu/v2/x/cvm/types"
)

const (
	OpWeightMsgDeploy = "op_weight_msg_deploy"
)

// WeightedOperations creates an operation with a weight for each type of message generators.
func WeightedOperations(appParams sim.AppParams, cdc codec.JSONCodec, k keeper.Keeper, bk types.BankKeeper) simulation.WeightedOperations {
	var weightMsgDeploy int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeploy, &weightMsgDeploy, nil,
		func(_ *rand.Rand) {
			weightMsgDeploy = simappparams.DefaultWeightMsgSend
		})

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeployHello55(k, bk)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeploySimple(k, bk)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeploySimpleEvent(k, bk)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeployStorage(k, bk)),
		simulation.NewWeightedOperation(weightMsgDeploy, SimulateMsgDeployStringTest(k, bk)),
	}
}

// SimulateMsgDeployHello55 creates a massage deploying /tests/hello55.sol contract.
func SimulateMsgDeployHello55(k keeper.Keeper, bk types.BankKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// deploy hello55.sol
		msg, contractAddr, err := DeployContract(caller, Hello55Code, Hello55Abi, k, bk, r, ctx, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}

		// check sayHi() ret
		data, err := hex.DecodeString(Hello55SayHi)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 32)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		if value != 55 {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgDeploySimple creates a massage deploying /tests/simple.sol contract.
func SimulateMsgDeploySimple(k keeper.Keeper, bk types.BankKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// deploy simple.sol
		msg, contractAddr, err := DeployContract(caller, SimpleCode, SimpleAbi, k, bk, r, ctx, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleGet)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		futureOperations := []sim.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallSimpleSet(k, bk, contractAddr, int(r.Uint32())),
			},
		}

		return sim.NewOperationMsg(&msg, true, "", nil), futureOperations, nil
	}
}

// SimulateMsgCallSimpleSet creates a message calling set() in /tests/simple.sol contract.
func SimulateMsgCallSimpleSet(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, varValue int) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}

		// call set()
		msg, _, err := CallFunction(caller, SimpleSetPrefix, hexStr, contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleGet)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgDeploySimpleEvent creates a massage deploying /tests/simpleevent.sol contract.
func SimulateMsgDeploySimpleEvent(k keeper.Keeper, bk types.BankKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// deploy simpleevent.sol
		msg, contractAddr, err := DeployContract(caller, SimpleeventCode, SimpleeventAbi, k, bk, r, ctx, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleeventGet)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		futureOperations := []sim.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallSimpleEventSet(k, bk, contractAddr, int(r.Uint32())),
			},
		}

		return sim.NewOperationMsg(&msg, true, "", nil), futureOperations, nil
	}
}

// SimulateMsgCallSimpleEventSet creates a message calling set() in /tests/simpleevent.sol contract.
func SimulateMsgCallSimpleEventSet(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, varValue int) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}

		// call set()
		msg, _, err := CallFunction(caller, SimpleeventSetPrefix, hexStr, contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check get() ret
		data, err := hex.DecodeString(SimpleeventGet)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgDeployStorage creates a massage deploying /tests/storage.sol contract.
func SimulateMsgDeployStorage(k keeper.Keeper, bk types.BankKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// deploy storage.sol
		msg, contractAddr, err := DeployContract(caller, StorageCode, StorageAbi, k, bk, r, ctx, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}

		// check retrieve() ret
		data, err := hex.DecodeString(StorageRetrieve)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		if value != 0 {
			panic("return value incorrect")
		}

		// check sayMyAddres() ret
		data, err = hex.DecodeString(StorageSayMyAddres)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		ret, err = k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}
		sender := sdk.AccAddress(ret[12:])
		if !sender.Equals(caller.Address) {
			panic("return value incorrect")
		}

		futureOperations := []sim.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallStorageStore(k, bk, contractAddr, int(r.Uint32())),
			},
		}

		return sim.NewOperationMsg(&msg, true, "", nil), futureOperations, nil
	}
}

// SimulateMsgCallStorageStore creates a message calling store() in /tests/storage.sol contract.
func SimulateMsgCallStorageStore(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, varValue int) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		hexStr := strconv.FormatInt(int64(varValue), 16)
		length := len(hexStr)
		for i := 0; i < 64-length; i++ {
			hexStr = "0" + hexStr
		}

		// call store()
		msg, _, err := CallFunction(caller, StorageStorePrefix, hexStr, contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check retrieve() ret
		data, err := hex.DecodeString(StorageRetrieve)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		ret, err := k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 64)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		if value != int64(varValue) {
			panic("return value incorrect")
		}

		// check sayMyAddres() ret
		data, err = hex.DecodeString(StorageSayMyAddres)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		ret, err = k.Tx(ctx, caller.Address, contractAddr, 0, data, []*payload.ContractMeta{}, true, false, false)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		sender := sdk.AccAddress(ret[12:])
		if !sender.Equals(caller.Address) {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgDeployStringTest creates a massage deploying /tests/stringtest.sol contract.
func SimulateMsgDeployStringTest(k keeper.Keeper, bk types.BankKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// deploy stringtest.sol
		msg, contractAddr, err := DeployContract(caller, StringtestCode, StringtestAbi, k, bk, r, ctx, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgDeploy, ""), nil, err
		}

		var ref string // hex str shared among future operations for checking purpose

		futureOperations := []sim.FutureOperation{
			{
				BlockHeight: int(ctx.BlockHeight()) + 1,
				Op:          SimulateMsgCallStringTestGetl(k, bk, contractAddr, &ref),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + 1,
				Op:          SimulateMsgCallStringTestGets(k, bk, contractAddr, &ref),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + 2,
				Op:          SimulateMsgCallStringTestChangeString(k, bk, contractAddr, &ref),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + 3,
				Op:          SimulateMsgCallStringTestGets(k, bk, contractAddr, &ref),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + 3,
				Op:          SimulateMsgCallStringTestGetl(k, bk, contractAddr, &ref),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallStringTestChangeGiven(k, bk, contractAddr),
			},
			{
				BlockHeight: int(ctx.BlockHeight()) + r.Intn(10),
				Op:          SimulateMsgCallStringTestTestStuff(k, bk, contractAddr),
			},
		}

		return sim.NewOperationMsg(&msg, true, "", nil), futureOperations, nil
	}
}

// SimulateMsgCallStringTestChangeString creates a message calling changeString() in /tests/stringtest.sol contract.
func SimulateMsgCallStringTestChangeString(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, ref *string) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// turn length into a hex string of length 64
		length := r.Intn(32) + 1
		hexLen := strconv.FormatInt(int64(length), 16)
		l := len(hexLen)
		for i := 0; i < 64-l; i++ {
			hexLen = "0" + hexLen
		}

		// turn string into a hex string of length 64
		*ref = sim.RandStringOfLength(r, length)
		hexStr := hex.EncodeToString([]byte(*ref))
		l = len(hexStr)
		for i := 0; i < 64-l; i++ {
			hexStr = hexStr + "0"
		}

		// call changeString()
		msg, ret, err := CallFunction(caller, StringtestChangeStringPrefix, hexLen+hexStr, contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check ret and update ref
		*ref = StringtestChangeStringPrefix[8:] + hexLen + hexStr
		if hex.EncodeToString(ret) != *ref {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgCallStringTestChangeGiven creates a message calling changeGiven() in /tests/stringtest.sol contract.
func SimulateMsgCallStringTestChangeGiven(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// assemble length into a hex string of length 64
		length := r.Intn(30) + 3
		hexLen := strconv.FormatInt(int64(length), 16)
		l := len(hexLen)
		for i := 0; i < 64-l; i++ {
			hexLen = "0" + hexLen
		}

		// assemble string into a hex string of length 64
		str := sim.RandStringOfLength(r, length)
		hexStr := hex.EncodeToString([]byte(str))
		l = len(hexStr)
		for i := 0; i < 64-l; i++ {
			hexStr = hexStr + "0"
		}

		// call changeGiven()
		msg, ret, err := CallFunction(caller, StringtestChangeGivenPrefix, hexLen+hexStr, contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check ret
		ref := StringtestChangeStringPrefix[8:] + hexLen + hex.EncodeToString([]byte("Abc")) + hexStr[6:]
		if hex.EncodeToString(ret) != ref {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgCallStringTestGets creates a message calling gets() in /tests/stringtest.sol contract.
func SimulateMsgCallStringTestGets(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, ref *string) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// call gets()
		msg, ret, err := CallFunction(caller, StringtestGets, "", contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check ret
		if *ref == "" && len(ret) != 64 {
			panic("return value incorrect")
		}
		if *ref != "" && hex.EncodeToString(ret) != *ref {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgCallStringTestGetl creates a message calling getl() in /tests/stringtest.sol contract.
func SimulateMsgCallStringTestGetl(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress, ref *string) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// call getl()
		msg, ret, err := CallFunction(caller, StringtestGetl, "", contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check ret
		length, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 32)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		if *ref == "" && length != 0 {
			panic("return value incorrect")
		}
		str := *ref
		if str != "" && str[64:128] != hex.EncodeToString(ret) {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}

// SimulateMsgCallStringTestTestStuff creates a message calling testStuff() in /tests/stringtest.sol contract.
func SimulateMsgCallStringTestTestStuff(k keeper.Keeper, bk types.BankKeeper, contractAddr sdk.AccAddress) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []sim.Account, chainID string) (
		sim.OperationMsg, []sim.FutureOperation, error) {
		caller, _ := sim.RandomAcc(r, accs)

		// call testStuff()
		msg, ret, err := CallFunction(caller, StringtestTestStuff, "", contractAddr, k, bk, ctx, r, chainID, app)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}

		// check ret
		value, err := strconv.ParseInt(hex.EncodeToString(ret), 16, 32)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName, types.TypeMsgCall, ""), nil, err
		}
		if value != 123123 {
			panic("return value incorrect")
		}

		return sim.NewOperationMsg(&msg, true, "", nil), nil, nil
	}
}
