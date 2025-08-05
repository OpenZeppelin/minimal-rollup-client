package handler

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"

	pacayaBindings "github.com/OpenZeppelin/minimal-rollup-client/client/bindings/pacaya"
	"github.com/OpenZeppelin/minimal-rollup-client/client/internal/testutils"
)

func (s *EventHandlerTestSuite) TestBatchesVerifiedHandle() {
	handler := NewBatchesVerifiedEventHandler(s.RPCClient)
	id := testutils.RandomHash().Big().Uint64()
	s.NotPanics(func() {
		s.NotNil(handler.HandlePacaya(context.Background(), &pacayaBindings.TaikoInboxClientBatchesVerified{
			BatchId: testutils.RandomHash().Big().Uint64(),
			Raw:     types.Log{BlockHash: testutils.RandomHash(), BlockNumber: id},
		}))
	})
}
