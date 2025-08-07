package blobencoding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OpenZeppelin/minimal-rollup-client/client/pkg/utils"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// HeaderSeparator is not the safest separator.
var HeaderSeparator = byte('|')

type MinimalRollupBlobData struct {
	Attributes   []Attribute // any attribute: version, L1/L2 chain/block params etc.
	Transactions types.Transactions
}

type Attribute struct {
	Key   string
	Value any
}

// EncodeMinimalRollupBlob encodes rollup blob data into a blob.
func EncodeMinimalRollupBlob(data *MinimalRollupBlobData) ([]*eth.Blob, error) {
	txB, err := rlp.EncodeToBytes(data.Transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to RLP encode transactions: %w", err)
	}
	attrsB, err := encodeAttributes(data.Attributes)
	if err != nil {
		return nil, fmt.Errorf("failed to encode rollup attrs: %w", err)
	}
	b := append(append(attrsB, HeaderSeparator), txB...)
	b, err = utils.Compress(b)
	if err != nil {
		return nil, err
	}
	return splitToBlobs(b)
}

// DecodeMinimalRollupBlob decodes a rollup blob
func DecodeMinimalRollupBlob(blobs []*eth.Blob) (*MinimalRollupBlobData, error) {
	b, err := stitchBlobs(blobs)
	if err != nil {
		return nil, err
	}
	b, err = utils.Decompress(b)
	if err != nil {
		return nil, err
	}
	sepIndex := bytes.IndexByte(b, HeaderSeparator)
	if sepIndex < 0 {
		return nil, errors.New("invalid blob: failed to find attr separator")
	}
	attrsB := b[:sepIndex]
	txB := b[sepIndex+1:]
	var txs types.Transactions
	err = rlp.DecodeBytes(txB, &txs)
	if err != nil {
		return nil, fmt.Errorf("failed to RLP decode transactions: %w", err)
	}
	attrs, err := decodeAttributes(attrsB)
	if err != nil {
		return nil, fmt.Errorf("failed to decode rollup attrs: %w", err)
	}
	return &MinimalRollupBlobData{
		Attributes:   attrs,
		Transactions: txs,
	}, nil
}

// splitToBlobs splits the txListBytes into multiple blobs.
func splitToBlobs(txListBytes []byte) ([]*eth.Blob, error) {
	var blobs []*eth.Blob
	for start := 0; start < len(txListBytes); start += eth.MaxBlobDataSize {
		end := start + eth.MaxBlobDataSize
		if end > len(txListBytes) {
			end = len(txListBytes)
		}

		var blob = &eth.Blob{}
		if err := blob.FromData(txListBytes[start:end]); err != nil {
			return nil, err
		}

		blobs = append(blobs, blob)
	}

	return blobs, nil
}

func stitchBlobs(blobs []*eth.Blob) ([]byte, error) {
	if len(blobs) == 0 {
		return []byte{}, nil
	}

	estimatedSize := len(blobs) * eth.MaxBlobDataSize
	result := make([]byte, 0, estimatedSize)

	for _, blob := range blobs {
		data, err := blob.ToData()
		if err != nil {
			return nil, fmt.Errorf("failed to extract data from blob: %w", err)
		}
		result = append(result, data...)
	}

	return result, nil
}

// We are probably not going to use JSON encoding for blobs
// and we are just picking a convenient encoding below for now.
// TODO: Use CBOR?

func encodeAttributes(attrs []Attribute) ([]byte, error) {
	return json.Marshal(attrs)
}

func decodeAttributes(b []byte) (attrs []Attribute, err error) {
	err = json.Unmarshal(b, &attrs)
	return
}
