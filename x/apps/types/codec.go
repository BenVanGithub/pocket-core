package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterStructure(MsgAppStake{}, "apps/MsgAppStake")
	cdc.RegisterStructure(LegacyMsgAppStake{}, "apps/LegacyMsgAppStake")
	cdc.RegisterStructure(MsgBeginAppUnstake{}, "apps/MsgAppBeginUnstake")
	cdc.RegisterStructure(MsgAppUnjail{}, "apps/MsgAppUnjail")
	cdc.RegisterImplementation((*sdk.Msg)(nil), &MsgAppStake{}, &MsgBeginAppUnstake{}, &MsgAppUnjail{}, LegacyMsgAppStake{})
	ModuleCdc = cdc
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.NewCodec(types.NewInterfaceRegistry())
	RegisterCodec(ModuleCdc)
	crypto.RegisterAmino(ModuleCdc.AminoCodec().Amino)
}
