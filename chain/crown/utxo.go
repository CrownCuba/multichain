package crown

import 
(
	"fmt"

	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/api/utxo"
)

type (
	Client        = bitcoin.Client
	ClientOptions = bitcoin.ClientOptions
)

var (
	NewClient            = bitcoin.NewClient
	DefaultClientOptions = bitcoin.DefaultClientOptions
)

type Tx struct {
	inputs []utxo.Input
	recipients  []utxo.Recipient
	msgTx *wire.MsgTx
	params *ChainParams
	expiryHeight uint32
	signed bool
}

type TxBuilder struct {
	params *ChainParams
}

func NewTxBuilder(params *ChainParams) TxBuilder {
	return TxBuilder{params: params}
}

func (txBuilder TxBuilder) BuildTx(inputs []utxo.Input, recipients []utxo.Recipient) (utxo.Tx, error){
	msgTx := wire.NewMsgTx(bitcoin.Version)

	// Inputs
	for _, input := range inputs {
		hash := chainhash.Hash{}
		copy(hash[:], input.Hash)
		index := input.Index.Uint32()
		msgTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&hash, index), nil, nil))
	}

	// Outputs
	for _, recipient := range recipients {
		addr, err := DecodeAddress(string(recipient.To))
		if err != nil {
			return nil, err
		}
		script, err := txscript.PayToAddrScript(addr)
		if err != nil {
			return nil, err
		}
		value := recipient.Value.Int().Int64()
		if value < 0 {
			return nil, fmt.Errorf("expected value >= 0, got value %v", value)
		}
		msgTx.AddTxOut(wire.NewTxOut(value, script))
	}

	return &Tx{inputs: inputs, recipients: recipients, msgTx: msgTx, signed: false}, nil
}
