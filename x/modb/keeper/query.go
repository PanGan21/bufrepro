package keeper

import (
	"bufrepro/x/modb/types"
)

var _ types.QueryServer = Keeper{}
