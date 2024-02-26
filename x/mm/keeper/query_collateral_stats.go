package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"cosmossdk.io/math"

	"github.com/kopi-money/kopi/x/mm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCollateralStats(ctx context.Context, req *types.GetCollateralStatsQuery) (*types.GetCollateralStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	totalUSD := math.LegacyZeroDec()
	var stats []*types.CollateralDenomStats

	for _, denom := range k.DenomKeeper.GetCollateralDenoms(ctx) {
		sum := k.getCollateralSum(ctx, denom.Denom)
		sumUSD, err := k.DexKeeper.GetValueInUSD(ctx, denom.Denom, sum)
		if err != nil {
			return nil, err
		}

		totalUSD = totalUSD.Add(sumUSD)

		var priceUSD math.LegacyDec
		if sum.GT(math.ZeroInt()) {
			priceUSD = sumUSD.Quo(sum.ToLegacyDec())
		} else {
			priceUSD, _ = k.DexKeeper.GetPriceInUSD(ctx, denom.Denom)
		}

		stats = append(stats, &types.CollateralDenomStats{
			Denom:     denom.Denom,
			Amount:    sum.String(),
			AmountUsd: sumUSD.String(),
			Ltv:       denom.Ltv.String(),
			PriceUsd:  priceUSD.String(),
		})
	}

	return &types.GetCollateralStatsResponse{Stats: stats, TotalUsd: totalUSD.String()}, nil
}

func (k Keeper) GetCollateralDenomStats(ctx context.Context, req *types.GetCollateralDenomStatsQuery) (*types.GetCollateralDenomStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	denom := k.DenomKeeper.GetCollateralDenom(ctx, req.Denom)
	if denom == nil {
		return nil, types.ErrInvalidCollateralDenom
	}

	sum := math.ZeroInt()
	collaterals := []*types.UserCollateral{}

	for _, collateral := range k.GetAllCollaterals(ctx, denom.Denom) {
		sum = sum.Add(collateral.Amount)

		collaterals = append(collaterals, &types.UserCollateral{
			Address: collateral.Address,
			Amount:  collateral.Amount.String(),
		})
	}

	sumUSD, err := k.DexKeeper.GetValueInUSD(ctx, denom.Denom, sum)
	if err != nil {
		return nil, err
	}

	return &types.GetCollateralDenomStatsResponse{
		UserCollateral: collaterals,
		Sum:            sum.String(),
		SumUsd:         sumUSD.String(),
	}, nil
}

func (k Keeper) GetCollateralUserStats(ctx context.Context, req *types.GetCollateralUserStatsQuery) (*types.GetCollateralStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	totalUSD := math.LegacyZeroDec()
	var stats []*types.CollateralDenomStats

	for _, denom := range k.DenomKeeper.GetCollateralDenoms(ctx) {
		amount := k.getCollateralDenomForAddress(ctx, denom.Denom, req.Address)
		sumUSD, err := k.DexKeeper.GetValueInUSD(ctx, denom.Denom, amount)
		if err != nil {
			continue
			//return nil, err
		}

		priceUSD, err := k.DexKeeper.GetPriceInUSD(ctx, denom.Denom)
		if err != nil {
			return nil, err
		}

		totalUSD = totalUSD.Add(sumUSD)

		stats = append(stats, &types.CollateralDenomStats{
			Denom:     denom.Denom,
			Amount:    amount.String(),
			AmountUsd: sumUSD.String(),
			Ltv:       denom.Ltv.String(),
			PriceUsd:  priceUSD.String(),
		})
	}

	return &types.GetCollateralStatsResponse{Stats: stats, TotalUsd: totalUSD.String()}, nil
}

func (k Keeper) GetCollateralDenomUserStats(ctx context.Context, req *types.GetCollateralDenomUserStatsQuery) (*types.GetCollateralDenomUserStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	amount, found := k.GetCollateral(ctx, req.Denom, req.Address)
	if !found {
		amount.Amount = math.ZeroInt()
	}

	return &types.GetCollateralDenomUserStatsResponse{Amount: amount.Amount.String()}, nil
}

func (k Keeper) GetWithdrawableCollateral(ctx context.Context, req *types.GetWithdrawableCollateralQuery) (*types.GetWithdrawableCollateralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	withdrawable, err := k.CalcWithdrawableCollateralAmount(ctx, req.Address, req.Denom)
	if err != nil {
		return nil, errors.Wrap(err, "could not calculate withdrawable amount")
	}

	withdrawableUSD, err := k.DexKeeper.GetValueInUSD(ctx, req.Denom, withdrawable)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert withdrawable amount to usd")
	}

	return &types.GetWithdrawableCollateralResponse{
		Amount:    withdrawable.String(),
		AmountUsd: withdrawableUSD.String(),
	}, nil
}
