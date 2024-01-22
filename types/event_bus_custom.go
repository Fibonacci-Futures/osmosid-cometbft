package types

import (
	"context"
	"time"
)

// mempoolv1 "github.com/cometbft/cometbft/mempool/v1" //nolint:staticcheck // SA1019 Priority mempool deprecated but still supported in this release.

const (
	UnConfirmedTxsKey = "utx.hash"
)

type UnconfirmedTx struct {
	TxData    Tx        // the original transaction data
	Hash      TxKey     // the transaction hash
	Height    int64     // height when this transaction was initially checked (for expiry)
	Timestamp time.Time // time when transaction was entered (for TTL)

	// gasWanted int64           // app: gas required to execute this transaction
	// priority  int64           // app: priority value for this transaction
	Sender string // app: assigned sender label
	// peers     map[uint16]bool // peer IDs who have sent us this transaction
}

func (b *EventBus) PublishUnconfirmedTx(data string) error {
	// no explicit deadline for publishing events
	ctx := context.Background()
	events := make(map[string][]string)

	events[UnConfirmedTxsKey] = append(events[UnConfirmedTxsKey], "1") //utx.hash=1 for now
	// data := UnconfirmedTx{Hash: hash}
	return b.pubsub.PublishWithEvents(ctx, data, events)
}

/*
func (b *EventBus) PublishUnconfirmedTx2(wtx *mempoolv1.WrappedTx) error {
	// no explicit deadline for publishing events
	ctx := context.Background()
	events := make(map[string][]string)

	events[UnConfirmedTxsKey] = append(events[UnConfirmedTxsKey], "1") //utx.hash=1 for now
	data := UTx{Hash: fmt.Sprintf("%X", wtx.Hash())}
	return b.pubsub.PublishWithEvents(ctx, data, events)
}
func (b *EventBus) PublishUnconfirmedTx1(wtx *mempoolv1.WrappedTx) error {
	data := UTx{Hash: fmt.Sprintf("%X", wtx.Hash())}
	return b.Publish(EventValidBlock, data)
}
*/
