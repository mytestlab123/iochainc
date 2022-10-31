package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochainc/testutil/keeper"
	"github.com/mytestlab123/iochainc/x/iochainc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochaincKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
