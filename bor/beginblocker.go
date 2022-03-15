package bor

import (
	"fmt"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/bor/consensus/bor"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/maticnetwork/heimdall/bor/client/rest"
	"github.com/maticnetwork/heimdall/helper"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {

	if ctx.BlockHeight() == int64(helper.SpanOverrideBlockHeight) {
		k.Logger(ctx).Info("overriding span BeginBlocker", "height", ctx.BlockHeight())
		j, ok := rest.SPAN_OVERRIDES[helper.GenesisDoc.ChainID]
		if !ok {
			k.Logger(ctx).Error("Error in fetching span overrides")
		}

		var spans []*bor.ResponseWithHeight
		if err := json.Unmarshal(j, &spans); err != nil {
			k.Logger(ctx).Error("Error Unmarshal spans", "error", err)
		}

		for _, span := range spans {
			fmt.Println("----- span", span)
			var heimdallSpan hmTypes.Span
			if err := json.Unmarshal(span.Result, &heimdallSpan); err != nil {
				k.Logger(ctx).Error("Error Unmarshal heimdallSpan", "error", err)
			}

			if err := k.AddNewRawSpan(ctx, heimdallSpan); err != nil {
				k.Logger(ctx).Error("Error AddNewRawSpan", "error", err)
			}
			k.UpdateLastSpan(ctx, heimdallSpan.ID)
		}
	}
}
