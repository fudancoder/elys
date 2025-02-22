package keeper

import (
	sdkmath "cosmossdk.io/math"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/perpetual/types"
)

func (k Keeper) FundingFeeDistribution(ctx sdk.Context, mtp *types.MTP, pool *types.Pool) (sdkmath.Int, error) {

	totalLongOpenInterest := pool.GetTotalLongOpenInterest()
	totalShortOpenInterest := pool.GetTotalShortOpenInterest()

	// Total fund collected should be
	long, short := k.GetFundingDistributionValue(ctx, mtp.LastFundingCalcBlock, pool.AmmPoolId)
	var totalFund sdkmath.LegacyDec
	// calc funding fee share
	var fundingFeeShare sdkmath.LegacyDec
	amountDistributed := sdkmath.ZeroInt()
	if mtp.Position == types.Position_LONG {
		// Ensure liabilitiesLong is not zero to avoid division by zero
		if totalLongOpenInterest.IsZero() {
			return amountDistributed, errors.New("totalLongOpenInterest in FundingFeeDistribution cannot be zero")
		}
		fundingFeeShare = mtp.Custody.ToLegacyDec().Quo(totalLongOpenInterest.ToLegacyDec())
		totalFund = short

		// if funding fee share is zero, skip mtp
		if fundingFeeShare.IsZero() || totalFund.IsZero() {
			return amountDistributed, nil
		}

		// calculate funding fee amount
		fundingFeeAmount := totalFund.Mul(fundingFeeShare)

		amountDistributed = fundingFeeAmount.TruncateInt()
		// update mtp custody
		mtp.Custody = mtp.Custody.Add(fundingFeeAmount.TruncateInt())

		// decrease fees collected
		err := pool.UpdateFeesCollected(mtp.CustodyAsset, fundingFeeAmount.TruncateInt(), false)
		if err != nil {
			return sdkmath.ZeroInt(), err
		}

		// update pool custody balance
		err = pool.UpdateCustody(mtp.CustodyAsset, fundingFeeAmount.TruncateInt(), true, mtp.Position)
		if err != nil {
			return sdkmath.ZeroInt(), err
		}

		// add payment to total funding fee paid in custody asset
		mtp.FundingFeeReceivedCustody = mtp.FundingFeeReceivedCustody.Add(fundingFeeAmount.TruncateInt())
	} else {
		// Ensure liabilitiesShort is not zero to avoid division by zero
		if totalShortOpenInterest.IsZero() {
			return amountDistributed, errors.New("totalShortOpenInterest in FundingFeeDistribution cannot be zero")
		}
		fundingFeeShare = mtp.Liabilities.ToLegacyDec().Quo(totalShortOpenInterest.ToLegacyDec())
		totalFund = long

		// if funding fee share is zero, skip mtp
		if fundingFeeShare.IsZero() || totalFund.IsZero() {
			return amountDistributed, nil
		}

		// calculate funding fee amount
		fundingFeeAmount := totalFund.Mul(fundingFeeShare).TruncateInt()

		// adding case for fundingFeeAmount being smaller tha 10^-18
		if fundingFeeAmount.IsZero() {
			return amountDistributed, nil
		}
		// decrease fees collected
		err := pool.UpdateFeesCollected(mtp.LiabilitiesAsset, fundingFeeAmount, false)
		if err != nil {
			return amountDistributed, err
		}

		tradingAssetPrice, err := k.GetAssetPrice(ctx, mtp.TradingAsset)
		if err != nil {
			return amountDistributed, err
		}

		// For short, fundingFeeAmount is in trading asset, need to convert to custody asset which is in usdc
		custodyAmt := fundingFeeAmount.ToLegacyDec().Mul(tradingAssetPrice).TruncateInt()

		amountDistributed = custodyAmt
		// update mtp Custody
		mtp.Custody = mtp.Custody.Add(custodyAmt)

		// update pool liability balance
		err = pool.UpdateCustody(mtp.CustodyAsset, custodyAmt, true, mtp.Position)
		if err != nil {
			return sdkmath.ZeroInt(), err
		}

		// add payment to total funding fee paid in custody asset
		mtp.FundingFeeReceivedCustody = mtp.FundingFeeReceivedCustody.Add(custodyAmt)
	}

	return amountDistributed, nil
}
