package keeper_test

import (
	"github.com/certikfoundation/shentu/x/shield/keeper"
	"github.com/certikfoundation/shentu/x/shield/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func TestKeeper_GetAllProviders(t *testing.T) {
	type args struct {
		ctx sdk.Context
	}
	tests := []struct {
		name          string
		keeper        keeper.Keeper
		args          args
		wantProviders []types.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.keeper
			if gotProviders := k.GetAllProviders(tt.args.ctx); !reflect.DeepEqual(gotProviders, tt.wantProviders) {
				t.Errorf("GetAllProviders() = %v, want %v", gotProviders, tt.wantProviders)
			}
		})
	}
}

func TestKeeper_GetProvider(t *testing.T) {
	type args struct {
		ctx       sdk.Context
		delegator sdk.AccAddress
	}
	tests := []struct {
		name      string
		keeper    keeper.Keeper
		args      args
		wantDt    types.Provider
		wantFound bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.keeper
			gotDt, gotFound := k.GetProvider(tt.args.ctx, tt.args.delegator)
			if !reflect.DeepEqual(gotDt, tt.wantDt) {
				t.Errorf("GetProvider() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if gotFound != tt.wantFound {
				t.Errorf("GetProvider() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestKeeper_GetProvidersIteratorPaginated(t *testing.T) {
	type args struct {
		ctx   sdk.Context
		page  uint
		limit uint
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
		want   sdk.Iterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.keeper
			if got := k.GetProvidersIteratorPaginated(tt.args.ctx, tt.args.page, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvidersIteratorPaginated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeeper_GetProvidersPaginated(t *testing.T) {
	type args struct {
		ctx   sdk.Context
		page  uint
		limit uint
	}
	tests := []struct {
		name          string
		keeper        keeper.Keeper
		args          args
		wantProviders []types.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.keeper
			if gotProviders := k.GetProvidersPaginated(tt.args.ctx, tt.args.page, tt.args.limit); !reflect.DeepEqual(gotProviders, tt.wantProviders) {
				t.Errorf("GetProvidersPaginated() = %v, want %v", gotProviders, tt.wantProviders)
			}
		})
	}
}

func TestKeeper_IterateProviders(t *testing.T) {
	type args struct {
		ctx      sdk.Context
		callback func(provider types.Provider) (stop bool)
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}

func TestKeeper_IterateProvidersPaginated(t *testing.T) {
	type args struct {
		ctx   sdk.Context
		page  uint
		limit uint
		cb    func(vote types.Provider) (stop bool)
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}

func TestKeeper_RemoveDelegation(t *testing.T) {
	type args struct {
		ctx     sdk.Context
		delAddr sdk.AccAddress
		valAddr sdk.ValAddress
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}

func TestKeeper_SetProvider(t *testing.T) {
	type args struct {
		ctx      sdk.Context
		delAddr  sdk.AccAddress
		provider types.Provider
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}

func TestKeeper_UpdateDelegationAmount(t *testing.T) {
	type args struct {
		ctx     sdk.Context
		delAddr sdk.AccAddress
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}

func TestKeeper_addProvider(t *testing.T) {
	type args struct {
		ctx  sdk.Context
		addr sdk.AccAddress
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
		want   types.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.keeper
			if got := k.AddProvider(tt.args.ctx, tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeeper_updateProviderForDelegationChanges(t *testing.T) {
	type args struct {
		ctx       sdk.Context
		delAddr   sdk.AccAddress
		stakedAmt sdk.Int
	}
	tests := []struct {
		name   string
		keeper keeper.Keeper
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.keeper
		})
	}
}
