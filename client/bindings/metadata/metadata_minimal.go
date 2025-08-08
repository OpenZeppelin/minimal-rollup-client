package metadata

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/OpenZeppelin/minimal-rollup-client/client/bindings/minimal"
	minimalBindings "github.com/OpenZeppelin/minimal-rollup-client/client/bindings/minimal"

	"github.com/ethereum/go-ethereum/log"
)

// Ensure MinimalDataPublication implements TaikoBlockMetaData.
var _ MinimalPublicationData = (*MinimalDataPublication)(nil)

type MinimalDataPublication struct {
	publicationHash common.Hash
	header          minimal.IInboxPublicationHeader
	attributes      MinimalPublicationAttributes
	types.Log
}

// NewMinimalDataBlockMetadata creates a new instance of NewMinimalDataBlockMetadata
// from the TaikoInbox.Published event.
func NewMinimalDataBlockMetadata(e *minimalBindings.TaikoInboxPublished) *MinimalDataPublication {

	var taikoPubAttr MinimalPublicationAttributes
	if len(e.Attributes) == 2 {
		publicationMetadata, err := DecodePublicationMetadata(e.Attributes[0])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		publicationBlobRef, err := DecodeBlobRef(e.Attributes[1])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		taikoPubAttr = MinimalPublicationAttributes{
			Metadata: *publicationMetadata,
			BlobRef:  *publicationBlobRef,
		}

	}

	return &MinimalDataPublication{
		publicationHash: e.PubHash,
		header:          e.Header,
		attributes:      taikoPubAttr,
		Log:             e.Raw,
	}
}

func (m *MinimalDataPublication) Header() minimal.IInboxPublicationHeader {
	return m.header
}

func (m *MinimalDataPublication) PublicationHash() common.Hash {
	return m.publicationHash
}

func (m *MinimalDataPublication) Attributes() MinimalPublicationAttributes {
	return m.attributes
}

// GetRawBlockHeight returns the raw L1 block height.
func (m *MinimalDataPublication) GetRawBlockHeight() *big.Int {
	return new(big.Int).SetUint64(m.BlockNumber)
}

// GetRawBlockHash returns the raw L1 block hash.
func (m *MinimalDataPublication) GetRawBlockHash() common.Hash {
	return m.BlockHash
}

// GetTxIndex returns the transaction index.
func (m *MinimalDataPublication) GetTxIndex() uint {
	return m.Log.TxIndex
}

// GetTxHash returns the transaction hash.
func (m *MinimalDataPublication) GetTxHash() common.Hash {
	return m.Log.TxHash
}
