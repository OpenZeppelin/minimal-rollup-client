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
  just clean &&
  just compile &&
  cd -

ABIGEN_BIN=$TAIKO_GETH_DIR/build/bin/abigen
PACKAGE_NAME=minimal

echo ""
echo "Start generating Go bindings for minimal-rollup interfaces..."
echo ""

mkdir -p $DIR/../bindings/${PACKAGE_NAME}

INTERFACES=$(find ../contracts/src/protocol -name "I*.sol" -type f -exec basename {} .sol \;)

for interface in $INTERFACES; do
  echo "Processing $interface..."

  JSON_FILES=$(find ../contracts/out/${interface}.sol -name "*.json" -type f)

  for json_file in $JSON_FILES; do
    contract_name=$(basename "$json_file" .json)

    echo "  Generating bindings for $contract_name..."

    output_name=$(echo ${contract_name#I} | tr '[:upper:]' '[:lower:]')

    cat "$json_file" |
      jq .abi |
      ${ABIGEN_BIN} --abi - --type ${contract_name} --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_${output_name}.go
  done
done


# this one is not found in src/protocol
cat ../contracts/out/IBlobRefRegistry.sol/IBlobRefRegistry.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type IBlobRefRegistry --pkg ${PACKAGE_NAME} --out $DIR/../bindings/${PACKAGE_NAME}/gen_blob_ref_registry.go

echo "🍻 Go contract bindings for minimal-rollup interfaces generated!"
