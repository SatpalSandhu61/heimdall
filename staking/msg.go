package staking

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	hmCommon "github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/types"
)

var cdc = codec.New()

const StakingRoute = "staking"

//
// Validator Join
//

var _ sdk.Msg = &MsgValidatorJoin{}

type MsgValidatorJoin struct {
	ValidatorAddress common.Address `json:"address"`
	SignerPubKey     types.PubKey   `json:"pubKey"`
	StartEpoch       uint64         `json:"startEpoch"`
	EndEpoch         uint64         `json:"endEpoch"`
	Amount           uint64         `json:"Amount"`
}

func NewMsgValidatorJoin(
	address common.Address,
	pubkey types.PubKey,
	startEpoch uint64,
	endEpoch uint64,
	amount uint64,
) MsgValidatorJoin {
	return MsgValidatorJoin{
		ValidatorAddress: address,
		SignerPubKey:     pubkey,
		StartEpoch:       startEpoch,
		EndEpoch:         endEpoch,
		Amount:           amount,
	}
}

func (msg MsgValidatorJoin) Type() string {
	return "validator-join"
}

func (msg MsgValidatorJoin) Route() string {
	return StakingRoute
}

func (msg MsgValidatorJoin) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, 0)
	return addrs
}

func (msg MsgValidatorJoin) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgValidatorJoin) ValidateBasic() sdk.Error {
	if bytes.Equal(msg.ValidatorAddress.Bytes(), helper.ZeroAddress.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator address %v", msg.ValidatorAddress.String())
	}

	if bytes.Equal(msg.SignerPubKey.Bytes(), helper.ZeroPubKey.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid pub key %v", msg.SignerPubKey.String())
	}

	return nil
}

func (msg MsgValidatorJoin) GetPower() uint64 {
	// add length checks
	return msg.Amount // TODO  Get power out of amount. Add 10^-18 here so that we dont overflow easily
}

//
// validator update
//

var _ sdk.Msg = &MsgSignerUpdate{}

// MsgSignerUpdate signer update struct
type MsgSignerUpdate struct {
	ValidatorAddress common.Address `json:"address"`
	NewSignerPubKey  types.PubKey   `json:"pubKey"`
	NewAmount        uint64         `json:"amount"`
}

func NewMsgValidatorUpdate(address common.Address, pubKey types.PubKey, amount uint64) MsgSignerUpdate {
	return MsgSignerUpdate{
		ValidatorAddress: address,
		NewSignerPubKey:  pubKey,
		NewAmount:        amount,
	}
}

func (msg MsgSignerUpdate) Type() string {
	return "validator-update"
}

func (msg MsgSignerUpdate) Route() string {
	return StakingRoute
}

func (msg MsgSignerUpdate) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, 0)
	return addrs
}

func (msg MsgSignerUpdate) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSignerUpdate) ValidateBasic() sdk.Error {
	if bytes.Equal(msg.ValidatorAddress.Bytes(), helper.ZeroAddress.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator address %v", msg.ValidatorAddress.String())
	}

	if bytes.Equal(msg.NewSignerPubKey.Bytes(), helper.ZeroPubKey.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid pub key %v", msg.NewSignerPubKey.String())
	}

	return nil
}

//
// validator exit
//

var _ sdk.Msg = &MsgValidatorExit{}

type MsgValidatorExit struct {
	ValidatorAddress common.Address
}

func NewMsgValidatorExit(address common.Address) MsgValidatorExit {
	return MsgValidatorExit{
		ValidatorAddress: address,
	}
}

func (msg MsgValidatorExit) Type() string {
	return "validator-exit"
}

func (msg MsgValidatorExit) Route() string {
	return StakingRoute
}

func (msg MsgValidatorExit) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, 0)
	return addrs
}

func (msg MsgValidatorExit) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgValidatorExit) ValidateBasic() sdk.Error {
	if bytes.Equal(msg.ValidatorAddress.Bytes(), helper.ZeroAddress.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator address %v", msg.ValidatorAddress.String())
	}

	return nil
}
