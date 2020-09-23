package crown

import (
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/pack"
)
type GasEstimator = bitcoin.GasEstimator

// Fixed fee amount since Crown doesn't allows too high fees

var Fee = pack.NewU64(100000)
