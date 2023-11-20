package types_test

import (
	"errors"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/elys-network/elys/x/margin/types"
	"github.com/stretchr/testify/assert"

	ptypes "github.com/elys-network/elys/x/parameter/types"
)

func TestCheckLongAssets_InvalidAssets(t *testing.T) {
	err := types.CheckLongAssets(ptypes.BaseCurrency, ptypes.BaseCurrency, ptypes.BaseCurrency)
	assert.True(t, errors.Is(err, sdkerrors.Wrap(types.ErrInvalidBorrowingAsset, "invalid borrowing asset")))

	err = types.CheckLongAssets(ptypes.ATOM, ptypes.BaseCurrency, ptypes.BaseCurrency)
	assert.True(t, errors.Is(err, sdkerrors.Wrap(types.ErrInvalidBorrowingAsset, "invalid borrowing asset")))
}

func TestCheckLongAssets_ValidAssets(t *testing.T) {
	err := types.CheckLongAssets(ptypes.BaseCurrency, ptypes.ATOM, ptypes.BaseCurrency)
	assert.Nil(t, err)

	err = types.CheckLongAssets(ptypes.ATOM, ptypes.ATOM, ptypes.BaseCurrency)
	assert.Nil(t, err)
}