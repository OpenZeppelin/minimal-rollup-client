package metadata

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
)

var (
	BlobRefComponent = []abi.ArgumentMarshaling{
		{Name: "blockNumber", Type: "uint256"},
		{Name: "blobhashes", Type: "bytes32[]"},
	}
	BlobRefType, _      = abi.NewType("tuple", "IInbox.BlobRef", BlobRefComponent)
	BlobRefArgumentType = abi.Arguments{{Name: "blobRef", Type: BlobRefType}}

	PublicationMetadataComponent = []abi.ArgumentMarshaling{
		{Name: "anchorBlockId", Type: "uint64"},
		{Name: "anchorBlockHash", Type: "bytes32"},
		{Name: "isDelayedInclusion", Type: "bool"},
	}
	PublicationMetadataType, _      = abi.NewType("tuple", "IInbox.PublicationMetadata", PublicationMetadataComponent)
	PublicationMetadataArgumentType = abi.Arguments{{Name: "publicationMetadata", Type: PublicationMetadataType}}
)

func convertToHashes(data [][32]uint8) []common.Hash {
	hashes := make([]common.Hash, len(data))
	for i, item := range data {
		hashes[i] = item
	}
	return hashes
}

func DecodeBlobRef(data []byte) (*BlobRef, error) {
	out, err := BlobRefArgumentType.Unpack(data)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(out[0])
	blockNumber := v.FieldByName("BlockNumber").Interface().(*big.Int)
	blobhashes := v.FieldByName("Blobhashes").Interface().([][32]uint8)

	result := &BlobRef{
		BlockNumber: blockNumber,
		Blobhashes:  convertToHashes(blobhashes),
	}

	return result, nil
}

func DecodePublicationMetadata(data []byte) (*MinimalInboxPublicationMetadata, error) {
	out, err := PublicationMetadataArgumentType.Unpack(data)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(out[0])
	anchorBlockId := v.FieldByName("AnchorBlockId").Interface().(uint64)
	anchorBlockHash := v.FieldByName("AnchorBlockHash").Interface().([32]uint8)
	isDelayedInclusion := v.FieldByName("IsDelayedInclusion").Interface().(bool)

	result := &MinimalInboxPublicationMetadata{
		AnchorBlockId:      big.NewInt(int64(anchorBlockId)),
		AnchorBlockHash:    anchorBlockHash,
		IsDelayedInclusion: isDelayedInclusion,
	}

	return result, nil
}
