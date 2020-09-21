package crown

import (
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/pack"
)
type GasEstimator = bitcoin.GasEstimator

var fee = pack.NewU64(100000)
