package handler

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	proofProducer "github.com/OpenZeppelin/minimal-rollup-client/client/prover/proof_producer"
)

func (s *EventHandlerTestSuite) TestAssignmentExpiredEventHandlerHandle() {
	handler := NewAssignmentExpiredEventHandler(
		s.RPCClient,
		common.Address{},
		s.d.ProverSetAddress,
		make(chan *proofProducer.ProofRequestBody, 1024),
	)
	s.Nil(handler.Handle(context.Background(), s.ProposeAndInsertValidBlock(s.proposer, s.eventSyncer)))
}
