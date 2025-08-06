package proof

import (
	"testing"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer"
	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/encoding"
	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/mock"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/go-playground/assert.v1"
)

func newTestProver() *Prover {
	return &Prover{
		blocker:     &mock.Blocker{},
		cacheOption: encoding.CACHE_BOTH,
	}
}

func Test_New(t *testing.T) {
	tests := []struct {
		name    string
		blocker blocker
		wantErr error
	}{
		{
			"success",
			&ethclient.Client{},
			nil,
		},
		{
			"noEthClient",
			nil,
			relayer.ErrNoEthClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.blocker, encoding.CACHE_BOTH)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
