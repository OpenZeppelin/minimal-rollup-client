package submitter

import (
	"context"
	"errors"

	"github.com/OpenZeppelin/minimal-rollup-client/client/bindings/metadata"
	proofProducer "github.com/OpenZeppelin/minimal-rollup-client/client/prover/proof_producer"
)

var (
	_               Submitter = (*ProofSubmitterPacaya)(nil)
	ErrInvalidProof           = errors.New("invalid proof found")
)

// Submitter is the interface for submitting proofs of the L2 blocks.
type Submitter interface {
	RequestProof(ctx context.Context, meta metadata.TaikoProposalMetaData) error
	BatchSubmitProofs(ctx context.Context, proofsWithHeaders *proofProducer.BatchProofs) error
	AggregateProofsByType(ctx context.Context, proofType proofProducer.ProofType) error
}
