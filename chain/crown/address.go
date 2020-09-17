package crown

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

type Address interface {
	btcutil.Address
	BitcoinAddress() btcutil.Address
}
type AddressPubKeyHash struct {
	*btcutil.AddressPubKeyHash
	params *ChainParams
}
type AddressScriptHash struct {
	*btcutil.AddressScriptHash
	params *ChainParams
}

func NewAddressPubKeyHash(hash []byte, params *ChainParams) (AddressPubKeyHash, error) {
	address, err := btcutil.NewAddressPubKeyHash(hash, params.Params)
	return AddressPubKeyHash{AddressPubKeyHash: address, params: params}, err
}
func NewAddressScriptHash(script []byte, params *ChainParams) (AddressScriptHash, error) {
	address, err := btcutil.NewAddressScriptHash(script, params.Params)
	return AddressScriptHash{AddressScriptHash: address, params: params}, err
}

//String method
func (address AddressScriptHash) String() string {
	return address.EncodeAddress()
}

func (address AddressPubKeyHash) String() string {
	return address.EncodeAddress()
}

//Encode
func (address AddressPubKeyHash) EncodeAddress() string {
	return encode((*address.AddressPubKeyHash.Hash160())[:], address.params.PubKeyHashAddrIDs)
}
func (address AddressScriptHash) EncodeAddress() string {
	return encode((*address.AddressScriptHash.Hash160())[:], address.params.ScriptHashAddrIDs)
}
func encode(hash, prefix []byte) string {
	body := append(prefix, hash...)
	chsum := checksum(body)
	body = append(body, chsum[:]...)
	return base58.Encode(body)
}

//Decode
func DecodeAddress(address string) (Address, error) {
	var decoded = base58.Decode(address)
	if len(decoded) != 27 && len(decoded) != 28 {
		fmt.Println("Inside the error " , address)
		return nil, base58.ErrInvalidFormat
	}
	var chsum [4]byte
	endbody := len(decoded) - 4
	body := decoded[:endbody]
	copy(chsum[:], decoded[endbody:])
	if checksum(body) != chsum {
		return nil, base58.ErrChecksum
	}
	return getAddress(decoded)

}
func getAddress(decoded []byte) (Address, error) {
	var (
		pubHash [20]byte
		params  *ChainParams
	)
	indPrefix := len(decoded) - 24
	prefix := decoded[:indPrefix]
	copy(pubHash[:], decoded[indPrefix:len(decoded)-4])
	params = &MainNetParams
	if indPrefix == 4 {
		params = &TesnetParams
	}
	if bytes.Equal(prefix, params.PubKeyHashAddrIDs) {
		return NewAddressPubKeyHash(pubHash[:], params)
	}
	if bytes.Equal(prefix, params.ScriptHashAddrIDs) {
		return NewAddressScriptHash(pubHash[:], params)
	}
	return nil, errors.New("Invalid address")
}

//

func checksum(body []byte) (chsum [4]byte) {
	//Doble hashing the checksum
	hash1 := sha256.Sum256(body)
	hash2 := sha256.Sum256(hash1[:])
	copy(chsum[:], hash2[:4])
	return
}

//IsForNet
func (address AddressPubKeyHash) IsForNet(params *chaincfg.Params) bool {
	return address.AddressPubKeyHash.IsForNet(params)
}
func (address AddressScriptHash) IsForNet(params *chaincfg.Params) bool {
	return address.AddressScriptHash.IsForNet(params)
}

//BitcoinAddress()
func (address AddressPubKeyHash) BitcoinAddress() btcutil.Address {
	return address.AddressPubKeyHash
}
func (address AddressScriptHash) BitcoinAddress() btcutil.Address {
	return address.AddressScriptHash
}

//ScriptAddress()
func (address AddressPubKeyHash) ScriptAddress() []byte {
	return address.AddressPubKeyHash.ScriptAddress()
}
func (address AddressScriptHash) ScriptAddress() []byte {
	return address.AddressScriptHash.ScriptAddress()
}
