package keeper

import (
	"microservice/x/microservice/types"
)

var _ types.QueryServer = Keeper{}
