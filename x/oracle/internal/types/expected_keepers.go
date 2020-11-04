package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

type CertKeeper interface {
	IsCertified(ctx sdk.Context, requestContentType string, content string, certType string) bool
}

type AuthKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account
}

type DistrKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type StakingKeeper interface {
	BondDenom(ctx sdk.Context) (res string)
}

type SupplyKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	SendCoinsFromAccountToModule(
		ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(
		ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
