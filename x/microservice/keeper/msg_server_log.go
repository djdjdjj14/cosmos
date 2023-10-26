package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"microservice/x/microservice/types"
)

func (k msgServer) CreateLog(goCtx context.Context, msg *types.MsgCreateLog) (*types.MsgCreateLogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var log = types.Log{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
		Time:    msg.Time,
	}

	id := k.AppendLog(
		ctx,
		log,
	)

	return &types.MsgCreateLogResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateLog(goCtx context.Context, msg *types.MsgUpdateLog) (*types.MsgUpdateLogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var log = types.Log{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
		Time:    msg.Time,
	}

	// Checks that the element exists
	val, found := k.GetLog(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetLog(ctx, log)

	return &types.MsgUpdateLogResponse{}, nil
}

func (k msgServer) DeleteLog(goCtx context.Context, msg *types.MsgDeleteLog) (*types.MsgDeleteLogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetLog(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLog(ctx, msg.Id)

	return &types.MsgDeleteLogResponse{}, nil
}
