package proof

import (
	"context"
	"testing"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/encoding"
	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/go-playground/assert.v1"
)

func Test_blockHeader(t *testing.T) {
	p := newTestProver()

	header, err := p.blockHeader(context.Background(), p.blocker, common.HexToHash("0x123"))
	assert.Equal(t, err, nil)
	assert.Equal(t, header, encoding.BlockToBlockHeader(types.NewBlockWithHeader(mock.Header)))
}

func Test_blockHeader_noHash(t *testing.T) {
	p := newTestProver()

	_, err := p.blockHeader(context.Background(), p.blocker, common.HexToHash("0x"))
	assert.Equal(t, err, nil)
}
