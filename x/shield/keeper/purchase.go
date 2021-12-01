package keeper

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/x/shield/types"
)

type pPPTriplet struct {
	poolID     uint64
	purchaseID uint64
	purchaser  sdk.AccAddress
}

// SetPurchaseList sets a purchase list.
func (k Keeper) SetPurchaseList(ctx sdk.Context, purchaseList types.PurchaseList) {
	store := ctx.KVStore(k.storeKey)

	purchaser, err := sdk.AccAddressFromBech32(purchaseList.Purchaser)
	if err != nil {
		panic(err)
	}
	bz := k.cdc.MustMarshalLengthPrefixed(&purchaseList)
	store.Set(types.GetPurchaseListKey(purchaseList.PoolId, purchaser), bz)
}

// AddPurchase sets a purchase of shield.
func (k Keeper) AddPurchase(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress, purchase types.Purchase) types.PurchaseList {
	purchaseList, found := k.GetPurchaseList(ctx, poolID, purchaser)
	if !found {
		purchaseList = types.NewPurchaseList(poolID, purchaser, []types.Purchase{purchase})
	} else {
		purchaseList.Entries = append(purchaseList.Entries, purchase)
	}
	k.SetPurchaseList(ctx, purchaseList)
	return purchaseList
}

// GetPurchaseList gets a purchase from store by txhash.
func (k Keeper) GetPurchaseList(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) (types.PurchaseList, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPurchaseListKey(poolID, purchaser))
	if bz != nil {
		var purchase types.PurchaseList
		k.cdc.MustUnmarshalLengthPrefixed(bz, &purchase)
		return purchase, true
	}
	return types.PurchaseList{}, false
}

// GetPurchase gets a purchase out of a purchase list
func (Keeper) GetPurchase(purchaseList types.PurchaseList, purchaseID uint64) (types.Purchase, bool) {
	for _, entry := range purchaseList.Entries {
		if entry.PurchaseId == purchaseID {
			return entry, true
		}
	}
	return types.Purchase{}, false
}

// DeletePurchaseList deletes a purchase of shield.
func (k Keeper) DeletePurchaseList(ctx sdk.Context, poolID uint64, purchaser sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	_, found := k.GetPurchaseList(ctx, poolID, purchaser)
	if !found {
		return types.ErrPurchaseNotFound
	}
	store.Delete(types.GetPurchaseListKey(poolID, purchaser))
	return nil
}

// DequeuePurchase removes a pool-purchaser pair at a given timestamp of the purchase queue.
func (k Keeper) DequeuePurchase(ctx sdk.Context, purchaseList types.PurchaseList, timestamp time.Time) {
	timeslice := k.GetExpiringPurchaseQueueTimeSlice(ctx, timestamp)
	for i, poolPurchaser := range timeslice {
		if (purchaseList.PoolId == poolPurchaser.PoolId) && purchaseList.Purchaser == poolPurchaser.Purchaser {
			if len(timeslice) > 1 {
				timeslice = append(timeslice[:i], timeslice[i+1:]...)
				k.SetExpiringPurchaseQueueTimeSlice(ctx, timestamp, timeslice)
				return
			}
			ctx.KVStore(k.storeKey).Delete(types.GetPurchaseExpirationTimeKey(timestamp))
			return
		}
	}
}

// PurchaseShield purchases shield of a pool.
func (k Keeper) purchaseShield(ctx sdk.Context, poolID uint64, shield sdk.Coins, description string, purchaser sdk.AccAddress, serviceFees sdk.Coins, stakingCoins sdk.Coins) (types.Purchase, error) {
	pool, found := k.GetPool(ctx, poolID)
	if !found {
		return types.Purchase{}, types.ErrNoPoolFound
	}
	if !pool.Active {
		return types.Purchase{}, types.ErrPoolInactive
	}
	if shield.Empty() {
		return types.Purchase{}, types.ErrNoShield
	}
	if serviceFees.Empty() && stakingCoins.Empty() {
		return types.Purchase{}, types.ErrNoShield
	}

	// Check available collaterals.
	bondDenom := k.sk.BondDenom(ctx)
	shieldAmt := shield.AmountOf(bondDenom)
	totalCollateral := k.GetTotalCollateral(ctx)
	totalWithdrawing := k.GetTotalWithdrawing(ctx)
	totalShield := k.GetTotalShield(ctx)
	totalClaimed := k.GetTotalClaimed(ctx)
	if totalShield.Add(shieldAmt).GT(totalCollateral.Sub(totalWithdrawing).Sub(totalClaimed)) {
		return types.Purchase{}, types.ErrNotEnoughCollateral
	}

	// Check pool shield limit.
	poolParams := k.GetPoolParams(ctx)
	protectionEndTime := ctx.BlockTime().Add(poolParams.ProtectionPeriod)
	maxShield := sdk.MinInt(pool.ShieldLimit, totalCollateral.Sub(totalWithdrawing).Sub(totalClaimed).ToDec().Mul(poolParams.PoolShieldLimit).TruncateInt())
	if shieldAmt.Add(pool.Shield).GT(maxShield) {
		return types.Purchase{}, types.ErrPoolShieldExceedsLimit
	}

	// get next purchase ID and set purchase ID after that
	purchaseID := k.GetNextPurchaseID(ctx)
	k.SetNextPurchaseID(ctx, purchaseID+1)
	if !serviceFees.Empty() {
		// Send service fees to the shield module account and update service fees.
		if err := k.bk.SendCoinsFromAccountToModule(ctx, purchaser, types.ModuleName, serviceFees); err != nil {
			return types.Purchase{}, err
		}
		totalServiceFees := k.GetServiceFees(ctx)
		totalServiceFees = totalServiceFees.Add(types.MixedDecCoins{Native: sdk.NewDecCoinsFromCoins(serviceFees...)})
		k.SetServiceFees(ctx, totalServiceFees)
		totalRemainingServiceFees := k.GetRemainingServiceFees(ctx)
		totalRemainingServiceFees = totalRemainingServiceFees.Add(types.MixedDecCoins{Native: sdk.NewDecCoinsFromCoins(serviceFees...)})
		k.SetRemainingServiceFees(ctx, totalRemainingServiceFees)
	} else {
		if err := k.AddStaking(ctx, poolID, purchaser, purchaseID, stakingCoins.AmountOf(bondDenom)); err != nil {
			return types.Purchase{}, err
		}
	}

	// Update global pool and project pool's shield.
	totalShield = totalShield.Add(shieldAmt)
	pool.Shield = pool.Shield.Add(shieldAmt)
	k.SetTotalShield(ctx, totalShield)
	k.SetPool(ctx, pool)

	// Set a new purchase.
	purchase := types.NewPurchase(purchaseID, protectionEndTime, protectionEndTime, description, shieldAmt, types.MixedDecCoins{Native: sdk.NewDecCoinsFromCoins(serviceFees...)})
	purchaseList := k.AddPurchase(ctx, poolID, purchaser, purchase)
	k.InsertExpiringPurchaseQueue(ctx, purchaseList, protectionEndTime)

	lastUpdateTime, found := k.GetLastUpdateTime(ctx)
	if !found || lastUpdateTime.IsZero() {
		k.SetLastUpdateTime(ctx, ctx.BlockTime())
	}

	return purchase, nil
}

// PurchaseShield purchases shield of a pool with standard fee rate.
func (k Keeper) PurchaseShield(ctx sdk.Context, poolID uint64, shield sdk.Coins, description string, purchaser sdk.AccAddress, staking bool) (types.Purchase, error) {
	poolParams := k.GetPoolParams(ctx)
	if poolParams.MinShieldPurchase.IsAnyGT(shield) {
		return types.Purchase{}, types.ErrPurchaseTooSmall
	}
	bondDenom := k.BondDenom(ctx)
	serviceFees := sdk.NewCoins()
	stakingCoins := sdk.NewCoins()
	if !staking {
		serviceFees = sdk.NewCoins(sdk.NewCoin(bondDenom, shield.AmountOf(bondDenom).ToDec().Mul(k.GetPoolParams(ctx).ShieldFeesRate).TruncateInt()))
	} else {
		// stake to the staking purchase pool
		stakingAmt := k.GetShieldStakingRate(ctx).MulInt(shield.AmountOf(bondDenom)).TruncateInt()
		stakingCoins = sdk.NewCoins(sdk.NewCoin(bondDenom, stakingAmt))
	}
	return k.purchaseShield(ctx, poolID, shield, description, purchaser, serviceFees, stakingCoins)
}

// RemoveExpiredPurchasesAndDistributeFees removes expired purchases and distributes fees for current block.
func (k Keeper) RemoveExpiredPurchasesAndDistributeFees(ctx sdk.Context) {
	lastUpdateTime, found := k.GetLastUpdateTime(ctx)
	if !found || lastUpdateTime.IsZero() {
		// Last update time will be set when a purchase is made.
		return
	}

	totalServiceFees := k.GetServiceFees(ctx)
	totalShield := k.GetTotalShield(ctx)
	serviceFees := types.InitMixedDecCoins()
	bondDenom := k.BondDenom(ctx)
	var stakeForShieldUpdateList []pPPTriplet

	// Check all purchases whose protection end time is before current block time.
	// 1) Update service fees for purchases whose protection end time is before current block time.
	// 2) Remove purchases whose deletion time is before current block time.
	iterator := k.ExpiringPurchaseQueueIterator(ctx, ctx.BlockTime())
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var timeslice types.PoolPurchaserPairs
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &timeslice)
		for _, poolPurchaser := range timeslice.Pairs {
			purchaser, err := sdk.AccAddressFromBech32(poolPurchaser.Purchaser)
			if err != nil {
				panic(err)
			}

			purchaseList, _ := k.GetPurchaseList(ctx, poolPurchaser.PoolId, purchaser)
			for i := 0; i < len(purchaseList.Entries); i++ {
				entry := purchaseList.Entries[i]

				// Skip entries that has not expired yet.
				if entry.ProtectionEndTime.After(ctx.BlockTime()) {
					continue
				}

				// If purchaseProtectionEndTime > previousBlockTime, update service fees.
				// Otherwise services fees were updated in the last block.
				if entry.ProtectionEndTime.After(lastUpdateTime) && entry.ServiceFees.Native.IsAllPositive() {
					// Add purchaseServiceFees * (purchaseProtectionEndTime - previousBlockTime) / protectionPeriod.
					serviceFees = serviceFees.Add(entry.ServiceFees.MulDec(
						sdk.NewDec(entry.ProtectionEndTime.Sub(lastUpdateTime).Nanoseconds()).Quo(
							sdk.NewDec(k.GetPoolParams(ctx).ProtectionPeriod.Nanoseconds()))))
					// Remove purchaseServiceFees from total service fees.
					totalServiceFees = totalServiceFees.Sub(entry.ServiceFees)
					// Set purchaseServiceFees to zero because it can be reached again.
					purchaseList.Entries[i].ServiceFees = types.InitMixedDecCoins()

					originalStaking := k.GetOriginalStaking(ctx, entry.PurchaseId)
					if !originalStaking.IsZero() {
						// keep track of the list to be updated to avoid overwriting the purchase list
						stakeForShieldUpdateList = append(stakeForShieldUpdateList, pPPTriplet{
							poolID:     poolPurchaser.PoolId,
							purchaseID: entry.PurchaseId,
							purchaser:  purchaser,
						})
					}
				}

				// If purchaseDeletionTime < currentBlockTime, remove the purchase.
				if entry.DeletionTime.Before(ctx.BlockTime()) {
					k.DequeuePurchase(ctx, purchaseList, entry.ProtectionEndTime)

					// If purchaseProtectionEndTime > previousBlockTime, calculate and set service fees before removing the purchase.
					purchaseList.Entries = append(purchaseList.Entries[:i], purchaseList.Entries[i+1:]...)
					// Update pool shield and total shield.
					pool, found := k.GetPool(ctx, purchaseList.PoolId)
					if !found {
						panic("cannot find the pool for an expired purchase")
					}
					totalShield = totalShield.Sub(entry.Shield)
					pool.Shield = pool.Shield.Sub(entry.Shield)
					k.SetPool(ctx, pool)
					// Minus one because the current entry is deleted.
					i--
				}
			}

			// purchaseList might have been updated in the loop.
			if len(purchaseList.Entries) == 0 {
				_ = k.DeletePurchaseList(ctx, purchaseList.PoolId, purchaser)
			} else {
				k.SetPurchaseList(ctx, purchaseList)
			}
		}
	}
	k.SetTotalShield(ctx, totalShield)
	k.SetServiceFees(ctx, totalServiceFees)
	for _, ppp := range stakeForShieldUpdateList {
		k.ProcessStakeForShieldExpiration(ctx, ppp.poolID, ppp.purchaseID, bondDenom,
			ppp.purchaser)
	}

	// Add service fees for this block from unexpired purchases.
	// totalServiceFees * (currentBlockTime - previousBlockTime) / protectionPeriodTime
	serviceFees = serviceFees.Add(totalServiceFees.MulDec(
		sdk.NewDec(ctx.BlockTime().Sub(lastUpdateTime).Nanoseconds())).QuoDec(
		sdk.NewDec(k.GetPoolParams(ctx).ProtectionPeriod.Nanoseconds())))

	// Limit service fees by remaining service fees.
	remainingServiceFees := k.GetRemainingServiceFees(ctx)
	if remainingServiceFees.Native.AmountOf(bondDenom).LT(serviceFees.Native.AmountOf(bondDenom)) {
		serviceFees.Native = remainingServiceFees.Native
	}

	// Add block service fees that need to be distributed for this block
	blockServiceFees := k.GetBlockServiceFees(ctx)
	serviceFees = serviceFees.Add(blockServiceFees)
	k.DeleteBlockServiceFees(ctx)

	// Distribute service fees.
	totalCollateral := k.GetTotalCollateral(ctx)
	providers := k.GetAllProviders(ctx)
	for _, provider := range providers {
		providerAddr, err := sdk.AccAddressFromBech32(provider.Address)
		if err != nil {
			panic(err)
		}

		// fees * providerCollateral / totalCollateral
		nativeFees := serviceFees.Native.MulDec(sdk.NewDecFromInt(provider.Collateral).QuoInt(totalCollateral))
		if nativeFees.AmountOf(bondDenom).GT(remainingServiceFees.Native.AmountOf(bondDenom)) {
			nativeFees = remainingServiceFees.Native
		}
		provider.Rewards = provider.Rewards.Add(types.MixedDecCoins{Native: nativeFees})
		k.SetProvider(ctx, providerAddr, provider)

		remainingServiceFees.Native = remainingServiceFees.Native.Sub(nativeFees)
	}
	// add back block service fees
	remainingServiceFees.Native = remainingServiceFees.Native.Add(blockServiceFees.Native...)
	k.SetRemainingServiceFees(ctx, remainingServiceFees)
	k.SetLastUpdateTime(ctx, ctx.BlockTime())
}

// GetPurchaserPurchases returns all purchases by a given purchaser.
func (k Keeper) GetPurchaserPurchases(ctx sdk.Context, address sdk.AccAddress) (res []types.PurchaseList) {
	pools := k.GetAllPools(ctx)
	for _, pool := range pools {
		pList, found := k.GetPurchaseList(ctx, pool.Id, address)
		if !found {
			continue
		}
		res = append(res, pList)
	}
	return
}

// GetPoolPurchaseLists returns all purchases in a given pool.
func (k Keeper) GetPoolPurchaseLists(ctx sdk.Context, poolID uint64) (purchases []types.PurchaseList) {
	k.IteratePoolPurchaseLists(ctx, poolID, func(purchaseList types.PurchaseList) bool {
		if purchaseList.PoolId == poolID {
			purchases = append(purchases, purchaseList)
		}
		return false
	})
	return purchases
}

// IteratePurchaseLists iterates through purchase lists in a pool
func (k Keeper) IteratePurchaseLists(ctx sdk.Context, callback func(purchase types.PurchaseList) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PurchaseListKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchaseList types.PurchaseList
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &purchaseList)

		if callback(purchaseList) {
			break
		}
	}
}

// IteratePoolPurchaseLists iterates through purchases in a pool
func (k Keeper) IteratePoolPurchaseLists(ctx sdk.Context, poolID uint64, callback func(purchaseList types.PurchaseList) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, poolID)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PurchaseListKey, bz...))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchaseList types.PurchaseList
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &purchaseList)

		if callback(purchaseList) {
			break
		}
	}
}

// GetAllPurchaseLists retrieves all purchase lists.
func (k Keeper) GetAllPurchaseLists(ctx sdk.Context) (purchases []types.PurchaseList) {
	k.IteratePurchaseLists(ctx, func(purchase types.PurchaseList) bool {
		purchases = append(purchases, purchase)
		return false
	})
	return
}

// InsertExpiringPurchaseQueue inserts a purchase into the expired purchase queue.
func (k Keeper) InsertExpiringPurchaseQueue(ctx sdk.Context, purchaseList types.PurchaseList, endTime time.Time) {
	timeSlice := k.GetExpiringPurchaseQueueTimeSlice(ctx, endTime)

	poolPurchaser := types.PoolPurchaser{PoolId: purchaseList.PoolId, Purchaser: purchaseList.Purchaser}
	if len(timeSlice) == 0 {
		k.SetExpiringPurchaseQueueTimeSlice(ctx, endTime, []types.PoolPurchaser{poolPurchaser})
		return
	}
	timeSlice = append(timeSlice, poolPurchaser)
	k.SetExpiringPurchaseQueueTimeSlice(ctx, endTime, timeSlice)
}

// GetExpiringPurchaseQueueTimeSlice gets a specific purchase queue timeslice,
// which is a slice of purchases corresponding to a given time.
func (k Keeper) GetExpiringPurchaseQueueTimeSlice(ctx sdk.Context, timestamp time.Time) []types.PoolPurchaser {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPurchaseExpirationTimeKey(timestamp))
	if bz == nil {
		return []types.PoolPurchaser{}
	}
	var ppPairs types.PoolPurchaserPairs
	k.cdc.MustUnmarshalLengthPrefixed(bz, &ppPairs)
	return ppPairs.Pairs
}

// SetExpiringPurchaseQueueTimeSlice sets a time slice for a purchase expiring at give time.
func (k Keeper) SetExpiringPurchaseQueueTimeSlice(ctx sdk.Context, timestamp time.Time, ppPairs []types.PoolPurchaser) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&types.PoolPurchaserPairs{Pairs: ppPairs})
	store.Set(types.GetPurchaseExpirationTimeKey(timestamp), bz)
}

// ExpiringPurchaseQueueIterator returns a iterator of purchases expiring before endTime
func (k Keeper) ExpiringPurchaseQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.PurchaseQueueKey,
		sdk.InclusiveEndBytes(types.GetPurchaseExpirationTimeKey(endTime)))
}

// SetNextPurchaseID sets the latest pool ID to store.
func (k Keeper) SetNextPurchaseID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextPurchaseIDKey(), bz)
}

// GetNextPurchaseID gets the latest pool ID from store.
func (k Keeper) GetNextPurchaseID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	opBz := store.Get(types.GetNextPurchaseIDKey())
	return binary.LittleEndian.Uint64(opBz)
}

// GetAllPurchases retrieves all purchases.
func (k Keeper) GetAllPurchases(ctx sdk.Context) (purchases []types.Purchase) {
	k.IteratePurchaseListEntries(ctx, func(purchase types.Purchase) bool {
		purchases = append(purchases, purchase)
		return false
	})
	return
}

// IteratePurchaseListEntries iterates through entries of
// all purchase lists.
func (k Keeper) IteratePurchaseListEntries(ctx sdk.Context, callback func(purchase types.Purchase) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PurchaseListKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var purchaseList types.PurchaseList
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &purchaseList)

		for _, entry := range purchaseList.Entries {
			if callback(entry) {
				break
			}
		}
	}
}
