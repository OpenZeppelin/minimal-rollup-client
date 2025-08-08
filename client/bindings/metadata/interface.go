package metadata

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	minimalBindings "github.com/OpenZeppelin/minimal-rollup-client/client/bindings/minimal"
	pacayaBindings "github.com/OpenZeppelin/minimal-rollup-client/client/bindings/pacaya"
)

// TaikoProposalMetaData defines all the metadata of a Taiko block.
type TaikoProposalMetaData interface {
	Pacaya() TaikoBatchMetaDataPacaya
	IsPacaya() bool
	GetRawBlockHeight() *big.Int
	GetRawBlockHash() common.Hash
	GetTxIndex() uint
	GetTxHash() common.Hash
	GetProposer() common.Address
	GetCoinbase() common.Address
	GetBlobCreatedIn() *big.Int
}

type TaikoBatchMetaDataPacaya interface {
	GetTxListHash() common.Hash
	GetExtraData() []byte
	GetCoinbase() common.Address
	GetBatchID() *big.Int
	GetGasLimit() uint32
	GetLastBlockTimestamp() uint64
	GetProposer() common.Address
	GetProposedAt() uint64
	GetProposedIn() uint64
	GetBlobCreatedIn() *big.Int
	GetTxListOffset() uint32
	GetTxListSize() uint32
	GetLastBlockID() uint64
	GetBlobHashes() []common.Hash
	GetAnchorBlockID() uint64
	GetAnchorBlockHash() common.Hash
	GetBlocks() []pacayaBindings.ITaikoInboxBlockParams
	GetBaseFeeConfig() *pacayaBindings.LibSharedDataBaseFeeConfig
	GetRawBlockHeight() *big.Int
	GetRawBlockHash() common.Hash
	GetTxIndex() uint
	GetTxHash() common.Hash
	InnerMetadata() *pacayaBindings.ITaikoInboxBatchMetadata
}

type MinimalPublicationData interface {
	Header() minimalBindings.IInboxPublicationHeader
	PublicationHash() common.Hash
	Attributes() MinimalPublicationAttributes
	GetRawBlockHeight() *big.Int
	GetRawBlockHash() common.Hash
	GetTxIndex() uint
	GetTxHash() common.Hash
}

type MinimalPublicationAttributes struct {
	Metadata MinimalInboxPublicationMetadata
	BlobRef  BlobRef
}

type BlobRef struct {
	BlockNumber *big.Int      `json:"blockNumber"` // uint256 → *big.Int
	Blobhashes  []common.Hash `json:"blobhashes"`  // bytes32[] → []common.Hash
}

type MinimalInboxPublicationMetadata struct {
	AnchorBlockId      *big.Int
	AnchorBlockHash    [32]byte
	IsDelayedInclusion bool
}
