package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"microservice/x/msg/types"
)

// SetPost set a specific post in the store from its index
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKeyPrefix))
	b := k.cdc.MustMarshal(&post)
	store.Set(types.PostKey(
		post.Index,
	), b)
}

// GetPost returns a post from its index
func (k Keeper) GetPost(
	ctx sdk.Context,
	index string,

) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKeyPrefix))

	b := store.Get(types.PostKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePost removes a post from the store
func (k Keeper) RemovePost(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKeyPrefix))
	store.Delete(types.PostKey(
		index,
	))
}

// GetAllPost returns all post
func (k Keeper) GetAllPost(ctx sdk.Context) (list []types.Post) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
