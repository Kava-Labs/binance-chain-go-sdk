package types

import (
	ntypes "github.com/kava-labs/binance-chain-go-sdk/common/types"
	"github.com/kava-labs/binance-chain-go-sdk/types/tx"
	"github.com/kava-labs/go-amino"
	types "github.com/kava-labs/tendermint/rpc/core/types"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}
