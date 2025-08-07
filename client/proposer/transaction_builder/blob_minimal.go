package builder

import (
	"context"

	"github.com/OpenZeppelin/minimal-rollup-client/client/blobencoding"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/core/types"
)

// BuildMinimalRollup implements the ProposeBatchTransactionBuilder interface.
// TODO: Move common blob builder logic from here to another package which
func (b *BlobTransactionBuilder) BuildMinimalRollup(
	ctx context.Context,
	txBatch []types.Transactions,
) (*txmgr.TxCandidate, error) {
	// TODO: Implement below:
	// 1. Encode the params included in the blob (as attributes - maybe the blobencoding package needs changes).
	// 2. Encode `publish()` call after adjusting BuildMinimalRollup args (needs encoding package changes)

	// TODO: Consider forced inclusion at the proposer layer and deal with it later.

	var allTxs types.Transactions
	for _, txs := range txBatch {
		allTxs = append(allTxs, txs...)
	}

	blobs, err := blobencoding.EncodeMinimalRollupBlob(&blobencoding.MinimalRollupBlobData{
		// Attrs: ?
		Transactions: allTxs,
	})
	if err != nil {
		return nil, err
	}

	return &txmgr.TxCandidate{
		// TxData:   data, // TODO: Need to encode this as explained blow
		Blobs:    blobs,
		To:       &b.taikoWrapperAddress, // TODO: Use different address?
		GasLimit: b.gasLimit,
	}, nil
}
