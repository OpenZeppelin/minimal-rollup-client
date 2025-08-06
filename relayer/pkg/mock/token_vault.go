package mock

import (
	"math/big"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TokenVault struct {
}

func (t *TokenVault) CanonicalToBridged(
	opts *bind.CallOpts,
	chainID *big.Int,
	canonicalAddress common.Address,
) (common.Address, error) {
	return relayer.ZeroAddress, nil
}
