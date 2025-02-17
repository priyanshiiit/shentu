package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/certikfoundation/shentu/v2/common"
)

// default parameter values
var (
	// default values for Shield pool's parameters
	DefaultProtectionPeriod  = time.Hour * 24 * 21                                                   // 21 days
	DefaultShieldFeesRate    = sdk.NewDecWithPrec(769, 5)                                            // 0.769%
	DefaultWithdrawPeriod    = time.Hour * 24 * 21                                                   // 21 days
	DefaultPoolShieldLimit   = sdk.NewDecWithPrec(50, 2)                                             // 50%
	DefaultMinShieldPurchase = sdk.NewCoins(sdk.NewCoin(common.MicroCTKDenom, sdk.NewInt(50000000))) // 50 CTK

	// default values for Shield claim proposal's parameters
	DefaultClaimPeriod              = time.Hour * 24 * 21                                                    // 21 days
	DefaultPayoutPeriod             = time.Hour * 24 * 56                                                    // 56 days
	DefaultMinClaimProposalDeposit  = sdk.NewCoins(sdk.NewCoin(common.MicroCTKDenom, sdk.NewInt(100000000))) // 100 CTK
	DefaultClaimProposalDepositRate = sdk.NewDecWithPrec(10, 2)                                              // 10%
	DefaultClaimProposalFeesRate    = sdk.NewDecWithPrec(1, 2)                                               // 1%

	// default value for staking-shield rate parameter
	DefaultStakingShieldRate = sdk.NewDec(2)
)

// parameter keys
var (
	ParamStoreKeyPoolParams          = []byte("shieldpoolparams")
	ParamStoreKeyClaimProposalParams = []byte("claimproposalparams")
	ParamStoreKeyStakingShieldRate   = []byte("stakingshieldrateparams")
)

// ParamKeyTable is the key declaration for parameters.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyPoolParams, PoolParams{}, validatePoolParams),
		paramtypes.NewParamSetPair(ParamStoreKeyClaimProposalParams, ClaimProposalParams{}, validateClaimProposalParams),
		paramtypes.NewParamSetPair(ParamStoreKeyStakingShieldRate, sdk.Dec{}, validateStakingShieldRateParams),
	)
}

// NewPoolParams creates a new PoolParams object.
func NewPoolParams(protectionPeriod, withdrawPeriod time.Duration, shieldFeesRate sdk.Dec, poolShieldLimit sdk.Dec, minShieldPurchase sdk.Coins) PoolParams {
	return PoolParams{
		ProtectionPeriod:  protectionPeriod,
		ShieldFeesRate:    shieldFeesRate,
		WithdrawPeriod:    withdrawPeriod,
		PoolShieldLimit:   poolShieldLimit,
		MinShieldPurchase: minShieldPurchase,
	}
}

// DefaultPoolParams returns a default PoolParams instance.
func DefaultPoolParams() PoolParams {
	return NewPoolParams(DefaultProtectionPeriod, DefaultWithdrawPeriod, DefaultShieldFeesRate, DefaultPoolShieldLimit, DefaultMinShieldPurchase)
}

func validatePoolParams(i interface{}) error {
	v, ok := i.(PoolParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	protectionPeriod := v.ProtectionPeriod
	shieldFeesRate := v.ShieldFeesRate
	withdrawPeriod := v.WithdrawPeriod
	poolShieldLimit := v.PoolShieldLimit
	minShieldPurchase := v.MinShieldPurchase

	if protectionPeriod <= 0 {
		return fmt.Errorf("protection period must be positive: %s", protectionPeriod)
	}
	if shieldFeesRate.IsNegative() || shieldFeesRate.GT(sdk.OneDec()) {
		return fmt.Errorf("shield fees rate should be positive and less or equal to one but is %s", shieldFeesRate)
	}
	if withdrawPeriod <= 0 {
		return fmt.Errorf("withdraw period must be positive: %s", withdrawPeriod)
	}
	if poolShieldLimit.IsNegative() || poolShieldLimit.GT(sdk.OneDec()) {
		return fmt.Errorf("pool shield limit should be positive and less or equal to one but is %s", poolShieldLimit)
	}
	if !minShieldPurchase.IsValid() {
		return fmt.Errorf("minimum shield purchase must be a valid sdk.Coins, is %s", minShieldPurchase.String())
	}

	return nil
}

// NewClaimProposalParams creates a new ClaimProposalParams instance.
func NewClaimProposalParams(claimPeriod, payoutPeriod time.Duration, minDeposit sdk.Coins, depositRate, feesRate sdk.Dec) ClaimProposalParams {
	return ClaimProposalParams{
		ClaimPeriod:  claimPeriod,
		PayoutPeriod: payoutPeriod,
		MinDeposit:   minDeposit,
		DepositRate:  depositRate,
		FeesRate:     feesRate,
	}
}

// DefaultClaimProposalParams returns a default ClaimProposalParams instance.
func DefaultClaimProposalParams() ClaimProposalParams {
	return NewClaimProposalParams(DefaultClaimPeriod, DefaultPayoutPeriod,
		DefaultMinClaimProposalDeposit, DefaultClaimProposalDepositRate, DefaultClaimProposalFeesRate)
}

func validateClaimProposalParams(i interface{}) error {
	v, ok := i.(ClaimProposalParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	claimPeriod := v.ClaimPeriod
	payoutPeriod := v.PayoutPeriod
	minDeposit := v.MinDeposit
	depositRate := v.DepositRate
	feesRate := v.FeesRate

	if claimPeriod <= 0 {
		return fmt.Errorf("claim period must be positive: %s", claimPeriod)
	}
	if payoutPeriod <= 0 {
		return fmt.Errorf("payout period must be positive: %s", payoutPeriod)
	}
	if !minDeposit.IsValid() {
		return fmt.Errorf("minimum deposit amount must be a valid sdk.Coins amount, is %s",
			minDeposit.String())
	}
	if depositRate.IsNegative() || depositRate.GT(sdk.OneDec()) {
		return fmt.Errorf("deposit rate should be positive and less or equal to one but is %s",
			depositRate.String())
	}
	if feesRate.IsNegative() || feesRate.GT(sdk.OneDec()) {
		return fmt.Errorf("fees rate should be positive and less or equal to one but is %s",
			feesRate.String())
	}

	return nil
}

// DefaultStakingShieldRateParams returns a default DefaultStakingShieldRateParams.
func DefaultStakingShieldRateParams() sdk.Dec {
	return sdk.NewDec(2)
}

func validateStakingShieldRateParams(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.LTE(sdk.ZeroDec()) {
		return fmt.Errorf("staking shield rate should be greater than 0")
	}
	return nil
}
