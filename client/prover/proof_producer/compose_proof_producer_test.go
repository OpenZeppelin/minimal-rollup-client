package producer

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/OpenZeppelin/minimal-rollup-client/client/bindings/metadata"
)

func TestComposeProducerRequestProof(t *testing.T) {
	var (
		producer = &ComposeProofProducer{
			Dummy:              true,
			DummyProofProducer: DummyProofProducer{},
			SgxGethProducer:    &SgxGethProofProducer{Dummy: true},
		}
		blockID = common.Big32
	)
	res, err := producer.RequestProof(
		context.Background(),
		&ProofRequestOptionsPacaya{},
		blockID,
		&metadata.TaikoDataBlockMetadataPacaya{},
		time.Now(),
	)
	require.Nil(t, err)

	require.Equal(t, res.BatchID, blockID)
	require.NotEmpty(t, res.Proof)
}
