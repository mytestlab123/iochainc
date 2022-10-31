package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochainc/testutil/keeper"
	"github.com/mytestlab123/iochainc/x/iochainc/keeper"
	"github.com/mytestlab123/iochainc/x/iochainc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IochaincKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
