package iochainc_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochainc/testutil/keeper"
	"github.com/mytestlab123/iochainc/testutil/nullify"
	"github.com/mytestlab123/iochainc/x/iochainc"
	"github.com/mytestlab123/iochainc/x/iochainc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochaincKeeper(t)
	iochainc.InitGenesis(ctx, *k, genesisState)
	got := iochainc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
