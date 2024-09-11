package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m Migrator) V2Migration(ctx sdk.Context) error {
	params := m.keeper.GetLegacyParams(ctx)
	m.keeper.SetParams(ctx, params)
	return nil
}