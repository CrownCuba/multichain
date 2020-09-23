package crown

import 
(
	"fmt"
	"bytes"
	"math/big"

	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/api/utxo"
	"github.com/renproject/pack"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/btcec"
)

type (
	Client        = bitcoin.Client
	ClientOptions = bitcoin.ClientOptions
)

var (
	NewClient            = bitcoin.NewClient
	DefaultClientOptions = bitcoin.DefaultClientOptions
)


// Version of Crown transactions supported by the multichain.

const Version = 1

// Tx represents a simple Crown transaction that implements the Bitcoin Compat
// API.

type Tx struct {
	inputs []utxo.Input
	recipients  []utxo.Recipient
	msgTx *wire.MsgTx
	params *ChainParams
	expiryHeight uint32
	signed bool
}

func (tx *Tx) Hash() (pack.Bytes, error) {
	txhash := tx.msgTx.TxHash()
	return pack.NewBytes(txhash[:]), nil
}

func (tx *Tx) Inputs() ([]utxo.Input, error) {
	return tx.inputs, nil
}

func (tx *Tx) Outputs() ([]utxo.Output, error) {
	hash, err := tx.Hash()
	if err != nil {
		return nil, fmt.Errorf("bad hash: %v", err)
	}
	outputs := make([]utxo.Output, len(tx.msgTx.TxOut))
	for i := range outputs {
		outputs[i].Outpoint = utxo.Outpoint{
			Hash:  hash,
			Index: pack.NewU32(uint32(i)),
		}
		outputs[i].PubKeyScript = pack.Bytes(tx.msgTx.TxOut[i].PkScript)
		if tx.msgTx.TxOut[i].Value < 0 {
			return nil, fmt.Errorf("bad output %v: value is less than zero", i)
		}
		outputs[i].Value = pack.NewU256FromU64(pack.NewU64(uint64(tx.msgTx.TxOut[i].Value)))
	}
	return outputs, nil
}

// Sighashes returns the digests that must be signed before the transaction
// can be submitted by the client.

func (tx *Tx) Sighashes() ([]pack.Bytes32, error) {
	sighashes := make([]pack.Bytes32, len(tx.inputs))
  
	for i, txin := range tx.inputs {
	  pubKeyScript := txin.PubKeyScript
	  sigScript := txin.SigScript
	  value := txin.Value.Int().Int64()
	  if value < 0 {
		return []pack.Bytes32{}, fmt.Errorf("expected value >= 0, got value %v", value)
	  }
  
	  var hash []byte
	  var err error
	  if sigScript == nil {
		hash, err = txscript.CalcSignatureHash(pubKeyScript, txscript.SigHashAll, tx.msgTx, i)
	  } else {
		hash, err = txscript.CalcSignatureHash(sigScript, txscript.SigHashAll, tx.msgTx, i)
		
	  }
	  if err != nil {
		return []pack.Bytes32{}, err
	  }
  
	  sighash := [32]byte{}
	  copy(sighash[:], hash)
	  sighashes[i] = pack.NewBytes32(sighash)
	}
  
	return sighashes, nil
}

// Signs the built transaction.

func (tx *Tx) Sign(signatures []pack.Bytes65, pubKey pack.Bytes) error {
	if tx.signed {
		return fmt.Errorf("already signed")
	}
	if len(signatures) != len(tx.msgTx.TxIn) {
		return fmt.Errorf("expected %v signatures, got %v signatures", len(tx.msgTx.TxIn), len(signatures))
	}

	for i, rsv := range signatures {
		var err error

		// Decode the signature and the pubkey script.
		r := new(big.Int).SetBytes(rsv[:32])
		s := new(big.Int).SetBytes(rsv[32:64])
		signature := btcec.Signature{
			R: r,
			S: s,
		}

		sigScript := tx.inputs[i].SigScript

		// Support non-segwit Crown desn't have segwit and it's not needed
		builder := txscript.NewScriptBuilder()
		builder.AddData(append(signature.Serialize(), byte(txscript.SigHashAll)))
		builder.AddData(pubKey)
		if sigScript != nil {
			builder.AddData(sigScript)
		}
		tx.msgTx.TxIn[i].SignatureScript, err = builder.Script()
		if err != nil {
			return err
		}
	}

	tx.signed = true
	return nil
}

func (tx *Tx) Serialize() (pack.Bytes, error) {	
	buf := new(bytes.Buffer)
	if err := tx.msgTx.Serialize(buf); err != nil {
		return pack.Bytes{}, err
	}
	return pack.NewBytes(buf.Bytes()), nil
}

type TxBuilder struct {
	params *ChainParams
}

// NewTxBuilder returns a transaction builder that builds UTXO-compatible
// Crown transactions for the given chain configuration.

func NewTxBuilder(params *ChainParams) TxBuilder {
	return TxBuilder{params: params}
}

// BuildTx returns a Crown transaction that consumes funds from the given
// inputs, and sends them to the given recipients. The fee is paid to
// the Crown network is fixed to 0.01 CRW since the network doesn't accepts high fees.
// Outputs produced for recipients will use P2PKH, P2SH
// scripts as the pubkey script, based on the format of the recipient address.

func (txBuilder TxBuilder) BuildTx(inputs []utxo.Input, recipients []utxo.Recipient) (utxo.Tx, error){
	msgTx := wire.NewMsgTx(Version)
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
		script, err := txscript.PayToAddrScript(addr.BitcoinAddress())
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
