package crown

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

type ChainParams struct {
	*chaincfg.Params
	PubKeyHashAddrIDs []byte
	ScriptHashAddrIDs []byte
}

var (
	genMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
		0x46, 0x14, 0x3d, 0xca, 0xe4, 0x50, 0xe5, 0xac, 0xfe,
		0x51, 0x92, 0x95, 0x8f, 0x95, 0x6d, 0xc3, 0x6c, 0x14,
		0x77, 0xef, 0x66, 0xdb, 0x92, 0xb1, 0x8d, 0xab, 0xa9,
		0x18, 0x61, 0x35, 0xad, 0x80,
	})

	genBlockHash = chainhash.Hash([chainhash.HashSize]byte{
		0xda, 0xc6, 0x1d, 0x9a, 0xe6, 0xd7, 0x4f, 0x81, 0xcb,
		0x13, 0x8d, 0x8c, 0xf7, 0x3d, 0xff, 0x14, 0x86, 0xc6,
		0x19, 0xab, 0xf4, 0x64, 0x2f, 0x12, 0x5e, 0xd, 0x37,
		0x85, 0x0, 0x0, 0x0, 0x0,
	})
	genBlock = wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:    1,
			PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
			MerkleRoot: genMerkleRoot,            // 80ad356118a9ab8db192db66ef77146cc36d958f959251feace550e4ca3d1446
			Timestamp:  time.Unix(1412760826, 0), // 8th Oct 2014 09:33:46
			Bits:       0x1d00ffff,
			Nonce:      1612467894,
		},
	}
	
	MainParams = chaincfg.Params{
		Name:        "mainnet",
		Net:         0xdfb3ebb8,
		DefaultPort: "9340",
		DNSSeeds: []chaincfg.DNSSeed{
			{Host: "dnsseed1.crowncoin.org", HasFiltering: false},
			{Host: "dnsseed2.crowncoin.org", HasFiltering: false},
		},
		GenesisBlock:             &genBlock,
		GenesisHash:              &genBlockHash,
		TargetTimespan:           time.Hour * 24 * 14, //two weeks
		TargetTimePerBlock:       time.Minute,
		SubsidyReductionInterval: 2100000,
		Checkpoints: []chaincfg.Checkpoint{	
			{Height: 0, Hash: strToHash("0000000085370d5e122f64f4ab19c68614ff3df78c8d13cb814fd7e69a1dc6da")},
			{Height: 100000, Hash: strToHash("000000000001f9595f38d13a62f030c877717db91659157eb732e2a89c8f9c1d")},
			{Height: 200000, Hash: strToHash("000000000000e0438b3279cecc380a96c8a7b40bceb94fc03d0a210fdda2b959")},
			{Height: 300000, Hash: strToHash("00000000000059901c2ea32bc486d91a355a8fe362e34fc3d10c45bb5e5ca79d")},
			{Height: 400000, Hash: strToHash("0000000000001b97fb8367b435b4554cb9d5438d28f03c4259df7ed8854fe946")},
			{Height: 500000, Hash: strToHash("b53d68a141c9ced04eeca5624b66665a58732c48d383f81a29cf80a8a57186ff")},
			{Height: 600000, Hash: strToHash("84e2277d1dc957ae41869498311937bdcedce7e48fa33f962b9b9e9c16df5410")},
			{Height: 700000, Hash: strToHash("75ec82017af651ac1383b623898d6f6b57d0500137913fd000f928b8bd409146")},
			{Height: 800000, Hash: strToHash("64412d45320a7f2394b009bcf00bc841e60c3e0680dbfc45176568699af5bdec")},
			{Height: 900000, Hash: strToHash("bdf0bcfe3ada671b64526854af8cb7f6f52c1489446e26268e70fbee1e72be5f")},
			{Height: 1000000, Hash: strToHash("cb324809eef485d0243d2210dd15300a941fece86a19e858383429c80cf37b0b")},
			{Height: 1100000, Hash: strToHash("f172bdb7a894b9e055eae4b1b2e8d91aba9404d24fa808985346cc7c8eea35a6")},
		},
		PrivateKeyID:     128,
		CoinbaseMaturity: 100,
	}
	TestParams = chaincfg.Params{
		Name:                     "testnet",
		DefaultPort:              "19340",
		GenesisBlock:             &genBlock,
		GenesisHash:              &genBlockHash,
		TargetTimespan:           time.Hour * 48,
		TargetTimePerBlock:       time.Second * 90,
		SubsidyReductionInterval: 130000,
		PrivateKeyID:             239,
	}

)
var MainNetParams = ChainParams{
	Params:            &MainParams,
	PubKeyHashAddrIDs: []byte{0x01, 0x75, 0x07},
	ScriptHashAddrIDs: []byte{0x01, 0x74, 0xF1},
}
var TesnetParams = ChainParams{
	Params: &TestParams,
	PubKeyHashAddrIDs: []byte{0x01,0x7A,0xCD,0x67},
	ScriptHashAddrIDs: []byte{0x01,0x7A,0xCD,0x51},
}
func strToHash(Str string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(Str)
	if err != nil {
		panic(err)		
	}
	return hash
}
