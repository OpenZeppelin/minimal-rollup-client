#!/bin/bash

# Generate go contract bindings for minimal-rollup interfaces.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "TAIKO_GETH_DIR: ${TAIKO_GETH_DIR}"
echo ""

cd ${TAIKO_GETH_DIR} &&
  make all &&
  cd -

cd ../contracts &&
  make clean &&
  make compile &&
  cd -

ABIGEN_BIN=$TAIKO_GETH_DIR/build/bin/abigen
PACKAGE_NAME=minimal

echo ""
echo "Start generating Go bindings for minimal-rollup interfaces..."
echo ""

mkdir -p $DIR/../bindings/${PACKAGE_NAME}

CONTRACTS=../contracts/out/

echo "Generating bindings for TaikoInbox..."
cat $CONTRACTS/TaikoInbox.sol/TaikoInbox.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type TaikoInbox --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_inbox.go

echo "Generating bindings for CheckpointTracker..."
cat $CONTRACTS/CheckpointTracker.sol/CheckpointTracker.json |
  jq .abi |
  ${ABIGEN_BIN} --abi - --type CheckpointTracker --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_checkpoint_tracker.go

echo "Generating bindings for DelayedInclusionStore..."
cat $CONTRACTS/DelayedInclusionStore.sol/DelayedInclusionStore.json |
  jq .abi |
  ${ABIGEN_BIN} --abi - --type DelayedInclusionStore --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_delayed_inclusion_store.go

echo "Generating bindings for BlobRefRegistry..."
cat ../contracts/out/BlobRefRegistry.sol/BlobRefRegistry.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type BlobRefRegistry --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_blob_ref_registry.go

echo "🍻 Go contract bindings for minimal-rollup interfaces generated!"
