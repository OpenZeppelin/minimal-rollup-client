package processor

import (
	"context"
	"testing"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func Test_waitForConfirmations(t *testing.T) {
	p := newTestProcessor(true)

	err := p.waitForConfirmations(context.TODO(), mock.SucceedTxHash)
	assert.Nil(t, err)
}
