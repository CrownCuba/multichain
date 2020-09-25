package crown

import (
	"context" 
	"math/big"  
	"github.com/renproject/pack"
)

// Fixed fee amount since Crown doesn't allows too high fees

var Gas = big.NewInt(10000)

type Estimator struct {} 

// Implementation to get the real Gas from fees

func (Estimator) EstimateGasPrice(ctx context.Context) (pack.U256, error) { 
return pack.NewU256FromInt(Gas), nil
}
