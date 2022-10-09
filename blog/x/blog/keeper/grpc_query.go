package keeper

import (
	"github.com/amedumer/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
