package keeper_test

import (
	"encoding/hex"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdksimapp "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/certikfoundation/shentu/v2/common"
	"github.com/certikfoundation/shentu/v2/simapp"
	"github.com/certikfoundation/shentu/v2/x/gov/testgov"
	"github.com/certikfoundation/shentu/v2/x/shield/testshield"
	"github.com/certikfoundation/shentu/v2/x/staking/teststaking"
)

// nextBlock calls staking, shield, and gov endblockers and updates
// their test helpers' contexts.
func nextBlock(ctx sdk.Context, tstaking *teststaking.Helper, tshield *testshield.Helper, tgov *testgov.Helper) sdk.Context {
	newTime := ctx.BlockTime().Add(time.Second * time.Duration(int64(common.SecondsPerBlock)))
	ctx = ctx.WithBlockTime(newTime).WithBlockHeight(ctx.BlockHeight() + 1)

	tstaking.TurnBlock(ctx)
	tshield.TurnBlock(ctx)
	tgov.TurnBlock(ctx)

	return ctx
}

func skipBlocks(ctx sdk.Context, numBlocks int64, tstaking *teststaking.Helper, tshield *testshield.Helper, tgov *testgov.Helper) sdk.Context {
	newTime := ctx.BlockTime().Add(time.Second * time.Duration(int64(common.SecondsPerBlock)*numBlocks))
	ctx = ctx.WithBlockTime(newTime).WithBlockHeight(ctx.BlockHeight() + 1)

	tstaking.TurnBlock(ctx)
	tshield.TurnBlock(ctx)
	tgov.TurnBlock(ctx)

	return ctx
}

func strAddrEqualsAccAddr(strAddr string, accAddr sdk.AccAddress) bool {
	convertedAddr, err := sdk.AccAddressFromBech32(strAddr)
	if err != nil {
		panic(err)
	}
	return convertedAddr.Equals(accAddr)
}

// TestWithdraw tests withdraws triggered by staking undelegation.
func TestWithdrawsByUndelegate(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now().UTC()})

	// create and add addresses
	pks := simapp.CreateTestPubKeys(4)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, sdk.NewInt(2e8))
	del1addr, del2addr := sdk.AccAddress(pks[0].Address()), sdk.AccAddress(pks[1].Address())
	val1pk, val2pk := pks[2], pks[3]
	val1addr, val2addr := sdk.ValAddress(val1pk.Address()), sdk.ValAddress(val2pk.Address())

	// set up testing helpers
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	tshield := testshield.NewHelper(t, ctx, app.ShieldKeeper, tstaking.Denom)
	tgov := testgov.NewHelper(t, ctx, app.GovKeeper, tstaking.Denom)

	// set up two validators
	tstaking.CreateValidatorWithValPower(val1addr, val1pk, 100, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.CheckValidator(val1addr, stakingtypes.Bonded, false)

	tstaking.CreateValidatorWithValPower(val2addr, val2pk, 100, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.CheckValidator(val2addr, stakingtypes.Bonded, false)

	// attempt depositing collateral
	tshield.DepositCollateral(del1addr, 75, false)

	// both delegators delegate 50 to each validator
	tstaking.CheckDelegator(del1addr, val1addr, false)
	tstaking.Delegate(del1addr, val1addr, 50)
	tstaking.CheckDelegator(del1addr, val1addr, true)
	tstaking.CheckDelegator(del1addr, val2addr, false)
	tstaking.Delegate(del1addr, val2addr, 50)
	tstaking.CheckDelegator(del1addr, val2addr, true)

	tstaking.CheckDelegator(del2addr, val1addr, false)
	tstaking.Delegate(del2addr, val1addr, 50)
	tstaking.CheckDelegator(del2addr, val1addr, true)
	tstaking.CheckDelegator(del2addr, val2addr, false)
	tstaking.Delegate(del2addr, val2addr, 50)
	tstaking.CheckDelegator(del2addr, val2addr, true)

	// both delegators deposit collateral of amount 75
	tshield.DepositCollateral(del1addr, 75, true)
	tshield.DepositCollateral(del2addr, 75, true)

	// undelegate total 50 to trigger total withdrawal of 25
	tstaking.Undelegate(del1addr, val1addr, 30, true)
	tstaking.Undelegate(del2addr, val2addr, 10, true)
	tstaking.Undelegate(del1addr, val2addr, 20, true)
	tstaking.Undelegate(del2addr, val2addr, 40, true)

	ctx = nextBlock(ctx, tstaking, tshield, tgov)

	withdraws := app.ShieldKeeper.GetAllWithdraws(ctx)
	numWithdraws := len(withdraws)
	require.True(t, withdraws[0].Amount.Equal(sdk.NewInt(5)))
	require.True(t, strAddrEqualsAccAddr(withdraws[0].Address, del1addr))
	require.True(t, withdraws[1].Amount.Equal(sdk.NewInt(20)))
	require.True(t, strAddrEqualsAccAddr(withdraws[1].Address, del1addr))
	require.True(t, withdraws[2].Amount.Equal(sdk.NewInt(25)))
	require.True(t, strAddrEqualsAccAddr(withdraws[2].Address, del2addr))

	// undelegate 5 and trigger another withdrawal of 5.
	tstaking.Undelegate(del1addr, val1addr, 5, true)

	numWithdraws++
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, len(withdraws) == numWithdraws)
	require.True(t, withdraws[numWithdraws-1].Amount.Equal(sdk.NewInt(5)))

	// must fail deposit of 10
	tshield.DepositCollateral(del1addr, 10, false)

	// delegate 25
	tstaking.Delegate(del1addr, val1addr, 25)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)

	// withdraw 5
	tshield.WithdrawCollateral(del1addr, 5, true)
	numWithdraws++
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, len(withdraws) == numWithdraws)

	// undelegate 25 without triggering withdrawal
	tstaking.Undelegate(del1addr, val1addr, 25, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, len(withdraws) == numWithdraws)
}

// TestWithdraw tests withdraws triggered by staking redelegation.
func TestWithdrawsByRedelegate(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now().UTC()})

	// create and add addresses
	pks := simapp.CreateTestPubKeys(4)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, sdk.NewInt(2e8))
	del1addr, _ := sdk.AccAddress(pks[0].Address()), sdk.AccAddress(pks[1].Address())
	val1pk, val2pk := pks[2], pks[3]
	val1addr, val2addr := sdk.ValAddress(val1pk.Address()), sdk.ValAddress(val2pk.Address())

	// set up testing helpers
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	tshield := testshield.NewHelper(t, ctx, app.ShieldKeeper, tstaking.Denom)
	tgov := testgov.NewHelper(t, ctx, app.GovKeeper, tstaking.Denom)

	// set up two validators
	tstaking.CreateValidatorWithValPower(val1addr, val1pk, 100, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.CheckValidator(val1addr, stakingtypes.Bonded, false)

	tstaking.CreateValidatorWithValPower(val2addr, val2pk, 100, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.CheckValidator(val2addr, stakingtypes.Bonded, false)

	// must fail at depositing collateral
	tshield.DepositCollateral(del1addr, 75, false)

	// delegate 100 to the validator
	tstaking.CheckDelegator(del1addr, val1addr, false)
	tstaking.Delegate(del1addr, val1addr, 100)
	tstaking.CheckDelegator(del1addr, val1addr, true)

	// deposit collateral of amount 75
	tshield.DepositCollateral(del1addr, 75, true)

	// redelegate 50 to trigger withdrawal of 25
	// remaining staking: 100, remaining deposit: 50
	tstaking.Redelegate(del1addr, val1addr, val2addr, 50, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	withdraws := app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, withdraws[0].Amount.Equal(sdk.NewInt(25)))
	numWithdraws := len(withdraws)

	// Redelegation hopping must fail
	tstaking.Redelegate(del1addr, val2addr, val1addr, 10, false)

	// Redelegate 30 but do not trigger withdrawal
	// Remaining staking: 100, remaining deposit: 50
	tstaking.Redelegate(del1addr, val1addr, val2addr, 30, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, len(withdraws) == numWithdraws)

	// must fail deposit of 60
	tshield.DepositCollateral(del1addr, 60, false)

	// must succeed deposit of 50
	tshield.DepositCollateral(del1addr, 50, true)
}

// TestClaimProposal tests a claim proposal process that involves
// withdrawal and unbonding delays.
func TestClaimProposal(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now().UTC()})

	// create and add addresses
	pks := simapp.CreateTestPubKeys(5)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, sdk.ZeroInt())

	shieldAdmin := sdk.AccAddress(pks[0].Address())
	err := sdksimapp.FundAccount(app.BankKeeper, ctx, shieldAdmin, sdk.Coins{sdk.NewInt64Coin("uctk", 250e9)})
	require.NoError(t, err)
	app.ShieldKeeper.SetAdmin(ctx, shieldAdmin)

	sponsorAddr := sdk.AccAddress(pks[1].Address())
	err = sdksimapp.FundAccount(app.BankKeeper, ctx, sponsorAddr, sdk.Coins{sdk.NewInt64Coin("uctk", 1)})
	require.NoError(t, err)

	purchaser := sdk.AccAddress(pks[2].Address())
	err = sdksimapp.FundAccount(app.BankKeeper, ctx, purchaser, sdk.Coins{sdk.NewInt64Coin("uctk", 10e9)})
	require.NoError(t, err)

	del1addr := sdk.AccAddress(pks[3].Address())
	err = sdksimapp.FundAccount(app.BankKeeper, ctx, del1addr, sdk.Coins{sdk.NewInt64Coin("uctk", 125e9)})
	require.NoError(t, err)

	val1pk, val1addr := pks[4], sdk.ValAddress(pks[4].Address())
	err = sdksimapp.FundAccount(app.BankKeeper, ctx, sdk.AccAddress(pks[4].Address()), sdk.Coins{sdk.NewInt64Coin("uctk", 100e6)})
	require.NoError(t, err)

	var adminDeposit int64 = 200e9
	var delegatorDeposit int64 = 125e9
	totalDeposit := adminDeposit + delegatorDeposit

	// set up testing helpers
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	bondDenom := tstaking.Denom
	tshield := testshield.NewHelper(t, ctx, app.ShieldKeeper, bondDenom)
	tgov := testgov.NewHelper(t, ctx, app.GovKeeper, bondDenom)

	// set up a validator
	tstaking.CreateValidatorWithValPower(val1addr, val1pk, 100, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.CheckValidator(val1addr, stakingtypes.Bonded, false)

	// shield admin deposit and create pool
	// $BondDenom pool with shield = 100,000 $BondDenom, limit = 500,000 $BondDenom, serviceFees = 200 $BondDenom
	tstaking.Delegate(shieldAdmin, val1addr, adminDeposit)
	tshield.DepositCollateral(shieldAdmin, adminDeposit, true)
	tshield.CreatePool(shieldAdmin, sponsorAddr, 200e6, 100e9, 500e9, "CertiK", "fake_description")

	pools := app.ShieldKeeper.GetAllPools(ctx)
	require.True(t, len(pools) == 1)
	require.True(t, strAddrEqualsAccAddr(pools[0].SponsorAddr, sponsorAddr))

	poolID := pools[0].Id

	// delegator deposits
	tstaking.CheckDelegator(del1addr, val1addr, false)
	tstaking.Delegate(del1addr, val1addr, delegatorDeposit)
	tstaking.CheckDelegator(del1addr, val1addr, true)
	tshield.DepositCollateral(del1addr, delegatorDeposit, true)

	// purchaser purhcases a shield
	var shield int64 = 50e9
	tshield.PurchaseShield(purchaser, shield, poolID, true)

	// delegator undelegates all delegations, triggering a withdrawal
	tstaking.Undelegate(del1addr, val1addr, 25e9, true)
	withdraw1End := ctx.BlockTime().Add(app.ShieldKeeper.GetPoolParams(ctx).WithdrawPeriod)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.Undelegate(del1addr, val1addr, 90e9, true)
	ctx = nextBlock(ctx, tstaking, tshield, tgov)
	tstaking.Undelegate(del1addr, val1addr, 10e9, true)

	withdraws := app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, len(withdraws) == 3)
	require.True(t, withdraws[0].Amount.Equal(sdk.NewInt(25e9)))
	require.True(t, withdraws[1].Amount.Equal(sdk.NewInt(90e9)))
	require.True(t, withdraws[2].Amount.Equal(sdk.NewInt(10e9)))

	delUBD := app.StakingKeeper.GetAllUnbondingDelegations(ctx, del1addr)[0]
	require.True(t, delUBD.Entries[0].Balance.Equal(sdk.NewInt(25e9)))
	require.True(t, delUBD.Entries[1].Balance.Equal(sdk.NewInt(90e9)))
	require.True(t, delUBD.Entries[2].Balance.Equal(sdk.NewInt(10e9)))

	// 20 days later (345,600 blocks)
	ctx = skipBlocks(ctx, 345600, tstaking, tshield, tgov)

	// the purchaser submits a claim proposal
	loss := shield
	tgov.ShieldClaimProposal(purchaser, loss, poolID, 2, true)
	var proposalID uint64 = 1 // TODO: unmarshal sdk.Result to obtain proposal ID

	// verify that the withdrawal and unbonding have been delayed
	// about 19e9 must be secured (two of three withdraws & ubds are delayed)
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	delayedWithdrawEnd := ctx.BlockTime().Add(app.GovKeeper.GetVotingParams(ctx).VotingPeriod * 2)
	require.True(t, withdraws[0].Amount.Equal(sdk.NewInt(25e9)))
	require.True(t, withdraws[0].CompletionTime.Equal(withdraw1End)) //25e9 not delayed
	require.True(t, withdraws[1].Amount.Equal(sdk.NewInt(10e9)))
	require.True(t, withdraws[1].CompletionTime.Equal(delayedWithdrawEnd)) // 10e9 delayed
	require.True(t, withdraws[2].Amount.Equal(sdk.NewInt(90e9)))
	require.True(t, withdraws[2].CompletionTime.Equal(delayedWithdrawEnd)) // 90e9 delayed

	delUBD = app.StakingKeeper.GetAllUnbondingDelegations(ctx, del1addr)[0]
	require.True(t, delUBD.Entries[0].Balance.Equal(sdk.NewInt(25e9)))
	require.True(t, delUBD.Entries[0].CompletionTime.Equal(withdraw1End)) //25e9 not delayed
	require.True(t, delUBD.Entries[1].Balance.Equal(sdk.NewInt(90e9)))
	require.True(t, delUBD.Entries[1].CompletionTime.Equal(delayedWithdrawEnd)) // 90e9 delayed
	require.True(t, delUBD.Entries[2].Balance.Equal(sdk.NewInt(10e9)))
	require.True(t, delUBD.Entries[2].CompletionTime.Equal(delayedWithdrawEnd)) // 10e9 delayed

	// create reimbursement
	lossCoins := sdk.NewCoins(sdk.NewInt64Coin(bondDenom, loss))
	err = app.ShieldKeeper.CreateReimbursement(ctx, proposalID, lossCoins, purchaser)
	require.NoError(t, err)
	reimbursement, err := app.ShieldKeeper.GetReimbursement(ctx, proposalID)
	require.NoError(t, err)
	require.True(t, reimbursement.Amount.IsEqual(lossCoins))

	// confirm admin delegation reduction
	lossRatio := float64(loss) / float64(totalDeposit)
	expected := adminDeposit - int64(math.Round(float64(adminDeposit)*lossRatio))
	if hex.EncodeToString(shieldAdmin) < hex.EncodeToString(del1addr) {
		expected -= 1 // adjust for discrepancy due to sorting
	}

	adminDels := app.StakingKeeper.GetAllDelegatorDelegations(ctx, shieldAdmin)
	validator, _ := app.StakingKeeper.GetValidator(ctx, val1addr)
	require.True(t, validator.TokensFromShares(adminDels[0].Shares).Equal(sdk.NewDec(expected)))

	// confirm delegator unbonding reduction
	expected = 25e9 + 10e9 + 90e9 - int64(math.Round(float64(125e9)*lossRatio))
	if hex.EncodeToString(shieldAdmin) < hex.EncodeToString(del1addr) {
		expected += 1 // adjust for discrepancy due to sorting
	}
	withdraws = app.ShieldKeeper.GetAllWithdraws(ctx)
	require.True(t, withdraws[0].Amount.Add(withdraws[1].Amount.Add(withdraws[2].Amount)).Equal(sdk.NewInt(expected)))
	delUBD = app.StakingKeeper.GetAllUnbondingDelegations(ctx, del1addr)[0]
	require.True(t, delUBD.Entries[0].Balance.Add(delUBD.Entries[1].Balance.Add(delUBD.Entries[2].Balance)).Equal(sdk.NewInt(expected)))

	// test withdraw reimbursement
	// 56 days later (967,680 blocks)
	ctx = skipBlocks(ctx, 967680, tstaking, tshield, tgov)

	beforeInt := app.BankKeeper.GetBalance(ctx, purchaser, bondDenom).Amount
	tshield.WithdrawReimbursement(purchaser, proposalID, true)
	afterInt := app.BankKeeper.GetBalance(ctx, purchaser, bondDenom).Amount
	require.True(t, beforeInt.Add(sdk.NewInt(loss)).Equal(afterInt))
}
