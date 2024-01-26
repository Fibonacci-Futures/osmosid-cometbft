// Deprecated: Priority mempool will be removed in the next major release.
package v1

import (
	"github.com/cometbft/cometbft/types"
)

func (txmp *TxMempool) PublishUnconfirmedTx(wtx *WrappedTx) {
	txmp.eventBus.PublishUnconfirmedTx(types.UnconfirmedTx{
		TxData:    wtx.tx,
		Hash:      wtx.hash,
		Height:    wtx.height,
		Timestamp: wtx.timestamp,
		Sender:    wtx.sender,
	})
}
