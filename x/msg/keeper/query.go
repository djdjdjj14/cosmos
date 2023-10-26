package keeper

import (
	"microservice/x/msg/types"
)

var _ types.QueryServer = Keeper{}
