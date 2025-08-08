package blobencoding

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeRollupBlob(t *testing.T) {
	// Create test data
	testCases := []struct {
		name string
		data *MinimalRollupBlobData
	}{
		{
			name: "Empty data",
			data: &MinimalRollupBlobData{
				Attributes:   []Attribute{},
				Transactions: types.Transactions{},
			},
		},
		{
			name: "With attributes only",
			data: &MinimalRollupBlobData{
				Attributes: []Attribute{
					{Key: "version", Value: "1.0"},
					{Key: "chainId", Value: 1337},
					{Key: "blockNumber", Value: 42},
				},
				Transactions: types.Transactions{},
			},
		},
		{
			name: "With transactions only",
			data: &MinimalRollupBlobData{
				Attributes: []Attribute{},
				Transactions: types.Transactions{
					types.NewTransaction(
						0,
						common.HexToAddress("0x1234567890123456789012345678901234567890"),
						common.Big0,
						21000,
						common.Big1,
						[]byte("test data"),
					),
				},
			},
		},
		{
			name: "With attributes and transactions",
			data: &MinimalRollupBlobData{
				Attributes: []Attribute{
					{Key: "version", Value: "1.0"},
					{Key: "chainId", Value: 1337},
				},
				Transactions: types.Transactions{
					types.NewTransaction(
						1,
						common.HexToAddress("0x1234567890123456789012345678901234567890"),
						common.Big1,
						21000,
						common.Big2,
						[]byte("transaction 1"),
					),
					types.NewTransaction(
						2,
						common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
						common.Big2,
						30000,
						common.Big1,
						[]byte("transaction 2"),
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encode the data
			encoded, err := EncodeMinimalRollupBlob(tc.data)
			require.NoError(t, err, "Failed to encode rollup blob")

			// Ensure encoded data is not empty
			require.NotEmpty(t, encoded, "Encoded data is empty")

			// Decode the data
			decoded, err := DecodeMinimalRollupBlob(encoded)
			require.NoError(t, err, "Failed to decode rollup blob")

			// Verify attributes
			require.Equal(t, len(tc.data.Attributes), len(decoded.Attributes),
				"Attribute count mismatch")

			// Compare attributes (note: this is a simplified comparison)
			for i, attr := range tc.data.Attributes {
				if i < len(decoded.Attributes) {
					require.Equal(t, attr.Key, decoded.Attributes[i].Key,
						"Attribute key mismatch at index %d", i)

					// For value comparison, we need to handle the fact that JSON unmarshal
					// converts integers to float64
					origValue := reflect.ValueOf(attr.Value)
					decodedValue := reflect.ValueOf(decoded.Attributes[i].Value)

					// Special handling for numeric types
					if origValue.Kind() == reflect.Int && decodedValue.Kind() == reflect.Float64 {
						// Compare int as float64
						require.Equal(t, float64(origValue.Int()), decodedValue.Float(),
							"Numeric attribute value mismatch at index %d", i)
					} else if origValue.Type() != decodedValue.Type() {
						// For other type mismatches, compare as strings
						require.Equal(t, fmt.Sprint(attr.Value), fmt.Sprint(decoded.Attributes[i].Value),
							"Attribute value mismatch at index %d", i)
					} else {
						// Direct comparison for same types
						require.Equal(t, attr.Value, decoded.Attributes[i].Value,
							"Attribute value mismatch at index %d", i)
					}
				}
			}

			// Verify transactions
			require.Equal(t, len(tc.data.Transactions), len(decoded.Transactions),
				"Transaction count mismatch")

			// Compare transactions
			for i, tx := range tc.data.Transactions {
				if i < len(decoded.Transactions) {
					decodedTx := decoded.Transactions[i]

					// Compare transaction fields
					require.Equal(t, tx.Hash(), decodedTx.Hash(),
						"Transaction hash mismatch at index %d", i)

					require.True(t, bytes.Equal(tx.Data(), decodedTx.Data()),
						"Transaction data mismatch at index %d", i)

					require.Equal(t, tx.To().Hex(), decodedTx.To().Hex(),
						"Transaction recipient mismatch at index %d", i)

					require.Equal(t, tx.Nonce(), decodedTx.Nonce(),
						"Transaction nonce mismatch at index %d", i)
				}
			}
		})
	}
}
