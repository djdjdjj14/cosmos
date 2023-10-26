package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"microservice/x/microservice/types"
)

// GetLogCount get the total number of log
func (k Keeper) GetLogCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LogCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLogCount set the total number of log
func (k Keeper) SetLogCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LogCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLog appends a log in the store with a new id and update the count
func (k Keeper) AppendLog(
	ctx sdk.Context,
	log types.Log,
) uint64 {
	// Create the log
	count := k.GetLogCount(ctx)

	// Set the ID of the appended value
	log.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LogKey))
	appendedValue := k.cdc.MustMarshal(&log)
	store.Set(GetLogIDBytes(log.Id), appendedValue)

	// Update log count
	k.SetLogCount(ctx, count+1)

	return count
}

// SetLog set a specific log in the store
func (k Keeper) SetLog(ctx sdk.Context, log types.Log) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LogKey))
	b := k.cdc.MustMarshal(&log)
	store.Set(GetLogIDBytes(log.Id), b)
}

// GetLog returns a log from its id
func (k Keeper) GetLog(ctx sdk.Context, id uint64) (val types.Log, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LogKey))
	b := store.Get(GetLogIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLog removes a log from the store
func (k Keeper) RemoveLog(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LogKey))
	store.Delete(GetLogIDBytes(id))
}

// GetAllLog returns all log
func (k Keeper) GetAllLog(ctx sdk.Context) (list []types.Log) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LogKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Log
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLogIDBytes returns the byte representation of the ID
func GetLogIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLogIDFromBytes returns ID in uint64 format from a byte array
func GetLogIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
