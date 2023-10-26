package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"microservice/x/microservice/types"
)

func (k Keeper) LogAll(goCtx context.Context, req *types.QueryAllLogRequest) (*types.QueryAllLogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var logs []types.Log
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	logStore := prefix.NewStore(store, types.KeyPrefix(types.LogKey))

	pageRes, err := query.Paginate(logStore, req.Pagination, func(key []byte, value []byte) error {
		var log types.Log
		if err := k.cdc.Unmarshal(value, &log); err != nil {
			return err
		}

		logs = append(logs, log)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLogResponse{Log: logs, Pagination: pageRes}, nil
}

func (k Keeper) Log(goCtx context.Context, req *types.QueryGetLogRequest) (*types.QueryGetLogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	log, found := k.GetLog(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLogResponse{Log: log}, nil
}
