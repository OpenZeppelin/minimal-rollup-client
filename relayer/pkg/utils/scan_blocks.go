package utils

import (
	"context"
	"sync"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type headSubscriber interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

func ScanBlocks(ctx context.Context, ethClient headSubscriber, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	headers := make(chan *types.Header)

	sub, err := ethClient.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case <-headers:
			relayer.BlocksScanned.Inc()
		}
	}
}
