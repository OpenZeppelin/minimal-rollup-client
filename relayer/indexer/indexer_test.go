package indexer

import (
	"context"
	"time"

	"github.com/OpenZeppelin/minimal-rollup-client/relayer"
	"github.com/OpenZeppelin/minimal-rollup-client/relayer/pkg/mock"
)

func newTestService(syncMode SyncMode, watchMode WatchMode) (*Indexer, relayer.Bridge) {
	b := &mock.Bridge{}

	return &Indexer{
		eventRepo:     &mock.EventRepository{},
		bridge:        b,
		destBridge:    b,
		signalService: &mock.SignalService{},
		srcEthClient:  &mock.EthClient{},
		numGoroutines: 10,

		latestIndexedBlockNumber: 0,
		blockBatchSize:           100,

		queue: &mock.Queue{},

		syncMode:  syncMode,
		watchMode: watchMode,

		ctx: context.Background(),

		srcChainId:  mock.MockChainID,
		destChainId: mock.MockChainID,

		ethClientTimeout: 10 * time.Second,
		eventName:        relayer.EventNameMessageSent,
	}, b
}
