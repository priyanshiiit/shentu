package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

type unbondingInfo struct {
	delegator      string
	validator      string
	completionTime time.Time
}

// InsertWithdrawQueue prepares a withdraw queue timeslice
// for insertion into the queue.
func (k Keeper) InsertWithdrawQueue(ctx sdk.Context, withdraw types.Withdraw) {
	timeSlice := k.GetWithdrawQueueTimeSlice(ctx, withdraw.CompletionTime)
	timeSlice = append(timeSlice, withdraw)
	k.SetWithdrawQueueTimeSlice(ctx, withdraw.CompletionTime, timeSlice)
}

// SetWithdrawQueueTimeSlice stores a withdraw queue timeslice
// using the timestamp as the key.
func (k Keeper) SetWithdrawQueueTimeSlice(ctx sdk.Context, timestamp time.Time, withdraws []types.Withdraw) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&types.Withdraws{Withdraws: withdraws})
	store.Set(types.GetWithdrawCompletionTimeKey(timestamp), bz)
}

// GetWithdrawQueueTimeSlice gets a specific withdraw queue timeslice,
// which is a slice of withdraws corresponding to a given time.
func (k Keeper) GetWithdrawQueueTimeSlice(ctx sdk.Context, timestamp time.Time) []types.Withdraw {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetWithdrawCompletionTimeKey(timestamp))
	if bz == nil {
		return []types.Withdraw{}
	}
	var withdraws types.Withdraws
	k.cdc.MustUnmarshalLengthPrefixed(bz, &withdraws)
	return withdraws.Withdraws
}

// WithdrawQueueIterator returns all the withdraw queue timeslices from time 0 until endTime.
func (k Keeper) WithdrawQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.WithdrawQueueKey, sdk.InclusiveEndBytes(types.GetWithdrawCompletionTimeKey(endTime)))
}

// IterateWithdraws iterates through all ongoing withdraws.
func (k Keeper) IterateWithdraws(ctx sdk.Context, callback func(withdraw []types.Withdraw) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.WithdrawQueueKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		timeslice := types.Withdraws{}
		value := iterator.Value()
		k.cdc.MustUnmarshalLengthPrefixed(value, &timeslice)

		if callback(timeslice.Withdraws) {
			break
		}
	}
}

// GetAllWithdraws gets all collaterals that are being withdrawn.
func (k Keeper) GetAllWithdraws(ctx sdk.Context) (result []types.Withdraw) {
	k.IterateWithdraws(ctx, func(withdraws []types.Withdraw) bool {
		result = append(result, withdraws...)
		return false
	})
	return result
}

// GetWithdrawsByProvider gets all withdraws of a provider.
func (k Keeper) GetWithdrawsByProvider(ctx sdk.Context, providerAddr string) []types.Withdraw {
	var result []types.Withdraw
	k.IterateWithdraws(ctx, func(withdraws []types.Withdraw) bool {
		for _, withdraw := range withdraws {
			if withdraw.Address == providerAddr {
				result = append(result, withdraw)
			}
		}
		return false
	})
	return result
}

// RemoveTimeSliceFromWithdrawQueue removes a time slice from the withdraw queue.
func (k Keeper) RemoveTimeSliceFromWithdrawQueue(ctx sdk.Context, timestamp time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetWithdrawCompletionTimeKey(timestamp))
}

// DequeueCompletedWithdrawQueue dequeues completed withdraws
// and processes their completions.
func (k Keeper) DequeueCompletedWithdrawQueue(ctx sdk.Context) {
	// retrieve completed withdraws from the queue
	store := ctx.KVStore(k.storeKey)
	withdrawTimesliceIterator := k.WithdrawQueueIterator(ctx, ctx.BlockHeader().Time)
	defer withdrawTimesliceIterator.Close()

	var withdraws []types.Withdraw
	for ; withdrawTimesliceIterator.Valid(); withdrawTimesliceIterator.Next() {
		var timeslice types.Withdraws
		value := withdrawTimesliceIterator.Value()
		k.cdc.MustUnmarshalLengthPrefixed(value, &timeslice)
		withdraws = append(withdraws, timeslice.Withdraws...)
		store.Delete(withdrawTimesliceIterator.Key())
	}

	// For each completed withdraw, process adjustments.
	totalCollateral := k.GetTotalCollateral(ctx)
	totalWithdrawing := k.GetTotalWithdrawing(ctx)
	for _, withdraw := range withdraws {
		providerAddr, err := sdk.AccAddressFromBech32(withdraw.Address)
		if err != nil {
			panic(err)
		}

		provider, found := k.GetProvider(ctx, providerAddr)
		if !found {
			panic("provider not found but its collaterals are being withdrawn")
		}
		provider.Collateral = provider.Collateral.Sub(withdraw.Amount)
		provider.Withdrawing = provider.Withdrawing.Sub(withdraw.Amount)
		k.SetProvider(ctx, providerAddr, provider)

		totalCollateral = totalCollateral.Sub(withdraw.Amount)
		totalWithdrawing = totalWithdrawing.Sub(withdraw.Amount)
	}
	k.SetTotalCollateral(ctx, totalCollateral)
	k.SetTotalWithdrawing(ctx, totalWithdrawing)
}

// ComputeWithdrawAmountByTime computes the amount of collaterals
// from a given provider that will be dequeued from the withdraw
// queue by a given time.
func (k Keeper) ComputeWithdrawAmountByTime(ctx sdk.Context, provider string, time time.Time) sdk.Int {
	withdrawTimesliceIterator := k.WithdrawQueueIterator(ctx, time)
	defer withdrawTimesliceIterator.Close()

	amount := sdk.ZeroInt()
	for ; withdrawTimesliceIterator.Valid(); withdrawTimesliceIterator.Next() {
		var timeslice types.Withdraws
		value := withdrawTimesliceIterator.Value()
		k.cdc.MustUnmarshalLengthPrefixed(value, &timeslice)

		for _, withdraw := range timeslice.Withdraws {
			if withdraw.Address == provider {
				amount = amount.Add(withdraw.Amount)
			}
		}
	}
	return amount
}

func (k Keeper) ComputeTotalUnbondingAmount(ctx sdk.Context, provider sdk.AccAddress) sdk.Int {
	unbondings := k.sk.GetAllUnbondingDelegations(ctx, provider)

	sum := sdk.ZeroInt()
	for _, ubd := range unbondings {
		for _, entry := range ubd.Entries {
			sum = sum.Add(entry.Balance)
		}
	}
	return sum
}

func (k Keeper) ComputeUnbondingAmountByTime(ctx sdk.Context, provider sdk.AccAddress, time time.Time) sdk.Int {
	dvPairs := k.getUnbondingsByProviderMaturingByTime(ctx, provider.String(), time)

	sum := sdk.ZeroInt()
	seen := make(map[string]bool)
	for _, dvPair := range dvPairs {
		valAddr := dvPair.validator
		if seen[valAddr] {
			continue
		}
		seen[valAddr] = true

		// obtain unbonding entries and iterate through them
		addr, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			panic(err)
		}
		ubd, found := k.sk.GetUnbondingDelegation(ctx, provider, addr)
		if !found {
			continue //TODO
		}
		for i := 0; i < len(ubd.Entries); i++ {
			entry := ubd.Entries[i]
			if !entry.IsMature(time) {
				break
			}
			sum = sum.Add(entry.Balance)
		}
	}
	return sum
}

func (k Keeper) getUnbondingsByProviderMaturingByTime(ctx sdk.Context, provider string, time time.Time) (results []unbondingInfo) {
	unbondingTimesliceIterator := k.sk.UBDQueueIterator(ctx, time)
	defer unbondingTimesliceIterator.Close()

	for ; unbondingTimesliceIterator.Valid(); unbondingTimesliceIterator.Next() {
		var timeslice stakingtypes.DVPairs
		value := unbondingTimesliceIterator.Value()
		k.cdc.MustUnmarshal(value, &timeslice)

		for _, ubd := range timeslice.Pairs {
			if ubd.DelegatorAddress == provider {
				completionTime, _ := sdk.ParseTimeBytes(unbondingTimesliceIterator.Key()[1:])
				ubdInfo := unbondingInfo{
					delegator:      ubd.DelegatorAddress,
					validator:      ubd.ValidatorAddress,
					completionTime: completionTime,
				}
				results = append(results, ubdInfo)
			}
		}
	}
	return results
}

// DelayWithdraws delays the given amount of withdraws maturing
// before the delay duration until the end of the delay duration.
func (k Keeper) DelayWithdraws(ctx sdk.Context, provider string, amount sdk.Int, delayedTime time.Time) error {
	// Retrieve delay candidates, which are withdraws
	// ending before the delay duration from now.
	withdrawTimesliceIterator := k.WithdrawQueueIterator(ctx, delayedTime)
	defer withdrawTimesliceIterator.Close()

	withdraws := []types.Withdraw{}
	for ; withdrawTimesliceIterator.Valid(); withdrawTimesliceIterator.Next() {
		var timeslice types.Withdraws
		value := withdrawTimesliceIterator.Value()
		k.cdc.MustUnmarshalLengthPrefixed(value, &timeslice)

		for _, withdraw := range timeslice.Withdraws {
			if withdraw.Address == provider {
				withdraws = append(withdraws, withdraw)
			}
		}
	}

	// Delay withdraws, starting with the candidates
	// with the oldest withdraw completion time.
	remaining := amount
	for i := len(withdraws) - 1; i >= 0; i-- {
		if !remaining.IsPositive() {
			break
		}
		// Remove from withdraw queue.
		if timeSlice := k.GetWithdrawQueueTimeSlice(ctx, withdraws[i].CompletionTime); len(timeSlice) > 1 {
			for j := len(timeSlice) - 1; j >= 0; j-- {
				if timeSlice[j].Address == provider && timeSlice[j].Amount.Equal(withdraws[i].Amount) {
					timeSlice = append(timeSlice[:j], timeSlice[j+1:]...)
					k.SetWithdrawQueueTimeSlice(ctx, withdraws[i].CompletionTime, timeSlice)
					break
				}
			}
		} else {
			k.RemoveTimeSliceFromWithdrawQueue(ctx, withdraws[i].CompletionTime)
		}

		// Adjust the withdraw completion time and re-insert.
		withdraws[i].CompletionTime = delayedTime
		k.InsertWithdrawQueue(ctx, withdraws[i])

		remaining = remaining.Sub(withdraws[i].Amount)
	}

	if remaining.IsPositive() {
		panic("failed to delay enough withdraws")
	}

	return nil
}

func (k Keeper) DelayUnbonding(ctx sdk.Context, provider sdk.AccAddress, amount sdk.Int, delayedTime time.Time) error {
	providerStr := provider.String()

	// Retrieve delay candidates, which are unbondings
	// ending before the delay duration from now.
	ubds := k.getUnbondingsByProviderMaturingByTime(ctx, providerStr, delayedTime)

	// Delay unbondings, starting with the candidates
	// with the oldest unbonding completion time.
	remaining := amount
	for i := len(ubds) - 1; i >= 0; i-- {
		if !remaining.IsPositive() {
			break
		}
		// Remove from unbonding queue.
		if timeSlice := k.sk.GetUBDQueueTimeSlice(ctx, ubds[i].completionTime); len(timeSlice) > 1 {
			for j := len(timeSlice) - 1; j >= 0; j-- {
				if timeSlice[j].DelegatorAddress == providerStr && timeSlice[j].ValidatorAddress == ubds[i].validator {
					timeSlice = append(timeSlice[:j], timeSlice[j+1:]...)
					k.sk.SetUBDQueueTimeSlice(ctx, ubds[i].completionTime, timeSlice)
					break
				}
			}
		} else {
			k.sk.RemoveUBDQueue(ctx, ubds[i].completionTime)
		}

		valAddr, err := sdk.ValAddressFromBech32(ubds[i].validator)
		if err != nil {
			panic(err)
		}
		unbondingDels, found := k.sk.GetUnbondingDelegation(ctx, provider, valAddr)
		if !found {
			panic("unbonding list was not found for the given provider-validator pair")
		}

		found = false
		amount := sdk.ZeroInt()
		for j := 0; j < len(unbondingDels.Entries); j++ {
			if !found {
				if unbondingDels.Entries[j].CompletionTime.Equal(ubds[i].completionTime) {
					unbondingDels.Entries[j].CompletionTime = delayedTime
					found = true
					amount = unbondingDels.Entries[j].Balance
				}
				continue
			}

			if unbondingDels.Entries[j].CompletionTime.Before(unbondingDels.Entries[j-1].CompletionTime) {
				unbondingDels.Entries[j-1], unbondingDels.Entries[j] = unbondingDels.Entries[j], unbondingDels.Entries[j-1]
				continue
			}

			break
		}
		if !found {
			panic("particular unbonding entry not found for the given timestamp")
		}

		k.sk.InsertUBDQueue(ctx, unbondingDels, delayedTime)
		k.sk.SetUnbondingDelegation(ctx, unbondingDels)

		remaining = remaining.Sub(amount)
	}

	if remaining.IsPositive() {
		panic("failed to delay enough unbondings")
	}

	return nil
}
