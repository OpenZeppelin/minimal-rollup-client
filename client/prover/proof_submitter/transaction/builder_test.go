package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/OpenZeppelin/minimal-rollup-client/client/bindings/metadata"
	pacayaBindings "github.com/OpenZeppelin/minimal-rollup-client/client/bindings/pacaya"
	"github.com/OpenZeppelin/minimal-rollup-client/client/internal/testutils"
	producer "github.com/OpenZeppelin/minimal-rollup-client/client/prover/proof_producer"
)

func (s *TransactionTestSuite) TestBuildTxs() {
	header, err := s.RPCClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.NotNil(header)

	builder := s.builder.BuildProveBatchesPacaya(&producer.BatchProofs{
		ProofResponses: []*producer.ProofResponse{{
			BatchID:   common.Big1,
			Meta:      metadata.NewTaikoDataBlockMetadataPacaya(&pacayaBindings.TaikoInboxClientBatchProposed{}),
			Proof:     testutils.RandomBytes(100),
			Opts:      &producer.ProofRequestOptionsPacaya{Headers: []*types.Header{header}},
			ProofType: producer.ProofTypeOp,
		}},
	})
	_, err = builder(&bind.TransactOpts{})
	s.Nil(err)
}
